# Season 0: The Toolbox

*One video per tool, each one standalone · Nov 6, 2026 · @Lion*

---

## What this season is

Each video takes **one tool or building block** (Nginx, Postgres, Redis, a message queue, and so on). You learn it properly, run it yourself, break it on purpose, and explain it. Every video stands on its own: no shared story, no "previously on". A viewer can start at any episode.

Before each video you **practise with the tool hands-on**. The practice is where your confidence comes from, and it's also where your best video moments come from: the surprises, the numbers you measured yourself, the thing you broke.

---

## How deep is "deep enough"?

| Level | Name | What it means | How you prove it |
|---|---|---|---|
| **L1** | Explain | What it is, why it exists, the main trade-off. | Explain it in 2 minutes without notes. |
| **L2** | Use | You've installed and run it, used its main features, know its important settings, and have seen how it fails. | You've finished the practice exercises and can predict what happens when it breaks. |
| **L3** | Internals | How it works inside, and why it's fast or slow. You can predict how it behaves under load and answer "why?" two levels deep. | You can explain it to a senior engineer and survive their follow-up questions. |

**The ceiling:** L3 doesn't mean being able to run the tool in production for a bank. It means understanding the mechanism well enough that no reasonable viewer question surprises you.

---

## Free learning resources

Every tool below has its own **Learn from (all free)** list:
- **Core, enough on its own:** 2-3 resources. If you read these and do the practice exercises, you have enough for the video.
- **Watch:** a video, when a good one exists.
- **Go further (optional):** for when something is still unclear, or you want an extra story for the video.

Everything listed is free to read or watch. The paid books (*Designing Data-Intensive Applications*, Alex Xu's *System Design Interview*, *Release It!*) are excellent but optional. Borrow them if you can; you don't need them for this season.

**Free resources that cover many tools (bookmark these once):**
- **Martin Kleppmann's *Distributed Systems* lectures** (University of Cambridge, free on YouTube) plus his free lecture notes PDF. This is the closest free substitute for *DDIA* for T12-T19.
- **CMU 15-445 *Database Systems*** by Andy Pavlo (free on YouTube, with slides on the course site), for T4-T6 and T16-T17.
- **MIT 6.5840 *Distributed Systems*** (free lectures and labs on the course site), for T16-T19 if you want more depth.
- **AWS Builders' Library** (aws.amazon.com/builders-library): short articles by Amazon engineers on caching, timeouts, retries, health checks and load shedding.
- **Microsoft Azure Architecture Center, Cloud Design Patterns** (learn.microsoft.com): one clear page per pattern (retry, circuit breaker, bulkhead, saga, cache-aside and more), product-neutral despite the name.
- **The System Design Primer** (github.com/donnemartin/system-design-primer): a free overview of almost every tool here, good for revision.
- **Google *SRE book* and *SRE Workbook*** (sre.google/books), for T6, T21 and T22.
- **martinfowler.com** and **microservices.io**, for T21, T23 and T24.
- **YouTube channels:** Hussein Nasser (backend fundamentals), ByteByteGo (short animated overviews), Arpit Bhayani (internals). Watch them *after* your own practice, so your explanation stays yours.
- **Official docs:** for Postgres, Redis, RabbitMQ, Kafka, Nginx and OpenSearch, the official docs are genuinely the best source. Get used to reading them; it's also what you'll tell viewers to do.

**How to find any resource:** search the exact title in quotes. Titles of blog posts and talks are more stable than URLs, which change when companies redesign their sites.

## How to practise (applies to every tool)

**Before Toolbox #1, set up your lab (half a day):**
- Install Docker and learn the basics of `docker compose`: `up`, `down`, `logs`, `exec`, `pause`. Almost every exercise runs in containers, so nothing pollutes your machine and every setup can be thrown away.
- Install the load-testing tools you'll reuse: **k6** (or `hey`) for HTTP load, plus `curl`, `jq` and `htop`.
- Create a GitHub repo, `toolbox-labs`, with one folder per tool. Each folder gets a `docker-compose.yml`, the scripts you wrote, and a `NOTES.md`. Viewers can clone it, and that's a good reason for them to subscribe.
- Write **the lab app**: a tiny HTTP API (~100 lines in whichever language you're most comfortable with) that you'll reuse as the test subject for many tools:
  - `GET /items/:id` reads one row from Postgres.
  - `POST /items` writes one row.
  - `GET /slow` waits 2 seconds, then responds.
  - `POST /counter` increments a number.
  - `GET /health` returns 200.

**The rules of practice:**
1. **Try before you read the answer.** For each exercise, attempt it from the official docs first. Use a full tutorial only when you're stuck for more than 30 minutes.
2. **Keep a lab notebook** (`NOTES.md`). Write down every surprise, error message and "oh, *that's* why". Surprises are your best video moments.
3. **Measure your own numbers.** Each tool has a *Measure it* exercise. On camera, "I measured 2 ms on my laptop" beats any number quoted from a blog.
4. **Break it on purpose.** Every tool has a *Break it* exercise. If you've seen the failure with your own eyes, you can explain it with conviction.
5. **Screen-record your practice sessions.** Run OBS while you work. Half the time you'll capture a better demo moment than you could stage later, and it's free B-roll.
6. **Explain it back.** After practice, explain the tool out loud for 5 minutes without notes. Anything you stumble on means one more exercise.

**Every tool's practice section has the same parts:**
- **Setup:** get it running.
- **Exercises:** the core features, in order.
- **Break it:** cause the main failure mode yourself.
- **Measure it:** get one or two real numbers for the video.
- **Mini-project (optional):** combine everything into something small and real. Do it for the L3 tools if you have time.
- **Ready to record when:** a self-check before you script.

---

## The 6-day prep routine for each video

- **Day 1, Read:** the tool's *Core* resources (watch the *Watch* item if there is one). Write one paragraph in your own words for each *Must know* item.
- **Day 2, Exercises:** Setup, plus the first half of the exercises.
- **Day 3, Exercises:** the rest of the exercises, plus *Break it*.
- **Day 4, Measure and go deeper:** *Measure it*, plus the mini-project if there's one. For L3 tools, answer "why?" twice for each *Must know* item. Write answers to the questions viewers are likely to ask.
- **Day 5, Script:** plan the whiteboard panels and the demo steps. Write the hook word for word and keep the rest as bullet points. Do one dry run.
- **Day 6, Record:** record the demo separately from the whiteboard explanation.

L2 tools often fit in 4-5 days. L3 tools may need a 7th day. That's fine.

## Every Toolbox video follows the same shape

1. **The problem (30-60 s):** a concrete situation where you need this tool ("the server forgets you're logged in as soon as you add a second one").
2. **The mental model (2 min):** one simple picture you keep coming back to.
3. **How it works inside (main part):** taken to the tool's depth level.
4. **Live demo (2-4 min):** the best exercise from your practice, in a terminal.
5. **Break it (1-2 min):** the failure, shown live.
6. **Your numbers (30 s):** what you measured.
7. **When to use it, and when not to (1 min).**

Aim for 10-15 minutes.

---

## The tool list

Ordered so that each tool only needs the ones before it.

| # | Tool / building block | Practice with | Depth |
|---|---|---|---|
| **Module 1: The web** | | | |
| T1 | How a request travels (DNS, TCP, TLS, HTTP) | `dig`, `curl`, browser DevTools | L2 |
| T2 | Web server and reverse proxy | Nginx | L2 |
| T3 | Load balancers and stateless servers | Nginx / HAProxy, k6 | L2 |
| **Module 2: The database** | | | |
| T4 | Relational database and indexes | PostgreSQL | **L3** |
| T5 | Transactions, isolation and locking | PostgreSQL | **L3** |
| T6 | Backups and recovery | `pg_dump`, WAL archiving | L2 |
| **Module 3: Making things fast** | | | |
| T7 | Redis | Redis | **L3** |
| T8 | Caching patterns and CDN | Redis, Cloudflare (free tier) | **L3** / L2 |
| T9 | Sessions, cookies, JWT and OTP login | Lab app + Redis | L2 |
| T10 | Object storage and file uploads | MinIO (S3-compatible) | L2 |
| T11 | Search engines | OpenSearch / Elasticsearch | L2 (L3 for the inverted index) |
| **Module 4: Working in the background** | | | |
| T12 | Message queues | RabbitMQ | **L3** |
| T13 | Event logs and pub/sub | Kafka or Redpanda | L2 |
| T14 | Background jobs, cron and distributed locks | cron, Redis | L2 |
| T15 | Idempotency and online payments | Payment gateway sandbox | **L3** |
| **Module 5: Scaling data** | | | |
| T16 | Database replication | PostgreSQL streaming replica | **L3** |
| T17 | Sharding and consistent hashing | Script, Citus (optional) | **L3** |
| T18 | Distributed ID generation | Script + Postgres | L2 |
| T19 | Consistency models and CAP | Paper, plus a 3-node demo | L2 (concept-heavy) |
| **Module 6: Staying up** | | | |
| T20 | Rate limiting | Redis, Nginx | **L3** |
| T21 | Timeouts, retries, circuit breakers, bulkheads | Toxiproxy + a resilience library | **L3** |
| T22 | Observability: logs, metrics, traces, SLOs | Prometheus, Grafana | L2 |
| T23 | Containers and safe deployments | Docker, kind (local Kubernetes) | L2 |
| **Module 7: Bigger architectures** | | | |
| T24 | Microservices, sagas, outbox and CDC | 2 small services + RabbitMQ, Debezium | L2 |
| T25 | Real-time: polling, SSE, WebSockets | WebSocket server + Redis pub/sub | L2 |
| T26 | Geospatial indexing | Redis GEO, PostGIS | L2 |

**Where to go deepest:** the L3 tools (Postgres, transactions, Redis, caching, queues, idempotency, replication, sharding, rate limiting, failure handling). They're the most-asked-about topics in system design and the ones viewers search for most.

---

# The tools in detail

---

## Module 1: The web

### T1 · How a request travels (DNS, TCP, TLS, HTTP), depth L2
**Why it exists:** every tool in this season sits somewhere on the path a request takes. Viewers who understand that path understand where each tool fits.
**Must know**
- DNS resolution: recursive resolver, A/CNAME records, TTL, and why DNS changes are slow.
- The TCP handshake, and why connections are reused (keep-alive, connection pools).
- The TLS handshake at headline level: certificates, and why HTTPS adds round trips.
- HTTP: methods, status codes, the headers that matter (`Cache-Control`, `Cookie`, `Retry-After`). HTTP/1.1 vs HTTP/2 at headline level.
- Round trips and latency, and why distance matters.

**Skip:** TCP congestion-control internals, QUIC details.
**Learn from (all free)**
- *Core, enough on its own:*
  1. *High Performance Browser Networking* by Ilya Grigorik (hpbn.co): Chapters 1-4 (latency, TCP, UDP headline, TLS).
  2. MDN Web Docs, "An overview of HTTP" and "HTTP caching" (developer.mozilla.org).
  3. Cloudflare Learning Center, "What is DNS?" and "What happens in a TLS handshake?" (cloudflare.com/learning).
- *Watch:* Julia Evans's *Mess With DNS* (messwithdns.net): a free playground where you make real DNS records and watch queries arrive.
- *Go further (optional):* The free *How DNS Works* comic (howdns.works).

**Practice**
- **Setup:** nothing to install beyond `dig`, `curl` and a browser.
- **Exercises**
  1. `dig` five popular sites. Note their TTLs and whether they use CNAMEs. Run `dig +trace` once and follow the chain from the root servers.
  2. Use `curl -v` on an HTTPS site and identify the TLS handshake, the request headers and the response headers in the output.
  3. Write an HTTP request **by hand** with `nc` (netcat) to an HTTP-only test site and read the raw response. Seeing HTTP as plain text clears up a lot of confusion.
  4. In browser DevTools → Network, load a big site and read the waterfall: DNS, connect, TLS, waiting (TTFB), download.
- **Break it:** point `/etc/hosts` at a wrong IP for a domain and see what the browser does. Request a site over HTTPS with the wrong hostname (`curl --resolve`) to trigger a certificate error.
- **Measure it:** use `curl -w` with `time_namelookup`, `time_connect`, `time_appconnect` and `time_starttransfer` against servers in three regions (for example Singapore, Europe, the US). Make a small table. That's your latency slide.
- **Ready to record when:** you can draw the whole path from typing a URL to seeing the page, and say roughly how many milliseconds each step costs.

**Video idea:** *What actually happens in the 300 ms after you tap a link*

### T2 · Web server and reverse proxy (Nginx), depth L2
**Why it exists:** app servers are good at running code and bad at handling thousands of slow connections, TLS and static files. A reverse proxy takes that job off them.
**Must know**
- Web server vs app server, and what a reverse proxy does.
- TLS termination, serving static files, `proxy_pass` / `upstream`, timeouts, gzip, request size limits.
- Reading access and error logs.

**Skip:** Nginx modules, OpenResty/Lua.
**Learn from (all free)**
- *Core, enough on its own:*
  1. The official Nginx "Beginner's Guide" (nginx.org/en/docs/beginners_guide.html).
  2. The NGINX Admin Guide sections "Web Server" and "Reverse Proxy" (docs.nginx.com).
  3. DigitalOcean's tutorial "Understanding Nginx Server and Location Block Selection Algorithms".
- *Watch:* Hussein Nasser's Nginx and reverse-proxy videos on YouTube.
- *Go further (optional):* The Nginx module reference for `ngx_http_proxy_module` (every directive explained).

**Practice**
- **Setup:** Nginx and the lab app in docker compose.
- **Exercises**
  1. Serve a static HTML page and an image directly from Nginx.
  2. Proxy `/api/` to the lab app with `proxy_pass`. Add the `X-Forwarded-For` header and print it in the app.
  3. Turn on HTTPS with a self-signed certificate (or `mkcert`), so Nginx terminates TLS and talks plain HTTP to the app.
  4. Turn on gzip and compare response sizes with `curl --compressed -w '%{size_download}'`.
  5. Set `client_max_body_size` and try uploading something bigger.
- **Break it:** stop the app container and request `/api/`: read the 502 and the matching error-log line. Call `GET /slow` with `proxy_read_timeout 1s` to get a 504. Knowing the difference between 502 and 504 is a classic viewer question.
- **Measure it:** load-test a static file served by Nginx vs the same file served by the lab app, using k6. Report requests per second.
- **Ready to record when:** you can write a basic reverse-proxy config from memory and explain 502 vs 504.

**Video idea:** *Why almost every website has Nginx in front of it*

### T3 · Load balancers and stateless servers, depth L2
**Why it exists:** one server has limits and can die. A load balancer spreads traffic and routes around dead servers, but only if the servers don't keep state.
**Must know**
- L4 vs L7 load balancing.
- Algorithms: round robin, least connections, IP/consistent hashing.
- Active and passive health checks.
- Sticky sessions and their downsides; what "stateless" really means.
- Vertical vs horizontal scaling.
- The load balancer as a single point of failure (pairs, floating IPs, managed load balancers). Autoscaling at headline level.

**Skip:** BGP/anycast (L1).
**Learn from (all free)**
- *Core, enough on its own:*
  1. Sam Rose, "Load Balancing" (samwho.dev/load-balancing): an interactive visual essay, the best intro there is.
  2. NGINX Admin Guide, "HTTP Load Balancing" (docs.nginx.com).
  3. AWS Builders' Library, "Implementing health checks" (aws.amazon.com/builders-library).
- *Watch:* Hussein Nasser's videos on L4 vs L7 load balancers.
- *Go further (optional):* The HAProxy Starter Guide (docs.haproxy.org) for active health checks and algorithms.

**Practice**
- **Setup:** Nginx (or HAProxy) in front of 3 copies of the lab app (`docker compose up --scale app=3`). Make the app print its container name in each response.
- **Exercises**
  1. Run round robin and watch responses rotate across servers. Switch to `least_conn` and hit `/slow` and `/items` together to see the difference.
  2. Make one server 3× weaker (CPU limit in compose) and try weighted balancing.
  3. Add a health check on `/health` (HAProxy does active checks for free).
  4. Store a value in the app's memory (for example a counter) and call it repeatedly: you get three different answers. That's statelessness, shown in one exercise.
- **Break it:** kill one server in the middle of a k6 test. Count the failed requests, then tune the health checks or retries until the count is near zero.
- **Measure it:** requests per second with 1, 2 and 3 servers. Is it linear? Usually not, and the "why not" (the database or the load balancer becomes the limit) is a great video moment.
- **Ready to record when:** you can explain why sticky sessions are a crutch and what a stateless server needs instead.

**Video idea:** *A server died and nobody noticed: load balancers explained*

---

## Module 2: The database

### T4 · Relational database and indexes (PostgreSQL), depth L3
**Why it exists:** most apps keep their important data in a relational database, and most "the app is slow" problems are a missing or wrong index.
**Must know**
- Tables, primary and foreign keys, normalisation vs denormalisation (headline level).
- **How a B-tree index works**, composite indexes and why column order matters, covering indexes, partial indexes.
- Why `LIKE 'abc%'` can use an index and `LIKE '%abc%'` can't.
- Reading `EXPLAIN ANALYZE`: sequential scan, index scan, index-only scan.
- Connection limits and pooling (PgBouncer).
- The N+1 query problem.
- MVCC at headline level (it sets up T5).

**Skip:** planner cost-model internals, VACUUM tuning.
**Learn from (all free)**
- *Core, enough on its own:*
  1. *Use The Index, Luke* by Markus Winand (use-the-index-luke.com): read the whole thing. This alone covers indexes to L3.
  2. PostgreSQL docs, chapters "Indexes" and "Using EXPLAIN" (postgresql.org/docs).
  3. CMU 15-445 *Database Systems* (Andy Pavlo, free on YouTube): the lectures on storage, B+trees and query execution.
- *Watch:* The CMU 15-445 B+tree lectures, for the internals behind the diagrams.
- *Go further (optional):* SQL practice at pgexercises.com; paste your plans into explain.depesz.com to read them more easily; the PgBouncer docs (pgbouncer.org).

**Practice**
- **Setup:** Postgres in Docker, plus `psql`. Turn on `\timing`.
- **Exercises**
  1. Create a `products` table and fill it with 5 million rows using `generate_series` and `random()`.
  2. Query by a non-indexed column, run `EXPLAIN ANALYZE`, add an index, and run it again. Read both plans line by line.
  3. Create a composite index `(category, price)`. Query by category only, by price only, and by both. Which ones use the index, and why?
  4. Compare `LIKE 'abc%'` vs `LIKE '%abc%'`. *Bonus:* try a `pg_trgm` index for the second.
  5. Build a covering index and get an index-only scan.
  6. Write an N+1 loop in the lab app (fetch 100 items one query at a time) and compare it with one `IN (...)` query.
- **Break it:** open connections in a loop until you hit `max_connections` and read the error. Add PgBouncer in front and repeat.
- **Measure it:** query time before and after the index on 5 million rows; index size vs table size (`pg_relation_size`); how much slower inserts get with 0 vs 5 indexes.
- **Mini-project:** take 5 slow queries, fix each one with the right index, and justify every choice.
- **Ready to record when:** you can draw a B-tree, show how a lookup walks it, and predict from a query whether an index will be used.

**Video idea:** *Why your query takes 8 seconds, and how one index makes it 2 ms*

### T5 · Transactions, isolation and locking, depth L3
**Why it exists:** when two users act at the same moment, "check, then save" logic quietly corrupts data. Transactions and locks are how databases stop that.
**Must know**
- ACID in practical terms.
- The check-then-act race condition.
- Isolation levels (read committed, repeatable read, serializable) and the anomalies each allows: dirty reads, lost updates, write skew.
- MVCC: why readers don't block writers.
- Three fixes for a race: `SELECT ... FOR UPDATE`, an atomic conditional `UPDATE ... WHERE stock > 0`, and optimistic locking with a version column.
- Deadlocks: how they happen and how the database resolves them.

**Skip:** lock-manager internals, implementing two-phase commit.
**Learn from (all free)**
- *Core, enough on its own:*
  1. PostgreSQL docs, chapter "Concurrency Control" (MVCC, isolation levels, explicit locking): short and exact.
  2. Martin Kleppmann's *Hermitage* project (github.com/ept/hermitage): ready-made two-session scripts that reproduce every isolation anomaly in Postgres. It *is* your exercise list.
  3. CMU 15-445 lectures on concurrency control and two-phase locking (YouTube).
- *Watch:* The CMU 15-445 concurrency-control lectures.
- *Go further (optional):* Jepsen's "Consistency Models" map (jepsen.io/consistency) for the precise definitions.

**Practice**
- **Setup:** Postgres, and **two `psql` terminals side by side**. That layout *is* the demo.
- **Exercises**
  1. Reproduce a lost update by hand: both sessions read `stock = 1`, both write `stock = 0`, and two sales are recorded.
  2. Fix it with `SELECT ... FOR UPDATE` and watch the second session block until the first commits.
  3. Fix it with an atomic `UPDATE ... SET stock = stock - 1 WHERE stock > 0` and check the affected-row count.
  4. Fix it with optimistic locking (a `version` column) and handle the retry.
  5. Run the same race under `SERIALIZABLE` and read the serialization error.
  6. Reproduce write skew (the classic "two doctors on call" example from DDIA).
- **Break it:** create a deadlock on purpose (two sessions locking two rows in opposite order) and read Postgres's deadlock report.
- **Measure it:** write a script that fires 100 concurrent purchases at an item with 10 in stock. Count how many oversell with no fix, and how many requests per second each fix handles.
- **Mini-project:** a tiny "ticket booking" endpoint in the lab app that never oversells under k6 load.
- **Ready to record when:** you can show a race on two terminals from memory, and say which fix you'd pick for which situation.

**Video idea:** *How two people bought the last ticket at the same second*

### T6 · Backups and recovery, depth L2
**Why it exists:** hardware dies, people run `DELETE` without `WHERE`, and ransomware happens. A backup you've never restored is a hope, not a plan.
**Must know**
- Logical backups (`pg_dump`) vs physical backups.
- WAL and point-in-time recovery (PITR).
- The 3-2-1 rule, off-site copies, RPO and RTO.
- Why restores must be tested, and why a replica isn't a backup.

**Skip:** comparing backup tools in depth (pgBackRest, Barman and others): just name them.
**Learn from (all free)**
- *Core, enough on its own:*
  1. PostgreSQL docs, chapter "Backup and Restore" (SQL dump, file-system backup, continuous archiving and PITR).
  2. GitLab, "Postmortem of database outage of January 31" (about.gitlab.com/blog, 2017): the story for your video.
  3. The pgBackRest User Guide (pgbackrest.org), to see how a real backup tool organises this.
- *Go further (optional):* Google SRE book chapter "Data Integrity: What You Read Is What You Wrote" (sre.google/books).

**Practice**
- **Setup:** Postgres with the 5-million-row table from T4.
- **Exercises**
  1. Take a `pg_dump`, drop the table, and restore it.
  2. Take a physical backup with `pg_basebackup`.
  3. Turn on WAL archiving, insert rows, run a deliberate `DELETE` with no `WHERE`, then do PITR to the minute before it.
  4. Upload a backup to object storage (MinIO or any free tier) to make it off-site.
  5. Write a tiny script that restores the latest backup into a fresh container and counts rows: an automated restore test.
- **Break it:** corrupt a backup file (truncate it) and try to restore it. That's why you test restores.
- **Measure it:** backup and restore times for the table. That restore time *is* your RTO, a strong on-camera moment.
- **Ready to record when:** you can explain RPO vs RTO with your own measured numbers.

**Video idea:** *Your backups probably don't work. Here's how GitLab found out*

---

## Module 3: Making things fast

### T7 · Redis, depth L3
**Why it exists:** an in-memory data store that's fast and atomic. It's used for caches, sessions, counters, locks, queues, leaderboards and location data.
**Must know**
- Why it's fast: in memory, single-threaded execution, so each command is atomic.
- Data types: strings, hashes, lists, sets, sorted sets, geo, streams (headline level).
- `INCR`, `SET key value NX EX`, TTLs, `MULTI` and Lua scripts for multi-step atomicity.
- Persistence (RDB vs AOF), and exactly what you lose in a crash.
- `maxmemory` and eviction policies.
- Replication, Sentinel and Cluster at headline level. Hot keys and big keys.

**Skip:** Redis Cluster slot internals; the Redlock debate (L1: know it exists and why it's contested).
**Learn from (all free)**
- *Core, enough on its own:*
  1. Redis docs (redis.io/docs): "Data types", "Persistence", "Key eviction" and "Pipelining". The official docs are very good.
  2. Redis University (university.redis.io): free courses; start with the one on Redis data structures.
  3. Redis docs, "Scripting with Lua" and "Transactions".
- *Watch:* Redis's own YouTube channel, for short data-structure walkthroughs.
- *Go further (optional):* The Redis command reference (redis.io/commands), which lists every command's time complexity. Great for the "why is `KEYS *` slow?" moment.

**Practice**
- **Setup:** Redis in Docker, plus `redis-cli`. Keep `MONITOR` running in a second terminal, which shows every command live and makes a great visual.
- **Exercises**
  1. Strings and counters: `SET`, `GET`, `INCR`, `EXPIRE`, `TTL`.
  2. A lock: `SET lock:x me NX EX 10`. Try to take it from a second client.
  3. A leaderboard with sorted sets: `ZADD`, `ZINCRBY`, `ZREVRANGE ... WITHSCORES`.
  4. A simple queue with lists: `LPUSH` / `BRPOP` with two consumers.
  5. Hashes for objects: `HSET` / `HGETALL`.
  6. A Lua script that does "check and decrement" atomically.
  7. Geo: `GEOADD` some points and `GEOSEARCH` by radius.
- **Break it**
  - Set `maxmemory 10mb`, fill Redis, and watch eviction under `allkeys-lru` vs `noeviction`.
  - Kill Redis with `docker kill` with persistence off, then with AOF on, and compare what survives.
  - Run `KEYS *` on 1 million keys and watch every other client freeze. That's single-threading, made visible.
- **Measure it:** `redis-benchmark` for `GET`/`SET`; latency from `redis-cli --latency`; memory used per 1 million small keys.
- **Mini-project:** a leaderboard API in the lab app with top-10 and "my rank".
- **Ready to record when:** you can explain why single-threaded makes Redis *faster* and atomic, and what persistence mode you'd pick for a cache vs a lock store.

**Video idea:** *Redis: the Swiss Army knife behind every fast app*

### T8 · Caching patterns and CDN, depth L3 (patterns) / L2 (CDN)
**Why it exists:** the fastest query is the one you don't run. Caching is the biggest performance lever, and the source of some of the strangest bugs.
**Must know**
- Cache-aside, read-through, write-through, write-behind.
- TTLs and invalidation, and why invalidation is hard.
- Hit-ratio maths: what a 95% vs 99% hit rate means for the database.
- The cache stampede (thundering herd) and its fixes: locking/request coalescing, jittered TTLs, stale-while-revalidate.
- What not to cache, or what to cache carefully.
- HTTP cache headers (`Cache-Control`, `ETag`). CDN basics: edge, origin, cache key, purge.

**Learn from (all free)**
- *Core, enough on its own:*
  1. AWS Builders' Library, "Caching challenges and strategies".
  2. MDN Web Docs, "HTTP caching" (Cache-Control, ETag, revalidation).
  3. Cloudflare Learning Center, "What is a CDN?" and "What is caching?".
- *Go further (optional):* "Scaling Memcache at Facebook" (USENIX NSDI 2013, free PDF): the section on leases explains stampede protection at huge scale.

**Practice**
- **Setup:** the lab app + Postgres + Redis. Make `GET /items/:id` artificially slow (`pg_sleep(0.2)`).
- **Exercises**
  1. Add cache-aside to `GET /items/:id` and compare latency with and without the cache.
  2. Update an item and see the stale read. Fix it with delete-on-write invalidation.
  3. Add a hit/miss counter and calculate the hit ratio under a k6 test with realistic, skewed key popularity.
  4. Set `Cache-Control` headers on a static file and watch the browser's behaviour in DevTools (200, 304, from cache).
  5. Put a static page behind Cloudflare's free tier and check `cf-cache-status` (MISS → HIT). Purge it.
- **Break it:** make one hot key expire while 500 requests per second hit it, and watch the database spike in `pg_stat_activity`. That's a stampede. Fix it with a lock or single-flight, then add jittered TTLs.
- **Measure it:** database queries per second at a 0%, 90% and 99% hit ratio.
- **Ready to record when:** you can show a stampede happening and explain two fixes.

**Video idea:** *The cache that took down the database*

### T9 · Sessions, cookies, JWT and OTP login, depth L2
**Why it exists:** HTTP forgets everything between requests. Sessions and tokens are how an app remembers who you are, and they're a security hotspot.
**Must know**
- Cookies and their flags (`HttpOnly`, `Secure`, `SameSite`, `Expires`).
- Server-side sessions vs JWT, and the real trade-off (revocation, size, where state lives).
- Where sessions live once there are several servers.
- The OTP flow: generate, store a hash, expire, limit attempts.
- Password hashing (bcrypt/argon2) at headline level.
- OAuth/OIDC at L1 ("Login with Google" exists, and roughly how it works).

**Learn from (all free)**
- *Core, enough on its own:*
  1. *The Copenhagen Book* (thecopenhagenbook.com): a free, practical guide to sessions, password hashing, OTP and email verification. This alone is nearly enough.
  2. OWASP Cheat Sheet Series (cheatsheetseries.owasp.org): "Session Management", "Authentication" and "Password Storage".
  3. MDN Web Docs, "Using HTTP cookies".
- *Go further (optional):* The jwt.io introduction to JSON Web Tokens; OWASP "Multifactor Authentication" cheat sheet for OTP rules.

**Practice**
- **Setup:** the lab app behind Nginx with 2 instances (reuse T3).
- **Exercises**
  1. Add login with an in-memory session and a cookie. Watch users get logged out at random as round robin switches servers.
  2. Move the sessions to Redis and the problem is gone. Implement logout (deleting the session).
  3. Swap to a JWT. Decode it at jwt.io to see that it's readable, not secret. Now try to "log out" a stolen token.
  4. Build an OTP flow: a 6-digit code stored hashed in Redis with a 5-minute TTL and at most 3 attempts.
  5. Set the cookie flags and check them in DevTools → Application.
- **Break it:** steal a cookie (copy it into another browser) and see that it works. Then try `HttpOnly` against `document.cookie`. Brute-force your OTP with no attempt limit to see how quickly 6 digits fall, then add the limit.
- **Measure it:** bcrypt hashing time at cost 10 vs 12 (why it's slow on purpose); JWT size vs session-ID cookie size.
- **Ready to record when:** you can argue both sides of "sessions vs JWT" and pick one with reasons.

**Video idea:** *How a website remembers that you're logged in*

### T10 · Object storage and file uploads, depth L2
**Why it exists:** a server's disk isn't shared, doesn't scale, and dies with the server. Object storage (S3 and similar) is where files belong.
**Must know**
- File system vs object storage: buckets, keys, objects you replace rather than edit, the durability claims.
- Presigned URLs, so clients upload straight to storage.
- Multipart and resumable uploads (the tus protocol).
- An upload processing pipeline (upload → event/queue → worker → thumbnails).
- A CDN in front of storage, storage classes and cost maths.

**Skip:** erasure-coding internals (L1).
**Learn from (all free)**
- *Core, enough on its own:*
  1. AWS S3 User Guide: "Uploading objects with presigned URLs" and "Uploading and copying objects using multipart upload" (docs.aws.amazon.com). They apply to MinIO too.
  2. The MinIO docs (min.io/docs): quickstart, `mc` client, bucket notifications.
  3. The tus resumable upload protocol (tus.io/protocols/resumable-upload): short and readable.
- *Go further (optional):* Cloudflare Learning Center, "What is object storage?".

**Practice**
- **Setup:** MinIO in Docker (S3-compatible), plus the AWS CLI or `mc`.
- **Exercises**
  1. Create a bucket and upload, list, download and delete objects.
  2. Generate a presigned PUT URL in the lab app and upload with `curl` or a small HTML page, so the file never touches your server.
  3. Upload a 1 GB file with multipart and look at the parts.
  4. Run a tus server (tusd) and resume an interrupted upload.
  5. Set up MinIO bucket notifications and write a worker that makes a thumbnail with ImageMagick.
- **Break it:** upload a 1 GB file *through* the lab app instead and watch its memory and the timeouts. Cut the network halfway through a normal upload (`docker network disconnect`) and try to resume.
- **Measure it:** upload time through the app vs a presigned direct upload; resume time vs starting from scratch.
- **Ready to record when:** you can explain why presigned uploads take load off your servers, and how resumable uploads work in chunks.

**Video idea:** *Why you should never upload files to your own server*

### T11 · Search engines, depth L2 (L3 for the inverted index)
**Why it exists:** databases are bad at "find text that roughly matches this, ranked by relevance". Search engines are built for exactly that.
**Must know**
- How an inverted index works.
- Analyzers: tokenising, lowercasing, stemming, synonyms, typo tolerance.
- Relevance at headline level (TF-IDF/BM25), filters and facets.
- Autocomplete: edge n-grams and tries.
- Keeping the index in sync with the database (dual writes vs CDC), and accepting that search is slightly behind.

**Skip:** cluster tuning and shard sizing beyond headline level.
**Learn from (all free)**
- *Core, enough on its own:*
  1. *Introduction to Information Retrieval* by Manning, Raghavan and Schütze (free at nlp.stanford.edu/IR-book): Chapters 1-2 (inverted index, tokenisation) and 6 (scoring). The theory behind every search engine.
  2. The OpenSearch docs (opensearch.org/docs): "Getting started", "Text analysis", "Query DSL".
  3. PostgreSQL docs, chapter "Full Text Search".
- *Go further (optional):* *Elasticsearch: The Definitive Guide* (free on elastic.co, written for an old version but the concepts are still the best explained); the Elastic docs on the `edge_ngram` tokenizer.

**Practice**
- **Setup:** OpenSearch (or Elasticsearch) plus Dashboards in Docker. Download a free product or movie dataset (Kaggle has plenty) of about 100,000 rows.
- **Exercises**
  1. **Build a tiny inverted index yourself** in ~40 lines of code over 1,000 documents, before touching the engine. This is what makes the L3 explanation easy.
  2. Index the dataset into OpenSearch and run match, phrase and fuzzy queries.
  3. Use the `_analyze` API to see exactly how a sentence is tokenised, then add a synonym filter.
  4. Add filters and facets (category, price range).
  5. Build autocomplete with an edge n-gram analyzer.
  6. Compare with Postgres `tsvector` + GIN full-text search on the same data.
- **Break it:** update a row in Postgres and search immediately: it isn't there yet (refresh interval, sync lag). Search with a typo and no fuzziness.
- **Measure it:** the same query with `LIKE '%x%'`, Postgres full-text search and OpenSearch, on 100,000 rows and on 1 million.
- **Ready to record when:** you can draw an inverted index and explain why it's fast, and when Postgres full-text search is enough.

**Video idea:** *How search finds what you meant among millions of items in 20 ms*

---

## Module 4: Working in the background

### T12 · Message queues, depth L3
**Why it exists:** some work doesn't need to happen while the user waits. Queues let you accept work fast, smooth out peaks, and keep going when a worker dies.
**Must know**
- Producers and consumers, and why async smooths peaks and decouples services.
- Acks/nacks, prefetch, visibility timeouts and redelivery.
- At-most-once vs at-least-once delivery, and why "exactly once" means at-least-once plus idempotency.
- Retries with backoff, dead-letter queues, ordering guarantees, priority queues, backpressure.
- Queue depth and consumer lag as the metrics to watch.

**Skip:** AMQP exchange types beyond direct, fanout and topic.
**Learn from (all free)**
- *Core, enough on its own:*
  1. RabbitMQ tutorials 1-6 (rabbitmq.com/tutorials), in your language.
  2. RabbitMQ docs: "Consumer Acknowledgements and Publisher Confirms", "Dead Letter Exchanges", "Consumer Prefetch".
  3. AWS SQS Developer Guide, "Visibility timeout" and "Dead-letter queues": the same ideas in a different product.
- *Go further (optional):* The CloudAMQP blog series "RabbitMQ for beginners"; Kleppmann's Cambridge lecture notes on messaging (see the cross-cutting list).

**Practice**
- **Setup:** RabbitMQ with the management UI in Docker. The UI's queue graphs are perfect video visuals.
- **Exercises**
  1. One producer, one consumer, then 3 competing consumers. Watch the work get shared.
  2. Manual acks vs auto-ack: kill a consumer mid-message under each and see which loses the message.
  3. Set prefetch to 1 vs 100 and see how fairly work spreads.
  4. Use a fanout exchange so one message reaches 3 queues (pub/sub).
  5. Retries with a delay, and a dead-letter queue after 3 failures.
  6. A priority queue: urgent messages jump the line.
- **Break it:** send a "poison" message that crashes the consumer and watch the infinite redelivery loop. Fix it with the DLQ. Stop all consumers, publish 1 million messages, and watch memory and the disk alarm.
- **Measure it:** throughput with 1, 3 and 10 consumers; how long a 100,000-message backlog takes to drain.
- **Mini-project:** move the T10 thumbnail worker onto RabbitMQ with retries and a DLQ.
- **Ready to record when:** you can explain at-least-once delivery and show a redelivery live.

**Video idea:** *Why big apps make your request wait in line*

### T13 · Event logs and pub/sub (Kafka), depth L2
**Why it exists:** a queue forgets a message once it's handled. A log keeps it, so many consumers can read the same events at their own pace and replay history.
**Must know**
- A log vs a queue: retention and replay.
- Topics, partitions, keys and ordering within a partition.
- Consumer groups, offsets and rebalancing.
- Pub/sub options: Redis pub/sub (fire-and-forget) vs Kafka (durable).
- When *not* to use Kafka.

**Skip:** broker internals, ISR tuning, Kafka Streams.
**Learn from (all free)**
- *Core, enough on its own:*
  1. Apache Kafka docs, "Introduction" and "Design" (kafka.apache.org/documentation).
  2. Confluent Developer, the free "Apache Kafka 101" course (developer.confluent.io).
  3. Jay Kreps, "The Log: What every software engineer should know about real-time data's unifying abstraction" (LinkedIn Engineering, 2013).
- *Watch:* The short videos in Confluent's Kafka 101 course.
- *Go further (optional):* The Redpanda docs (docs.redpanda.com) if you use Redpanda for the lab.

**Practice**
- **Setup:** Redpanda (simpler) or Kafka in KRaft mode in Docker, plus a UI (Redpanda Console or Kafka UI).
- **Exercises**
  1. Create a topic with 3 partitions and produce keyed messages. See which partition each key lands in.
  2. Run one consumer group with 3 consumers, then add a 4th (it sits idle). Explain why.
  3. Run a second consumer group on the same topic: both get every message.
  4. Reset a group's offset to the beginning and replay everything.
  5. Compare Redis pub/sub: a subscriber that's offline misses messages forever.
- **Break it:** kill a consumer and watch the group rebalance. Produce without keys and show that the per-user order is lost.
- **Measure it:** messages per second for one producer on your laptop; consumer lag while a consumer is paused.
- **Ready to record when:** you can say exactly when you'd pick Kafka over RabbitMQ, and the other way round.

**Video idea:** *Kafka vs RabbitMQ: do you need a queue or a log?*

### T14 · Background jobs, cron and distributed locks, depth L2
**Why it exists:** apps need scheduled work (reports, cleanups, reminders). Once you have more than one server, "run this once" gets surprisingly hard.
**Must know**
- Cron syntax and its limits; worker pools.
- The "job ran twice" problem when several servers run the same cron.
- Leases with TTLs, leader election at headline level, fencing tokens.
- Delayed and scheduled jobs with Redis sorted sets.
- Making jobs idempotent.

**Learn from (all free)**
- *Core, enough on its own:*
  1. Martin Kleppmann, "How to do distributed locking" (martin.kleppmann.com, 2016): fencing tokens, clearly explained.
  2. Redis docs, "Distributed Locks with Redis".
  3. crontab.guru, for cron syntax.
- *Go further (optional):* PostgreSQL docs on `SELECT ... FOR UPDATE SKIP LOCKED` (the "Locking Clause" section of the `SELECT` page), which is how many job queues run on Postgres.

**Practice**
- **Setup:** 3 lab-app containers + Redis + Postgres.
- **Exercises**
  1. Put a cron job in all 3 containers that inserts a "report" row every minute. Count the rows: 3 per minute.
  2. Add a Redis lease (`SET NX EX`) so only one runs.
  3. Build a delayed-job scheduler with a sorted set (score = run time) and a polling worker.
  4. Make the job idempotent with a unique `(job_name, run_minute)` constraint.
- **Break it:** `docker pause` the leaseholder past its TTL. Another container takes the lease, then you unpause the first: now both run. Fix it with a fencing token checked in the database.
- **Measure it:** how late the delayed jobs run (actual vs scheduled time) at different polling intervals.
- **Ready to record when:** you can show the pause-and-resume double run and explain fencing tokens.

**Video idea:** *The cron job that ran three times*

### T15 · Idempotency and online payments, depth L3
**Why it exists:** networks fail in the middle of requests. When a client retries "charge ₹500", the server must be sure it only charges once.
**Must know**
- How a payment gateway flow works: create the payment → redirect → customer pays → callback/webhook → server-side verification.
- Why a timeout doesn't tell you whether the payment happened.
- Idempotency keys: store, lock, replay the saved response, expiry.
- The payment state machine (created → pending → succeeded / failed / refunded).
- Webhooks arrive at least once (and sometimes out of order).
- Reconciliation and refunds.

**Skip:** card-network internals, PCI DSS details (L1).
**Learn from (all free)**
- *Core, enough on its own:*
  1. Stripe blog, "Designing robust and predictable APIs with idempotency" (stripe.com/blog/idempotency, 2017).
  2. Brandur Leach, "Implementing Stripe-like Idempotency Keys in Postgres" (brandur.org/idempotency-keys): a complete, practical implementation. Your exercise 3 in detail.
  3. Stripe docs, "Idempotent requests" and "Webhooks" (signatures, retries, ordering).
- *Go further (optional):* The IETF draft "The Idempotency-Key HTTP Header Field"; the bKash developer portal and SSLCommerz developer docs for local gateway flows.

**Practice**
- **Setup:** a sandbox or test-mode account (Stripe test mode is the easiest to start with; also try bKash or SSLCommerz sandbox), the lab app, and a tunnel (ngrok or cloudflared) so webhooks can reach your laptop.
- **Exercises**
  1. Create a test payment and complete it.
  2. Receive the webhook and verify its signature.
  3. Build a `POST /charge` that takes an `Idempotency-Key` header and stores key → status → response in Postgres with a unique constraint.
  4. Send the same request twice at the same time: exactly one charge, and both callers get the same response.
  5. Model the payment state machine as a table and refuse illegal transitions (for example refunding something that failed).
- **Break it:** add a random 50% timeout between "charge" and "save" in your code, then retry without idempotency: double charges in the sandbox dashboard. Replay a webhook 3 times. Send webhooks out of order.
- **Measure it:** how many duplicates 100 retried requests produce with and without the key.
- **Mini-project:** a nightly reconciliation script that compares your database with the gateway's transaction list and flags mismatches.
- **Ready to record when:** you can explain the "did it go through?" problem and walk through your idempotency table row by row.

**Video idea:** *Why you got charged twice, and how payment systems prevent it*

---

## Module 5: Scaling data

### T16 · Database replication, depth L3
**Why it exists:** copies of the database give you more read capacity and survive a machine dying, but the copies are always slightly behind.
**Must know**
- Primary/replica setups, synchronous vs asynchronous replication.
- Replication lag and the anomalies it causes: read-your-own-writes, monotonic reads.
- Failover, split brain, and why automatic failover is hard.
- Read scaling, and why a replica isn't a backup.

**Learn from (all free)**
- *Core, enough on its own:*
  1. PostgreSQL docs, chapter "High Availability, Load Balancing, and Replication".
  2. Martin Kleppmann's Cambridge *Distributed Systems* lectures (YouTube) and free lecture notes (PDF on his website): the parts on replication.
  3. GitHub, "October 21 post-incident analysis" (github.blog, 2018): a real failover and split-brain-style story.
- *Watch:* Kleppmann's replication lectures.
- *Go further (optional):* The Patroni docs (patroni.readthedocs.io), to see how automatic failover is done in practice.

**Practice**
- **Setup:** a Postgres primary + 2 streaming replicas in docker compose. Use the Bitnami images if you want this easy, or set it up by hand for depth (recommended).
- **Exercises**
  1. Write on the primary and read it on the replicas. Check `pg_stat_replication`.
  2. Try to write to a replica and read the error.
  3. Set `recovery_min_apply_delay = '5s'` on one replica, write, then read immediately: the row is missing. That's replication lag you can see.
  4. Implement "read your own writes" in the lab app: route a user to the primary for 5 seconds after they write.
  5. Switch one replica to synchronous and compare write latency.
- **Break it:** stop the primary and promote a replica by hand. Then restart the old primary without reconfiguring it: two primaries, which is split brain.
- **Measure it:** lag under a heavy write load; write latency with async vs sync replication.
- **Ready to record when:** you can explain sync vs async in terms of what each one risks losing.

**Video idea:** *"I just saved it, why isn't it there?" Replication lag explained*

### T17 · Sharding and consistent hashing, depth L3
**Why it exists:** at some size, one machine can't hold all the data or handle all the writes. You split the data across machines, and everything gets harder.
**Must know**
- Vertical vs horizontal partitioning; partitioning within one database vs across machines.
- Choosing the shard key, range vs hash sharding, hot spots.
- Cross-shard queries and joins, and why resharding is painful.
- Consistent hashing with virtual nodes; directory-based sharding.

**Learn from (all free)**
- *Core, enough on its own:*
  1. The Amazon Dynamo paper (2007, free at allthingsdistributed.com): Section 4 on partitioning and consistent hashing with virtual nodes.
  2. Notion, "Herding elephants: lessons learned from sharding Postgres at Notion" (notion.com/blog, 2021).
  3. PostgreSQL docs, chapter "Table Partitioning".
- *Watch:* The MIT 6.5840 *Distributed Systems* lectures on sharded key/value stores (free on YouTube and the course site).
- *Go further (optional):* Tom White, "Consistent Hashing" (2007 blog post), short with diagrams; Instagram Engineering, "Sharding & IDs at Instagram"; the Citus docs.

**Practice**
- **Setup:** Python (or your language), plus 3 Postgres containers.
- **Exercises**
  1. Use Postgres declarative partitioning (by range and by hash) on one database and look at the partition-pruning plans.
  2. Write a `hash(key) % N` router across 3 Postgres containers and insert 100,000 users.
  3. Write a consistent-hash ring in ~30 lines, with virtual nodes.
  4. Add a 4th node. Count the moved keys for `% N` (about 75%) vs the ring (about 25%), and plot it.
  5. Run a "find all orders over ₹1,000" query: it has to hit every shard (scatter-gather).
- **Break it:** shard by a skewed key (for example country, where one value has 80% of the rows) and watch one shard overload.
- **Measure it:** the share of keys moved when adding a node, for 10, 100 and 1,000 virtual nodes; how evenly keys spread.
- **Optional:** the Citus extension for real distributed Postgres.
- **Ready to record when:** you can explain consistent hashing on a whiteboard and argue for a shard key.

**Video idea:** *When one database isn't enough: sharding explained*

### T18 · Distributed ID generation, depth L2
**Why it exists:** auto-increment needs one central counter. Across many machines you need IDs that are unique without coordination, and ideally sortable.
**Must know**
- Why auto-increment breaks across shards.
- UUID v4 vs v7, and index locality.
- The Snowflake layout (timestamp / machine / sequence) and clock skew.
- Ticket servers and ID ranges.
- Base62 encoding for short codes, and why guessable IDs are a security risk.

**Learn from (all free)**
- *Core, enough on its own:*
  1. Twitter Engineering, "Announcing Snowflake" (2010), and the original `twitter-archive/snowflake` repo README on GitHub.
  2. Instagram Engineering, "Sharding & IDs at Instagram" (2012): IDs generated inside Postgres.
  3. RFC 9562, *Universally Unique IDentifiers (UUIDs)* (2024): the sections on UUID v4 and v7.

**Practice**
- **Setup:** Postgres and a script.
- **Exercises**
  1. Write a Snowflake generator and decode an ID back into its time, machine and sequence.
  2. Generate 1 million IDs across 4 simulated machines and check there are no duplicates.
  3. Write a Base62 encoder and decoder for short links.
  4. Generate UUID v4 and v7 and compare their sort order.
- **Break it:** move a machine's clock backwards and see duplicates or out-of-order IDs. Add a guard.
- **Measure it:** insert 1 million UUIDv4 primary keys vs 1 million sequential ones (or UUIDv7). Compare insert time and index size. The difference surprises most people.
- **Ready to record when:** you can design an ID scheme on the whiteboard and defend each bit.

**Video idea:** *How big apps generate billions of unique IDs without one central counter*

### T19 · Consistency models and CAP, depth L2 (concept-heavy)
**Why it exists:** once data lives on several machines, "what does a read return?" stops having one answer. Consistency models name the possible answers.
**Must know**
- Strong vs eventual consistency; read-your-writes, monotonic reads, linearizability at headline level.
- **CAP stated correctly:** the choice only applies during a network partition. PACELC.
- Quorums (N, R, W).
- Choosing consistency per feature: a bank balance needs strong consistency; a view counter can be eventual.

**Learn from (all free)**
- *Core, enough on its own:*
  1. Martin Kleppmann, "Please stop calling databases CP or AP" (martin.kleppmann.com, 2015).
  2. Martin Kleppmann's Cambridge lectures and notes: the parts on consistency, quorums and linearizability.
  3. Jepsen, "Consistency Models" (jepsen.io/consistency): one page per model with clear definitions.
- *Watch:* Kleppmann's Cambridge consistency lectures.
- *Go further (optional):* Eric Brewer, "CAP Twelve Years Later: How the 'Rules' Have Changed" (InfoQ, 2012); Werner Vogels, "Eventually Consistent – Revisited" (allthingsdistributed.com, 2008); the Cassandra docs on consistency levels for your lab.

**Practice**
- **Setup:** paper, plus Docker for a small multi-node database (a 3-node Cassandra or ScyllaDB cluster).
- **Exercises**
  1. Draw timeline diagrams (two clients, two replicas) for a stale read, a non-monotonic read and a lost update, until you can draw each from memory.
  2. In the 3-node cluster, write and read at consistency `ONE` vs `QUORUM` vs `ALL`.
  3. Work out quorum overlap (R + W > N) for N = 3 and N = 5 on paper.
- **Break it:** cut one node off the network (`docker network disconnect`). At `ALL`, writes fail (choosing consistency); at `ONE`, they succeed but reads may be stale (choosing availability). That's CAP, live.
- **Measure it:** read and write latency at `ONE` vs `QUORUM` (PACELC's "else latency").
- **Ready to record when:** you can state CAP precisely and correct the most common misunderstanding of it.

**Video idea:** *The CAP theorem is misunderstood. Here's what it really says*

---

## Module 6: Staying up

### T20 · Rate limiting, depth L3
**Why it exists:** without limits, one bot, one buggy client or one attacker can use up resources (or money) meant for everyone.
**Must know**
- Fixed window, sliding-window log, sliding-window counter, token bucket, leaky bucket, with the trade-offs of each.
- Where to limit (edge/CDN, gateway, app) and what to key on (IP, user, API key, phone number).
- Distributed limiting with Redis, and the race conditions involved.
- `429` and `Retry-After`; returning limit headers.
- Abuse defences beyond limits: CAPTCHA, making abuse expensive.

**Learn from (all free)**
- *Core, enough on its own:*
  1. Stripe blog, "Scaling your API with rate limiters" (2017), with its linked code samples.
  2. Cloudflare blog, "How we built rate limiting capable of scaling to millions of domains" (2017): the sliding-window counter.
  3. Nginx docs, `ngx_http_limit_req_module` (burst and nodelay explained).
- *Go further (optional):* The Redis docs tutorial on rate limiting; IETF draft "RateLimit header fields for HTTP".

**Practice**
- **Setup:** the lab app + Redis + Nginx, plus k6.
- **Exercises**
  1. A fixed-window limiter with `INCR` + `EXPIRE`.
  2. Show the window-edge burst: 2× the limit across a window boundary.
  3. A sliding-window log with sorted sets.
  4. A token bucket in a Lua script (atomic).
  5. Nginx `limit_req` with `burst` and `nodelay`.
  6. Return `429` with `Retry-After` and `X-RateLimit-*` headers.
- **Break it:** implement the fixed window as GET-then-SET instead of `INCR` and hammer it: the limit leaks because of a race condition (a callback to T5).
- **Measure it:** plot allowed vs rejected requests over time for each algorithm under the same bursty k6 load. Put the charts side by side.
- **Ready to record when:** you can draw each algorithm's behaviour over time and say which one you'd use for an API vs for login attempts.

**Video idea:** *Four ways to rate limit an API (and which one to pick)*

### T21 · Timeouts, retries, circuit breakers and bulkheads, depth L3
**Why it exists:** in a system of services, one slow dependency can take down everything that calls it. These patterns contain the damage.
**Must know**
- Why a missing timeout cascades: thread and connection pools fill up.
- Choosing timeouts from latency percentiles.
- Exponential backoff with jitter, retry budgets, retry storms.
- Circuit breaker states (closed / open / half-open), bulkheads, fallbacks, load shedding.

**Learn from (all free)**
- *Core, enough on its own:*
  1. AWS Builders' Library, "Timeouts, retries, and backoff with jitter" and "Using load shedding to avoid overload".
  2. Microsoft Azure Architecture Center: the "Retry", "Circuit Breaker" and "Bulkhead" pattern pages (learn.microsoft.com). Free, short, and very clear.
  3. Martin Fowler, "CircuitBreaker" (martinfowler.com).
- *Go further (optional):* Marc Brooker, "Exponential Backoff And Jitter" (AWS Architecture Blog, 2015), with the simulation charts; the Toxiproxy README on GitHub.

**Practice**
- **Setup:** two copies of the lab app (A calls B), with **Toxiproxy** between them to inject latency and errors, plus k6.
- **Exercises**
  1. Add 5 seconds of latency to B through Toxiproxy with no timeout in A. Watch A's *other* endpoints die as its pool fills up.
  2. Add a timeout in A. Things improve, but A still waits up to the full timeout each time.
  3. Add retries without jitter, then with jitter, and compare the load on B.
  4. Add a circuit breaker (use a library in your language, then write a ~40-line one yourself) and watch the open → half-open → closed transitions in the logs.
  5. Add a bulkhead: a separate small pool for calls to B, so other endpoints survive.
- **Break it:** retries at 3 layers (3 × 3 × 3 = 27 times the load) against a struggling B: a retry storm.
- **Measure it:** A's success rate and p99 latency in each of the 5 setups. That's one powerful chart for the video.
- **Ready to record when:** you can explain how one slow dependency cascades, and how each pattern stops it.

**Video idea:** *How one slow API can take down an entire app*

### T22 · Observability: logs, metrics, traces, SLOs, depth L2
**Why it exists:** you can't fix what you can't see. Observability is how you find out *what* is broken and *where*, quickly.
**Must know**
- Logs vs metrics vs traces, and when to use each.
- The RED method (rate, errors, duration) and the USE method (utilisation, saturation, errors).
- Percentiles (p50, p95, p99) vs averages.
- SLIs, SLOs and error budgets; alerting on symptoms, not causes.
- Distributed tracing at headline level (OpenTelemetry).

**Learn from (all free)**
- *Core, enough on its own:*
  1. Google *SRE book* (free at sre.google/books): Chapters "Service Level Objectives" and "Monitoring Distributed Systems"; *SRE Workbook*: "Implementing SLOs" and "Alerting on SLOs".
  2. Prometheus docs, "Getting started" and "Histograms and summaries" (prometheus.io/docs).
  3. Tom Wilkie, "The RED Method: How to instrument your services" (grafana.com/blog, 2018) and Brendan Gregg, "The USE Method" (brendangregg.com).
- *Watch:* Gil Tene, "How NOT to Measure Latency" (conference talk on YouTube). It's the source of the "averages lie" lesson.
- *Go further (optional):* The OpenTelemetry docs "Getting started" for your language; the Grafana docs on dashboards.

**Practice**
- **Setup:** Prometheus + Grafana in Docker; add a metrics endpoint to the lab app (every language has a Prometheus client). *Optional:* Jaeger for traces.
- **Exercises**
  1. Expose a request counter and a latency histogram, labelled by endpoint and status.
  2. Build a RED dashboard in Grafana.
  3. Write PromQL for error rate and p99 latency (`histogram_quantile`).
  4. Define an SLO (for example 99.5% of requests under 300 ms) and chart the error budget burn.
  5. Add one alert rule and trigger it.
  6. *Optional:* add OpenTelemetry tracing across the two apps from T21 and find the slow hop in Jaeger.
- **Break it:** make 2% of requests take 5 seconds. The average barely moves; p99 jumps. That's the core lesson, live.
- **Measure it:** average vs p50 vs p99 under that load.
- **Ready to record when:** you can explain why averages lie, and what an error budget is used for.

**Video idea:** *Average latency lies: p99 and SLOs explained*

### T23 · Containers and safe deployments, depth L2
**Why it exists:** most outages happen right after a change. Safe deployment practices let teams ship often without breaking things.
**Must know**
- Docker images, layers and containers; what Kubernetes does (L1-L2).
- Rolling, blue-green and canary deploys; readiness vs liveness probes.
- Feature flags.
- **Expand/contract database migrations**, such as renaming a column without downtime.
- Rollbacks, and the stages of a CI/CD pipeline.

**Learn from (all free)**
- *Core, enough on its own:*
  1. Docker docs, "Get started" and "Multi-stage builds" (docs.docker.com).
  2. Kubernetes docs, the "Kubernetes Basics" tutorial and "Configure Liveness, Readiness and Startup Probes" (kubernetes.io).
  3. martinfowler.com: "BlueGreenDeployment", "CanaryRelease", "ParallelChange" (expand/contract) and Pete Hodgson's "Feature Toggles".
- *Go further (optional):* Stripe blog, "Online migrations at scale" (2017): a real expand/contract migration on huge tables; the kind quick start (kind.sigs.k8s.io).

**Practice**
- **Setup:** **kind** (Kubernetes in Docker) or minikube, the lab app image, and a k6 test running the whole time.
- **Exercises**
  1. Write a Dockerfile for the lab app and get the image small (a multi-stage build). Compare sizes.
  2. Deploy 3 replicas to kind and do a rolling update to v2 while k6 runs.
  3. Add readiness probes and compare the error counts with and without them.
  4. Blue-green: run v1 and v2 side by side and switch the Service selector.
  5. Canary: send ~10% of traffic to v2 (by replica ratio, or with an ingress).
  6. Add a simple feature flag (an environment variable or a database row) and turn a feature on without deploying.
  7. Roll back with `kubectl rollout undo`.
- **Break it:** rename a column the naive way (deploy code and migration together) while k6 runs and count the errors. Then do it with expand → migrate → contract and get zero errors.
- **Measure it:** errors during a rolling deploy with and without readiness probes; rollback time.
- **Ready to record when:** you can do a zero-downtime column rename from memory and explain canary vs blue-green.

**Video idea:** *How teams deploy 40 times a day without breaking anything*

---

## Module 7: Bigger architectures

### T24 · Microservices, sagas, outbox and CDC, depth L2
**Why it exists:** big teams split big apps into services. You lose the database transaction that used to keep everything consistent, and these patterns are how you get by without it.
**Must know**
- When to split a monolith (and when not to), and how to draw service boundaries.
- Synchronous calls (REST/gRPC) vs asynchronous events.
- Why distributed transactions (2PC) are usually avoided.
- Sagas (orchestration vs choreography) and compensating actions.
- The dual-write problem, the transactional outbox, and CDC (Debezium).

**Learn from (all free)**
- *Core, enough on its own:*
  1. microservices.io by Chris Richardson: the "Saga", "Transactional outbox" and "Database per service" pattern pages.
  2. martinfowler.com: "Microservices", "MonolithFirst" and "Microservice Prerequisites".
  3. The Debezium tutorial and the "Outbox Event Router" docs (debezium.io/documentation).
- *Watch:* Caitie McCaffrey, "Applying the Saga Pattern" (GOTO conference talk on YouTube).
- *Go further (optional):* Microsoft Azure Architecture Center, "Saga distributed transactions pattern".

**Practice**
- **Setup:** two tiny services (for example "orders" and "inventory"), each with its own Postgres, plus RabbitMQ. Later add Kafka + Debezium.
- **Exercises**
  1. Make "place order" call inventory synchronously. Stop inventory and see orders fail.
  2. Switch to events through RabbitMQ.
  3. Implement a 3-step saga (reserve → pay → confirm) with a compensation for each step.
  4. Implement the transactional outbox: write the order and an outbox row in one transaction, and have a relay publish it.
  5. Replace the relay with Debezium CDC reading the outbox table.
- **Break it:** in the dual-write version (save to the database, then publish), kill the service between the two steps: an order with no event, so a stuck saga. Show that the outbox survives the same kill.
- **Measure it:** delay from order to event for the polling relay vs Debezium.
- **Ready to record when:** you can explain the dual-write problem in one sentence and draw the outbox fix.

**Video idea:** *Saved to the database but the event never sent: the dual-write problem*

### T25 · Real-time: polling, SSE, WebSockets, depth L2
**Why it exists:** chat, live tracking, notifications and live counters need the server to push data to the client, which plain HTTP request/response doesn't do.
**Must know**
- Short polling, long polling, SSE and WebSockets, with the trade-offs of each.
- What each open connection costs (memory, file descriptors), and load balancing long-lived connections.
- Scaling across servers with a pub/sub backplane.
- Heartbeats, reconnects with backoff, delivery acks and receipts, flaky mobile networks.

**Learn from (all free)**
- *Core, enough on its own:*
  1. javascript.info: the chapters "Long polling", "WebSocket" and "Server Sent Events". Short, with working code.
  2. MDN Web Docs, "The WebSocket API" and "Using server-sent events".
  3. Discord blog, "How Discord Scaled Elixir to 5,000,000 Concurrent Users" (2017): a real story for the video.
- *Watch:* Hussein Nasser's videos on WebSockets, SSE and long polling.
- *Go further (optional):* The Nginx docs page "WebSocket proxying".

**Practice**
- **Setup:** the lab app with a WebSocket endpoint, 2 instances behind Nginx (with WebSocket proxying configured), plus Redis.
- **Exercises**
  1. Build the same "live clock" four ways: short polling, long polling, SSE, WebSocket. Compare them in the DevTools Network tab.
  2. Build a chat room over WebSocket.
  3. Run 2 instances: users connected to different instances can't see each other's messages.
  4. Fix it with a Redis pub/sub backplane.
  5. Add heartbeats and client-side reconnects with backoff.
  6. Add delivery acks ("sent ✓", "delivered ✓✓").
- **Break it:** restart one instance: every client on it disconnects at once and they all reconnect together (a thundering herd). Add reconnect jitter and compare.
- **Measure it:** open 10,000 connections with k6 (WebSocket mode) and record server memory per connection; requests per minute from polling vs a single socket for the same updates.
- **Ready to record when:** you can say when SSE is enough, and how WebSockets scale beyond one server.

**Video idea:** *Polling vs SSE vs WebSockets: how live apps actually work*

### T26 · Geospatial indexing, depth L2
**Why it exists:** "find the nearest X" is everywhere (rides, delivery, maps, dating apps), and a normal index can't answer it efficiently.
**Must know**
- Why a normal index on latitude and longitude is a poor fit for "nearest" queries.
- Geohash, its prefix property, and the edge problem (you also check the neighbouring cells).
- Quadtrees; S2 and H3 at headline level.
- Redis GEO vs PostGIS.
- The cost of moving objects that update their location every few seconds.

**Learn from (all free)**
- *Core, enough on its own:*
  1. Redis docs, the geospatial data type and `GEOSEARCH`.
  2. The PostGIS workshop "Introduction to PostGIS" (postgis.net/workshops): the chapters on geometries, spatial indexing and nearest-neighbour search.
  3. Uber Engineering, "H3: Uber's Hexagonal Hierarchical Spatial Index" (2018), plus the docs at h3geo.org.
- *Go further (optional):* The Wikipedia article "Geohash" (good diagrams of the encoding); the S2 Geometry docs (s2geometry.io) for the headline idea.

**Practice**
- **Setup:** Redis + PostGIS (the `postgis/postgis` image), plus a geohash map viewer online.
- **Exercises**
  1. Encode a few locations to geohash by hand with a script, and see nearby places share prefixes.
  2. Find a point near a cell edge whose closest neighbour has a completely different geohash. That's the edge problem.
  3. `GEOADD` 50,000 random points in a city and run `GEOSEARCH` by radius.
  4. Do the same in PostGIS with `ST_DWithin` and a GiST index, with and without the index.
  5. Write a ~50-line quadtree that splits when a cell holds more than N points.
  6. Simulate 5,000 moving points updating every 3 seconds in Redis.
- **Break it:** query with a plain B-tree index on `(lat, lng)` and read the `EXPLAIN` to see why it's slow.
- **Measure it:** nearest-neighbour query time with no index, a B-tree, Redis GEO and a PostGIS GiST index; the update throughput Redis sustains for moving points.
- **Ready to record when:** you can explain geohash and its edge problem on a whiteboard, and when you'd choose PostGIS over Redis.

**Video idea:** *How apps find what's nearest to you in milliseconds*

---

## Schedule

- **Pace:** with the 6-day routine, plan for **one video a week**. 26 videos is about 6 months. Splitting into modules gives natural playlists and natural breaks.
- **Buffer rule:** don't publish T1 until T1 and T2 are recorded. Always keep one finished video ready, so a bad week never breaks the schedule.
- **Order:** follow the module order. Each tool's exercises reuse earlier setups (the lab app, Nginx, Postgres, Redis), so practice gets faster as the season goes on.
- **After each module:** take one lighter week. Re-watch your own videos, read the comments, and note which explanations confused people. It's the cheapest way to improve the next module.

### Before Toolbox #1
- [ ] Mic, OBS and Excalidraw template set up; a throwaway pilot recorded and reviewed.
- [ ] Docker and k6 installed; the `toolbox-labs` repo created.
- [ ] The lab app written and running in docker compose.
- [ ] T1 and T2 practice finished and both videos recorded.
