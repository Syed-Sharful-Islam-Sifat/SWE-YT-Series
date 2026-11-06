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

Every tool below has a **Study path**: a numbered list of exact pages, chapters, sections and lecture numbers, in the order to study them, each with a time estimate. Everything on it is free. Finish the study path and the practice exercises, and you have enough for the video. Some tools also have an *If something is still unclear* line: use it only if you need it.

All links were checked on Nov 6, 2026. Most opened directly. A few sites (Cloudflare, X/Twitter, Uber, LinkedIn, Instagram's blog) block automated checks, so those were confirmed by title search instead; they open normally in a browser. The Stanford IR book PDFs (T11) couldn't be reached from the checking machine, but the book is still listed at nlp.stanford.edu/IR-book. If a link ever breaks, search its exact title in quotes.

The paid books (*Designing Data-Intensive Applications*, Alex Xu's *System Design Interview*, *Release It!*) are excellent but optional. Nothing in this season needs them.

**The three free courses used most across the season** (bookmark them once):
- [Martin Kleppmann, *Distributed Systems* lecture series](https://www.youtube.com/playlist?list=PLeKd45zvjcDFUEv_ohr_HdUFe97RItdiB) (Cambridge, 23 short lectures) and its [lecture notes (PDF)](https://www.cl.cam.ac.uk/teaching/2122/ConcDisSys/dist-sys-notes.pdf). Used in T16 and T19.
- [CMU 15-445 *Intro to Database Systems*, Fall 2024](https://www.youtube.com/playlist?list=PLSE8ODhjZXjYDBpQnSymaectKjxCy6BYq) by Andy Pavlo, with slides on the [course schedule](https://15445.courses.cs.cmu.edu/fall2024/schedule.html). Used in T4 and T5.
- [Google SRE book and SRE Workbook](https://sre.google/books/). Used in T6 and T22.

**Good for revision once a module is done:** [The System Design Primer](https://github.com/donnemartin/system-design-primer), [Azure Cloud Design Patterns](https://learn.microsoft.com/en-us/azure/architecture/patterns/), and the [AWS Builders' Library](https://builder.aws.com/learn/topics/builders-library).

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

- **Day 1, Read:** the tool's *Study path*, in order. Write one paragraph in your own words for each *Must know* item.
- **Day 2, Exercises:** Setup, plus the first half of the exercises.
- **Day 3, Exercises:** the rest of the exercises, plus *Break it*.
- **Day 4, Measure and go deeper:** *Measure it*, plus the mini-project if there's one. For L3 tools, answer "why?" twice for each *Must know* item. Write answers to the questions viewers are likely to ask.
- **Day 5, Script:** turn the bullet points you wrote for each *Video breakdown* row into whiteboard panels (roughly one panel per row), then plan the demo steps. Write the hook word for word and keep the rest as bullet points. Do one dry run.
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

Aim for 10-15 minutes. Part 3 is split into sub-parts (3a, 3b, …), one per concept, at about 1-2 minutes each.

### How to use each tool's Video breakdown

Every tool has a **Video breakdown** table just before its video idea. Each row is one part of the video:
- **What you cover:** what that part says or shows.
- **Learn from:** the numbered steps of that tool's *Study path* that teach it ("Study 2, 5" means steps 2 and 5).
- **Practise with:** the matching exercise in that tool's *Practice* section ("Ex 3", "Break it", "Measure it", "Mini-project").

Work through it **part by part, not resource by resource**:
1. Pick a row. Do its *Learn from* steps, then its *Practise with* exercise.
2. Straight away, write 3-5 bullet points for that part of the script, plus one sentence on what you'll draw or show.
3. Tick the row and move to the next.

When every row is ticked, your script outline is already written, and every part of it is backed by something you've read *and* done. Rows with "—" in *Learn from* come from your own practice (demos and numbers). Rows with "—" in *Practise with* are explanation-only, so rehearse them out loud.

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
| T10 | Object storage and file uploads | Garage (S3-compatible) | L2 |
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
**Study path (free, in this order, ~5 h)**
1. [HPBN, Ch. 1 "Primer on Latency and Bandwidth"](https://hpbn.co/primer-on-latency-and-bandwidth/): the whole chapter. *(40 min)*
2. [HPBN, Ch. 2 "Building Blocks of TCP"](https://hpbn.co/building-blocks-of-tcp/): read "Three-Way Handshake" and "Head-of-Line Blocking" closely; skim the congestion-control sections. *(45 min)*
3. [HPBN, Ch. 4 "Transport Layer Security (TLS)"](https://hpbn.co/transport-layer-security-tls/): read "Encryption, Authentication, and Integrity", "TLS Handshake" and "Chain of Trust and Certificate Authorities". *(45 min)*
4. [MDN, "An overview of HTTP"](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Overview): the whole page. *(30 min)*
5. [Cloudflare, "What is DNS?"](https://www.cloudflare.com/learning/dns/what-is-dns/): the whole page, including the 8 steps of a lookup. *(20 min)*
6. [Mess With DNS](https://messwithdns.net/): do the built-in experiments. You create real records and watch queries arrive live. *(1 h)*

*If something is still unclear:* ["How DNS Works" comic](https://dnsimple.com/comics/how-dns-works/); [Cloudflare, "What happens in a TLS handshake?"](https://www.cloudflare.com/learning/ssl/what-happens-in-a-tls-handshake/).

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | You tap a link and a page appears ~300 ms later. What happened in those milliseconds? | Study 1 | Measure it (your numbers make the hook) |
| 2 · Mental model | A request is a chain of round trips: DNS → TCP → TLS → HTTP → response. | Study 1 | Ex 4 (the waterfall *is* the picture) |
| 3a · DNS | Resolvers, A/CNAME records, TTL, caching, why DNS changes are slow. | Study 5, 6 | Ex 1 |
| 3b · TCP | The three-way handshake; why connections are reused (keep-alive). | Study 2 | Ex 4 |
| 3c · TLS | The handshake, certificates, the chain of trust, the extra round trips. | Study 3 | Ex 2 |
| 3d · HTTP | Request and response as plain text: methods, status codes, key headers. | Study 4 | Ex 3 |
| 4 · Live demo | `dig +trace` → raw HTTP with `nc` → `curl -v`. | — | Ex 1-3 |
| 5 · Break it | Wrong IP in `/etc/hosts`; a certificate error from the wrong hostname. | Study 3 (chain of trust) | Break it |
| 6 · Your numbers | Your DNS / connect / TLS / first-byte table for 3 regions. | Study 1 | Measure it |
| 7 · When it matters | Why CDNs, connection reuse and HTTP/2 exist: they cut round trips. | Study 1, 2 | — |

**Video idea:** *What actually happens in the 300 ms after you tap a link*

### T2 · Web server and reverse proxy (Nginx), depth L2
**Why it exists:** app servers are good at running code and bad at handling thousands of slow connections, TLS and static files. A reverse proxy takes that job off them.
**Must know**
- Web server vs app server, and what a reverse proxy does.
- TLS termination, serving static files, `proxy_pass` / `upstream`, timeouts, gzip, request size limits.
- Reading access and error logs.

**Skip:** Nginx modules, OpenResty/Lua.
**Study path (free, in this order, ~3 h)**
1. [Nginx, "Beginner's Guide"](https://nginx.org/en/docs/beginners_guide.html): the whole page. *(30 min)*
2. [NGINX Admin Guide, "Serving Static Content"](https://docs.nginx.com/nginx/admin-guide/web-server/serving-static-content/). *(20 min)*
3. [NGINX Admin Guide, "NGINX Reverse Proxy"](https://docs.nginx.com/nginx/admin-guide/web-server/reverse-proxy/): the whole page (headers, buffers, binding). *(30 min)*
4. [DigitalOcean, "Understanding Nginx Server and Location Block Selection Algorithms"](https://www.digitalocean.com/community/tutorials/understanding-nginx-server-and-location-block-selection-algorithms): how Nginx picks which config applies. *(45 min)*
5. Keep open while practising: [`ngx_http_proxy_module` reference](https://nginx.org/en/docs/http/ngx_http_proxy_module.html). Look up `proxy_pass`, `proxy_set_header`, `proxy_read_timeout` and `proxy_connect_timeout`. *(as needed)*

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | An app server handling slow clients, TLS and images by itself wastes its capacity. | Study 3 | — |
| 2 · Mental model | Nginx is the front desk: it handles the door and passes real work to the app. | Study 1 | — |
| 3a · Config structure | `http` / `server` / `location` blocks, and how Nginx picks the block for a request. | Study 1, 4 | Ex 1 |
| 3b · Static files | Serving files directly without touching the app. | Study 2 | Ex 1 |
| 3c · Reverse proxy | `proxy_pass`, forwarding headers (`X-Forwarded-For`), timeouts. | Study 3, 5 | Ex 2 |
| 3d · TLS, gzip, limits | TLS termination, compression, `client_max_body_size`. | Study 3, 5 | Ex 3-5 |
| 4 · Live demo | Static file → proxied API → HTTPS, all in one config. | — | Ex 1-3 |
| 5 · Break it | 502 (app down) vs 504 (app too slow): what each one means. | Study 5 (timeouts) | Break it |
| 6 · Your numbers | Requests per second for a static file: Nginx vs the app. | — | Measure it |
| 7 · When not to use it | When a managed load balancer or CDN already does the job; alternatives (Caddy, Traefik, HAProxy) at L1. | Study 3 | — |

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
**Study path (free, in this order, ~3 h)**
1. [Sam Rose, "Load Balancing"](https://samwho.dev/load-balancing/): the whole interactive essay, playing with every simulation. *(45 min)*
2. [NGINX Admin Guide, "HTTP Load Balancing"](https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/): load-balancing methods, server weights, passive health checks. *(45 min)*
3. [AWS Builders' Library, "Implementing health checks"](https://builder.aws.com/content/3Ev53O39izHCtWLzp4XU6t8PC1O/implementing-health-checks): the whole article. *(45 min)*
4. [HAProxy docs, "Health checks"](https://www.haproxy.com/documentation/haproxy-configuration-tutorials/reliability/health-checks/): for the active health checks in exercise 3. *(20 min)*

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | One server hits its limit, or dies, and the whole site goes down. | Study 1 | — |
| 2 · Mental model | A traffic officer spreading cars across lanes, and closing a lane that's blocked. | Study 1 | — |
| 3a · Algorithms | Round robin, least connections, weighted, hashing. | Study 1, 2 | Ex 1-2 |
| 3b · Health checks | Passive vs active checks; what a good health check tests. | Study 3, 4 | Ex 3 |
| 3c · Stateless servers | Why state inside a server breaks balancing; sticky sessions as a crutch. | Study 1, 2 | Ex 4 |
| 3d · L4 vs L7, and the LB as SPOF | What each layer sees; pairs, floating IPs, managed load balancers. | Study 2, 3 | — |
| 4 · Live demo | Responses rotating across 3 servers, then the in-memory counter giving 3 different answers. | — | Ex 1, 4 |
| 5 · Break it | Kill a server mid-load-test and count the failed requests. | Study 3 | Break it |
| 6 · Your numbers | Requests per second with 1, 2, 3 servers, and why it isn't linear. | — | Measure it |
| 7 · When to use it | When one bigger server is simpler; DNS round robin; managed load balancers. | Study 1, 3 | — |

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
**Study path (free, in this order, ~10 h)**
1. [Use The Index, Luke, Ch. 1 "Anatomy of an Index"](https://use-the-index-luke.com/sql/anatomy): all three sections (leaf nodes, the B-tree, slow indexes). *(1 h)*
2. [Use The Index, Luke, Ch. 2 "The Where Clause"](https://use-the-index-luke.com/sql/where-clause): every section. It includes ["LIKE filters"](https://use-the-index-luke.com/sql/where-clause/searching-for-ranges/like-performance-tuning), which is exercise 4. *(2 h)*
3. [Use The Index, Luke, Ch. 3 "Performance and Scalability"](https://use-the-index-luke.com/sql/testing-scalability). *(45 min)*
4. [Use The Index, Luke, Ch. 5 "Clustering Data"](https://use-the-index-luke.com/sql/clustering): index-only scans and covering indexes. *(45 min)*
5. [Use The Index, Luke, Ch. 8 "Modifying Data"](https://use-the-index-luke.com/sql/dml): why every index slows down inserts. *(30 min)*
6. [PostgreSQL docs, "Indexes"](https://www.postgresql.org/docs/current/indexes.html): the sections "Introduction", "Index Types", "Multicolumn Indexes", "Partial Indexes" and "Index-Only Scans and Covering Indexes". *(1 h)*
7. [PostgreSQL docs, "Using EXPLAIN"](https://www.postgresql.org/docs/current/using-explain.html): "EXPLAIN Basics" and "EXPLAIN ANALYZE". *(1 h)*
8. Watch: [CMU 15-445 (Fall 2024)](https://www.youtube.com/playlist?list=PLSE8ODhjZXjYDBpQnSymaectKjxCy6BYq), **Lecture #08 "Indexes & Filters I"** (B+trees). Slides are on the [course schedule](https://15445.courses.cs.cmu.edu/fall2024/schedule.html). *(1.5 h)*
9. [PgBouncer, "Usage"](https://www.pgbouncer.org/usage.html): the pooling modes (session, transaction, statement). *(20 min)*

*If something is still unclear:* SQL warm-up at [pgexercises.com](https://pgexercises.com/) (the "Basic" and "Joins" sets); paste plans into [explain.depesz.com](https://explain.depesz.com/) to read them more easily; CMU Lecture #03 "Database Storage I" for how rows sit on disk.

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | A simple query on 5 million rows takes seconds. | — | Ex 1-2 |
| 2 · Mental model | The index at the back of a book: sorted, so you can skip straight to the page. | Study 1 | — |
| 3a · How a B-tree works | Leaf nodes, the tree above them, and how a lookup walks down it. | Study 1, 8 | — |
| 3b · Reading EXPLAIN | Sequential scan vs index scan; reading cost, rows and time. | Study 7 | Ex 2 |
| 3c · Composite indexes | Why column order matters. | Study 2, 6 | Ex 3 |
| 3d · LIKE | Why `'abc%'` can use an index and `'%abc%'` can't. | Study 2 (LIKE filters) | Ex 4 |
| 3e · Covering indexes | Index-only scans: answering from the index alone. | Study 4, 6 | Ex 5 |
| 3f · The cost of indexes | Slower writes and extra disk space. | Study 5 | Measure it |
| 3g · Beyond indexes | N+1 queries and connection pooling. | Study 3, 9 | Ex 6, Break it |
| 4 · Live demo | The same query before and after an index, with both plans side by side. | — | Ex 2-3 |
| 5 · Break it | Hit `max_connections`, then fix it with PgBouncer. | Study 9 | Break it |
| 6 · Your numbers | Query time before/after; index size vs table size; insert slowdown. | — | Measure it |
| 7 · When not to index | Tiny tables, columns with few distinct values, write-heavy tables. | Study 3, 5 | Mini-project |

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
**Study path (free, in this order, ~8 h)**
1. [PostgreSQL docs, "Concurrency Control: Introduction"](https://www.postgresql.org/docs/current/mvcc.html). *(15 min)*
2. [PostgreSQL docs, "Transaction Isolation"](https://www.postgresql.org/docs/current/transaction-iso.html): the whole page (Read Committed, Repeatable Read, Serializable). *(1.5 h)*
3. [PostgreSQL docs, "Explicit Locking"](https://www.postgresql.org/docs/current/explicit-locking.html): the sections "Row-Level Locks" and "Deadlocks". *(45 min)*
4. [Hermitage](https://github.com/ept/hermitage): read the README, then run **every** test in [postgres.md](https://github.com/ept/hermitage/blob/master/postgres.md) in two `psql` windows. This doubles as your exercises 1-6. *(2-3 h)*
5. Watch: [CMU 15-445 (Fall 2024)](https://www.youtube.com/playlist?list=PLSE8ODhjZXjYDBpQnSymaectKjxCy6BYq), **Lecture #16 "Concurrency Control Theory"** and **#19 "Multi-Version Concurrency Control"**; add #17 "Two-Phase Locking" if you have time. *(3-4.5 h)*

*If something is still unclear:* [Jepsen, "Consistency Models"](https://jepsen.io/consistency): click Read Committed, Repeatable Read and Serializable on the map for precise definitions.

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Two buyers, one ticket left, both succeed. | — | Ex 1 |
| 2 · Mental model | A transaction is a bubble; the isolation level is how thin its walls are. | Study 1, 5 (#16) | — |
| 3a · ACID in practice | What each letter actually promises you. | Study 5 (#16) | — |
| 3b · Race conditions | Check-then-act and the lost update. | Study 2, 4 | Ex 1 |
| 3c · Isolation levels | Read committed, repeatable read, serializable, and the anomalies each allows (including write skew). | Study 2, 4 | Ex 5-6 |
| 3d · MVCC | Why readers don't block writers. | Study 1, 5 (#19) | — |
| 3e · Three fixes | `FOR UPDATE`, an atomic conditional `UPDATE`, optimistic locking. | Study 3 | Ex 2-4 |
| 3f · Deadlocks | How they form and how Postgres breaks them. | Study 3 (Deadlocks) | Break it |
| 4 · Live demo | Two `psql` windows side by side: the race, then the `FOR UPDATE` fix. | — | Ex 1-2 |
| 5 · Break it | A deliberate deadlock and Postgres's report. | Study 3 | Break it |
| 6 · Your numbers | 100 concurrent purchases: oversells with no fix, and throughput of each fix. | — | Measure it |
| 7 · Which fix when | Atomic update for counters, `FOR UPDATE` for multi-step logic, optimistic for low contention, `SERIALIZABLE` with retries. | Study 2, 3 | Mini-project |

**Video idea:** *How two people bought the last ticket at the same second*

### T6 · Backups and recovery, depth L2
**Why it exists:** hardware dies, people run `DELETE` without `WHERE`, and ransomware happens. A backup you've never restored is a hope, not a plan.
**Must know**
- Logical backups (`pg_dump`) vs physical backups.
- WAL and point-in-time recovery (PITR).
- The 3-2-1 rule, off-site copies, RPO and RTO.
- Why restores must be tested, and why a replica isn't a backup.

**Skip:** comparing backup tools in depth (pgBackRest, Barman and others): just name them.
**Study path (free, in this order, ~4 h)**
1. [PostgreSQL docs, "Backup and Restore"](https://www.postgresql.org/docs/current/backup.html) (chapter intro), then ["SQL Dump"](https://www.postgresql.org/docs/current/backup-dump.html). *(45 min)*
2. [PostgreSQL docs, "Continuous Archiving and Point-in-Time Recovery (PITR)"](https://www.postgresql.org/docs/current/continuous-archiving.html): the whole page. This is exercise 3. *(1.5 h)*
3. [GitLab, "Postmortem of database outage of January 31"](https://about.gitlab.com/blog/postmortem-of-database-outage-of-january-31/): the whole post. It's your video's story. *(30 min)*
4. [Google SRE book, "Data Integrity: What You Read Is What You Wrote"](https://sre.google/sre-book/data-integrity/): the first half, up to the case studies. *(45 min)*

*If something is still unclear:* [pgBackRest User Guide](https://pgbackrest.org/user-guide.html): skim the concepts section to see how a real backup tool organises full, differential and incremental backups.

**Practice**
- **Setup:** Postgres with the 5-million-row table from T4.
- **Exercises**
  1. Take a `pg_dump`, drop the table, and restore it.
  2. Take a physical backup with `pg_basebackup`.
  3. Turn on WAL archiving, insert rows, run a deliberate `DELETE` with no `WHERE`, then do PITR to the minute before it.
  4. Upload a backup to object storage (Garage from T10, or any free tier) to make it off-site.
  5. Write a tiny script that restores the latest backup into a fresh container and counts rows: an automated restore test.
- **Break it:** corrupt a backup file (truncate it) and try to restore it. That's why you test restores.
- **Measure it:** backup and restore times for the table. That restore time *is* your RTO, a strong on-camera moment.
- **Ready to record when:** you can explain RPO vs RTO with your own measured numbers.

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Someone runs `DELETE` without `WHERE`; the GitLab story. | Study 3 | — |
| 2 · Mental model | A snapshot plus a diary of every change (WAL): replay to any minute. | Study 2 | — |
| 3a · Logical vs physical | `pg_dump` vs `pg_basebackup`, and when to use each. | Study 1 | Ex 1-2 |
| 3b · WAL and PITR | Recovering to the minute before a mistake. | Study 2 | Ex 3 |
| 3c · 3-2-1 and off-site | Three copies, two media, one off-site. | Study 4 | Ex 4 |
| 3d · RPO and RTO | How much data you can lose, and how long recovery takes. | Study 4 | Measure it |
| 3e · Testing restores | Why untested backups fail; why a replica isn't a backup. | Study 3, 4 | Ex 5 |
| 4 · Live demo | A PITR to one minute before a deliberate `DELETE`. | — | Ex 3 |
| 5 · Break it | Restoring a corrupted backup file. | Study 3 | Break it |
| 6 · Your numbers | Your backup and restore times: your real RTO. | — | Measure it |
| 7 · Which approach when | `pg_dump` for small databases, physical + WAL for bigger ones, managed database backups. | Study 1, 2 | — |

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
**Study path (free, in this order, ~6 h)**
1. [Redis docs, "Data types"](https://redis.io/docs/latest/develop/data-types/): the overview, then the sub-pages for Strings, Lists, Sets, Hashes and Sorted sets. *(2 h)*
2. [Redis docs, "Transactions"](https://redis.io/docs/latest/develop/using-commands/transactions/) (`MULTI`/`EXEC`/`WATCH`). *(30 min)*
3. [Redis docs, "Scripting with Lua"](https://redis.io/docs/latest/develop/programmability/eval-intro/): this is exercise 6. *(45 min)*
4. [Redis docs, "Redis pipelining"](https://redis.io/docs/latest/develop/using-commands/pipelining/). *(20 min)*
5. [Redis docs, "Persistence"](https://redis.io/docs/latest/operate/oss_and_stack/management/persistence/): RDB vs AOF and their trade-offs. *(45 min)*
6. [Redis docs, "Key eviction"](https://redis.io/docs/latest/develop/reference/eviction/): `maxmemory` and the eviction policies. *(30 min)*
7. [Redis docs, `KEYS` command](https://redis.io/docs/latest/commands/keys/): read the warning and the time complexity before the "break it" exercise. *(5 min)*

*If something is still unclear:* [Redis command reference](https://redis.io/docs/latest/commands/): every command lists its time complexity, which helps you predict what's slow.

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Counters, sessions and leaderboards at thousands of operations a second are too much for the database. | — | — |
| 2 · Mental model | One huge in-memory dictionary run by a single worker who does one command at a time. | Study 1 | — |
| 3a · Why it's fast and atomic | Memory plus a single thread; time complexity per command; pipelining. | Study 4, 7 | Break it (`KEYS *`) |
| 3b · Data types | Strings, lists, sets, hashes, sorted sets, geo, and what each is for. | Study 1 | Ex 1, 3, 4, 5, 7 |
| 3c · Atomic multi-step logic | `MULTI`/`EXEC` and Lua scripts. | Study 2, 3 | Ex 2, 6 |
| 3d · Persistence | RDB vs AOF, and what you lose in a crash. | Study 5 | Break it (kill) |
| 3e · Memory and eviction | `maxmemory` and eviction policies. | Study 6 | Break it (`maxmemory`) |
| 4 · Live demo | `MONITOR` in one window; build a leaderboard and a lock in the other. | — | Ex 2-3 |
| 5 · Break it | `KEYS *` on a million keys freezing every client. | Study 7 | Break it |
| 6 · Your numbers | `redis-benchmark` ops/sec, latency, memory per million keys. | — | Measure it |
| 7 · When not to use it | As the only copy of data you can't lose; datasets bigger than RAM; complex queries. | Study 5, 6 | Mini-project |

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

**Study path (free, in this order, ~4 h)**
1. [AWS Builders' Library, "Caching challenges and strategies"](https://builder.aws.com/content/3K9TeXp6strMiJ4vhakpRD7BQWO/caching-challenges-and-strategies): the whole article. *(1 h)*
2. [Azure Architecture Center, "Cache-Aside pattern"](https://learn.microsoft.com/en-us/azure/architecture/patterns/cache-aside). *(20 min)*
3. [MDN, "HTTP caching"](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Caching): the whole page (`Cache-Control`, freshness, `ETag`/revalidation). *(1 h)*
4. [Cloudflare, "What is a CDN?"](https://www.cloudflare.com/learning/cdn/what-is-a-cdn/) and ["What is caching?"](https://www.cloudflare.com/learning/cdn/what-is-caching/). *(30 min)*

*If something is still unclear:* ["Scaling Memcache at Facebook" (NSDI 2013, PDF)](https://www.usenix.org/system/files/conference/nsdi13/nsdi13-final170_update.pdf): Section 3.2 "Reducing Load", the part on leases (stampede protection at huge scale).

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | The same slow query runs a thousand times a second. | — | — |
| 2 · Mental model | A notepad on your desk: check it before walking to the archive. | Study 2 | — |
| 3a · Patterns | Cache-aside, read-through, write-through, write-behind. | Study 1, 2 | Ex 1 |
| 3b · TTL and invalidation | Stale data and how to remove it on write. | Study 1, 2 | Ex 2 |
| 3c · Hit-ratio maths | What 90% vs 99% means for the database. | Study 1 | Ex 3 |
| 3d · Stampede | The thundering herd, and its fixes: locking, jitter, stale-while-revalidate. | Study 1 (+ "still unclear" link) | Break it |
| 3e · HTTP caching and CDN | `Cache-Control`, `ETag`, edge vs origin, purge. | Study 3, 4 | Ex 4-5 |
| 4 · Live demo | Latency with and without cache; then `cf-cache-status` going MISS → HIT. | — | Ex 1, 5 |
| 5 · Break it | A hot key expires under load and the database spikes. | Study 1 | Break it |
| 6 · Your numbers | Database queries per second at 0%, 90% and 99% hit ratio. | — | Measure it |
| 7 · What not to cache | Stock at checkout, fast-changing or personal data; measure before you cache. | Study 1 | — |

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

**Study path (free, in this order, ~4 h)**
1. [The Copenhagen Book, "Sessions"](https://copenhagen.pilcrowonpaper.com/sessions). *(45 min)*
2. [The Copenhagen Book, "Server-side tokens"](https://copenhagen.pilcrowonpaper.com/server-side-tokens). *(20 min)*
3. [The Copenhagen Book, "Password authentication"](https://copenhagen.pilcrowonpaper.com/password-authentication). *(30 min)*
4. [The Copenhagen Book, "Generating random values"](https://copenhagen.pilcrowonpaper.com/random-values): how to generate OTP codes safely. *(15 min)*
5. [The Copenhagen Book, "Multi-factor authentication (MFA)"](https://copenhagen.pilcrowonpaper.com/mfa): OTP and TOTP. *(30 min)*
6. [MDN, "Using HTTP cookies"](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides/Cookies): especially the section on security attributes. *(30 min)*
7. [OWASP, "Session Management Cheat Sheet"](https://cheatsheetseries.owasp.org/cheatsheets/Session_Management_Cheat_Sheet.html): the cookie attributes and session expiration sections. *(45 min)*
8. [jwt.io, "Introduction to JSON Web Tokens"](https://www.jwt.io/introduction). *(15 min)*

*If something is still unclear:* [OWASP Authentication Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html); [OWASP Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html).

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | HTTP forgets you between requests; add a second server and users get logged out. | — | Ex 1 |
| 2 · Mental model | A cloakroom ticket (session ID) vs a signed passport (JWT). | Study 1, 8 | — |
| 3a · Cookies | `HttpOnly`, `Secure`, `SameSite`, expiry. | Study 6, 7 | Ex 5 |
| 3b · Server-side sessions | Session IDs, and where sessions live with many servers. | Study 1, 2 | Ex 1-2 |
| 3c · JWT | What's inside, and the revocation trade-off. | Study 8, 1 | Ex 3 |
| 3d · Passwords | Why hashing is slow on purpose. | Study 3 | Measure it (bcrypt) |
| 3e · OTP | Generate, store hashed, expire, limit attempts. | Study 4, 5 | Ex 4 |
| 4 · Live demo | Random logouts with 2 servers, fixed by moving sessions to Redis. | — | Ex 1-2 |
| 5 · Break it | A stolen cookie works in another browser; brute-forcing an OTP without limits. | Study 7 | Break it |
| 6 · Your numbers | bcrypt cost 10 vs 12; JWT vs session-cookie size. | — | Measure it |
| 7 · Which to choose | Sessions vs JWT with reasons; use a proven library rather than writing your own. | Study 1, 8 | — |

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
**Study path (free, in this order, ~3 h)**
1. [Cloudflare, "What is object storage?"](https://www.cloudflare.com/learning/cloud/what-is-object-storage/). *(15 min)*
2. [AWS S3 docs, "Uploading objects with presigned URLs"](https://docs.aws.amazon.com/AmazonS3/latest/userguide/PresignedUrlUploadObject.html). *(30 min)*
3. [AWS S3 docs, "Uploading and copying objects using multipart upload"](https://docs.aws.amazon.com/AmazonS3/latest/userguide/mpuoverview.html). *(45 min)*
4. [tus, "Resumable upload protocol"](https://tus.io/protocols/resumable-upload): the "Core Protocol" section. *(45 min)*
5. [Garage, "Quick start"](https://garagehq.deuxfleurs.fr/documentation/quick-start/) to set up your lab, and ["S3 compatibility status"](https://garagehq.deuxfleurs.fr/documentation/reference-manual/s3-compatibility/) to see which S3 features it supports. *(30 min)*

**Practice**
- **Setup:** Garage in Docker (a free, open-source S3-compatible server), plus the AWS CLI. MinIO's free edition is no longer maintained, so use Garage instead; everything you learn applies to S3 too.
- **Exercises**
  1. Create a bucket and upload, list, download and delete objects.
  2. Generate a presigned PUT URL in the lab app and upload with `curl` or a small HTML page, so the file never touches your server.
  3. Upload a 1 GB file with multipart and look at the parts.
  4. Run a tus server (tusd) and resume an interrupted upload.
  5. After an upload finishes, have the client call `POST /uploads/:id/complete`. The app then puts a job on a list in Redis, and a worker makes a thumbnail with ImageMagick.
- **Break it:** upload a 1 GB file *through* the lab app instead and watch its memory and the timeouts. Cut the network halfway through a normal upload (`docker network disconnect`) and try to resume.
- **Measure it:** upload time through the app vs a presigned direct upload; resume time vs starting from scratch.
- **Ready to record when:** you can explain why presigned uploads take load off your servers, and how resumable uploads work in chunks.

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Uploads saved on one server's disk: other servers can't see them, and big uploads time out. | Study 1 | — |
| 2 · Mental model | A huge valet locker: you hand over whole objects and get them back by key. | Study 1 | — |
| 3a · File system vs object storage | Buckets, keys, objects you replace rather than edit, durability. | Study 1 | Ex 1 |
| 3b · Presigned URLs | Letting the client upload straight to storage. | Study 2 | Ex 2 |
| 3c · Multipart and resumable uploads | Chunks, parallel parts, resuming after a failure. | Study 3, 4 | Ex 3-4 |
| 3d · After the upload | A processing pipeline (thumbnails) and a CDN in front. | Study 1 | Ex 5 |
| 4 · Live demo | A presigned upload, then a resumed interrupted upload. | Study 5 (setup) | Ex 2, 4 |
| 5 · Break it | A 1 GB upload through the app's memory; a dropped network mid-upload. | Study 3 | Break it |
| 6 · Your numbers | Through the app vs direct upload; resume vs restart. | — | Measure it |
| 7 · When not to use it | Tiny files, or apps that truly need file-system behaviour. | Study 1 | — |

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
**Study path (free, in this order, ~8 h)**
1. [*Introduction to Information Retrieval*, Ch. 1 "Boolean retrieval" (PDF)](https://nlp.stanford.edu/IR-book/pdf/01bool.pdf): Sections 1.1-1.3 (building an inverted index and processing queries with it). *(1.5 h)*
2. [Same book, Ch. 2 "The term vocabulary and postings lists" (PDF)](https://nlp.stanford.edu/IR-book/pdf/02voc.pdf): Section 2.2 (tokenisation, stop words, normalisation, stemming). *(1 h)*
3. [Same book, Ch. 6 "Scoring, term weighting and the vector space model" (PDF)](https://nlp.stanford.edu/IR-book/pdf/06vect.pdf): Section 6.2 (tf-idf). *(45 min)*
4. [OpenSearch, "Getting started"](https://docs.opensearch.org/latest/getting-started/): every page in this section. *(1 h)*
5. [OpenSearch, "Text analysis"](https://docs.opensearch.org/latest/analyzers/). *(1 h)*
6. [OpenSearch, "Query DSL"](https://docs.opensearch.org/latest/query-dsl/): the full-text queries (`match`, `match_phrase`, `multi_match`), fuzziness, and bool filters. *(1 h)*
7. [Elasticsearch, "Edge n-gram tokenizer"](https://www.elastic.co/docs/reference/text-analysis/analysis-edgengram-tokenizer): the autocomplete example. OpenSearch has the same tokenizer. *(20 min)*
8. [PostgreSQL docs, "Full Text Search"](https://www.postgresql.org/docs/current/textsearch.html): the sections "Introduction" and "Tables and Indexes". *(45 min)*

*If something is still unclear:* [*Elasticsearch: The Definitive Guide*](https://www.elastic.co/guide/en/elasticsearch/guide/current/index.html): written for an old version, but its "Search in Depth" and relevance chapters still explain the concepts better than anything else.

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | `LIKE '%x%'` on millions of rows: slow, no ranking, no typo tolerance. | — | Measure it |
| 2 · Mental model | The index at the back of a book: word → list of pages. | Study 1 | — |
| 3a · Inverted index | Building it and answering queries with it. | Study 1 | Ex 1 |
| 3b · Text analysis | Tokenising, lowercasing, stemming, synonyms. | Study 2, 5 | Ex 3 |
| 3c · Relevance | tf-idf, and BM25 at headline level. | Study 3, 6 | Ex 2 |
| 3d · Filters, facets, fuzzy | Narrowing results and tolerating typos. | Study 6 | Ex 4 |
| 3e · Autocomplete | Edge n-grams. | Study 7 | Ex 5 |
| 3f · Keeping it in sync | Dual writes vs CDC; search is slightly behind the database. | Study 4 | Break it |
| 4 · Live demo | Your 40-line inverted index, then the same search in OpenSearch. | — | Ex 1-3 |
| 5 · Break it | A just-updated row that search can't find yet; a typo with no fuzziness. | Study 6 | Break it |
| 6 · Your numbers | `LIKE` vs Postgres full-text vs OpenSearch at 100K and 1M rows. | — | Measure it |
| 7 · When Postgres is enough | Postgres full-text search for small and medium needs. | Study 8 | Ex 6 |

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
**Study path (free, in this order, ~5 h)**
1. [RabbitMQ Tutorials](https://www.rabbitmq.com/tutorials): tutorials 1-6, in your language. They're also your exercises 1-4. *(3 h)*
2. [RabbitMQ, "Consumer Acknowledgements and Publisher Confirms"](https://www.rabbitmq.com/docs/confirms): the consumer-acknowledgement half. *(45 min)*
3. [RabbitMQ, "Consumer Prefetch"](https://www.rabbitmq.com/docs/consumer-prefetch). *(15 min)*
4. [RabbitMQ, "Dead Letter Exchanges"](https://www.rabbitmq.com/docs/dlx). *(30 min)*
5. [AWS SQS, "Visibility timeout"](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html): the same redelivery idea in a different product. *(20 min)*

*If something is still unclear:* [CloudAMQP, "RabbitMQ for beginners"](https://www.cloudamqp.com/blog/part1-rabbitmq-for-beginners-what-is-rabbitmq/).

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Users wait for slow work, and traffic peaks overwhelm the workers. | Study 1 | — |
| 2 · Mental model | Order tickets on a kitchen rail: cooks take the next one when they're free. | Study 1 | — |
| 3a · Producers and consumers | Why async smooths peaks and decouples services. | Study 1 (tutorials 1-2) | Ex 1 |
| 3b · Acks and redelivery | At-most-once vs at-least-once; why "exactly once" needs idempotency. | Study 2, 5 | Ex 2 |
| 3c · Fair dispatch | Prefetch and how work spreads. | Study 3 | Ex 3 |
| 3d · Pub/sub with exchanges | One message to many queues. | Study 1 (tutorials 3-5) | Ex 4 |
| 3e · Retries and dead letters | Backoff, DLQs, poison messages. | Study 4 | Ex 5 |
| 3f · Priority and backpressure | Urgent work first; what happens when consumers fall behind. | Study 1 | Ex 6 |
| 4 · Live demo | The management UI while you kill a consumer mid-message and watch redelivery. | — | Ex 2 |
| 5 · Break it | A poison message looping forever, fixed with a DLQ. | Study 4 | Break it |
| 6 · Your numbers | Throughput with 1, 3, 10 consumers; drain time for 100K messages. | — | Measure it |
| 7 · When not to use it | When the user needs the answer now; when you need replay (that's T13). | Study 1 | Mini-project |

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
**Study path (free, in this order, ~5 h)**
1. [Apache Kafka, "Introduction"](https://kafka.apache.org/intro/). *(30 min)*
2. [Confluent Developer, "Apache Kafka 101" course](https://developer.confluent.io/courses/apache-kafka/events/): all modules (short videos plus text). *(2 h)*
3. [Jay Kreps, "The Log: What every software engineer should know about real-time data's unifying abstraction"](https://engineering.linkedin.com/distributed-systems/log-what-every-software-engineer-should-know-about-real-time-datas-unifying): Part One "What Is a Log?" and Part Two "Data Integration". *(1.5 h)*
4. [Apache Kafka docs, "Design"](https://kafka.apache.org/documentation/#design): "Persistence", "The Consumer" and "Message Delivery Semantics". *(1 h)*

*If something is still unclear:* if you use Redpanda for the lab, the [Redpanda docs](https://docs.redpanda.com/home/) quick start.

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Several systems need the same events, but a queue deletes each message once it's handled. | Study 3 | — |
| 2 · Mental model | An append-only notebook where every reader keeps their own bookmark. | Study 3 | — |
| 3a · Log vs queue | Retention and replay. | Study 1, 3 | — |
| 3b · Topics, partitions, keys | Ordering within a partition. | Study 1, 2 | Ex 1 |
| 3c · Consumer groups and offsets | Sharing work, and independent readers. | Study 2, 4 | Ex 2-4 |
| 3d · Delivery semantics | At-least-once, retention, committing offsets. | Study 4 | — |
| 3e · Other pub/sub options | Redis pub/sub (fire-and-forget) vs a durable log. | — | Ex 5 |
| 4 · Live demo | Two consumer groups reading the same topic, then a replay from offset 0. | — | Ex 3-4 |
| 5 · Break it | A rebalance when a consumer dies; lost ordering without keys. | Study 2 | Break it |
| 6 · Your numbers | Messages/sec on your laptop; lag while a consumer is paused. | — | Measure it |
| 7 · When not to use it | Small apps and simple task queues (use RabbitMQ, T12). | Study 1 | — |

**Video idea:** *Kafka vs RabbitMQ: do you need a queue or a log?*

### T14 · Background jobs, cron and distributed locks, depth L2
**Why it exists:** apps need scheduled work (reports, cleanups, reminders). Once you have more than one server, "run this once" gets surprisingly hard.
**Must know**
- Cron syntax and its limits; worker pools.
- The "job ran twice" problem when several servers run the same cron.
- Leases with TTLs, leader election at headline level, fencing tokens.
- Delayed and scheduled jobs with Redis sorted sets.
- Making jobs idempotent.

**Study path (free, in this order, ~2 h)**
1. [Martin Kleppmann, "How to do distributed locking"](https://martin.kleppmann.com/2016/02/08/how-to-do-distributed-locking.html): the whole post. Fencing tokens are the key idea. *(45 min)*
2. [Redis docs, "Distributed Locks with Redis"](https://redis.io/docs/latest/develop/clients/patterns/distributed-locks/): the single-instance section. *(30 min)*
3. [crontab.guru](https://crontab.guru/): play with cron expressions. *(10 min)*
4. [PostgreSQL docs, `SELECT`, "The Locking Clause"](https://www.postgresql.org/docs/current/sql-select.html#SQL-FOR-UPDATE-SHARE): `FOR UPDATE SKIP LOCKED`, which is how many job queues run on Postgres. *(20 min)*

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | A report gets generated three times because three servers run the same cron. | — | Ex 1 |
| 2 · Mental model | A talking stick with a timer: only the holder may speak, until time runs out. | Study 2 | — |
| 3a · Cron and workers | Cron syntax and its limits. | Study 3 | Ex 1 |
| 3b · Leases | A lock with a TTL in Redis. | Study 2 | Ex 2 |
| 3c · Why locks fail | Process pauses and clocks; fencing tokens. | Study 1 | Break it |
| 3d · Delayed jobs | Sorted sets; `SKIP LOCKED` job queues in Postgres. | Study 4 | Ex 3 |
| 3e · Idempotent jobs | Making a double run harmless. | — | Ex 4 |
| 4 · Live demo | Three runs, fixed with a lease. | — | Ex 1-2 |
| 5 · Break it | Pause the lease holder past its TTL so both run; fix it with a fencing token. | Study 1 | Break it |
| 6 · Your numbers | How late delayed jobs run at different polling intervals. | — | Measure it |
| 7 · When to use a real scheduler | When a dedicated job system or managed scheduler is worth it (L1). | Study 1 | — |

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
**Study path (free, in this order, ~4.5 h)**
1. [Stripe, "Designing robust and predictable APIs with idempotency"](https://stripe.com/blog/idempotency). *(30 min)*
2. [Brandur Leach, "Implementing Stripe-like Idempotency Keys in Postgres"](https://brandur.org/idempotency-keys): the whole post. It's exercise 3 in full detail. *(1.5 h)*
3. [Stripe API docs, "Idempotent requests"](https://docs.stripe.com/api/idempotent_requests). *(10 min)*
4. [Stripe docs, "Webhooks"](https://docs.stripe.com/webhooks): signature verification, retries, duplicate events and event ordering. *(1 h)*
5. A local gateway's flow: the [bKash developer portal](https://developer.bka.sh/) (Checkout / Tokenized checkout) or the [SSLCommerz developer docs](https://developer.sslcommerz.com/) (payment flow and IPN). *(1 h)*

*If something is still unclear:* [IETF draft, "The Idempotency-Key HTTP Header Field"](https://datatracker.ietf.org/doc/draft-ietf-httpapi-idempotency-key-header/).

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | A timeout, a retry, and the customer is charged twice. | Study 1 | — |
| 2 · Mental model | A receipt number on every request: the cashier checks it before charging again. | Study 1 | — |
| 3a · How a payment flow works | Create → redirect → pay → callback/webhook → verify. | Study 5, 4 | Ex 1 |
| 3b · "Did it go through?" | Why a timeout tells you nothing about the outcome. | Study 1 | — |
| 3c · Idempotency keys | Store, lock, replay the saved response, expiry. | Study 1, 2, 3 | Ex 3-4 |
| 3d · State machine | Legal and illegal payment transitions. | Study 2 | Ex 5 |
| 3e · Webhooks | Signatures, duplicates, out-of-order delivery. | Study 4 | Ex 2 |
| 3f · Reconciliation | Comparing your records with the gateway's. | Study 2 | Mini-project |
| 4 · Live demo | The same request sent twice at once: one charge, the same response. | — | Ex 4 |
| 5 · Break it | Retry without a key → a double charge in the sandbox dashboard; replayed webhooks. | Study 4 | Break it |
| 6 · Your numbers | Duplicates from 100 retried requests, with and without keys. | — | Measure it |
| 7 · Beyond payments | Anywhere a retry could repeat a side effect (emails, orders, SMS). | Study 1 | — |

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

**Study path (free, in this order, ~5 h)**
1. Watch: [Kleppmann, *Distributed Systems*](https://www.youtube.com/playlist?list=PLeKd45zvjcDFUEv_ohr_HdUFe97RItdiB), **lectures 5.1 "Replication"** and **5.2 "Quorums"**. The [lecture notes (PDF)](https://www.cl.cam.ac.uk/teaching/2122/ConcDisSys/dist-sys-notes.pdf) cover the same material in Section 5. *(1 h)*
2. [PostgreSQL docs, "High Availability, Load Balancing, and Replication"](https://www.postgresql.org/docs/current/high-availability.html): the intro and the comparison of solutions. *(45 min)*
3. [PostgreSQL docs, "Log-Shipping Standby Servers"](https://www.postgresql.org/docs/current/warm-standby.html): "Streaming Replication", "Synchronous Replication" and the failover section. This is your lab guide. *(1.5 h)*
4. [GitHub, "October 21 post-incident analysis"](https://github.blog/news-insights/company-news/oct21-post-incident-analysis/): a real failover across data centres going wrong. *(45 min)*

*If something is still unclear:* [Patroni docs](https://patroni.readthedocs.io/en/latest/), to see how automatic failover is done in practice.

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | "I just saved it, why isn't it there?" Or: the only database machine dies. | Study 1 | — |
| 2 · Mental model | The primary writes in a diary; replicas copy each page a moment later. | Study 1 | — |
| 3a · Primary and replicas | Read scaling; replicas refuse writes. | Study 1, 2 | Ex 1-2 |
| 3b · Sync vs async | What each one risks losing, and what it costs. | Study 3 | Ex 5 |
| 3c · Lag anomalies | Read-your-writes and monotonic reads, and how to fix them. | Study 1 | Ex 3-4 |
| 3d · Failover and split brain | Promoting a replica; two primaries. | Study 3, 4 | Break it |
| 4 · Live demo | A delayed replica that "loses" the row you just wrote. | — | Ex 3 |
| 5 · Break it | Manual promotion, then the old primary comes back: split brain. | Study 4 | Break it |
| 6 · Your numbers | Lag under load; write latency async vs sync. | — | Measure it |
| 7 · When to use it | When replicas help; why a replica isn't a backup (callback to T6). | Study 2 | — |

**Video idea:** *"I just saved it, why isn't it there?" Replication lag explained*

### T17 · Sharding and consistent hashing, depth L3
**Why it exists:** at some size, one machine can't hold all the data or handle all the writes. You split the data across machines, and everything gets harder.
**Must know**
- Vertical vs horizontal partitioning; partitioning within one database vs across machines.
- Choosing the shard key, range vs hash sharding, hot spots.
- Cross-shard queries and joins, and why resharding is painful.
- Consistent hashing with virtual nodes; directory-based sharding.

**Study path (free, in this order, ~4.5 h)**
1. [PostgreSQL docs, "Table Partitioning"](https://www.postgresql.org/docs/current/ddl-partitioning.html): "Overview", "Declarative Partitioning" and "Partition Pruning". This is exercise 1. *(1 h)*
2. [The Amazon Dynamo paper (PDF)](https://www.allthingsdistributed.com/files/amazon-dynamo-sosp2007.pdf): Section 4.2 "Partitioning Algorithm" and 4.3 "Replication" (consistent hashing with virtual nodes). *(1 h)*
3. [Tom White, "Consistent Hashing"](https://tom-e-white.com/2007/11/consistent-hashing.html): short, with diagrams and code. *(20 min)*
4. [Notion, "Herding elephants: Lessons learned from sharding Postgres at Notion"](https://www.notion.com/blog/sharding-postgres-at-notion). *(45 min)*
5. [Instagram Engineering, "Sharding & IDs at Instagram"](https://instagram-engineering.com/sharding-ids-at-instagram-1cf5a71e5a5c): the sharding part. *(20 min)*

*If something is still unclear:* [Citus docs](https://docs.citusdata.com/) for the optional lab; the [MIT 6.5840 schedule](https://pdos.csail.mit.edu/6.824/schedule.html) for its lectures on sharded key/value systems.

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | One machine can't hold all the data or handle all the writes. | Study 4 | — |
| 2 · Mental model | Splitting one phone book into volumes A-M and N-Z. | Study 1 | — |
| 3a · Partitioning vs sharding | Splitting within one database vs across machines. | Study 1 | Ex 1 |
| 3b · Choosing the shard key | Range vs hash; what a good key looks like. | Study 4, 5 | Ex 2 |
| 3c · Hot spots | When one shard gets most of the traffic. | Study 4 | Break it |
| 3d · Consistent hashing | The ring and virtual nodes. | Study 2, 3 | Ex 3-4 |
| 3e · The pain | Cross-shard queries, joins and resharding. | Study 4 | Ex 5 |
| 4 · Live demo | Add a node: keys moved with `% N` vs with the ring. | — | Ex 4 |
| 5 · Break it | A skewed shard key overloading one shard. | Study 4 | Break it |
| 6 · Your numbers | Keys moved and spread for 10, 100, 1,000 virtual nodes. | — | Measure it |
| 7 · When not to shard | Bigger machine, replicas and partitioning come first. | Study 4 | — |

**Video idea:** *When one database isn't enough: sharding explained*

### T18 · Distributed ID generation, depth L2
**Why it exists:** auto-increment needs one central counter. Across many machines you need IDs that are unique without coordination, and ideally sortable.
**Must know**
- Why auto-increment breaks across shards.
- UUID v4 vs v7, and index locality.
- The Snowflake layout (timestamp / machine / sequence) and clock skew.
- Ticket servers and ID ranges.
- Base62 encoding for short codes, and why guessable IDs are a security risk.

**Study path (free, in this order, ~1.5 h)**
1. [Twitter, "Announcing Snowflake" (2010)](https://blog.x.com/engineering/en_us/a/2010/announcing-snowflake). *(10 min)*
2. [Snowflake README, 2010 version](https://github.com/twitter-archive/snowflake/tree/snowflake-2010): the ID layout (timestamp, worker, sequence). *(15 min)*
3. [Instagram Engineering, "Sharding & IDs at Instagram"](https://instagram-engineering.com/sharding-ids-at-instagram-1cf5a71e5a5c): the ID part (41 bits time, 13 bits shard, 10 bits sequence). *(20 min)*
4. [RFC 9562, *UUIDs*](https://www.rfc-editor.org/rfc/rfc9562): Section 5.4 (UUIDv4) and Section 5.7 (UUIDv7). *(45 min)*

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Auto-increment on several shards gives duplicate IDs. | Study 3 | — |
| 2 · Mental model | An ID = timestamp + who made it + a counter. | Study 1, 2 | — |
| 3a · Why auto-increment breaks | One counter can't live on many machines. | Study 3 | — |
| 3b · UUIDs | v4 vs v7, and index locality. | Study 4 | Ex 4 |
| 3c · Snowflake | The bit layout. | Study 1, 2 | Ex 1-2 |
| 3d · Instagram's variant | Generating IDs inside Postgres. | Study 3 | — |
| 3e · Clocks | Clock skew and going backwards. | Study 2 | Break it |
| 3f · Short codes | Base62, and why guessable IDs are a risk. | — | Ex 3 |
| 4 · Live demo | Generate an ID and decode it back into time, machine and sequence. | — | Ex 1 |
| 5 · Break it | Move the clock backwards and get duplicates. | Study 2 | Break it |
| 6 · Your numbers | Insert speed and index size, UUIDv4 vs sequential. | — | Measure it |
| 7 · Which scheme when | UUIDv7 vs Snowflake vs in-database generation. | Study 3, 4 | — |

**Video idea:** *How big apps generate billions of unique IDs without one central counter*

### T19 · Consistency models and CAP, depth L2 (concept-heavy)
**Why it exists:** once data lives on several machines, "what does a read return?" stops having one answer. Consistency models name the possible answers.
**Must know**
- Strong vs eventual consistency; read-your-writes, monotonic reads, linearizability at headline level.
- **CAP stated correctly:** the choice only applies during a network partition. PACELC.
- Quorums (N, R, W).
- Choosing consistency per feature: a bank balance needs strong consistency; a view counter can be eventual.

**Study path (free, in this order, ~5.5 h)**
1. Watch: [Kleppmann, *Distributed Systems*](https://www.youtube.com/playlist?list=PLeKd45zvjcDFUEv_ohr_HdUFe97RItdiB), **lectures 7.2 "Linearizability"** and **7.3 "Eventual consistency"** (and 5.2 "Quorums" if you skipped it in T16). Notes: Section 7 of the [PDF](https://www.cl.cam.ac.uk/teaching/2122/ConcDisSys/dist-sys-notes.pdf). *(1.5 h)*
2. [Martin Kleppmann, "Please stop calling databases CP or AP"](https://martin.kleppmann.com/2015/05/11/please-stop-calling-databases-cp-or-ap.html). *(45 min)*
3. [Jepsen, "Consistency Models"](https://jepsen.io/consistency): the map, then the pages for Linearizable, Sequential, Causal, Read Your Writes and Monotonic Reads. *(1 h)*
4. [Eric Brewer, "CAP Twelve Years Later: How the 'Rules' Have Changed"](https://www.infoq.com/articles/cap-twelve-years-later-how-the-rules-have-changed/). *(45 min)*
5. [Werner Vogels, "Eventually Consistent – Revisited"](https://www.allthingsdistributed.com/2008/12/eventually_consistent.html). *(30 min)*
6. For the lab: [Cassandra docs, "Dynamo"](https://cassandra.apache.org/doc/latest/cassandra/architecture/dynamo.html): the section on tunable consistency. *(30 min)*

**Practice**
- **Setup:** paper, plus Docker for a small multi-node database (a 3-node Cassandra or ScyllaDB cluster).
- **Exercises**
  1. Draw timeline diagrams (two clients, two replicas) for a stale read, a non-monotonic read and a lost update, until you can draw each from memory.
  2. In the 3-node cluster, write and read at consistency `ONE` vs `QUORUM` vs `ALL`.
  3. Work out quorum overlap (R + W > N) for N = 3 and N = 5 on paper.
- **Break it:** cut one node off the network (`docker network disconnect`). At `ALL`, writes fail (choosing consistency); at `ONE`, they succeed but reads may be stale (choosing availability). That's CAP, live.
- **Measure it:** read and write latency at `ONE` vs `QUORUM` (PACELC's "else latency").
- **Ready to record when:** you can state CAP precisely and correct the most common misunderstanding of it.

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Two replicas give two different answers to the same question. | Study 1 | — |
| 2 · Mental model | Two copies of a notebook in two cities, with a bad phone line between them. | Study 1 | — |
| 3a · Strong vs eventual | What each one promises. | Study 1, 5 | — |
| 3b · Client guarantees | Read-your-writes, monotonic reads. | Study 3 | Ex 1 |
| 3c · Linearizability | Headline level. | Study 1, 3 | Ex 1 |
| 3d · CAP, stated correctly | It's only a choice during a partition. | Study 2, 4 | Break it |
| 3e · PACELC | The everyday trade-off is latency. | Study 2 | Measure it |
| 3f · Quorums | N, R, W and why R + W > N overlaps. | Study 6 | Ex 2-3 |
| 4 · Live demo | Cut a node off the network and write at `ALL` vs `ONE`. | — | Break it |
| 5 · Break it | The partition itself (same as the demo). | Study 2 | Break it |
| 6 · Your numbers | Latency at `ONE` vs `QUORUM`. | — | Measure it |
| 7 · Choosing per feature | Bank balance strong, view counter eventual. | Study 4, 5 | — |

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

**Study path (free, in this order, ~2.5 h)**
1. [Stripe, "Scaling your API with rate limiters"](https://stripe.com/blog/rate-limiters). *(30 min)*
2. [Cloudflare, "How we built rate limiting capable of scaling to millions of domains"](https://blog.cloudflare.com/counting-things-a-lot-of-different-things/): the sliding-window counter, explained with numbers. *(45 min)*
3. [Azure Architecture Center, "Rate Limiting pattern"](https://learn.microsoft.com/en-us/azure/architecture/patterns/rate-limiting-pattern). *(20 min)*
4. [Nginx, `ngx_http_limit_req_module`](https://nginx.org/en/docs/http/ngx_http_limit_req_module.html): `limit_req_zone`, `burst`, `nodelay`. This is exercise 5. *(20 min)*

*If something is still unclear:* [IETF draft, "RateLimit header fields for HTTP"](https://datatracker.ietf.org/doc/draft-ietf-httpapi-ratelimit-headers/), for exercise 6.

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | A bot hammers your login or OTP endpoint. | Study 1 | — |
| 2 · Mental model | A bucket of tokens that refills at a steady rate. | Study 1 | — |
| 3a · Fixed window | How it works, and the burst at the window edge. | Study 1 | Ex 1-2 |
| 3b · Sliding windows | The log and the counter. | Study 2 | Ex 3 |
| 3c · Token and leaky buckets | Smooth limits with bursts allowed. | Study 1, 3 | Ex 4 |
| 3d · Where and what to limit | Edge, gateway or app; IP, user, key, phone number. | Study 1, 3 | — |
| 3e · Distributed limiting | Redis-backed limits and their race conditions. | Study 1 | Break it |
| 3f · Telling the client | `429`, `Retry-After`, limit headers. | Study 4 (+ IETF link) | Ex 6 |
| 4 · Live demo | A k6 burst against your token bucket, with the allowed/rejected chart. | — | Ex 4-5 |
| 5 · Break it | GET-then-SET leaking past the limit. | Study 1 | Break it |
| 6 · Your numbers | The four algorithms' charts side by side. | — | Measure it |
| 7 · When limits aren't enough | CAPTCHA and making abuse expensive. | Study 3 | — |

**Video idea:** *Four ways to rate limit an API (and which one to pick)*

### T21 · Timeouts, retries, circuit breakers and bulkheads, depth L3
**Why it exists:** in a system of services, one slow dependency can take down everything that calls it. These patterns contain the damage.
**Must know**
- Why a missing timeout cascades: thread and connection pools fill up.
- Choosing timeouts from latency percentiles.
- Exponential backoff with jitter, retry budgets, retry storms.
- Circuit breaker states (closed / open / half-open), bulkheads, fallbacks, load shedding.

**Study path (free, in this order, ~4 h)**
1. [AWS Builders' Library, "Timeouts, retries, and backoff with jitter"](https://builder.aws.com/content/3EumjoZascWd1oZiEgL8ORlv3qE/timeouts-retries-and-backoff-with-jitter). *(45 min)*
2. [Marc Brooker, "Exponential Backoff And Jitter"](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/): the simulation charts are great for your video. *(20 min)*
3. Azure Architecture Center: ["Retry pattern"](https://learn.microsoft.com/en-us/azure/architecture/patterns/retry), ["Circuit Breaker pattern"](https://learn.microsoft.com/en-us/azure/architecture/patterns/circuit-breaker) and ["Bulkhead pattern"](https://learn.microsoft.com/en-us/azure/architecture/patterns/bulkhead). *(1.5 h)*
4. [Martin Fowler, "CircuitBreaker"](https://martinfowler.com/bliki/CircuitBreaker.html). *(20 min)*
5. [AWS Builders' Library, "Using load shedding to avoid overload"](https://builder.aws.com/content/3Eun1EEyX6p2e3VYNyRLSJzLuMV/using-load-shedding-to-avoid-overload). *(45 min)*
6. For the lab: the [Toxiproxy README](https://github.com/Shopify/toxiproxy). *(15 min)*

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | One slow dependency takes down every endpoint. | Study 1 | Ex 1 |
| 2 · Mental model | Electrical fuses and a ship's watertight compartments. | Study 3, 4 | — |
| 3a · How cascades happen | Pools fill up while waiting. | Study 1 | Ex 1 |
| 3b · Timeouts | Choosing them from latency percentiles. | Study 1 | Ex 2 |
| 3c · Retries | Backoff, jitter, retry budgets. | Study 1, 2, 3 (Retry) | Ex 3 |
| 3d · Circuit breakers | Closed, open, half-open. | Study 3, 4 | Ex 4 |
| 3e · Bulkheads | Isolating pools per dependency. | Study 3 (Bulkhead) | Ex 5 |
| 3f · Load shedding | Rejecting work early to survive. | Study 5 | — |
| 4 · Live demo | Toxiproxy latency with no protection, then with a breaker. | Study 6 (setup) | Ex 1, 4 |
| 5 · Break it | Retries at three layers: a retry storm. | Study 1 | Break it |
| 6 · Your numbers | Success rate and p99 across the five setups, in one chart. | — | Measure it |
| 7 · The cost of each pattern | Complexity, tuning, and failing fast vs waiting. | Study 3 | — |

**Video idea:** *How one slow API can take down an entire app*

### T22 · Observability: logs, metrics, traces, SLOs, depth L2
**Why it exists:** you can't fix what you can't see. Observability is how you find out *what* is broken and *where*, quickly.
**Must know**
- Logs vs metrics vs traces, and when to use each.
- The RED method (rate, errors, duration) and the USE method (utilisation, saturation, errors).
- Percentiles (p50, p95, p99) vs averages.
- SLIs, SLOs and error budgets; alerting on symptoms, not causes.
- Distributed tracing at headline level (OpenTelemetry).

**Study path (free, in this order, ~6 h)**
1. [Google SRE book, Ch. 4 "Service Level Objectives"](https://sre.google/sre-book/service-level-objectives/). *(45 min)*
2. [Google SRE book, Ch. 6 "Monitoring Distributed Systems"](https://sre.google/sre-book/monitoring-distributed-systems/): especially "The Four Golden Signals". *(45 min)*
3. [SRE Workbook, "Implementing SLOs"](https://sre.google/workbook/implementing-slos/) and ["Alerting on SLOs"](https://sre.google/workbook/alerting-on-slos/). *(1.5 h)*
4. [Grafana, "The RED Method: How to instrument your services"](https://grafana.com/blog/the-red-method-how-to-instrument-your-services/) and [Brendan Gregg, "The USE Method"](https://www.brendangregg.com/usemethod.html). *(40 min)*
5. Watch: [Gil Tene, "How NOT to Measure Latency" (Strange Loop)](https://www.youtube.com/watch?v=lJ8ydIuPFeU): the source of the "averages lie" lesson. *(45 min)*
6. [Prometheus, "Getting started"](https://prometheus.io/docs/prometheus/latest/getting_started/) and ["Histograms and summaries"](https://prometheus.io/docs/practices/histograms/). *(1.25 h)*

*If something is still unclear:* [OpenTelemetry, "Getting started"](https://opentelemetry.io/docs/getting-started/) for the optional tracing exercise.

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Users complain while the dashboard says everything's fine. | Study 5 | — |
| 2 · Mental model | Logs are a diary, metrics are dashboard gauges, traces are a GPS route. | Study 2 | — |
| 3a · Logs, metrics, traces | What each is for. | Study 2 | — |
| 3b · What to measure | Golden signals, RED, USE. | Study 2, 4 | Ex 1-2 |
| 3c · Percentiles | Why averages hide pain. | Study 5, 6 | Ex 3 |
| 3d · SLIs, SLOs, error budgets | Defining "good enough" and spending the budget. | Study 1, 3 | Ex 4 |
| 3e · Alerting | Alert on symptoms, not causes. | Study 2, 3 | Ex 5 |
| 3f · Tracing | Finding the slow hop across services (headline). | "Still unclear" link | Ex 6 |
| 4 · Live demo | The RED dashboard during a load test. | — | Ex 2 |
| 5 · Break it | 2% slow requests: the average barely moves, p99 jumps. | Study 5 | Break it |
| 6 · Your numbers | Average vs p50 vs p99 under that load. | — | Measure it |
| 7 · What not to alert on | Causes, noise, and anything nobody will act on. | Study 2 | — |

**Video idea:** *Average latency lies: p99 and SLOs explained*

### T23 · Containers and safe deployments, depth L2
**Why it exists:** most outages happen right after a change. Safe deployment practices let teams ship often without breaking things.
**Must know**
- Docker images, layers and containers; what Kubernetes does (L1-L2).
- Rolling, blue-green and canary deploys; readiness vs liveness probes.
- Feature flags.
- **Expand/contract database migrations**, such as renaming a column without downtime.
- Rollbacks, and the stages of a CI/CD pipeline.

**Study path (free, in this order, ~6.5 h)**
1. [Docker docs, "Get started"](https://docs.docker.com/get-started/). *(1.5 h)*
2. [Docker docs, "Multi-stage builds"](https://docs.docker.com/build/building/multi-stage/): this is exercise 1. *(20 min)*
3. [kind, "Quick Start"](https://kind.sigs.k8s.io/docs/user/quick-start/): to set up your local cluster. *(20 min)*
4. [Kubernetes, "Learn Kubernetes Basics"](https://kubernetes.io/docs/tutorials/kubernetes-basics/): all modules, ending with ["Performing a Rolling Update"](https://kubernetes.io/docs/tutorials/kubernetes-basics/update/update-intro/). *(2 h)*
5. [Kubernetes, "Configure Liveness, Readiness and Startup Probes"](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/). *(30 min)*
6. Martin Fowler: ["BlueGreenDeployment"](https://martinfowler.com/bliki/BlueGreenDeployment.html), ["CanaryRelease"](https://martinfowler.com/bliki/CanaryRelease.html), ["ParallelChange"](https://martinfowler.com/bliki/ParallelChange.html) (expand/contract). *(45 min)*
7. [Pete Hodgson, "Feature Toggles (aka Feature Flags)"](https://martinfowler.com/articles/feature-toggles.html): up to and including "Categories of toggles". *(45 min)*
8. [Stripe, "Online migrations at scale"](https://stripe.com/blog/online-migrations): a real expand/contract migration. *(30 min)*

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Most outages happen right after a change. | Study 8 | — |
| 2 · Mental model | Changing tyres on a moving car: swap one at a time, keep the old one close. | Study 6 | — |
| 3a · Containers and images | Layers and small images. | Study 1, 2 | Ex 1 |
| 3b · What Kubernetes does | Headline level: desired state, replicas, services. | Study 3, 4 | — |
| 3c · Rolling updates and probes | Readiness vs liveness. | Study 4, 5 | Ex 2-3 |
| 3d · Blue-green and canary | Switching all at once vs a small slice first. | Study 6 | Ex 4-5 |
| 3e · Feature flags | Turning features on without deploying. | Study 7 | Ex 6 |
| 3f · Database migrations | Expand/contract (parallel change). | Study 6, 8 | Break it |
| 3g · Rollback | Getting back fast. | Study 4 | Ex 7 |
| 4 · Live demo | A naive column rename vs expand/contract under load. | — | Break it |
| 5 · Break it | The naive rename (same as the demo). | Study 8 | Break it |
| 6 · Your numbers | Errors with/without readiness probes; rollback time. | — | Measure it |
| 7 · Which strategy when | Rolling vs blue-green vs canary vs flags. | Study 6, 7 | — |

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

**Study path (free, in this order, ~5 h)**
1. Martin Fowler: ["MonolithFirst"](https://martinfowler.com/bliki/MonolithFirst.html) and ["MicroservicePrerequisites"](https://martinfowler.com/bliki/MicroservicePrerequisites.html). *(20 min)*
2. [James Lewis and Martin Fowler, "Microservices"](https://martinfowler.com/articles/microservices.html). *(1 h)*
3. microservices.io: ["Database per service"](https://microservices.io/patterns/data/database-per-service.html), ["Saga"](https://microservices.io/patterns/data/saga.html) and ["Transactional outbox"](https://microservices.io/patterns/data/transactional-outbox.html). *(1 h)*
4. Watch: [Caitie McCaffrey, "Applying the Saga Pattern" (GOTO 2015)](https://www.youtube.com/watch?v=xDuwrtwYHu8). *(45 min)*
5. [Debezium, "Tutorial"](https://debezium.io/documentation/reference/stable/tutorial.html), then ["Outbox Event Router"](https://debezium.io/documentation/reference/stable/transformations/outbox-event-router.html): your lab guide for exercise 5. *(2 h)*

*If something is still unclear:* [Azure Architecture Center, "Saga distributed transactions pattern"](https://learn.microsoft.com/en-us/azure/architecture/patterns/saga).

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | The order is saved but its event is never sent. | Study 3 | — |
| 2 · Mental model | No single transaction any more: a checklist of steps, each with an undo. | Study 4 | — |
| 3a · When to split | Monolith first; the prerequisites. | Study 1, 2 | — |
| 3b · Sync vs async | REST/gRPC calls vs events. | Study 2 | Ex 1-2 |
| 3c · Why not 2PC | The cost of distributed transactions. | Study 4 | — |
| 3d · Sagas | Compensations; orchestration vs choreography. | Study 3, 4 | Ex 3 |
| 3e · Dual writes and the outbox | Why "save, then publish" fails; the outbox fix. | Study 3 | Ex 4 |
| 3f · CDC | Debezium reading the outbox. | Study 5 | Ex 5 |
| 4 · Live demo | Kill the service between save and publish; then the same kill with the outbox. | — | Break it, Ex 4 |
| 5 · Break it | The dual-write failure (same as the demo). | Study 3 | Break it |
| 6 · Your numbers | Event delay: polling relay vs Debezium. | — | Measure it |
| 7 · When not to use microservices | Small teams and young products. | Study 1 | — |

**Video idea:** *Saved to the database but the event never sent: the dual-write problem*

### T25 · Real-time: polling, SSE, WebSockets, depth L2
**Why it exists:** chat, live tracking, notifications and live counters need the server to push data to the client, which plain HTTP request/response doesn't do.
**Must know**
- Short polling, long polling, SSE and WebSockets, with the trade-offs of each.
- What each open connection costs (memory, file descriptors), and load balancing long-lived connections.
- Scaling across servers with a pub/sub backplane.
- Heartbeats, reconnects with backoff, delivery acks and receipts, flaky mobile networks.

**Study path (free, in this order, ~3 h)**
1. javascript.info: ["Long polling"](https://javascript.info/long-polling), ["WebSocket"](https://javascript.info/websocket) and ["Server Sent Events"](https://javascript.info/server-sent-events). Short, with working code. *(1.5 h)*
2. MDN: ["The WebSocket API"](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API) and ["Using server-sent events"](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events). *(45 min)*
3. [Nginx, "WebSocket proxying"](https://nginx.org/en/docs/http/websocket.html). *(10 min)*
4. [Discord, "How Discord Scaled Elixir to 5,000,000 Concurrent Users"](https://discord.com/blog/how-discord-scaled-elixir-to-5-000-000-concurrent-users). *(30 min)*

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Live updates without the user pressing refresh. | Study 1 | — |
| 2 · Mental model | Asking "are we there yet?" over and over vs keeping a phone line open. | Study 1 | — |
| 3a · Polling and long polling | How they work and what they cost. | Study 1 | Ex 1 |
| 3b · Server-sent events | One-way server push over HTTP. | Study 1, 2 | Ex 1 |
| 3c · WebSockets | Two-way, long-lived connections; proxying them. | Study 1, 2, 3 | Ex 2 |
| 3d · Scaling across servers | A pub/sub backplane. | Study 4 | Ex 3-4 |
| 3e · Reliability | Heartbeats, reconnects with backoff, delivery acks. | Study 1 | Ex 5-6 |
| 4 · Live demo | Two instances that can't see each other's messages, fixed with Redis pub/sub. | — | Ex 3-4 |
| 5 · Break it | A restart makes every client reconnect at once; jitter fixes it. | Study 4 | Break it |
| 6 · Your numbers | Memory per connection; requests from polling vs one socket. | — | Measure it |
| 7 · Which to choose | When SSE is enough, when you need WebSockets, when polling is fine. | Study 1, 2 | — |

**Video idea:** *Polling vs SSE vs WebSockets: how live apps actually work*

### T26 · Geospatial indexing, depth L2
**Why it exists:** "find the nearest X" is everywhere (rides, delivery, maps, dating apps), and a normal index can't answer it efficiently.
**Must know**
- Why a normal index on latitude and longitude is a poor fit for "nearest" queries.
- Geohash, its prefix property, and the edge problem (you also check the neighbouring cells).
- Quadtrees; S2 and H3 at headline level.
- Redis GEO vs PostGIS.
- The cost of moving objects that update their location every few seconds.

**Study path (free, in this order, ~4.5 h)**
1. [Wikipedia, "Geohash"](https://en.wikipedia.org/wiki/Geohash): how the encoding works and its limitations for proximity searches (the edge problem). *(45 min)*
2. [Redis docs, "Redis geospatial"](https://redis.io/docs/latest/develop/data-types/geospatial/). *(20 min)*
3. PostGIS workshop: ["Geometries"](https://postgis.net/workshops/postgis-intro/geometries.html), ["Spatial Indexing"](https://postgis.net/workshops/postgis-intro/indexing.html) and ["Nearest-Neighbour Searching"](https://postgis.net/workshops/postgis-intro/knn.html). Set up with the [workshop intro](https://postgis.net/workshops/postgis-intro/). *(2 h)*
4. [Uber, "H3: Uber's Hexagonal Hierarchical Spatial Index"](https://www.uber.com/blog/h3/), plus the [H3 docs introduction](https://h3geo.org/docs/). *(1 h)*

*If something is still unclear:* [S2 Geometry](https://s2geometry.io/), for the headline idea of S2 cells.

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

**Video breakdown**

| Part | What you cover | Learn from | Practise with |
|---|---|---|---|
| 1 · Problem | Find the nearest of 50,000 points in milliseconds. | — | — |
| 2 · Mental model | A grid of map squares: search only your square and its neighbours. | Study 1 | — |
| 3a · Why a normal index fails | Two separate ranges don't make a circle. | Study 3 (Spatial Indexing) | Break it |
| 3b · Geohash | Encoding and shared prefixes. | Study 1 | Ex 1 |
| 3c · The edge problem | Close points with different geohashes; checking neighbours. | Study 1 | Ex 2 |
| 3d · Quadtrees | Splitting cells that get too full. | Study 3 | Ex 5 |
| 3e · S2 and H3 | Headline level: better cells. | Study 4 | — |
| 3f · Tools | Redis GEO vs PostGIS. | Study 2, 3 | Ex 3-4 |
| 3g · Moving objects | Updating thousands of positions every few seconds. | Study 2 | Ex 6 |
| 4 · Live demo | Radius search in Redis, with the points shown on a map. | — | Ex 3 |
| 5 · Break it | A B-tree on (lat, lng) and its `EXPLAIN`. | Study 3 | Break it |
| 6 · Your numbers | No index vs B-tree vs Redis GEO vs PostGIS GiST; update throughput. | — | Measure it |
| 7 · Which to choose | PostGIS for rich queries, Redis for fast-moving points. | Study 2, 3 | — |

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
