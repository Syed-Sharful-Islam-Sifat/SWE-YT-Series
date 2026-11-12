// N+1 lab: one HTTP server, one endpoint per strategy.
// Every response carries X-Query-Count, X-Rows and X-Duration-Ms headers,
// and every request prints one log line, so you can compare strategies.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ---------- per-request query counter (this is how you DETECT N+1) ----------

type ctxKey struct{}

type stats struct {
	queries atomic.Int64
	rows    atomic.Int64
}

func statsFrom(ctx context.Context) *stats {
	s, _ := ctx.Value(ctxKey{}).(*stats)
	return s
}

// tracer is called by the driver around every query.
type tracer struct {
	latency time.Duration // simulated network round trip to the database
	logSQL  bool
}

func (t tracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, d pgx.TraceQueryStartData) context.Context {
	if s := statsFrom(ctx); s != nil {
		s.queries.Add(1)
	}
	if t.logSQL {
		log.Printf("  SQL: %s %v", strings.Join(strings.Fields(d.SQL), " "), d.Args)
	}
	if t.latency > 0 {
		// The connection is held during this sleep, exactly like a real network round trip.
		time.Sleep(t.latency)
	}
	return ctx
}

func (t tracer) TraceQueryEnd(ctx context.Context, _ *pgx.Conn, d pgx.TraceQueryEndData) {
	if s := statsFrom(ctx); s != nil {
		s.rows.Add(d.CommandTag.RowsAffected())
	}
}

// ---------- HTTP wiring ----------

type strategy func(ctx context.Context, db *pgxpool.Pool, limit int) (any, error)

func handle(db *pgxpool.Pool, name string, fn strategy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil || limit <= 0 {
			limit = 50
		}
		if limit > 1000 {
			limit = 1000 // always cap page size: it caps the N in N+1
		}

		st := &stats{}
		ctx := context.WithValue(r.Context(), ctxKey{}, st)
		start := time.Now()
		res, err := fn(ctx, db, limit)
		dur := time.Since(start)
		if err != nil {
			log.Printf("%s: %v", name, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ms := float64(dur.Microseconds()) / 1000
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Query-Count", strconv.FormatInt(st.queries.Load(), 10))
		w.Header().Set("X-Rows", strconv.FormatInt(st.rows.Load(), 10))
		w.Header().Set("X-Duration-Ms", fmt.Sprintf("%.1f", ms))
		_ = json.NewEncoder(w).Encode(res)

		log.Printf("%-22s limit=%-4d queries=%-5d rows=%-6d time=%9.1fms",
			name, limit, st.queries.Load(), st.rows.Load(), ms)
	}
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func newPool(ctx context.Context) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(env("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/shop"))
	if err != nil {
		return nil, err
	}
	latencyMs, _ := strconv.Atoi(env("DB_LATENCY_MS", "0"))
	poolSize, _ := strconv.Atoi(env("POOL_SIZE", "10"))
	cfg.MaxConns = int32(poolSize)
	cfg.ConnConfig.Tracer = tracer{
		latency: time.Duration(latencyMs) * time.Millisecond,
		logSQL:  env("LOG_SQL", "") == "1",
	}
	log.Printf("pool size=%d, simulated DB latency=%dms", poolSize, latencyMs)
	return pgxpool.NewWithConfig(ctx, cfg)
}

var routes = []struct {
	path string
	fn   strategy
}{
	{"/v1/naive", naive},                  // Step 1: the N+1
	{"/v2/join", joinAll},                 // Step 3: one JOIN
	{"/v3/batch", batchIn},                // Step 4: two queries, WHERE IN / ANY (eager loading)
	{"/v4/json", jsonAgg},                 // Step 5: one query, nested JSON built by the DB
	{"/v5/top3-naive", top3Naive},         // Step 6: "latest 3 reviews" the N+1 way
	{"/v5/top3-lateral", top3Lateral},     // Step 6: LATERAL join
	{"/v5/top3-window", top3Window},       // Step 6: window function
	{"/v6/counts-naive", countsNaive},     // Step 7: COUNT per product
	{"/v6/counts-groupby", countsGroupBy}, // Step 7: one GROUP BY
	{"/v6/counts-column", countsColumn},   // Step 7: denormalized counter column
	{"/v7/nested-naive", nestedNaive},     // Step 8: products -> reviews -> users, 1+N+M
	{"/v7/nested-loader", nestedLoader},   // Step 8: DataLoader
	{"/v8/parallel-naive", parallelNaive}, // Step 9: "just run them in parallel"
}

func main() {
	ctx := context.Background()
	db, err := newPool(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	for _, rt := range routes {
		mux.HandleFunc("GET "+rt.path, handle(db, rt.path, rt.fn))
	}
	addr := env("ADDR", ":8080")
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
