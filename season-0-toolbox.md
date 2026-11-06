# Season 0: The Toolbox

*One video per component, before (and alongside) Building Haat · Nov 6, 2026 · @Lion*

---

## Why this season exists

Designing Haat in Season 1 means reasoning about load balancers, Redis, queues, replicas, shards and more, often several in one episode. Explaining them convincingly on a whiteboard is hard if you've never run them yourself. Season 0 fixes that. Each video takes **one component**: you learn it deeply, run it with your own hands, break it on purpose, and explain it.

When Season 1 starts, every box on the Haat diagram is one you've already used and already made a video about. Haat episodes can then say *"I explained how Redis works in Toolbox #8; today we'll see what happens when it's under Eid traffic"* and spend their time on the design, not on definitions.

---

## How deep is "deep enough"?

You're building a system design channel, not running these tools in production. Each component in this document gets one of three depth levels:

| Level | Name | What it means | How you prove it |
|---|---|---|---|
| **L1** | Explain | What it is, why it exists, the main trade-off, and where it sits on the diagram. | Explain it in 2 minutes without notes. |
| **L2** | Use | You've installed and run it, used its main features, know its important settings, and have seen how it fails. | You've done the lab, and you can predict what happens when it breaks. |
| **L3** | Internals | How it works inside, and why it's fast or slow. You can predict how it behaves under load and answer "why?" two levels deep. | You can explain it to a senior engineer and survive their follow-up questions. |

**The ceiling:** L3 doesn't mean tuning Kafka brokers or running Postgres for a bank. It means understanding the mechanism well enough that no reasonable viewer question surprises you.

**Where to go deepest:** the relational database, transactions and locking, Redis, caching, queues, replication, sharding and failure handling. These come back in almost every Season 1 episode, so time spent there pays off the most.

---

## The full component map

Every component Haat needs, with how deep to go and where it first matters in Season 1.

| # | Component | Practice with | Depth | First needed in Haat |
|---|---|---|---|---|
| T1 | How a request travels (DNS, TCP, TLS, HTTP) | `dig`, `curl -v`, browser DevTools | L2 | Ep 1 |
| T2 | Web server and reverse proxy | Nginx | L2 | Ep 1 |
| T3 | Relational database and indexes | PostgreSQL | **L3** | Ep 1, 3, 5 |
| T4 | Transactions, isolation and locking | PostgreSQL | **L3** | Ep 1 (planted), 7, 8, 10 |
| T5 | Backups and recovery | `pg_dump`, WAL archiving | L2 | Ep 1, 21 |
| T6 | Sessions, cookies, JWT and OTP login | Small app + Redis | L2 | Ep 1, 2, 6 |
| T7 | Load balancers and stateless servers | Nginx / HAProxy, k6 | L2 | Ep 2 |
| T8 | Redis | Redis | **L3** | Ep 2 onward (8 episodes) |
| T9 | Caching patterns and CDN | Redis, Cloudflare (free tier) | **L3** / L2 | Ep 3 |
| T10 | Object storage and file uploads | MinIO (S3-compatible) | L2 | Ep 2, 4 |
| T11 | Message queues | RabbitMQ | **L3** | Ep 4, 7, 9, 10, 14, 19 |
| T12 | Event logs and pub/sub | Kafka or Redpanda | L2 | Ep 5, 14, 17 |
| T13 | Background jobs, cron and leases | cron, Redis | L2 | Ep 1, 10 |
| T14 | Search engines | OpenSearch / Elasticsearch | L2 (L3 for the inverted index) | Ep 5 |
| T15 | Rate limiting | Redis, Nginx | **L3** | Ep 6 |
| T16 | Idempotency and online payments | Payment gateway sandbox | **L3** | Ep 8, 9, 10, 14 |
| T17 | Database replication | PostgreSQL streaming replica | **L3** | Ep 11, 13, 21 |
| T18 | Sharding and consistent hashing | Script + Postgres (Citus optional) | **L3** | Ep 11, 12, 18 |
| T19 | Distributed ID generation | Script | L2 | Ep 12 |
| T20 | Consistency models and CAP | Paper timelines | L2 (concept-heavy) | Ep 13, 21 |
| T21 | Microservices, sagas, outbox and CDC | 2 small services + RabbitMQ | L2 | Ep 14 |
| T22 | Timeouts, retries, circuit breakers, bulkheads | Toxiproxy + a resilience library | **L3** | Ep 15, 20, 21 |
| T23 | Containers and safe deployments | Docker, docker compose / kind | L2 | Ep 20 |
| T24 | Observability and SLOs | Prometheus, Grafana | L2 | Ep 15, 20, 21 |
| T25 | Real-time: polling, SSE, WebSockets | WebSocket server + Redis pub/sub | L2 | Ep 17, 18 |
| T26 | Geospatial indexing | Redis GEO, PostGIS | L2 | Ep 16, 17 |

**Covered inside Season 1 only (L1 is enough, no Toolbox video):**
- Push notifications (FCM/APNs) and SMS gateways: Ep 9
- Virtual waiting rooms: Ep 7
- Hot keys, sharded counters and HyperLogLog: Ep 18
- News feed fan-out and ranking: Ep 19
- Multi-region, DNS failover and RPO/RTO in practice: Ep 21
- API gateways and service meshes: a short mention in Ep 14 and 20

These are design *patterns* built from the components above, so they're best taught inside the scenario that needs them.

---

## Every Toolbox video follows the same shape

1. **The problem (30-60 s):** a small, Haat-flavoured situation this component solves ("the server forgets you're logged in").
2. **The mental model (2 min):** one simple picture that you'll keep reusing.
3. **How it works inside (main part):** taken to the component's depth level.
4. **Live demo (2-4 min):** the lab you did, shown in a terminal. Season 1 has no coding on screen; Season 0 is where viewers watch you actually run things, and that builds their trust in you.
5. **Break it (1-2 min):** show the failure mode on purpose.
6. **Trade-offs and when *not* to use it (1 min).**
7. **Where it shows up in Haat (15 s):** "You'll see this break in Haat Episode 7."

Aim for 10-15 minutes.

## The 5-day prep routine for a Toolbox video

- **Day 1, Read:** the *Read* resources listed for the component. Write one paragraph in your own words for each *Must know* item.
- **Day 2, Lab:** do the lab from scratch, without copying a finished tutorial end to end. Write down every surprise; surprises make the best video moments.
- **Day 3, Break it and go deeper:** run the failure scenario. For L3 components, answer "why?" twice for each *Must know* item. Write answers to the likely viewer questions.
- **Day 4, Script:** plan the whiteboard panels and the demo steps. Write the hook and the Haat connection word for word, and keep everything else as bullet points. Do one dry run.
- **Day 5, Record:** record the demo separately from the whiteboard parts.

**Reuse your lab environment:** keep one `docker-compose.yml` that grows with the season (Nginx → app ×3 → Postgres → Redis → RabbitMQ → MinIO → OpenSearch …). By T26 it's a small Haat in Docker, and it can be a GitHub repo viewers clone. That's a strong reason for them to subscribe.

---

# The components in detail

Each component lists why Haat needs it, what you must know, what you can skip, what to read, the hands-on lab, and a video title idea.

---

## Block A: Foundations (needed for Haat Episode 1)

### T1 · How a request travels (DNS, TCP, TLS, HTTP), depth L2
**Why Haat needs it:** every request-flow diagram in the season; latency maths; the CDN (Ep 3); DNS failover (Ep 21).
**Must know**
- DNS resolution: recursive resolver, A/CNAME records, TTL and why DNS changes are slow.
- The TCP handshake, and why connections are reused (keep-alive, connection pools).
- The TLS handshake at headline level: certificates and why HTTPS costs extra round trips.
- HTTP: methods, status codes, and the headers that matter later (`Cache-Control`, `Cookie`, `Retry-After`). HTTP/1.1 vs HTTP/2 at headline level.
- Round trips and latency: why Dhaka → Singapore → Dhaka costs ~60 ms every time.

**Skip:** TCP congestion-control internals, QUIC details.
**Read:** *High Performance Browser Networking* by Ilya Grigorik (free at hpbn.co), chapters on latency, TCP and TLS.
**Lab:** `dig` a few domains and watch the TTLs. Run `curl -v` and `curl -w` with timing variables to measure DNS, connect, TLS and first-byte time for a server in Singapore vs the US. Open the DevTools Network waterfall on any shop website.
**Video idea:** *What actually happens when you tap a shop's link on Facebook*

### T2 · Web server and reverse proxy (Nginx), depth L2
**Why Haat needs it:** Ep 1's single server, Ep 2's load balancing, TLS termination, static files.
**Must know**
- Web server vs app server, and why a reverse proxy sits in front.
- TLS termination, serving static files, `proxy_pass` / `upstream`, timeouts, gzip.
- Reading access and error logs.

**Skip:** Nginx modules, OpenResty/Lua.
**Read:** the Nginx "Beginner's Guide" (nginx.org docs); a Hussein Nasser video on reverse proxies.
**Lab:** in Docker, put Nginx in front of two tiny app containers. Serve a static file directly from Nginx. Kill one app container and watch what users see, then read the error log.
**Video idea:** *Why almost every website has Nginx in front of it*

### T3 · Relational database and indexes (PostgreSQL), depth L3
**Why Haat needs it:** the single database (Ep 1), slow pages (Ep 3), `LIKE` search (Ep 5), the 300M-row orders table (Ep 11).
**Must know**
- Tables, primary and foreign keys, normalisation vs denormalisation (headline level).
- **How a B-tree index works**, composite indexes and why column order matters, covering indexes.
- Why `LIKE 'jama%'` can use an index and `LIKE '%jama%'` can't.
- Reading `EXPLAIN ANALYZE`: sequential scan vs index scan.
- Connection limits and pooling (PgBouncer). This is what actually breaks first in Ep 2-3.
- The N+1 query problem.
- MVCC at headline level (it sets up T4).

**Skip:** planner cost-model internals, VACUUM tuning.
**Read:** *Use The Index, Luke* (use-the-index-luke.com), free and excellent; DDIA Chapter 3 (the B-tree section).
**Lab:** generate 5 million product rows with `generate_series`. Time a query, add an index, time it again, and compare the `EXPLAIN` output. Try both `LIKE` forms. Open connections until you hit `max_connections`, then add PgBouncer.
**Video idea:** *Why your database query takes 8 seconds, and how one index makes it 2 ms*

### T4 · Transactions, isolation and locking, depth L3
**Why Haat needs it:** overselling (Ep 7), double charges (Ep 8), double-released stock (Ep 10), and why sagas exist (Ep 14).
**Must know**
- ACID in practical terms.
- The check-then-act race condition.
- Isolation levels (read committed, repeatable read, serializable) and the anomalies each allows: lost update, write skew.
- MVCC: why readers don't block writers.
- Three fixes for a race: `SELECT ... FOR UPDATE`, an atomic conditional `UPDATE ... WHERE stock > 0`, and optimistic locking with a version column.
- Deadlocks: how they happen and what the database does about them.

**Skip:** lock-manager internals, implementing two-phase commit.
**Read:** DDIA Chapter 7 "Transactions"; the PostgreSQL docs chapter "Transaction Isolation".
**Lab:** in two `psql` sessions, reproduce a lost update by hand. Fix it all three ways. Create a deadlock on purpose. Write a small script that fires 100 concurrent "checkouts" at an item with 10 in stock: watch it oversell, then fix it.
**Video idea:** *How two people bought the last phone at the same second*

### T5 · Backups and recovery, depth L2
**Why Haat needs it:** Ep 1's backup plan; RPO/RTO in the finale (Ep 21).
**Must know**
- Logical backups (`pg_dump`) vs physical backups.
- WAL and point-in-time recovery.
- The 3-2-1 rule, off-site copies, RPO and RTO.
- Why an untested backup isn't a backup.
- Why a replica isn't a backup.

**Skip:** comparing backup tools in depth (pgBackRest and others): just name them.
**Read:** the PostgreSQL docs "Backup and Restore"; GitLab's "Postmortem of database outage of January 31" (2017).
**Lab:** take a `pg_dump`, drop a table, and restore it, timing each step. *Stretch:* enable WAL archiving and recover to the minute before a deliberate "oops" `DELETE`.
**Video idea:** *Your backups probably don't work. Here's how GitLab found out*

### T6 · Sessions, cookies, JWT and OTP login, depth L2
**Why Haat needs it:** OTP login (Ep 1), users logged out by in-memory sessions (Ep 2), OTP abuse (Ep 6).
**Must know**
- Cookies and their flags (`HttpOnly`, `Secure`, `SameSite`).
- Server-side sessions vs JWT, and the real trade-off: revocation and size.
- Where sessions live once you have 3 servers.
- The OTP flow: generate, store a hash, expire, limit attempts.
- Password hashing (bcrypt/argon2) at headline level, and HTTPS everywhere.

**Skip:** OAuth/OIDC flows in depth (L1: "Login with Facebook" exists and roughly how it works).
**Read:** the OWASP Session Management and Authentication Cheat Sheets.
**Lab:** build a tiny app that keeps sessions in memory. Run two copies behind Nginx round-robin and watch users get logged out at random. Move the sessions to Redis and the problem is gone (a preview of T8).
**Video idea:** *How a website remembers that you're logged in*

---

## Block B: Scaling one shop (Haat Episodes 2-6)

### T7 · Load balancers and stateless servers, depth L2
**Why Haat needs it:** Ep 2; health checks and autoscaling later; the global load balancer in Ep 21.
**Must know**
- L4 vs L7 load balancing.
- Algorithms: round robin, least connections, hashing.
- Health checks.
- Sticky sessions, and why stateless servers are better.
- Vertical vs horizontal scaling.
- The load balancer as a single point of failure (pairs, managed load balancers).
- Autoscaling at headline level.

**Skip:** BGP/anycast (L1, for Ep 21).
**Read:** Alex Xu Vol. 1, Chapter 1; the Nginx docs "HTTP Load Balancing"; AWS Builders' Library "Implementing health checks".
**Lab:** extend the T2 setup to 3 app servers. Load-test 1 server vs 3 with k6 or `hey`. Kill one server mid-test and watch the health check take it out of rotation.
**Video idea:** *A server died and nobody noticed: load balancers explained*

### T8 · Redis, depth L3 (the most reused component in the season)
**Why Haat needs it:** sessions (2), cache (3), rate limits (6), stock counter (7), idempotency store (8), scheduler (10), rider locations (16), like counters (18).
**Must know**
- Why it's fast: in memory, single-threaded execution, so each command is atomic.
- Data types: strings, hashes, lists, sets, sorted sets, geo.
- `INCR`, `SET key value NX EX`, TTLs, `MULTI` and Lua scripts for multi-step atomicity.
- Persistence (RDB vs AOF), and exactly what you lose in a crash.
- `maxmemory` and eviction policies.
- Replication, Sentinel and Cluster at headline level. Hot keys.

**Skip:** Redis Cluster slot internals; the Redlock debate (L1: know it exists and why it's contested).
**Read:** the Redis docs "Data types" and "Persistence"; *Redis in Action* by Josiah Carlson (free online from Redis).
**Lab:** in `redis-cli`, build a counter, a lock with `SET NX EX`, a leaderboard with `ZADD`/`ZRANGE`, and a rate limiter, and try `GEOADD`/`GEOSEARCH`. Run `redis-benchmark`. Set a small `maxmemory`, fill it up and watch keys get evicted. Kill Redis with and without AOF and compare what survives.
**Video idea:** *Redis: the Swiss Army knife behind every fast app*

### T9 · Caching patterns and CDN, depth L3 (patterns) / L2 (CDN)
**Why Haat needs it:** the Eid campaign (Ep 3), and the feed cache later (Ep 19).
**Must know**
- Cache-aside, write-through and write-behind.
- TTLs and invalidation, and why invalidation is the hard part.
- Hit-ratio maths: how a 95% hit rate changes database load.
- The cache stampede and its fixes: locking/request coalescing, jittered TTLs, stale-while-revalidate.
- What never to cache blindly (stock, prices at checkout).
- HTTP cache headers. CDN basics: edge, origin, cache key, purge.

**Read:** AWS Builders' Library "Caching challenges and strategies"; the Cloudflare Learning Center "What is a CDN?"; "Scaling Memcache at Facebook" (2013), the leases section.
**Lab:** put a Redis cache in front of a slow Postgres query and measure the hit ratio. Expire a hot key under load and watch the database spike (that's a stampede), then fix it with a lock. Put a static page behind Cloudflare's free tier and check the `cf-cache-status` header.
**Video idea:** *The cache that took down the database*

### T10 · Object storage and file uploads, depth L2
**Why Haat needs it:** the local image-folder hack (Ep 2), uploads failing at 95% (Ep 4).
**Must know**
- File system vs object storage: buckets, keys, objects you replace rather than edit, durability.
- Presigned URLs, so the phone uploads straight to storage and not through your servers.
- Multipart and resumable uploads (the tus protocol).
- An image pipeline: upload → queue → thumbnail worker.
- A CDN in front of storage, and storage cost maths.

**Skip:** erasure-coding internals (L1).
**Read:** the AWS S3 docs on presigned URLs and multipart upload; the tus.io protocol overview.
**Lab:** run MinIO in Docker. Upload with a presigned URL using `curl`. Upload a 1 GB file with multipart, cut the network halfway and resume it.
**Video idea:** *Why uploading videos to your own server is a bad idea*

### T11 · Message queues, depth L3
**Why Haat needs it:** media processing (4), order writes (7), notifications (9), scheduled jobs (10), saga events (14), feed fan-out (19).
**Must know**
- Producers and consumers, and why going async smooths out traffic peaks.
- Acks, visibility timeouts and redelivery.
- At-most-once vs at-least-once delivery, and why "exactly once" really means at-least-once plus idempotency.
- Retries with backoff, dead-letter queues, ordering guarantees, priority queues, backpressure.
- Queue depth as the metric to watch.

**Skip:** AMQP exchange types beyond direct and fanout.
**Read:** RabbitMQ tutorials 1-6 (rabbitmq.com); the AWS SQS docs on visibility timeout; DDIA Chapter 11, the messaging section.
**Lab:** run RabbitMQ in Docker (use the management UI). Run a producer and 3 consumers. Kill a consumer mid-message and watch the message get redelivered. Send a "poison" message and watch it land in the DLQ. Burst 100,000 messages and watch the queue drain.
**Video idea:** *Why big apps make your request wait in line*

### T12 · Event logs and pub/sub (Kafka), depth L2
**Why Haat needs it:** CDC to search (Ep 5), saga events (Ep 14), live updates (Ep 17).
**Must know**
- A log vs a queue: retention and replay.
- Topics, partitions, keys and ordering within a partition.
- Consumer groups and offsets.
- Pub/sub fan-out: Redis pub/sub (fire-and-forget) vs Kafka (durable).
- When *not* to use Kafka.

**Skip:** broker internals, ISR tuning, Kafka Streams.
**Read:** the Kafka docs "Introduction"; Jay Kreps, "The Log: What every software engineer should know about real-time data's unifying abstraction" (2013).
**Lab:** run Kafka (KRaft mode) or Redpanda in Docker. Create a 3-partition topic and two consumer groups. Replay from offset 0. Show that messages with the same key stay in order.
**Video idea:** *Kafka vs RabbitMQ: do you need a queue or a log?*

### T13 · Background jobs, cron and leases, depth L2
**Why Haat needs it:** Ep 1's cron job; the same cron cloned onto 6 servers (Ep 10).
**Must know**
- Cron basics and worker pools.
- The "job ran twice" problem.
- Leases with TTLs, leader election at headline level, fencing tokens.
- Delayed jobs with sorted sets.
- Making jobs idempotent.

**Read:** Martin Kleppmann, "How to do distributed locking" (2016).
**Lab:** run the same cron job in two containers and watch the work happen twice. Fix it with a Redis lease (`SET NX EX`). Then `docker pause` the leaseholder past its TTL: both run again. Add a fencing token check in the database to stop it.
**Video idea:** *The cron job that ran six times*

### T14 · Search engines, depth L2 (L3 for the inverted index)
**Why Haat needs it:** the 8-second search box (Ep 5).
**Must know**
- How an inverted index works.
- Analyzers: tokenising, stemming, synonyms, and handling Banglish spellings ("jama", "jamah").
- Relevance at headline level (TF-IDF/BM25), filters and facets.
- Autocomplete: edge n-grams and tries.
- Keeping the index in sync: dual writes vs CDC. Search results being slightly behind is fine.

**Skip:** cluster tuning and shard sizing beyond headline level.
**Read:** the Elastic docs "Text analysis" and the relevance section; the PostgreSQL docs "Full Text Search".
**Lab:** index 100,000 products in OpenSearch. Compare `LIKE`, Postgres full-text search (`tsvector`) and OpenSearch on the same queries. Build autocomplete with edge n-grams.
**Video idea:** *How search finds "jama" among 2 million products in 20 ms*

### T15 · Rate limiting, depth L3
**Why Haat needs it:** bots, scraping and the OTP SMS bill (Ep 6).
**Must know**
- Fixed window, sliding-window log and counter, token bucket, leaky bucket, with the trade-offs of each.
- Where to limit (edge, gateway, app) and what to key on (IP, user, phone number).
- Distributed limiting with Redis.
- `429` and `Retry-After`.
- Bot defences beyond limits: CAPTCHA, making abuse expensive.

**Read:** Alex Xu Vol. 1, Chapter 4; the Stripe blog "Scaling your API with rate limiters" (2017); the Cloudflare blog on how they built rate limiting (2017).
**Lab:** implement a fixed-window limiter and a token-bucket limiter in Redis (the token bucket in Lua). Load-test with k6 and plot allowed vs rejected requests. Try Nginx's `limit_req`.
**Video idea:** *How a bot can run up a huge SMS bill overnight: rate limiting explained*

---

## Block C: The big sale (Haat Episodes 7-10)

T4 (locking) and T11 (queues) carry most of Act 2. Re-watch your own videos before Episode 7.

### T16 · Idempotency and online payments, depth L3
**Why Haat needs it:** double charges (Ep 8), duplicate notifications (Ep 9), duplicate jobs (Ep 10), compensations (Ep 14).
**Must know**
- How a payment gateway flow works: redirect → customer pays → callback/webhook → you verify.
- Why networks create "did it go through?" moments: timeouts don't tell you whether the payment happened.
- Idempotency keys: store, lock, replay the saved response.
- The payment state machine.
- Webhooks arrive at least once, so they need deduplication.
- Reconciliation jobs, and refunds.

**Skip:** card-network internals, PCI DSS details (L1).
**Read:** the Stripe blog "Designing robust and predictable APIs with idempotency" (2017); Alex Xu Vol. 2 "Payment system"; the developer docs of a local gateway (bKash, SSLCommerz).
**Lab:** in a gateway's sandbox or test mode, create a payment and receive the webhook (through a tunnel such as ngrok or cloudflared). Replay the webhook twice. Add an idempotency table with a unique constraint and replay again.
**Video idea:** *Why you got charged twice, and how payment systems prevent it*

---

## Block D: Growing up (Haat Episodes 11-15)

Before this block, watch Martin Kleppmann's distributed systems lectures on replication and consistency.

### T17 · Database replication, depth L3
**Why Haat needs it:** read replicas (Ep 11), "My Orders is empty" (Ep 13), failover (Ep 21).
**Must know**
- Primary/replica setups, synchronous vs asynchronous replication.
- Replication lag and the anomalies it causes: read-your-own-writes and monotonic reads.
- Failover and split brain.
- Read scaling, and why a replica isn't a backup.

**Read:** DDIA Chapter 5; the PostgreSQL docs "High Availability, Load Balancing, and Replication".
**Lab:** set up a Postgres primary with a streaming replica in docker compose. Write to the primary and immediately read from the replica. Add `recovery_min_apply_delay` so the lag is visible: that's Ep 13 live. Promote the replica.
**Video idea:** *"I placed an order but it's not there": replication lag explained*

### T18 · Sharding and consistent hashing, depth L3
**Why Haat needs it:** splitting the orders table (Ep 11), order IDs (Ep 12), hot keys (Ep 18).
**Must know**
- Vertical vs horizontal partitioning.
- Choosing the shard key (`user_id` vs `order_id`), range vs hash sharding, hot spots.
- Cross-shard queries and joins, and why resharding hurts.
- Consistent hashing with virtual nodes; directory-based sharding.

**Read:** DDIA Chapter 6; Alex Xu Vol. 1, Chapter 5; the Notion blog "Herding elephants: lessons learned from sharding Postgres at Notion" (2021); Instagram Engineering "Sharding & IDs at Instagram" (2012).
**Lab:** write a ~30-line consistent-hash ring script. Add and remove a node and count how many keys move, compared with `hash % N`. *Optional:* shard a table with the Citus extension.
**Video idea:** *When one database isn't enough: sharding explained*

### T19 · Distributed ID generation, depth L2
**Why Haat needs it:** two customers with the same order number (Ep 12).
**Must know**
- Why auto-increment breaks across shards.
- UUID v4 vs v7 (and why index locality matters).
- The Snowflake layout (timestamp / machine / sequence) and clock skew.
- Base62 short codes, and why guessable IDs are a security issue.

**Read:** Alex Xu Vol. 1, Chapter 7; Twitter's "Announcing Snowflake" (2010).
**Lab:** write a Snowflake generator. Insert 1 million UUIDv4 keys and 1 million sequential keys, then compare index size and insert speed.
**Video idea:** *How big apps generate billions of unique IDs without a single counter*

### T20 · Consistency models and CAP, depth L2 (concept-heavy)
**Why Haat needs it:** consistency tiers (Ep 13), the CAP decision in the finale (Ep 21).
**Must know**
- Strong vs eventual consistency; read-your-writes, monotonic reads, linearizability at headline level.
- **CAP stated correctly:** the choice only applies during a network partition. PACELC.
- Quorums (N, R, W).
- Choosing consistency per feature: stock must be strong; like counts can be eventual.

**Read:** DDIA Chapter 9; Martin Kleppmann, "Please stop calling databases CP or AP" (2015).
**Lab:** this one is on paper. Draw timelines for each anomaly with two clients and two replicas until you can draw them from memory. This video is mostly whiteboard.
**Video idea:** *The CAP theorem is misunderstood. Here's what it really says*

### T21 · Microservices, sagas, outbox and CDC, depth L2
**Why Haat needs it:** paid, never shipped (Ep 14).
**Must know**
- When to split a monolith, and how to draw service boundaries.
- Synchronous calls (REST/gRPC) vs asynchronous events.
- Why distributed transactions (2PC) are avoided.
- Sagas (orchestration vs choreography) and compensating actions.
- The dual-write problem, the transactional outbox, and CDC (Debezium).

**Read:** microservices.io on the Saga and Transactional Outbox patterns; Martin Fowler, "Microservice Prerequisites"; the Debezium docs on the outbox event router.
**Lab:** build two tiny services connected by RabbitMQ. Kill the order service between "save to database" and "publish event" to show the dual-write bug. Fix it with an outbox table and a relay.
**Video idea:** *Paid but never shipped: the dual-write problem*

### T22 · Timeouts, retries, circuit breakers and bulkheads, depth L3
**Why Haat needs it:** the slow courier API (Ep 15), and again in Ep 20 and 21.
**Must know**
- Why a missing timeout cascades: thread and connection pools fill up.
- Choosing timeouts from p99 latency.
- Exponential backoff with jitter, and retry storms.
- Circuit breaker states, bulkheads, fallbacks, load shedding.

**Read:** *Release It!* (2nd ed.), the chapters on stability antipatterns and stability patterns; AWS Builders' Library "Timeouts, retries, and backoff with jitter"; Martin Fowler, "CircuitBreaker".
**Lab:** have a service call a dependency through Toxiproxy and inject latency. Watch the caller's pool run out and every endpoint die. Add a timeout, then a circuit breaker (a library in your language, or write a simple one), and run the test again.
**Video idea:** *How one slow API can take down an entire app*

### T24 · Observability and SLOs, depth L2
*(Listed here because Episode 15 is where you first need it to tell the story.)*
**Why Haat needs it:** spotting a cascading failure (Ep 15), deploy safety (Ep 20), finale decisions (Ep 21).
**Must know**
- Logs vs metrics vs traces.
- The RED and USE methods.
- Percentiles (p50, p99) vs averages.
- SLIs, SLOs and error budgets; alerting on symptoms, not causes.
- Distributed tracing at headline level.

**Read:** the Google SRE book, chapters "Service Level Objectives" and "Monitoring Distributed Systems"; the Prometheus "Getting started" docs.
**Lab:** run Prometheus and Grafana in Docker scraping your lab app. Build a dashboard of request rate, error rate and p99 latency, and add one alert rule. Re-run the T22 slow-dependency test and watch it on the dashboard.
**Video idea:** *Average latency lies: p99 and SLOs explained*

---

## Block E: Haat comes alive (Haat Episodes 16-19)

### T26 · Geospatial indexing, depth L2
**Why Haat needs it:** the nearest rider (Ep 16), live tracking (Ep 17).
**Must know**
- Why a normal index on latitude and longitude is a poor fit for "nearest" queries.
- Geohash and its edge problem (you also check the neighbouring cells).
- Quadtrees; S2 and H3 at headline level.
- Redis GEO vs PostGIS.
- How often moving objects update, and what that costs.

**Read:** Alex Xu Vol. 2 "Proximity service"; the Uber Engineering post "H3: Uber's Hexagonal Hierarchical Spatial Index" (2018); the Redis GEO docs.
**Lab:** `GEOADD` 50,000 random points around Dhaka and run radius searches. Do the same in PostGIS with `ST_DWithin` and a GiST index. Look at geohash cells on an online geohash map.
**Video idea:** *How apps find the nearest rider in milliseconds*

### T25 · Real-time: polling, SSE and WebSockets, depth L2
**Why Haat needs it:** "Where is my rider?" plus chat (Ep 17), the live stream (Ep 18).
**Must know**
- Short polling, long polling, SSE and WebSockets, with the trade-offs of each.
- What each connection costs (memory, file descriptors).
- Routing connections to the right server; a gateway with a pub/sub backplane.
- Heartbeats and reconnects, delivery receipts and acks, flaky mobile networks.

**Read:** the MDN WebSockets guide; Hussein Nasser's videos on WebSockets and SSE; the Discord blog "How Discord Scaled Elixir to 5,000,000 Concurrent Users" (2017).
**Lab:** build a WebSocket chat server and run 2 instances: messages don't reach users on the other instance. Add a Redis pub/sub backplane and they do. Open 10,000 connections with k6 and watch memory.
**Video idea:** *How "Where is my rider?" updates live on your phone*

---

## Block F: Never go down (Haat Episodes 20-21)

### T23 · Containers and safe deployments, depth L2
**Why Haat needs it:** the one-line change that broke checkout (Ep 20), deploying in two regions (Ep 21).
**Must know**
- Docker images and containers; what Kubernetes does (L1).
- Rolling, blue-green and canary deploys.
- Feature flags.
- **Expand/contract database migrations** (how to rename a column without downtime).
- Rollbacks, and the stages of a CI/CD pipeline.

**Read:** Martin Fowler's articles "BlueGreenDeployment", "CanaryRelease" and "ParallelChange"; Pete Hodgson, "Feature Toggles" (martinfowler.com).
**Lab:** with a load test running, rename a column the naive way and watch requests fail. Then do it with expand/contract and no errors. Do a rolling update with docker compose or a local kind cluster.
**Video idea:** *How to deploy 40 times a day without breaking checkout*

---

## Schedule

### Recommended: build the toolbox in blocks
Each block goes out right before the Haat act that uses it. You learn every component deeply before you design with it, and Haat starts in about 6 weeks rather than 6 months.

| Phase | Toolbox videos | Then publish | Approx. time |
|---|---|---|---|
| 1 | T1-T6 (foundations) | Haat Ep 1 | ~6 weeks |
| 2 | T7, T8, T9, T10, T11, T14, T15, T13 | Haat Ep 2-6, each one after its Toolbox video (T7+T8 → Ep 2, T9 → Ep 3, T10+T11 → Ep 4, T14 → Ep 5, T15 → Ep 6) | ~10-12 weeks |
| 3 | T16, T12 | Haat Ep 7-10 (T16 before Ep 8) | ~6 weeks |
| 4 | T17, T18, T19, T20, T21, T22, T24 | Haat Ep 11-15 | ~12 weeks |
| 5 | T26, T25 | Haat Ep 16-19 | ~6 weeks |
| 6 | T23 | Haat Ep 20-21 | ~3 weeks |

Viewers get a natural rhythm ("learn the tool → watch it break in Haat"), and the playlist works both ways: Haat episodes link back to their Toolbox videos.

**Alternative:** publish all 26 Toolbox videos first. At one a week that's about 6 months before Haat Episode 1. It works, but the series' big hook (the story) waits a long time.

### Before Toolbox #1
- Set up the mic, OBS and Excalidraw template, and record a throwaway pilot (see the pre-season plan in the Season 1 doc).
- Create the lab repo with an empty `docker-compose.yml`.
- Have T1 and T2 recorded before publishing T1, so you always keep a one-video buffer.

### Ready for Haat Episode 1 when
- [ ] T1-T6 are published, and each lab ran on your machine.
- [ ] You can draw the request flow, explain a B-tree index, and reproduce a race condition without notes.
- [ ] You've restored a database from a backup yourself.
- [ ] You can explain why in-memory sessions break once there are 2 servers.
