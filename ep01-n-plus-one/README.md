# Ep 1 Lab: A Page Runs 1,000 Queries (N+1)

A hands-on exercise that covers every common fix for the N+1 problem. You build it once, measure everything, and the video comes from your own numbers.

**What you need:** Go 1.22+, Docker, `curl`. Time: about 4-5 hours for all steps (the 1.5-2 day "full depth" plan), or about 2 hours for Steps 0-4 + 10 only.

**The rule for the whole lab:** never just say "faster". After every step, write down **queries, rows, and milliseconds** in the [results table](#results-table) at the bottom. Those numbers are your video.

---

## The map: every solution, in one picture

| # | Solution | Queries | Best when | Main trade-off | Step |
|---|---|---|---|---|---|
| 0 | Naive loop (the problem) | 1 + N | never in a list | grows with page size | 1 |
| 1 | **JOIN** | 1 | simple 1-to-1 or small 1-to-many | parent columns repeated on every row; LIMIT trap | 3 |
| 2 | **Batch: `WHERE id IN (...)`** (eager loading, preload) | 2 | the default fix, works everywhere (SQL, Mongo, HTTP APIs) | very large ID lists need chunking | 4 |
| 3 | **Database builds JSON** (`json_agg`) | 1 | read-heavy API endpoints | SQL gets complex; DB does more CPU work | 5 |
| 4 | **LATERAL / window function** | 1-2 | "top N children per parent" | more advanced SQL | 6 |
| 5 | **GROUP BY aggregate** | 1 | you need counts/sums, not rows | - | 7 |
| 6 | **Denormalized column** | 1 | the aggregate is read far more than written | must keep it correct on every write | 7 |
| 7 | **DataLoader** (batch + per-request cache) | ~1 per level | code is split into resolvers / layers (GraphQL, services) | adds a small wait + complexity | 8 |
| ✗ | Run N queries in parallel | 1 + N | not a fix | eats the connection pool (→ Ep 2) | 9 |
| ✗ | Cache each item | 1 + N on a cold cache | not a fix alone | hides the problem; stampede risk (→ Ep 3) | 9 |
| ✓ | **Supporting fixes:** index on the foreign key, page-size cap, query-count test | - | always | - | 4, 1, 10 |

---

## Step 0: Setup (15 min)

```bash
cd ep01-n-plus-one
docker compose up -d        # Postgres 16 with 500 users, 1,000 products, 50,000 reviews
go run .                    # API on :8080
```

Check it works:

```bash
curl -s "localhost:8080/v1/naive?limit=2" | head -c 400
```

Every response has three headers. The server also prints one log line per request (the numbers below only show the format):

```
/v1/naive    limit=50   queries=51    rows=2500   time=   210.3ms
```

**Settings (environment variables):**

| Variable | Default | What it does |
|---|---|---|
| `DB_LATENCY_MS` | 0 | adds a fake network delay to every query (see Step 2) |
| `POOL_SIZE` | 10 | max DB connections |
| `LOG_SQL` | off | `LOG_SQL=1` prints every SQL statement |

---

## Step 1: Build the problem and SEE it (20 min)

Open [strategies.go](strategies.go) and read `naive()`. It's the code every developer writes on day one:

```go
products := listProducts(limit)      // 1 query
for each product {
    product.Reviews = reviewsFor(id) // N queries
}
```

1. Run with SQL logging so you can watch it happen:
   ```bash
   LOG_SQL=1 go run .
   curl -s -o /dev/null "localhost:8080/v1/naive?limit=10"
   ```
   You'll see 1 `SELECT ... FROM products` followed by 10 identical `SELECT ... FROM reviews WHERE product_id = $1` lines. **That pattern in a log is the fingerprint of N+1.** Screenshot it for the video.

2. Measure how it grows. Stop the server, run it without `LOG_SQL`, then:
   ```bash
   ./bench.sh "/v1/naive" "10 50 200 1000"
   ```
   Write the numbers down. Queries grow as **1 + N**. Time grows with it.

3. Note what the page-size cap in `main.go` does (`limit > 1000`): **pagination is the first defense.** It doesn't fix N+1, but it limits the damage to a known maximum.

> **Why it's slow even though each query is fast:** each query pays a fixed cost (network round trip, parsing, planning, a connection hop) on top of the actual work. N+1 multiplies the fixed cost by N.

---

## Step 2: Make it realistic: add network latency (10 min)

On your laptop the database is on the same machine (round trip ~0.1 ms). In production it's on another server: commonly ~0.5-2 ms in the same data center, much more across regions.

```bash
DB_LATENCY_MS=1 go run .
./bench.sh "/v1/naive /v3/batch" "50 200 1000"
```

Then try `DB_LATENCY_MS=5` (a cross-region or busy-network case).

**What to notice:** naive time ≈ `(1 + N) × latency + work`. Batch time barely changes, because it pays the latency **twice**, not 1,001 times. This is the single most important chart for the video: **"same code, same data, the only difference is the distance to the database."**

> Be honest in the video that this is simulated with a sleep in the driver hook ([main.go](main.go), `TraceQueryStart`). The sleep holds the connection, just like a real round trip does.

---

## Step 3: Fix #1, the JOIN (20 min)

Read `joinAll()` in [strategies.go](strategies.go).

```bash
./bench.sh "/v1/naive /v2/join" "50 200"
```

1. **Queries:** 1, whatever the page size.
2. **Look at `rows`:** a JOIN returns one row per *review*, with the product's name and price **repeated on every row**. For 50 products × ~50 reviews that's ~2,500 rows. With a wide parent table (a long description, many columns), that duplication is real network and memory cost. This is why ORMs often prefer two queries (Step 4) over one big JOIN.
3. **The LIMIT trap.** Try this in `psql` (`docker compose exec db psql -U postgres shop`):
   ```sql
   -- WRONG: limits review rows, not products. You get 10 rows, maybe all from product 1.
   SELECT p.id, r.id FROM products p LEFT JOIN reviews r ON r.product_id = p.id
   ORDER BY p.id LIMIT 10;
   ```
   The code puts the LIMIT inside a subquery (`FROM (SELECT ... LIMIT $1) p`). This is a classic real bug when people "fix" N+1 with a join and pagination breaks.
4. **Multiple children trap (think, don't build):** if a product had reviews *and* images and you JOIN both, you get reviews × images rows per product (a "cartesian explosion"). Joining two independent 1-to-many relations in one query is usually wrong. Use batching instead.

---

## Step 4: Fix #2, batching with `IN` / `ANY`, plus the missing index (40 min)

Read `batchIn()`. Query 1 gets the products. Query 2 gets **all** their reviews at once:

```sql
SELECT product_id, ... FROM reviews WHERE product_id = ANY($1)   -- $1 = [1,2,3,...,50]
```

Then the Go code groups them by `product_id` in a map. **This is exactly what ORMs do** when you ask for eager loading:

| Tool | Eager loading call |
|---|---|
| GORM (Go) | `db.Preload("Reviews").Find(&products)` |
| Prisma (Node) | `findMany({ include: { reviews: true } })` |
| Sequelize / TypeORM | `include: [Review]` / `relations: ["reviews"]` |
| Mongoose (MongoDB) | `.populate("reviews")` (runs a `$in` query under the hood) |
| Django / Rails | `prefetch_related` / `includes` |

Turn on `LOG_SQL=1` and run your ORM's version once in your own side project if you can. Seeing the ORM generate the `IN (...)` query is a great video moment.

**Now the hidden second problem: the index.**

1. In `psql`:
   ```sql
   EXPLAIN ANALYZE SELECT * FROM reviews WHERE product_id = 42;
   ```
   You'll see `Seq Scan on reviews`: it reads all 50,000 rows to find ~50. **Postgres does not index foreign keys automatically** (MySQL/InnoDB does). So the naive endpoint was actually doing 50 full table scans.
2. Add the index:
   ```sql
   CREATE INDEX reviews_product_created_idx ON reviews (product_id, created_at DESC);
   ```
3. Run `EXPLAIN ANALYZE` again (now `Index Scan`), and re-run:
   ```bash
   ./bench.sh "/v1/naive /v2/join /v3/batch /v4/json" "50 200 1000"
   ```

**What to notice:** the index makes *every* strategy faster, including naive. But with `DB_LATENCY_MS=1` naive is still slow, because the index fixes the work per query, **not the number of round trips.** Two different problems, two different fixes. Say this clearly in the video.

The index is `(product_id, created_at DESC)`, not just `(product_id)`. It also serves the `ORDER BY created_at DESC` and the "latest 3" queries in Step 6.

**Very large ID lists:** with thousands of IDs, split them into chunks (for example 500-1,000 per query). Some systems have hard limits (Firestore `in` queries accept only a small number of values per query, so check the current docs; SQL Server caps parameters at about 2,100). Postgres `= ANY($1)` with an array is one parameter, so it handles big lists better than `IN ($1,$2,...)`.

---

## Step 5: Fix #3, let the database build the JSON (20 min)

Read `jsonAgg()`. One query, and Postgres returns each product with its reviews already nested as a JSON array (`json_agg` + `json_build_object`). MySQL has `JSON_ARRAYAGG`; MongoDB's equivalent is `$lookup` in an aggregation pipeline.

```bash
./bench.sh "/v3/batch /v4/json" "50 1000"
```

**What to notice:**
- 1 query and no duplicated parent rows (compare `rows` with the JOIN).
- The correlated subquery still runs once per product, **but inside the database**, with no network round trip and using the index. "N lookups" isn't the problem. "N network round trips" is.
- Trade-off: the SQL is harder to read and test, and the DB does the JSON work. Good for hot read endpoints. Usually not worth it everywhere.

---

## Step 6: A new requirement: "only the latest 3 reviews per product" (30 min)

This is where the simple fixes stop working, which makes it good video material.

- The batch query from Step 4 would fetch **all** reviews and throw most away in Go.
- `LIMIT 3` on the batch query gives 3 reviews **total**, not 3 per product.

```bash
./bench.sh "/v5/top3-naive /v5/top3-lateral /v5/top3-window" "50 1000"
```

Two real solutions:
1. **`LATERAL` join** (`top3Lateral`): "for each product, run this small subquery". With the `(product_id, created_at DESC)` index, each one reads exactly 3 index entries.
2. **Window function** (`top3Window`): `row_number() OVER (PARTITION BY product_id ORDER BY created_at DESC)`, then keep `rn <= 3`. Works in Postgres, MySQL 8+, SQL Server, SQLite.

Try both **with and without** the index from Step 4 (`DROP INDEX reviews_product_created_idx;`). The window version has to read all reviews of those products, while LATERAL with the index reads only 3 per product. That's a nice "it depends, and here's how to measure it" moment.

---

## Step 7: Counts instead of lists (30 min)

The product list usually shows "★ 4.2 (57 reviews)", not the reviews themselves.

```bash
./bench.sh "/v6/counts-naive /v6/counts-groupby /v6/counts-column" "50 1000"
```

1. `counts-naive`: `SELECT count(*)` per product. Still N+1, just smaller rows. **N+1 is about the pattern, not the payload.**
2. `counts-groupby`: one `LEFT JOIN ... GROUP BY`. The normal fix.
3. `counts-column`: reads a precomputed `products.review_count`. Run it now: **every count is 0.** That's the lesson: a denormalized column is a cache you must maintain. Backfill it:
   ```sql
   UPDATE products p SET review_count = (SELECT count(*) FROM reviews r WHERE r.product_id = p.id);
   ```
   Then keep it correct on every new review. Two options to discuss:
   - Same transaction in the app: `INSERT INTO reviews ...; UPDATE products SET review_count = review_count + 1 WHERE id = $1;`
   - A database trigger (`AFTER INSERT OR DELETE ON reviews`)

   Trade-offs: every write gets slightly more expensive, and very popular products cause row-lock contention on that one counter row (a teaser for the "hot counters" topic). Use it when reads massively outnumber writes.

---

## Step 8: Fix #4, DataLoader, for when the loop is hidden (45 min)

Real code is rarely one loop. The N+1 hides across layers: a GraphQL resolver for `Product.reviews`, a `ReviewService.forProduct(id)` called from a template, a React component that fetches its own data. Each piece looks innocent on its own.

```bash
./bench.sh "/v7/nested-naive /v7/nested-loader" "10 50"
```

1. `nested-naive`: products → reviews → **user name per review**. That's 1 + N + M queries: 50 products × ~50 reviews ≈ **2,500+ queries**. Watch it with `LOG_SQL=1`.
2. `nested-loader`: the resolvers still say "load reviews for *this* product" and "load the user for *this* review", but they go through [loader.go](loader.go):
   - **Batching:** every `Load(id)` that arrives within 2 ms is collected, then **one** `WHERE id = ANY(...)` runs.
   - **Per-request cache:** the same user appearing in 30 reviews is fetched once.
   - Result: about **3 queries** (products, reviews, users), one per level of nesting, whatever the size.
3. Read `loader.go` (~70 lines). Make sure you can explain three things:
   - **Why the wait window:** you collect keys from concurrent callers.
   - **Why the loader is created per request:** a shared cache would leak data between users and serve stale data.
   - **Why the result slice is copied:** the cache shares it.

   This is the same idea as [graphql/dataloader](https://github.com/graphql/dataloader) and the Shopify case study.

**Where else this pattern appears:**
- **Microservices:** `for each order: GET /users/{id}` is N+1 over HTTP, and each call costs milliseconds, not microseconds. The fix is identical: a batch endpoint `GET /users?ids=1,2,3`, or a DataLoader in front of the client.
- **Frontend:** a list of 50 cards where each card fetches its own data = 50 HTTP requests. Fix: fetch the list with its data in one request, or batch on the server.

---

## Step 9: The two "fixes" that aren't (30 min)

**A. "Just run the queries in parallel."**
```bash
POOL_SIZE=5 DB_LATENCY_MS=2 go run .
./bench.sh "/v1/naive /v8/parallel-naive /v3/batch" "200"
```
Parallel is faster than naive, but it still sends 201 queries and grabs **every** connection in the pool for one request. Now fire many requests at once:
```bash
# with hey (go install github.com/rakyll/hey@latest), or k6 / ab
hey -n 200 -c 20 "http://localhost:8080/v8/parallel-naive?limit=200"
hey -n 200 -c 20 "http://localhost:8080/v3/batch?limit=200"
```
Compare the latency distributions. One request type starving the pool for everyone is exactly **Episode 2's topic**. End the video with this teaser.

**B. "Just cache it."** Put a cache in front of `reviewsFor(id)` and the warm case looks great. But the first visit, every cache expiry, and every deploy still run 1 + N queries, all at once, often for many users together. That's **Episode 3 (cache stampede)**. Caching is a good layer on top of a correct query pattern, not a replacement for one.

(You don't need to code B. Explain it with a diagram.)

---

## Step 10: Make sure it never comes back (30 min)

Fixing N+1 once is easy. Keeping it fixed in a team is the real engineering.

1. **A query-budget test.** [n1_test.go](n1_test.go) runs each strategy at `limit=10` and `limit=100` and fails if the query count is over budget. The key idea: **the query count must not grow with the page size.**
   ```bash
   go test -run TestQueryBudget -v
   ```
   Uncomment the `naive` line and watch it fail. That failing test is a strong closing shot for the video.
2. **Find it in production with `pg_stat_statements`.** After running some traffic:
   ```sql
   SELECT calls, round(mean_exec_time::numeric, 3) AS avg_ms, query
   FROM pg_stat_statements ORDER BY calls DESC LIMIT 5;
   ```
   An N+1 shows up as a cheap query with a **huge `calls` count**: `SELECT ... FROM reviews WHERE product_id = $1` called 50,000 times. Slow-query logs never catch it, because each query is fast. That's why it survives in production.
3. **Other detectors to mention:**
   - APM tools (Datadog, New Relic, Sentry) flag repeated identical spans in one trace.
   - ORM warnings, such as Rails' Bullet gem or Laravel's `preventLazyLoading()`.
   - Per-request query counts in logs, like this lab's log line.

---

## Step 11: Same problem in your own stack (optional, 30-60 min)

Write a few lines in each, so you can answer comments confidently.

**MongoDB (Node):**
```js
// N+1
for (const p of products) p.reviews = await Review.find({ productId: p._id });
// Batch
const reviews = await Review.find({ productId: { $in: products.map(p => p._id) } });
// One round trip: $lookup in an aggregation pipeline (with $limit inside for "top N")
```

**Firestore:** no joins at all. The options are:
- batched `in` queries (small value limit per query, so chunk the IDs)
- `getAll()` for many document IDs
- **denormalizing**, e.g. storing `reviewCount` and the latest 3 reviews on the product document. In Firestore this isn't a hack, it's the intended data model. A good contrast with Step 7.

**HTTP between services:**
- add a batch endpoint
- or put a DataLoader in the client

---

## Results table

Fill this in as you go. Don't publish numbers you didn't measure.

| Endpoint | Index? | Latency sim | limit | Queries | Rows | ms |
|---|---|---|---|---|---|---|
| /v1/naive | no | 0 | 50 | | | |
| /v1/naive | no | 1 ms | 50 | | | |
| /v1/naive | yes | 1 ms | 50 | | | |
| /v1/naive | yes | 1 ms | 1000 | | | |
| /v2/join | yes | 1 ms | 1000 | | | |
| /v3/batch | yes | 1 ms | 1000 | | | |
| /v4/json | yes | 1 ms | 1000 | | | |
| /v5/top3-lateral | yes | 1 ms | 1000 | | | |
| /v6/counts-groupby | yes | 1 ms | 1000 | | | |
| /v7/nested-naive | yes | 1 ms | 50 | | | |
| /v7/nested-loader | yes | 1 ms | 50 | | | |

---

## Check yourself before recording

You're ready when you can answer these without notes:

1. Why is N+1 slow when every single query takes under 1 ms?
2. When is a JOIN worse than two queries?
3. Why didn't the index alone fix the naive endpoint?
4. How do you get "top 3 per parent" in one query?
5. What breaks when you add a denormalized counter?
6. Why must a DataLoader be per request?
7. Why doesn't running the N queries in parallel fix it?
8. How would you find N+1 in a production system you've never seen before?

## What to show in the video (10-12 min)

You did 11 steps, but the video needs about 5 moments:

1. **Hook:** the SQL log scrolling 1,000 identical queries (Step 1).
2. **The "aha":** the latency chart. Same code, only the database distance changes (Step 2).
3. **The fixes:** the map table at the top of this file. Show JOIN / batch / DataLoader with their query counts (Steps 3, 4, 8).
4. **The twist:** "the index made it faster but didn't fix it", or "parallel made it worse for everyone" (Steps 4, 9).
5. **The close:** the failing query-budget test, plus "next episode: what happens when the pool runs dry".
