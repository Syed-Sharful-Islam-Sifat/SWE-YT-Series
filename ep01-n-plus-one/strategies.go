package main

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Review struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name,omitempty"`
	Rating    int       `json:"rating"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	PriceCents  int      `json:"price_cents"`
	ReviewCount *int     `json:"review_count,omitempty"`
	Reviews     []Review `json:"reviews,omitempty"`
}

const reviewCols = `id, user_id, rating, body, created_at`

// ---------- shared helpers ----------

func listProducts(ctx context.Context, db *pgxpool.Pool, limit int) ([]Product, error) {
	rows, err := db.Query(ctx, `SELECT id, name, price_cents FROM products ORDER BY id LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.PriceCents); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

func reviewsFor(ctx context.Context, db *pgxpool.Pool, productID int) ([]Review, error) {
	rows, err := db.Query(ctx,
		`SELECT `+reviewCols+` FROM reviews WHERE product_id = $1 ORDER BY created_at DESC`, productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Review
	for rows.Next() {
		var r Review
		if err := rows.Scan(&r.ID, &r.UserID, &r.Rating, &r.Body, &r.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, rows.Err()
}

func productIDs(ps []Product) []int {
	ids := make([]int, len(ps))
	for i, p := range ps {
		ids[i] = p.ID
	}
	return ids
}

// reviewsByProduct runs ONE query for many products and groups the result.
// The SQL is passed in so the batch and window strategies can share it.
func reviewsByProduct(ctx context.Context, db *pgxpool.Pool, sql string, ids []int) (map[int][]Review, error) {
	rows, err := db.Query(ctx, sql, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := make(map[int][]Review, len(ids))
	for rows.Next() {
		var pid int
		var r Review
		if err := rows.Scan(&pid, &r.ID, &r.UserID, &r.Rating, &r.Body, &r.CreatedAt); err != nil {
			return nil, err
		}
		out[pid] = append(out[pid], r)
	}
	return out, rows.Err()
}

// scanJoined turns flat "product x review" rows back into nested products.
// Used by both JOIN strategies. Review columns are NULL for products without reviews.
func scanJoined(ctx context.Context, db *pgxpool.Pool, sql string, limit int) ([]Product, error) {
	rows, err := db.Query(ctx, sql, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Product
	for rows.Next() {
		var p Product
		var rid, uid, rating *int
		var body *string
		var created *time.Time
		if err := rows.Scan(&p.ID, &p.Name, &p.PriceCents, &rid, &uid, &rating, &body, &created); err != nil {
			return nil, err
		}
		if len(out) == 0 || out[len(out)-1].ID != p.ID {
			out = append(out, p)
		}
		if rid != nil {
			last := &out[len(out)-1]
			last.Reviews = append(last.Reviews, Review{ID: *rid, UserID: *uid, Rating: *rating, Body: *body, CreatedAt: *created})
		}
	}
	return out, rows.Err()
}

// ---------- Step 1: the N+1 ----------

func naive(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	products, err := listProducts(ctx, db, limit) // 1 query
	if err != nil {
		return nil, err
	}
	for i := range products { // + N queries
		if products[i].Reviews, err = reviewsFor(ctx, db, products[i].ID); err != nil {
			return nil, err
		}
	}
	return products, nil
}

// ---------- Step 3: one JOIN ----------
// The LIMIT goes in a subquery: "LIMIT 50" on the joined rows would give
// 50 reviews, not 50 products. A classic bug with JOIN + pagination.

func joinAll(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	return scanJoined(ctx, db, `
		SELECT p.id, p.name, p.price_cents, r.id, r.user_id, r.rating, r.body, r.created_at
		FROM (SELECT id, name, price_cents FROM products ORDER BY id LIMIT $1) p
		LEFT JOIN reviews r ON r.product_id = p.id
		ORDER BY p.id, r.created_at DESC`, limit)
}

// ---------- Step 4: batch with WHERE IN / ANY (what ORMs call eager loading / preload) ----------

func batchIn(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	products, err := listProducts(ctx, db, limit) // query 1
	if err != nil {
		return nil, err
	}
	byProduct, err := reviewsByProduct(ctx, db, // query 2
		`SELECT product_id, `+reviewCols+` FROM reviews
		 WHERE product_id = ANY($1) ORDER BY product_id, created_at DESC`, productIDs(products))
	if err != nil {
		return nil, err
	}
	for i := range products {
		products[i].Reviews = byProduct[products[i].ID]
	}
	return products, nil
}

// ---------- Step 5: one query, the database builds the nested JSON ----------

func jsonAgg(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	rows, err := db.Query(ctx, `
		SELECT p.id, p.name, p.price_cents,
		  COALESCE((
		    SELECT json_agg(json_build_object(
		      'id', r.id, 'user_id', r.user_id, 'rating', r.rating,
		      'body', r.body, 'created_at', r.created_at) ORDER BY r.created_at DESC)
		    FROM reviews r WHERE r.product_id = p.id
		  ), '[]'::json)
		FROM products p ORDER BY p.id LIMIT $1`, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Product
	for rows.Next() {
		var p Product
		var raw []byte
		if err := rows.Scan(&p.ID, &p.Name, &p.PriceCents, &raw); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(raw, &p.Reviews); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}

// ---------- Step 6: only the latest 3 reviews per product ----------

func top3Naive(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	products, err := listProducts(ctx, db, limit)
	if err != nil {
		return nil, err
	}
	for i := range products {
		rows, err := db.Query(ctx, `SELECT `+reviewCols+` FROM reviews
			WHERE product_id = $1 ORDER BY created_at DESC LIMIT 3`, products[i].ID)
		if err != nil {
			return nil, err
		}
		for rows.Next() {
			var r Review
			if err := rows.Scan(&r.ID, &r.UserID, &r.Rating, &r.Body, &r.CreatedAt); err != nil {
				rows.Close()
				return nil, err
			}
			products[i].Reviews = append(products[i].Reviews, r)
		}
		rows.Close()
		if err := rows.Err(); err != nil {
			return nil, err
		}
	}
	return products, nil
}

func top3Lateral(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	return scanJoined(ctx, db, `
		SELECT p.id, p.name, p.price_cents, r.id, r.user_id, r.rating, r.body, r.created_at
		FROM (SELECT id, name, price_cents FROM products ORDER BY id LIMIT $1) p
		LEFT JOIN LATERAL (
		  SELECT `+reviewCols+` FROM reviews
		  WHERE product_id = p.id ORDER BY created_at DESC LIMIT 3
		) r ON true
		ORDER BY p.id, r.created_at DESC`, limit)
}

func top3Window(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	products, err := listProducts(ctx, db, limit)
	if err != nil {
		return nil, err
	}
	byProduct, err := reviewsByProduct(ctx, db, `
		SELECT product_id, `+reviewCols+` FROM (
		  SELECT r.*, row_number() OVER (PARTITION BY product_id ORDER BY created_at DESC) AS rn
		  FROM reviews r WHERE product_id = ANY($1)
		) t WHERE rn <= 3 ORDER BY product_id, created_at DESC`, productIDs(products))
	if err != nil {
		return nil, err
	}
	for i := range products {
		products[i].Reviews = byProduct[products[i].ID]
	}
	return products, nil
}

// ---------- Step 7: counts instead of full lists ----------

func countsNaive(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	products, err := listProducts(ctx, db, limit)
	if err != nil {
		return nil, err
	}
	for i := range products {
		var n int
		if err := db.QueryRow(ctx, `SELECT count(*) FROM reviews WHERE product_id = $1`, products[i].ID).Scan(&n); err != nil {
			return nil, err
		}
		products[i].ReviewCount = &n
	}
	return products, nil
}

func scanCounts(ctx context.Context, db *pgxpool.Pool, sql string, limit int) (any, error) {
	rows, err := db.Query(ctx, sql, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Product
	for rows.Next() {
		var p Product
		var n int
		if err := rows.Scan(&p.ID, &p.Name, &p.PriceCents, &n); err != nil {
			return nil, err
		}
		p.ReviewCount = &n
		out = append(out, p)
	}
	return out, rows.Err()
}

func countsGroupBy(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	return scanCounts(ctx, db, `
		SELECT p.id, p.name, p.price_cents, count(r.id)
		FROM (SELECT id, name, price_cents FROM products ORDER BY id LIMIT $1) p
		LEFT JOIN reviews r ON r.product_id = p.id
		GROUP BY p.id, p.name, p.price_cents
		ORDER BY p.id`, limit)
}

func countsColumn(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	return scanCounts(ctx, db,
		`SELECT id, name, price_cents, review_count FROM products ORDER BY id LIMIT $1`, limit)
}

// ---------- Step 8: nested data, products -> reviews -> user names ----------
// The code is written as small independent "resolvers", like a GraphQL server
// or a service layer where each function loads what it needs.

func userName(ctx context.Context, db *pgxpool.Pool, id int) (string, error) {
	var name string
	err := db.QueryRow(ctx, `SELECT name FROM users WHERE id = $1`, id).Scan(&name)
	return name, err
}

func nestedNaive(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	products, err := listProducts(ctx, db, limit)
	if err != nil {
		return nil, err
	}
	for i := range products {
		reviews, err := reviewsFor(ctx, db, products[i].ID) // N
		if err != nil {
			return nil, err
		}
		for j := range reviews {
			if reviews[j].UserName, err = userName(ctx, db, reviews[j].UserID); err != nil { // M
				return nil, err
			}
		}
		products[i].Reviews = reviews
	}
	return products, nil
}

func nestedLoader(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	// Loaders live for ONE request only (their cache must not leak between users).
	reviewLoader := NewLoader(2*time.Millisecond, func(ctx context.Context, ids []int) (map[int][]Review, error) {
		return reviewsByProduct(ctx, db,
			`SELECT product_id, `+reviewCols+` FROM reviews
			 WHERE product_id = ANY($1) ORDER BY product_id, created_at DESC`, ids)
	})
	userLoader := NewLoader(2*time.Millisecond, func(ctx context.Context, ids []int) (map[int]string, error) {
		rows, err := db.Query(ctx, `SELECT id, name FROM users WHERE id = ANY($1)`, ids)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		out := make(map[int]string, len(ids))
		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				return nil, err
			}
			out[id] = name
		}
		return out, rows.Err()
	})

	products, err := listProducts(ctx, db, limit)
	if err != nil {
		return nil, err
	}

	// The resolvers look exactly like the naive ones: "load reviews for THIS product",
	// "load the user for THIS review". The loader turns them into batches.
	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstErr error
	setErr := func(err error) {
		mu.Lock()
		if firstErr == nil {
			firstErr = err
		}
		mu.Unlock()
	}
	for i := range products {
		wg.Add(1)
		go func(p *Product) {
			defer wg.Done()
			reviews, err := reviewLoader.Load(ctx, p.ID)
			if err != nil {
				setErr(err)
				return
			}
			reviews = append([]Review(nil), reviews...) // copy: the loader's slice is shared
			var inner sync.WaitGroup
			for j := range reviews {
				inner.Add(1)
				go func(r *Review) {
					defer inner.Done()
					name, err := userLoader.Load(ctx, r.UserID)
					if err != nil {
						setErr(err)
						return
					}
					r.UserName = name
				}(&reviews[j])
			}
			inner.Wait()
			p.Reviews = reviews
		}(&products[i])
	}
	wg.Wait()
	return products, firstErr
}

// ---------- Step 9: "just run the N queries in parallel" ----------

func parallelNaive(ctx context.Context, db *pgxpool.Pool, limit int) (any, error) {
	products, err := listProducts(ctx, db, limit)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	errs := make([]error, len(products))
	for i := range products {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			products[i].Reviews, errs[i] = reviewsFor(ctx, db, products[i].ID)
		}(i)
	}
	wg.Wait()
	for _, e := range errs {
		if e != nil {
			return nil, e
		}
	}
	return products, nil
}
