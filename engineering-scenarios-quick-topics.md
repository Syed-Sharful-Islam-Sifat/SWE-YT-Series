# "How Would You Do It?": A Growth Path Through Real Engineering Problems

*General software engineering problems, how real teams solved them, ordered so you grow a level at a time · Nov 13, 2026 · @Lion*

Each topic is a situation real teams face: **"We have problem X. What's going on, and how do we fix it without breaking anything?"** Most of the video goes to the solution.

**Three rules for every topic:**
1. **It must be interesting** to someone scrolling past: a relatable pain, a surprising number, or a debate.
2. **The problem and solution must be general.** They should apply whatever language, framework or cloud the viewer uses. A company is only the real-world example.
3. **It must make you grow.** Topics start inside what you know and move step by step into things you've never done. Every level is a little beyond the last.

> **How this list was researched:** the case-study topics and facts come from web searches of engineering blogs, papers, conference talks and newsletters (Nov 2026). I worked from search excerpts and summaries, **not the full articles**. Before recording, read the original post (the first link in each topic) and take every number from it, not from this doc.

---

## The growth path: 43 episodes in 4 levels

| Level | Episodes | Theme | Where it sits for you |
|---|---|---|---|
| 1 | 1-13 | The bugs every backend developer meets | Inside your experience. Fast prep while you build the recording habit and learn to explain clearly |
| 2 | 14-25 | Changing things without breaking them | One step up. You know the tools; the new part is the *safe process* senior engineers use |
| 3 | 26-35 | Systems under pressure | Stretch. Load, outages, reliability and architecture decisions you haven't owned yet |
| 4 | 36-43 | Scale | Advanced. Large-scale data, real-time systems and infrastructure. This is where most of your growth happens |

### How to grow from this, not just talk about it

- **Learn in public.** You don't have to pretend to be an expert. From Level 3 on, frame videos honestly: *"I've never run a system at this scale, so I studied how Discord did it. Here's what I learned."* That builds more trust than faking experience, and it shows interviewers you learn fast.
- **Build one small demo per level.** Reading teaches the idea; building makes it yours. Each demo is also a GitHub project to show employers:
  - **Level 1:** a small Go API that reproduces an N+1 problem, pool exhaustion and a cache stampede, measured with a load-testing tool (e.g. k6) so you can show p50/p99 before and after.
  - **Level 2:** an expand-and-contract migration on PostgreSQL with zero downtime, plus a versioned API.
  - **Level 3:** a Redis-backed rate limiter and job queue in Go that survives a worker crash.
  - **Level 4:** a WebSocket chat server that fans out to thousands of fake clients, or a tiny sharded key-value store.
- **Use your own experience only as seasoning.** Where you've done something related, add a 1-minute "from my own work" segment. It adds credibility without the video being about you. The **Your angle** column marks where this fits.

---

## Episode order

### Level 1: "The bugs every backend developer meets"

| Ep | ID | Title | Prep | Builds on | Your angle |
|---|---|---|---|---|---|
| 1 | B15 | A page runs 1,200 queries (N+1) | 1 day | Starting point | You've tuned slow queries |
| 2 | 1.2 | The database is almost idle, but the API keeps timing out | 1 day | Ep 1: slow queries hold connections | — |
| 3 | 1.3 | A popular cache key expires and the database melts | 1 day | Ep 1-2: caching to reduce database load | You've designed caching layers |
| 4 | 1.1 | Your average response time is lying to you | 1 day | Ep 1-3: measuring slowness properly | — |
| 5 | B1 | An endpoint got 10× slower after a release | 1 day | Ep 4: percentiles to catch a regression | — |
| 6 | B2 | Memory keeps growing until the server restarts | 1 day | Ep 5: profiling, now for memory | Node.js memory leaks |
| 7 | B9 | The report button times out | 1 day | Ep 2: long requests hold resources, so move work to a queue | — |
| 8 | 4.3 | A customer paid ৳1 for a ৳500 order: never trust the client | 1 day | Ep 7: what the server must check | You've closed a gap like this |
| 9 | B13 | Money stored as float | 1 day | Ep 8: money bugs, part 2 | — |
| 10 | B14 | All timestamps stored in local time | 1 day | Ep 9: data stored the wrong way | Reservation systems |
| 11 | B7 | An API key was pushed to GitHub | 1 day | A mistake everyone fears | — |
| 12 | 2.1 | One regex took down Cloudflare's network | 1 day | Ep 11: small changes, huge damage | Your algorithm background helps explain backtracking |
| 13 | 9.6 | Does competitive programming help in a real job? (debate) | 1 day | End of Level 1 break: an opinion episode | You have both backgrounds |

### Level 2: "Changing things without breaking them"

| Ep | ID | Title | Prep | Builds on | Your angle |
|---|---|---|---|---|---|
| 14 | Patterns | How big companies change things with zero downtime (the 6 patterns) | 1 day | Ep 12: "every change needs a safe rollout". From here on, refer back to "pattern #N" | — |
| 15 | B10 | Switching payment providers | 1 day | Ep 14: pattern 4 (branch by abstraction) | Stripe; local gateways (bKash, SSLCommerz) |
| 16 | 7.1 | Change your API without breaking anyone | 1 day | Ep 14: expand and contract for APIs | You build REST APIs |
| 17 | 10.1 | Split a huge table into a new one while the app is live | 1 day | Ep 14: pattern 2 (dual write) in full | — |
| 18 | 4.2 | Our password hashing is weak. Fix it without forcing resets | 1 day | Ep 16-17: migrate gradually, now for passwords | — |
| 19 | B6 | Add "Login with Google" to an existing app | 1 day | Ep 18: changing how users log in | Firebase Auth |
| 20 | B4 + 5.2 | The web page takes 8 seconds on a phone | 1 day | Ep 4: percentiles, now measured in the browser | Next.js image optimization |
| 21 | B11 | Changing code that has no tests | 1 day | Ep 14: pattern 5 (shadow comparison) for code | — |
| 22 | 9.1 | Upgrade a core dependency years out of date | 1 day | Ep 21: changing old code safely | A Next.js major upgrade works as an example |
| 23 | 9.2 | Big bang or step by step? Changing millions of lines (debate) | 1 day | Ep 22: when incremental is *not* the answer | — |
| 24 | B12 | CI takes 45 minutes | 1 day | Ep 22-23: big changes need fast feedback | — |
| 25 | 9.5 | Writing specs so AI coding agents don't break your code | 1 day | Ep 21-24: safe change, with AI writing the code | You've written specs for AI agents |

### Level 3: "Systems under pressure"

| Ep | ID | Title | Prep | Builds on | Your angle |
|---|---|---|---|---|---|
| 26 | 6.1 | Our job queue ran out of memory, and everything stopped | 1 day | Ep 7: what happens when the queue itself fails | — |
| 27 | 7.2 | Some customers are flooding our API | 1 day | Ep 26: rate limits and load shedding | — |
| 28 | 2.2 | Everyone came back from holiday and autoscaling made it worse | 1 day | Ep 27: when "add more servers" doesn't help | — |
| 29 | B3 | The dashboard is green, but users are angry (SLOs and alerting) | 1 day | Ep 28: noticing problems from the user's side | — |
| 30 | 2.3 | Break your own servers on purpose (Chaos Monkey) | 1 day | Ep 29: finding failures before users do | — |
| 31 | 8.1 | Replace an old system piece by piece | 1 day | Ep 14: patterns 3 and 5 at full scale | Your Node.js → Go migration |
| 32 | 7.3 | Replace the API behind a mobile app with no downtime | 1 day | Ep 31: replay testing and canaries | — |
| 33 | 5.3 + B5 | Ship a risky change to an app on millions of devices | 1 day | Ep 32: when you can't simply roll back | — |
| 34 | 9.3 | One repo or many? (debate) | 1 day | Ep 31-33: how code structure shapes change | — |
| 35 | 8.2 | Microservices: in, and back out again (debate) | **2 days** | Ep 34: splitting code vs splitting services | Your Go services |

### Level 4: "Scale"

| Ep | ID | Title | Prep | Builds on | Your angle |
|---|---|---|---|---|---|
| 36 | 8.3 | One customer could see another customer's data (multi-tenant isolation) | 1 day | Ep 8: trust boundaries, now between customers | You've worked on a multi-tenant SaaS |
| 37 | 6.2 | One chat room with 30,000 people online | 1 day | Ep 26-27: queues and load, now with live connections | Your Socket.io project |
| 38 | 10.4 | Change the schema of a huge table without locking it | 1 day | Ep 17: same goal, a billion rows | — |
| 39 | 10.3 | Move trillions of messages to a new database with zero downtime | 1 day | Ep 17 and 38: dual writes at massive scale | — |
| 40 | 10.2 | Our database is too big for one machine. Shard it | **2 days** | Ep 39: when one database isn't enough | — |
| 41 | B8 | "Delete my data" across 30 services | 1 day | Ep 35 (many services) + Ep 26 (events) | — |
| 42 | 6.3 | Two people edit the same design at the same time | **2 days** | Ep 37: real-time, plus conflicts | — |
| 43 | 3.1 + 3.2 | Our cloud bill is too high. Should we leave the cloud? (finale debate) | 1 day | Ep 40-42: the money side of scale | You've projected cloud costs |

**Not in the plan:** GitHub's 2FA rollout (more product than engineering for this audience), Airbnb dropping React Native (mobile, outside your direction) and Uber's stale feature flags (niche). You can add them back later.

**Tip:** if a topic doesn't interest you when its turn comes, swap it with the next one in the same level. The levels matter more than the exact order.

---

## The 1-day prep routine

| Time | Task |
|---|---|
| Morning (2-3 h) | Read the original case study or reference fully. Note: the problem, the constraints, the steps they took, what went wrong. |
| Midday (1 h) | Generalize it: draw the "before" and "after", and write the solution as steps your viewer could apply to their own system. |
| Afternoon (2 h) | Whiteboard panels, hook, bullet-point script. |
| Evening (1-2 h) | Record and do a light edit. |

## The video shape (8-12 minutes)

1. **The situation (1 min):** "Imagine your team has this problem ..." Pose it as *your* problem first.
2. **The naive fix (1-2 min):** what most people would try first, and why it doesn't work.
3. **What's really going on (2 min):** the concept behind the problem. Ask viewers "what would you check first?"
4. **The solution (4-5 min):** step by step, including how to roll it out safely. The main part.
5. **How [Company] actually did it (1-2 min):** the real story and numbers.
6. **Checklist (30 s):** 3-4 rules to remember. This also makes a good LinkedIn post.

---

## The 6 patterns behind almost every safe change

This is **Episode 14**, the start of Level 2. From then on, refer back to them ("this is pattern #5 again").

1. **Expand and contract (parallel change):** add the new thing next to the old, move everyone over, then remove the old. ([Martin Fowler: Parallel Change](https://www.martinfowler.com/bliki/ParallelChange.html))
2. **Dual write → backfill → compare → switch reads → stop old writes:** the standard way to move data. (Stripe, topic 10.1)
3. **Strangler fig:** replace an old system piece by piece, routing one feature at a time to the new one. ([Wikipedia: Strangler fig pattern](https://en.wikipedia.org/wiki/Strangler_fig_pattern))
4. **Branch by abstraction:** put an interface in front of the thing you're replacing, so the swap happens behind it.
5. **Shadow mode / side-by-side comparison:** run old and new together, compare results, keep serving the old one until they match. (Khan Academy, Netflix, Stripe)
6. **Gradual rollout with an escape hatch:** staged groups, feature flags, canaries, and a way to go back instantly. (Dropbox, Cloudflare's lesson)

---

# Part A: Topics based on real case studies

## 1. Performance

**1.1 "Your average response time is lying to you" (tail latency)**
- **Situation:** your dashboard says the average response is 80 ms, yet users complain the app is slow. One page calls 20 backend services, and it's only as fast as the slowest one.
- **Real case:** Google's paper "The Tail at Scale" (Dean & Barroso, 2013) explains why rare slow responses become common when one request fans out to many servers, and gives techniques like **hedged requests**: if a reply hasn't arrived after the 95th-percentile time, send the same request to another replica and use whichever answers first. Summaries report a large drop in tail latency for only a few percent extra load.
- **What you teach:** averages vs percentiles (p50, p95, p99), why fan-out multiplies slowness, timeouts, hedged requests.
- **Keep it simple:** percentiles plus one fix. Don't try to cover the whole paper.
- **Twist:** hedging doubles work for slow requests. When can it make an overloaded system worse?
- **Read:** "The Tail at Scale", Communications of the ACM, Feb 2013 (search the title for the PDF) · [The Morning Paper summary](https://blog.acolyer.org/2015/01/15/the-tail-at-scale/)

**1.2 "The database is almost idle, but the API keeps timing out" (connection pool exhaustion)**
- **Situation:** requests fail with "timeout waiting for a connection", yet the database CPU is low.
- **Real cases:**
  - **Layers:** a surge of background jobs took every connection in a pool shared by the web app and the workers, so web requests couldn't get one. Fix: separate connection pools for the app and the workers.
  - **Decagon:** a few stuck connections quietly shrank the usable pool, so unrelated requests queued and then failed with pool exhaustion.
- **What you teach:** what a connection pool is, why slow queries starve fast ones, isolating workloads (the bulkhead idea), what to monitor (waiting clients, not just database CPU).
- **Twist:** "just make the pool bigger": why that can move the problem into the database instead of fixing it.
- **Read:** [Layers: Using multiple PgBouncers to isolate application and worker connections](https://www.uselayers.com/articles/using-multiple-pgbouncers-to-isolate-application-and-worker-connections) · [Decagon: Debugging a latent PgBouncer bug across four layers](https://decagon.ai/blog/debugging-a-latent-pgbouncer-bug-across-four-layers)

**1.3 "A popular cache key expires and the database melts" (caching: Facebook memcache, Shopify)**
- **Situation:** a heavily read cache key is deleted or expires, and thousands of requests miss at once and hit the database (thundering herd). Also: a slow request writes old data into the cache after a newer one (stale set).
- **Real case:** Facebook's memcache paper introduced **leases**: on a miss, the cache gives one client a token to fetch from the database; others wait or retry. Tokens are handed out at most once every 10 seconds per key, and a delete invalidates the lease so stale writes are rejected. One summary reports peak database queries dropping from 17K/s to 1.3K/s.
- **Second real case:** Shopify's rewritten storefront uses layered caching (a local data cache on the same machine plus full-page caching in Redis) and reports average responses nearly 5× faster than before.
- **What you teach:** cache-aside, cache layers, thundering herd, stale sets, and how one small mechanism (leases) fixes two problems.
- **Keep it simple:** explain the stampede and "only one request refills the cache". Skip the rest of the paper.
- **Read:** "Scaling Memcache at Facebook" (NSDI 2013; search the title for the PDF) · [MIT 6.824 FAQ on the paper](https://pdos.csail.mit.edu/6.824/papers/memcache-faq.txt) · [Engineer's Codex summary](https://read.engineerscodex.com/p/how-facebook-scaled-memcached) · [Shopify: Simplify, batch, and cache: storefront rendering](https://shopify.engineering/simplify-batch-cache-optimized-server-side-storefront-rendering)

## 2. Reliability and incidents

**2.1 "One regex took down Cloudflare's network" (Cloudflare, 2019)**
- **Situation:** a small configuration change is deployed and CPU hits 100% everywhere.
- **Real case:** a new WAF rule (to block a SharePoint exploit) contained a regex that caused catastrophic backtracking in the PCRE engine. The WAF rules shipped globally in one step, skipping the usual staged rollout, causing a ~27-minute global outage. Fixes included moving to a regex engine with run-time guarantees (re2 or Rust's regex) and changing testing and deployment.
- **What you teach:** regex backtracking, why "config is code", and why every change needs staged rollout.
- **Read:** [Cloudflare: Details of the Cloudflare outage on July 2, 2019](https://blog.cloudflare.com/details-of-the-cloudflare-outage-on-july-2-2019/) · [InfoQ summary](https://www.infoq.com/news/2019/07/waf-causes-outage/) · [Surfing Complexity analysis](https://surfingcomplexity.blog/2019/08/02/contributors-mitigators-risks-cloudflare-2019-07-02-outage/)

**2.2 "Everyone came back from holiday and autoscaling made it worse" (Slack, Jan 4 2021)**
- **Situation:** traffic spikes on the first workday of the year; you add servers, but things don't improve.
- **Real case:** one of Slack's AWS Transit Gateways (network infrastructure) became overloaded as users returned. Slack tried to add 1,200 servers to its web tier, but with the network saturated it didn't help. AWS manually added capacity.
- **What you teach:** saturation, how autoscaling can amplify a problem, hidden dependencies in the network layer, reading a postmortem.
- **Read:** [Slack: Slack's outage on January 4th 2021](https://slack.engineering/slacks-outage-on-january-4th-2021/) · [The Register summary](https://www.theregister.com/2021/02/02/slack_outage_aws_autoscaling/) · [Surfing Complexity: a tale of saturation](https://surfingcomplexity.blog/2021/02/08/slacks-jan-2021-outage-a-tale-of-saturation/)

**2.3 "Break your own servers on purpose" (Netflix Chaos Monkey)**
- **Situation:** your system is "redundant" on paper, but you've never seen what happens when a server dies.
- **Real case:** during Netflix's move from its own data centers to AWS, cloud instances could disappear without warning. Netflix built Chaos Monkey, which randomly terminates instances in production so teams must build services that survive it; it later grew into the Simian Army.
- **What you teach:** designing for failure, chaos engineering, starting small (game days, staging first).
- **Twist:** sources disagree on exact dates, so verify against Netflix's own posts.
- **Read:** [Gremlin: The origin of Chaos Monkey](https://gremlin.com/chaos-monkey/the-origin-of-chaos-monkey) · [IEEE Spectrum: Chaos engineering saved your Netflix](https://spectrum.ieee.org/chaos-engineering-saved-your-netflix) · Netflix Tech Blog (search "Chaos Monkey" and "Simian Army")

## 3. Infrastructure and cost

**3.1 "Our cloud bill is too high. Should we leave the cloud?" (37signals)**
- **Situation:** a debate video. Your cloud bill keeps growing for a stable, predictable workload.
- **Real case:** 37signals (Basecamp, HEY) moved its apps from AWS to its own hardware. DHH reported savings expected to top $10 million over five years; the last piece was moving storage off S3. Coverage notes the headline figures may not include hardware refreshes, staff or data center running costs.
- **What you teach:** when cloud makes sense (unpredictable growth, small teams) and when it doesn't (stable workload, strong ops team); how to compare costs honestly.
- **Read:** [DHH: Our cloud-exit savings will now top ten million over five years](https://world.hey.com/dhh/our-cloud-exit-savings-will-now-top-ten-million-over-five-years-c7d9b5bd) · [The Register (Oct 2024)](https://www.theregister.com/2024/10/21/37signals_aws_savings) · [The Register (May 2025)](https://www.theregister.com/2025/05/09/37signals_cloud_repatriation_storage_savings)

**3.2 "Move half an exabyte off S3 without losing a file" (Dropbox Magic Pocket)** *(used in Episode 43 with 3.1)*
- **Real case:** Dropbox kept metadata in its own data centers but file data in S3. Project "Magic Pocket" moved data to custom storage across three regions; the 90% target was hit in October 2015. During the move, data lived in both places but reads came only from Magic Pocket, so they could switch back if something went wrong.
- **What you teach:** large-scale data migration, verifying data, the "keep a way back" principle.
- **Read:** [InfoQ: Dropbox moves off AWS](https://www.infoq.com/news/2016/03/Dropbox-AWS/) · [InfoQ podcast with Preslav Le](https://www.infoq.com/podcasts/preslav-le-dropbox)

## 4. Security and authentication

**4.2 "Our password hashing is weak. Fix it without forcing everyone to reset" (Dropbox + OWASP)**
- **Situation:** you discover passwords are stored with a fast hash.
- **Real case:** Dropbox's 2012 breach exposed credentials, some hashed with SHA-1 and others with bcrypt. Its later design is layered: SHA-512 first (to avoid bcrypt's 72-byte input limit), then bcrypt with a per-user salt, then AES-256 encryption with a separately stored key.
- **What you teach:** fast vs slow hashes, salt and pepper, and the migration trick: wrap old hashes now (`bcrypt(old_hash)`), re-hash to the new scheme at each user's next login.
- **Read:** [Dropbox: How Dropbox securely stores your passwords](https://dropbox.tech/security/how-dropbox-securely-stores-your-passwords) · OWASP Password Storage Cheat Sheet (search the title)

**4.3 "A customer paid ৳1 for a ৳500 order": never trust the client**
- **Situation:** the app sends the order with prices, discounts and quantities, and the server saves what it receives. Someone edits the request and orders for almost nothing.
- **Real case:** PortSwigger's Web Security Academy teaches this as "excessive trust in client-side controls": its lab has you buy a jacket by changing the price in the request with a proxy tool. The root cause is that the store never checks the price against its database. Business logic flaws like this are common in real apps.
- **What you teach:** client-side checks are only for user experience; the server recalculates every price, discount, quantity and permission from its own data; also negative quantities, reused coupons, and changing someone else's order ID.
- **Twist:** the price changed between when the user saw it and when they paid. Which price wins, and how do you show it fairly?
- **Read:** [PortSwigger: Examples of business logic vulnerabilities](https://portswigger.net/web-security/logic-flaws/examples) · [PortSwigger lab: Excessive trust in client-side controls](https://portswigger.net/web-security/logic-flaws/examples/lab-logic-flaws-excessive-trust-in-client-side-controls)

## 5. Frontend, mobile and client apps

**5.2 "Our mobile website is too slow for users on cheap phones" (Pinterest PWA)** *(used in Episode 20 with B4)*
- **Real case:** Pinterest rebuilt its mobile web as a Progressive Web App: lean JavaScript bundles, service workers for network resilience, dominant-color placeholders for images, app shell, push notifications and caching. A later post reports a large year-over-year increase in mobile web weekly active users in India.
- **What you teach:** web performance basics (bundle size, code splitting, image optimization, caching), measuring performance as a business metric.
- **Read:** [Pinterest: A one year PWA retrospective](https://medium.com/pinterest-engineering/a-one-year-pwa-retrospective-f4a2f4129e05) · Addy Osmani: "A Pinterest Progressive Web App Performance Case Study" (search the title) · [Pinterest: Driving user growth with performance improvements](https://medium.com/pinterest-engineering/driving-user-growth-with-performance-improvements-cfc50dafadd7)

**5.3 "Ship a risky change to an app installed on millions of devices" (Dropbox desktop app)** *(Episode 33, with B5)*
- **Situation:** you can't redeploy software installed on users' devices, so a bad release can't simply be rolled back on a server. Applies to any mobile or desktop app; the change itself is just the example.
- **Real case:** Dropbox moved its desktop app from Python 2 to 3, starting at a 2015 hack week. A system called **Hydra** let the app choose the old or new version at startup, so users could be switched back quickly. It rolled out to internal users, then Beta, then everyone.
- **What you teach:** client-side releases, staged rollouts to user groups, remote kill switches, an escape hatch (pattern 6).
- **Read:** [Dropbox: How we rolled out one of the largest Python 3 migrations ever](https://dropbox.tech/application/how-we-rolled-out-one-of-the-largest-python-3-migrations-ever) · [Dropbox: Incrementally migrating over one million lines of code](https://dropbox.tech/application/incrementally-migrating-over-one-million-lines-of-code-from-python-2-to-python-3)

## 6. Real-time and messaging

**6.1 "Our job queue ran out of memory, and everything stopped" (Slack)**
- **Situation:** your async job queue (Redis) is a single point of failure; when it fills up, every feature that enqueues jobs fails.
- **Real case:** Slack's Redis job queue ran out of memory, so new jobs couldn't be enqueued. Slack put **Kafka in front of Redis** as a durable buffer, with a relay service controlling how fast jobs move from Kafka into Redis, leaving the app's queue logic unchanged. One summary reports over 1.4 billion jobs on busy days, peaking at 33,000/s.
- **What you teach:** queues, back-pressure, durability vs speed, adding a buffer without rewriting the system.
- **Keep it simple:** focus on "the queue filled up and everything stopped". Kafka only needs a one-line explanation ("a durable log that can hold far more").
- **Read:** [Slack: Scaling Slack's job queue](https://slack.engineering/scaling-slacks-job-queue/)

**6.2 "One chat room with 30,000 people online" (Discord)**
- **Situation:** a real-time chat platform where one message must reach tens of thousands of connected users, and everyone reconnects at once after a blip.
- **Real case:** Discord reached nearly 5 million concurrent users. Big servers (e.g. one with ~30,000 concurrent users) made the parts handling them fall behind on their message queues at peak, and bursts of users reconnecting at the same moment were another problem.
- **What you teach:** fan-out of one message to many connections, why one huge room is harder than many small ones, reconnect storms and how to spread them out (backoff with jitter).
- **Read:** [Discord: How Discord scaled Elixir to 5,000,000 concurrent users](https://discord.com/blog/how-discord-scaled-elixir-to-5-000-000-concurrent-users) · [Elixir blog: Real-time communication at scale at Discord](https://elixir-lang.org/blog/2020/10/08/real-time-communication-at-scale-with-elixir-at-discord/)

**6.3 "Two people edit the same design at the same time" (Figma multiplayer)**
- **Situation:** you're adding real-time collaboration to an editor. What happens when two users change the same thing?
- **Real case:** Figma clients talk to servers over WebSockets. Figma considered operational transformation (OT) and CRDTs, ruled out OT, and built something *inspired by* CRDTs: the server keeps the last value any client sent for each property of each object (last-writer-wins per property), and clients create object IDs that include their client ID so they never collide. Because it relies on a central server, it's not a pure CRDT.
- **What you teach:** conflict resolution, OT vs CRDT at a high level, why a central server simplifies things, how undo works with other people's edits.
- **Read:** [Evan Wallace: How Figma's multiplayer technology works](https://madebyevan.com/figma/how-figmas-multiplayer-technology-works/) · [Liveblocks: Understanding sync engines](https://www.liveblocks.io/blog/understanding-sync-engines-how-figma-linear-and-google-docs-work)

## 7. APIs

**7.1 "Change your API without breaking anyone who uses it" (Stripe versioning)**
- **Situation:** thousands of businesses call your API, and any change can break their code.
- **Real case:** Stripe has kept compatibility with every API version since 2011. Each account is pinned to the version it first used, and changes are released as dated versions. Stripe still works hard to get the design right first, so it rarely needs a new version.
- **What you teach:** additive changes, version pinning, translating responses between versions inside the server, deprecation.
- **Twist:** your mobile app is a client of your own API, and old app versions stay installed for years.
- **Read:** [Stripe: APIs as infrastructure: future-proofing Stripe with versioning](https://stripe.com/ae/blog/api-versioning) · [Stripe docs: API versioning](https://docs.stripe.com/api/versioning?lang=curl)

**7.2 "Some customers are flooding our API, and real payments are failing" (Stripe rate limiters)**
- **Situation:** one client's script hammers your API, and important traffic suffers.
- **Real case:** Stripe runs four types of limiters in production: two rate limiters and two **load shedders**. One load shedder reserves part of the fleet for the most important requests; another limits damage during incidents already happening. Advice includes clear 429/503 errors for users and switches to turn limiters off.
- **What you teach:** rate limiting vs load shedding, prioritizing critical traffic, designing what clients see when you say "no".
- **Read:** [Stripe: Scaling your API with rate limiters](https://stripe.com/blog/rate-limiters) · [Stripe docs: rate limits](https://docs.stripe.com/rate-limits)

**7.3 "Replace the API behind a mobile app with no downtime" (Netflix → GraphQL)**
- **Situation:** you want to replace the API layer your mobile app depends on, with millions of active users.
- **Real case:** Netflix moved its mobile apps from its internal Falcor API to GraphQL with zero downtime. Phase 1: a GraphQL shim on top of Falcor, validated with A/B tests. Phase 2: a new implementation, validated with **replay testing** (compare outputs of old and new for read-only queries) and **sticky canaries** (route some users persistently to the new cluster, which also works for writes).
- **What you teach:** pattern 5 (shadow comparison) for APIs, canaries, why read-only and write traffic need different testing. No GraphQL knowledge needed; the method is the point.
- **Read:** [The New Stack: Netflix's testing strategies for migrating to GraphQL](https://thenewstack.io/netflixs-testing-strategies-for-migrating-to-graphql/) · Netflix Tech Blog: "Migrating Netflix to GraphQL Safely" (search the title)

## 8. Architecture

**8.1 "Replace an old system piece by piece, without a big-bang rewrite" (Khan Academy)**
- **Situation:** an old system has to be replaced, but the product can't stop. The languages (Python → Go) are just the example; the method works for any rewrite.
- **Real case:** Khan Academy's project "Goliath" moved a ~1M-line Python 2 monolith to Go services behind GraphQL. When Go could handle a request, the gateway called **both** Python and Go, compared results, logged differences and still returned Python's answer; then traffic moved to Go with Python as fallback; then Python was deleted. By August 2021 the new backend handled 95% of traffic. A talk reports ~3.5 years and ~40 person-years.
- **What you teach:** patterns 3 and 5 together, the clearest real example of each.
- **Read:** [Khan Academy: Go + Services = One Goliath Project](https://blog.khanacademy.org/go-services-one-goliath-project) · [Incremental rewrites with GraphQL](https://blog.khanacademy.org/incremental-rewrites-with-graphql) · [Beating the odds](https://blog.khanacademy.org/beating-the-odds-khan-academys-successful-monolith%e2%86%92services-rewrite) · [Pragmatic Engineer: Breaking up a monolith](https://newsletter.pragmaticengineer.com/p/real-world-eng-8)

**8.2 "Microservices: in, and back out again" (Airbnb, Segment, Prime Video, Shopify)**
- **Situation:** a debate video. Should your team split the monolith?
- **Real cases:**
  - **Airbnb** moved from its Rails monolith to services when deploys backed up past 100 engineers; payments had to run two parallel systems through most of the migration.
  - **Segment** went the other way: 100+ services hurt developer productivity, so they merged them back into one service ("Centrifuge").
  - **Prime Video** moved one monitoring tool from serverless pieces into a single process and cut its cost by 90%. The debate was about that workload, not microservices in general.
  - **Shopify** chose a **modular monolith**: one codebase with enforced boundaries.
- **What you teach:** the real trade-offs, and why "it depends on your stage" is the honest answer.
- **Read:** [InfoQ: Airbnb monolith → SOA](https://www.infoq.com/news/2019/02/airbnb-monolith-migration-soa/) · [Segment: Goodbye Microservices](https://twilio.com/en-us/blog/developers/best-practices/goodbye-microservices) · [InfoQ: Prime Video](https://infoq.com/news/2023/05/prime-ec2-ecs-saves-costs) · [Shopify: Deconstructing the Monolith](https://shopify.engineering/deconstructing-monolith-designing-software-maximizes-developer-productivity)

**8.3 "One customer could see another customer's data" (multi-tenant isolation)**
- **Situation:** a SaaS keeps all customers' data in the same database, separated by a `tenant_id` column. One forgotten filter, or one shared cache key, and a customer sees someone else's data.
- **Real cases:** public reports describe cross-tenant exposures at several SaaS companies, caused by things like a missing tenant filter, shared caches and shared processing logic. Treat individual incident details as second-hand and use them only as examples of *how* it happens.
- **What you teach:** the three models (shared tables, schema per tenant, database per tenant); enforcing isolation in one place (a tenant-aware data layer or database row-level security) instead of trusting every query; tenant ID in every cache key; tests that try to read across tenants.
- **Twist:** your biggest customer demands their own database. How to move one tenant out without downtime (callback to Level 2 patterns).
- **Read:** [PostgreSQL docs: Row Security Policies](https://www.postgresql.org/docs/current/ddl-rowsecurity.html) · search "multi-tenant SaaS data isolation" for the AWS SaaS whitepapers · [LoginRadius: Multi-tenant authorization patterns](https://www.loginradius.com/blog/identity/what-is-multi-tenant-authorization)

## 9. Codebase and team

**9.1 "Upgrade a core dependency years out of date, without stopping feature work" (GitHub)**
- **Situation:** you're stuck on an old version of your main framework or library, and the business won't pause features. GitHub's Rails upgrade is the example; no Rails knowledge needed.
- **Real case:** GitHub took a year and a half to go from Rails 3.2 to 5.2. The key technique was **dual booting**: the app could run on either version and CI tested both. They shipped once production ran 30 minutes at peak traffic with no visible impact. Their own fork of Rails had made it harder.
- **What you teach:** compatibility with both versions, small steps on the main branch, staying close to upstream.
- **Read:** [GitHub: Upgrading GitHub from Rails 3.2 to 5.2](https://github.blog/engineering/upgrading-github-from-rails-3-2-to-5-2/)

**9.2 "Big bang or step by step? Changing millions of lines of code" (Stripe vs Airbnb)**
- **Situation:** a debate video. A change must touch most of a huge codebase. Do it all at once, or gradually? The example is a TypeScript migration, but the question applies to any large codebase change.
- **Real cases:** **Stripe** converted 3.7M+ lines from Flow to TypeScript **in a single pull request** (merged March 2022) using a codemod. **Airbnb** built **ts-migrate** to convert projects automatically, reporting 86% of a ~6M-line frontend converted.
- **What you teach:** codemods (automated code changes), when big bang is *safer* than incremental, and what preparation makes it possible.
- **Twist:** Pinterest also reported a 3.7M-line migration. Don't mix up the numbers.
- **Read:** [Stripe: Migrating millions of lines of code to TypeScript](https://stripe.com/blog/migrating-to-typescript) · [airbnb/ts-migrate](https://github.com/airbnb/ts-migrate) · [InfoQ: Airbnb TypeScript migration](https://www.infoq.com/news/2020/08/airbnb-typescript-migration)

**9.3 "One repo or many?" (Google's monorepo)**
- **Situation:** a debate video. Your company has 40 repositories and cross-repo changes are painful. Should you merge them?
- **Real case:** Google's 2016 paper describes one repository with about two billion lines of code in nine million files (2015 numbers), about 15 million lines changed per week in 2014, hierarchical code ownership, and a custom version control system. The paper discusses both benefits and costs.
- **What you teach:** monorepo vs many repos: atomic changes, shared tooling, code ownership, and the tooling a monorepo needs.
- **Keep it simple:** stay with the trade-offs; skip Google's internal tools.
- **Read:** [Google Research: Why Google stores billions of lines of code in a single repository](https://research.google/pubs/why-google-stores-billions-of-lines-of-code-in-a-single-repository/) · [ACM full text](https://cacmb4.acm.org/magazines/2016/7/204032-why-google-stores-billions-of-lines-of-code-in-a-single-repository/fulltext)

**9.5 "Writing specs so AI coding agents don't break your code"**
- **Situation:** the team uses AI coding agents. They're fast, but they make confident mistakes, ignore conventions and change things nobody asked for.
- **Real cases:** GitHub open-sourced **Spec Kit** (Sept 2025) for "spec-driven development": write a spec, turn it into a technical plan, split it into small tasks an agent can implement and test one at a time. Anthropic's Claude Code best-practice guidance stresses project instruction files, keeping context small, and test-first workflows so the agent can check its own work.
- **What you teach:** why vague prompts fail; what a good spec contains (goal, constraints, examples, what *not* to change, how to verify); small verifiable tasks; human review checklists. This is classic engineering practice (clear requirements, small changes, tests) applied to a new tool.
- **Twist:** who is responsible when agent-written code breaks production? Reviewing AI code vs reviewing human code.
- **Read:** [GitHub: Spec-driven development with AI](https://github.blog/ai-and-ml/generative-ai/spec-driven-development-with-ai-get-started-with-a-new-open-source-toolkit/) · [Claude Code docs: Best practices](https://code.claude.com/docs/en/best-practices)

**9.6 "Does competitive programming help in a real job?" (debate)**
- **Situation:** a huge question for CS students, especially in Bangladesh: is grinding Codeforces worth it?
- **What you teach:** where contest skills directly help (complexity thinking, edge cases, debugging under pressure, data structures behind caches, rate limiters and pagination); where they don't (reading large codebases, design trade-offs, communication, testing, operating systems in production); how to balance both.
- **Format:** opinion episode backed by concrete examples. Show 2-3 real backend problems where an algorithm idea mattered, and 2-3 where it didn't.
- **Read:** no single source. Use your own contest history and the episodes before it for examples.

## 10. Data and databases

**10.1 "Split a huge table into a new one while the app is live" (Stripe)**
- **Real case:** Stripe moved hundreds of millions of Subscriptions out of Customer records with a four-step dual-writing pattern, using the Scientist library to compare reads from both.
- **What you teach:** pattern 2 in full, on a small example first.
- **Read:** [Stripe: Online migrations at scale](https://stripe.com/se/blog/online-migrations) · [Gusto: splitting a large table](https://engineering.gusto.com/old-write)

**10.2 "Our database is too big for one machine. Shard it without a rewrite" (Figma, Notion)**
- **Real cases:** **Figma** went from one Postgres on AWS's largest instance to sharding, doing **logical sharding first** (routing as if sharded while data stays in one database) and physical sharding later, through a custom proxy. **Notion** sharded by workspace ID, the block table first, using Postgres logical replication.
- **What you teach:** shard keys, why "change the routing before moving the data" reduces risk.
- **Read:** [Figma: How Figma's databases team lived to tell the scale](https://www.figma.com/en-jp/blog/how-figmas-databases-team-lived-to-tell-the-scale/) · [Notion: Herding elephants](https://notion.com/blog/sharding-postgres-at-notion)

**10.3 "Move trillions of messages to a new database with zero downtime" (Discord)**
- **Real case:** Discord went MongoDB → Cassandra → ScyllaDB (177 Cassandra nodes became 72 ScyllaDB nodes). ScyllaDB speaks Cassandra's protocol, so app logic didn't need rewriting. They used dual writes plus a bulk migrator.
- **What you teach:** hot partitions (a data-modeling problem, familiar from any NoSQL database), dual writes + bulk copy.
- **Read:** [Discord: How Discord stores trillions of messages](https://discord.com/blog/how-discord-stores-trillions-of-messages)

**10.4 "Change the schema of a huge table without locking it" (GitHub gh-ost)**
- **Real case:** GitHub's gh-ost copies data to a "ghost" table and follows ongoing changes through MySQL's binary log instead of triggers; it can pause and throttle on load and replication lag.
- **What you teach:** shadow table + catch up + atomic swap. The idea applies to any database.
- **Read:** [GitHub: gh-ost announcement](https://github.blog/2016-08-01-gh-ost-github-s-online-migration-tool-for-mysql) · [gh-ost README](https://github.com/github/gh-ost)

---

# Part B: Everyday scenarios (no single famous case study)

Common problems in normal jobs. There's no one big write-up, but the solution is standard and documented. Use the reference to check your plan.

| # | Area | Scenario | The solution in one line | Reference to check |
|---|---|---|---|---|
| B1 | Performance | **An endpoint got 10× slower after a release** | Compare before/after → profile → flame graph → find the hot path → fix and add a performance test | Brendan Gregg's website (flame graphs, USE method) |
| B2 | Performance | **Memory keeps growing until the server restarts** | Watch memory over time → heap snapshots at two moments → compare → find what keeps growing (caches, listeners, globals) | Your language's profiler docs (e.g. Node.js heap snapshots) |
| B3 | Reliability | **The dashboard is green, but users are angry** | Alert on user-facing symptoms, not causes → define SLOs → delete or downgrade noisy alerts | Google SRE book (sre.google/books): "Monitoring Distributed Systems", "Service Level Objectives" |
| B4 | Frontend | **The web page takes 8 seconds on a phone** | Measure (Core Web Vitals) → shrink the bundle → code splitting / lazy loading → optimize images → cache | web.dev |
| B5 | Mobile | **A crash shipped to the app store and you can't roll back** | Staged rollout → remote config / kill switch → hotfix release → crash monitoring | Google Play staged rollouts docs; Apple phased release docs |
| B6 | Security | **Add "Login with Google" to an existing app** | OpenID Connect → link to existing accounts by verified email → keep password login during transition | OpenID Connect docs; OWASP cheat sheets |
| B7 | Security | **An API key was pushed to GitHub** | Rotate first → check logs → clean history → secret manager + secret scanning | GitHub docs: "Removing sensitive data from a repository" |
| B8 | Security / data | **"Delete my data" across 30 services** | Data inventory → `UserDeleted` event → per-service tracking → crypto-shredding for backups | Search "crypto shredding" |
| B9 | Architecture | **The report button times out** | Async job + job ID + polling/notification + file in object storage | Microsoft Azure Architecture Center: "Asynchronous Request-Reply pattern" |
| B10 | Architecture | **Switching payment providers** | Your own `PaymentProvider` interface → new provider behind a flag → old payments finish on the old one | Martin Fowler: "BranchByAbstraction" |
| B11 | Codebase | **Changing code that has no tests** | Characterization tests → small refactors → shadow-run old vs new | Michael Feathers, *Working Effectively with Legacy Code*; GitHub's Scientist library |
| B12 | Codebase | **CI takes 45 minutes** | Measure → cache → parallelize → run affected tests only → quarantine flaky tests | Google Testing Blog posts on flaky tests |
| B13 | Data | **Money stored as float** | Integer minor units (paisa) or `DECIMAL`; dual write; documented rounding; reconcile totals | Martin Fowler: "Money" pattern |
| B14 | Data | **All timestamps stored in local time** | Classify columns (instant vs date vs wall time) → UTC for instants → convert only in the UI | Jon Skeet: "Storing UTC is not a silver bullet" |
| B15 | Data / performance | **A page runs 1,200 queries (N+1)** | Measure queries per request → eager/batch loading → test that caps query count | Your ORM's or database driver's docs on batching |

---

## Where to find more topics

**Collections and newsletters**
- [Pragmatic Engineer: Real-World Engineering Challenges](https://newsletter.pragmaticengineer.com/p/real-world-4): real engineering stories across many areas. Some content is paid.
- [kilimchoi/engineering-blogs](https://github.com/kilimchoi/engineering-blogs): 400+ company engineering blogs. Not updated recently, so some links may be dead.
- [Simon Willison's link blog](https://simonwillison.net/): links and summarizes many of the best engineering posts.
- **InfoQ** (infoq.com): conference talks and news summaries.
- **Hacker News** (search old discussions at hn.algolia.com): the comments add counter-arguments and twists.
- **PortSwigger Web Security Academy:** free, with many interesting security topics that suit backend developers.

**Where to search by area**
- **Performance:** Shopify, Netflix, classic papers like "The Tail at Scale", Brendan Gregg's blog.
- **Reliability:** Cloudflare blog (detailed postmortems), Slack Engineering, Google SRE book, Surfing Complexity (incident analysis).
- **Real-time:** Discord, Slack, Figma.
- **APIs:** Stripe blog and docs.
- **Code and team:** GitHub, Google, Stripe.

**Patterns and books**
- [Martin Fowler: Parallel Change](https://www.martinfowler.com/bliki/ParallelChange.html), Strangler Fig, Branch by Abstraction, Feature Toggles.
- *Designing Data-Intensive Applications* (Kleppmann): the best single book for Levels 3 and 4.
- *Release It!* (Nygard), *Site Reliability Engineering* (Google, free online), *Working Effectively with Legacy Code* (Feathers).

**Summaries vs originals:** ByteByteGo, Medium, dev.to and YouTube summaries are good for *finding* topics, but some are AI-generated or wrong. Take your facts from the company's original post.

---

## Sources used for this list

- **Performance:** [The Morning Paper: The Tail at Scale](https://blog.acolyer.org/2015/01/15/the-tail-at-scale/) · [Layers: multiple PgBouncers](https://www.uselayers.com/articles/using-multiple-pgbouncers-to-isolate-application-and-worker-connections) · [Decagon: latent PgBouncer bug](https://decagon.ai/blog/debugging-a-latent-pgbouncer-bug-across-four-layers) · [Shopify: storefront rendering](https://shopify.engineering/simplify-batch-cache-optimized-server-side-storefront-rendering) · [MIT 6.824 memcache FAQ](https://pdos.csail.mit.edu/6.824/papers/memcache-faq.txt) · [Engineer's Codex: Facebook memcached](https://read.engineerscodex.com/p/how-facebook-scaled-memcached)
- **Reliability:** [Cloudflare: July 2, 2019 outage](https://blog.cloudflare.com/details-of-the-cloudflare-outage-on-july-2-2019/) · [InfoQ: Cloudflare WAF outage](https://www.infoq.com/news/2019/07/waf-causes-outage/) · [Surfing Complexity: Cloudflare](https://surfingcomplexity.blog/2019/08/02/contributors-mitigators-risks-cloudflare-2019-07-02-outage/) · [Slack: Jan 4 2021 outage](https://slack.engineering/slacks-outage-on-january-4th-2021/) · [The Register: Slack outage](https://www.theregister.com/2021/02/02/slack_outage_aws_autoscaling/) · [Surfing Complexity: Slack](https://surfingcomplexity.blog/2021/02/08/slacks-jan-2021-outage-a-tale-of-saturation/) · [Gremlin: origin of Chaos Monkey](https://gremlin.com/chaos-monkey/the-origin-of-chaos-monkey) · [IEEE Spectrum: chaos engineering](https://spectrum.ieee.org/chaos-engineering-saved-your-netflix)
- **Infrastructure:** [DHH: cloud-exit savings](https://world.hey.com/dhh/our-cloud-exit-savings-will-now-top-ten-million-over-five-years-c7d9b5bd) · [The Register (2024)](https://www.theregister.com/2024/10/21/37signals_aws_savings) · [The Register (2025)](https://www.theregister.com/2025/05/09/37signals_cloud_repatriation_storage_savings) · [InfoQ: Dropbox off AWS](https://www.infoq.com/news/2016/03/Dropbox-AWS/) · [InfoQ podcast: Preslav Le](https://www.infoq.com/podcasts/preslav-le-dropbox)
- **Security:** [Dropbox: password storage](https://dropbox.tech/security/how-dropbox-securely-stores-your-passwords) · [PortSwigger: business logic vulnerability examples](https://portswigger.net/web-security/logic-flaws/examples) · [PortSwigger lab: excessive trust in client-side controls](https://portswigger.net/web-security/logic-flaws/examples/lab-logic-flaws-excessive-trust-in-client-side-controls)
- **Frontend and client apps:** [Pinterest: one year PWA retrospective](https://medium.com/pinterest-engineering/a-one-year-pwa-retrospective-f4a2f4129e05) · [Pinterest: performance and growth](https://medium.com/pinterest-engineering/driving-user-growth-with-performance-improvements-cfc50dafadd7) · [Dropbox: Python 3 rollout](https://dropbox.tech/application/how-we-rolled-out-one-of-the-largest-python-3-migrations-ever) · [Dropbox: incremental migration](https://dropbox.tech/application/incrementally-migrating-over-one-million-lines-of-code-from-python-2-to-python-3)
- **Real-time:** [Slack: job queue](https://slack.engineering/scaling-slacks-job-queue/) · [Discord: Elixir 5M users](https://discord.com/blog/how-discord-scaled-elixir-to-5-000-000-concurrent-users) · [Elixir blog: Discord](https://elixir-lang.org/blog/2020/10/08/real-time-communication-at-scale-with-elixir-at-discord/) · [Figma multiplayer (archived)](https://madebyevan.com/figma/how-figmas-multiplayer-technology-works/) · [Liveblocks: sync engines](https://www.liveblocks.io/blog/understanding-sync-engines-how-figma-linear-and-google-docs-work)
- **APIs:** [Stripe: API versioning](https://stripe.com/ae/blog/api-versioning) · [Stripe docs: versioning](https://docs.stripe.com/api/versioning?lang=curl) · [Stripe: rate limiters](https://stripe.com/blog/rate-limiters) · [Stripe docs: rate limits](https://docs.stripe.com/rate-limits) · [The New Stack: Netflix GraphQL](https://thenewstack.io/netflixs-testing-strategies-for-migrating-to-graphql/)
- **Architecture:** [Khan Academy: Goliath](https://blog.khanacademy.org/go-services-one-goliath-project) · [Incremental rewrites with GraphQL](https://blog.khanacademy.org/incremental-rewrites-with-graphql) · [Beating the odds](https://blog.khanacademy.org/beating-the-odds-khan-academys-successful-monolith%e2%86%92services-rewrite) · [Pragmatic Engineer #8](https://newsletter.pragmaticengineer.com/p/real-world-eng-8) · [InfoQ: Airbnb SOA](https://www.infoq.com/news/2019/02/airbnb-monolith-migration-soa/) · [Segment: Goodbye Microservices](https://twilio.com/en-us/blog/developers/best-practices/goodbye-microservices) · [InfoQ: Prime Video](https://infoq.com/news/2023/05/prime-ec2-ecs-saves-costs) · [Shopify: Deconstructing the Monolith](https://shopify.engineering/deconstructing-monolith-designing-software-maximizes-developer-productivity) · [PostgreSQL: Row Security Policies](https://www.postgresql.org/docs/current/ddl-rowsecurity.html) · [LoginRadius: multi-tenant authorization](https://www.loginradius.com/blog/identity/what-is-multi-tenant-authorization)
- **Codebase and team:** [GitHub: Rails upgrade](https://github.blog/engineering/upgrading-github-from-rails-3-2-to-5-2/) · [Stripe: TypeScript](https://stripe.com/blog/migrating-to-typescript) · [airbnb/ts-migrate](https://github.com/airbnb/ts-migrate) · [InfoQ: Airbnb TypeScript](https://www.infoq.com/news/2020/08/airbnb-typescript-migration) · [Google: monorepo paper](https://research.google/pubs/why-google-stores-billions-of-lines-of-code-in-a-single-repository/) · [GitHub: spec-driven development](https://github.blog/ai-and-ml/generative-ai/spec-driven-development-with-ai-get-started-with-a-new-open-source-toolkit/) · [Claude Code docs: best practices](https://code.claude.com/docs/en/best-practices)
- **Data:** [Stripe: online migrations](https://stripe.com/se/blog/online-migrations) · [Gusto: old write](https://engineering.gusto.com/old-write) · [Figma: databases team](https://www.figma.com/en-jp/blog/how-figmas-databases-team-lived-to-tell-the-scale/) · [Notion: Herding elephants](https://notion.com/blog/sharding-postgres-at-notion) · [Discord: trillions of messages](https://discord.com/blog/how-discord-stores-trillions-of-messages) · [GitHub: gh-ost](https://github.blog/2016-08-01-gh-ost-github-s-online-migration-tool-for-mysql) · [gh-ost repo](https://github.com/github/gh-ost)
- **General:** [Pragmatic Engineer: Real-World Engineering Challenges](https://newsletter.pragmaticengineer.com/p/real-world-4) · [Martin Fowler: Parallel Change](https://www.martinfowler.com/bliki/ParallelChange.html) · [Wikipedia: Strangler fig pattern](https://en.wikipedia.org/wiki/Strangler_fig_pattern) · [kilimchoi/engineering-blogs](https://github.com/kilimchoi/engineering-blogs)
