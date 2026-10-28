# SWE YouTube Series, Season 1: Building Haat

*System design through one app that keeps growing · Oct 28, 2026 · @Lion*

---

## How the series works

The first 20 videos are scenario-based system design. Something goes wrong ("we sold 12,000 phones but only had 10,000"), and you solve it on the whiteboard by comparing several approaches. There's no coding, just reasoning, diagrams and back-of-envelope numbers. AI topics start in Season 2, after Episode 20.

## One app, one story

The whole season follows **Haat**, a fictional Bangladeshi marketplace where small sellers (the kind who sell on Facebook pages today) open proper online shops. **You are Haat's first engineer.** The viewer grows with the company, from launch day on one server to a nationwide platform with its own delivery riders and live shopping.

Haat **never changes what it is**. Every new feature is something a real marketplace adds as it grows: seller uploads, search, flash sales, payments, its own delivery fleet, live shopping, shop feeds. Viewers always know which app they're looking at.

### The two rules of the season

1. **Every episode starts from the previous one.** The new problem is caused either by last episode's fix (the side effect) or by the growth that fix made possible (the success). Never by "and now we also do X".
2. **The numbers only go up.** Users, orders, products and team size grow every episode, and they never contradict each other. The Continuity Bible at the end is the source of truth.

### Two recurring on-screen devices

- **The Status Board** (top corner, first 30 seconds): month since launch, users, orders per day, engineers. Viewers feel the company growing.
- **The Architecture Diagram**: one diagram for the whole season. Each episode adds or changes one or two boxes, highlighted in a new colour. Open on last episode's diagram and end on the new one. By Episode 20 it's the full system, and that final diagram is your season thumbnail.

### The opening callback (30 seconds, every episode)

> "Last time, [problem] almost killed Haat, and we fixed it with [solution]. That worked. Maybe too well. Because today, [solution / its success] is exactly what causes our next disaster..."

Each episode below includes a **Callback** line written for this purpose.

---

## Every episode follows the same shape

1. **Callback + Status Board (30 s):** where Haat is now and what we fixed last time.
2. **The scenario (1-2 min):** a story with real stakes, like money lost, angry users or servers on fire.
3. **Requirements and numbers (2 min):** users, requests per second, data size. This is what makes it engineering, not guessing.
4. **Approaches A, B, C (main part):** start with the simplest idea, show exactly where it breaks, then improve. Each approach gets pros, cons and "when I'd choose this".
5. **The twist (2-3 min):** a follow-up problem that makes even the best approach struggle. This keeps viewers to the end.
6. **Verdict (1 min):** what Haat actually ships, and the honest "it depends".
7. **The cliffhanger (15 s):** a teaser of what this fix will cause next episode.

## The 4-day prep routine

- **Day 1: Learn** the concepts behind the approaches from the study list.
- **Day 2: Solve it yourself** on paper before looking at anyone else's solution. Write every approach you can think of, then compare with sources.
- **Day 3: Numbers and the twist.** Do the capacity math, stress-test each approach, and prepare answers for likely viewer questions.
- **Day 4: Script and rehearse.** Plan the whiteboard panel by panel, update the Architecture Diagram, write the callback and the hook, and do a full dry run.

## Core study resources for the whole season

- **Designing Data-Intensive Applications** by Martin Kleppmann: the deep "why" behind almost every episode.
- **System Design Interview Vol. 1 and 2** by Alex Xu: many of these scenarios in interview form.
- **Engineering blogs** (Stripe, Uber, Discord, Netflix, Shopify, Meta, Pathao, bKash when available): real-world stories to quote in your videos.
- **Release It!** by Michael Nygard: for the reliability episodes (14, 19, 20).

Read the relevant chapter for each episode rather than reading these cover to cover first.

---

## Season at a glance

| # | What breaks | Topic | Caused by |
|---|---|---|---|
| 1 | A viral post kills the server; adding a 2nd server logs users out | Load balancing, stateless servers | Launch day |
| 2 | TV ad, 20× traffic, the database melts | Caching, cache stampede | Ep 1's servers all hit one database |
| 3 | Sellers' photo/video uploads choke the servers and fail at 95% | Object storage, resumable uploads | Ep 2 made Haat fast, so sellers flooded in |
| 4 | 2M products and search takes 8 seconds | Search engines, autocomplete | Ep 3 filled the catalog |
| 5 | Bots scrape search and spam OTP login, leaving a huge SMS bill | Rate limiting | Ep 4's public search API, plus growth |
| 6 | 11.11 flash sale: 12,000 phones sold, only 10,000 in stock | Race conditions, locking, atomic counters | Ep 5's Redis counters become stock counters |
| 7 | A customer is charged twice for a reservation | Idempotency, payment state machine | Ep 6's reservations must be paid |
| 8 | "Sale starts in 10 min" push would take 3 days, then self-DDoS | Message queues, fan-out | Ep 6's sale needs promoting |
| 9 | Expired reservations released twice, so oversold again | Job scheduling, leases | Ep 6's timeouts run as cron on 2 servers |
| 10 | The orders table hits 300M rows and the database is at 100% | Replication, sharding, consistent hashing | Everything so far shares one database |
| 11 | Two customers get the same order number | Distributed ID generation, Base62 links | Ep 10's shards each auto-increment |
| 12 | "I placed an order but My Orders is empty" | Replication lag, consistency models | Ep 10's read replicas |
| 13 | Microservices: order paid, never shipped | Sagas, transactional outbox | Team outgrew the monolith |
| 14 | The courier partner's API slows down and takes down the whole app | Cascading failures, circuit breakers | Ep 13's chain of services |
| 15 | Haat launches its own riders: who is nearest? | Geospatial indexing | Ep 14: the courier can't be trusted |
| 16 | "Where is my rider?" Live tracking + chat | WebSockets, pub/sub, delivery receipts | Ep 15's riders stream locations |
| 17 | Celebrity seller goes live: 500K viewers, 100K likes/sec | Hot keys, sharded counters | Ep 16's WebSockets enabled live shopping |
| 18 | A 5M-follower shop posts and feeds lag for an hour | News feed fan-out, ranking | Ep 17 made following shops popular |
| 19 | 40 services deploy daily; a column rename breaks checkout | Zero-downtime deployment | Ep 13 onward: many services, many deploys |
| 20 | **Finale:** the Dhaka data center goes dark during the Eid sale | Multi-region, disaster recovery, CAP | Everything we built |

---

# Act 1: Launch (Episodes 1-5)

*Haat goes from 0 to 1.5 million users. One database, a few servers, and every lesson learned the hard way.*

## Episode 1: The Facebook post that killed our server

**Status Board:** Launch week · 20K users → 200K visitors in a day · 2 engineers (you and the CTO)
**Callback:** None; this is the origin story. Open on the single-server diagram: one box running the app and the database.
**Prep time:** 4-5 days (first episode, so add time for learning to record).

**Scenario:** Haat launches with 300 sellers. A popular Facebook page shares it, and 200,000 people arrive in two hours. The single server running both the app and the database hits 100% CPU and stops responding. You rush to add a second app server behind a load balancer, and it works. Then support messages pour in: "I keep getting logged out", "my cart is empty again".

**Numbers:** 200K visitors in 2 hours, peak ~150 requests/sec, ~10 page views per visit. How many requests can one server handle? Where's the first bottleneck: CPU, memory or database?

**Approaches to compare**
- **Vertical scaling:** buy a bigger server. Fastest fix, but it has a ceiling, a big bill and is still a single point of failure.
- **Separate the database from the app server:** the first and cheapest split; each can now scale on its own.
- **Horizontal scaling + load balancer:** many app servers behind one entry point. Compare round robin, least connections and IP hash.
- **Fixing the logouts:**
  - *Sticky sessions:* the load balancer always sends a user to the same server. Easy, but if that server dies, the user's session dies too.
  - *Shared session store (Redis):* every server reads sessions from one fast store. **Redis enters the series here.**
  - *Stateless tokens (JWT):* the session lives in a signed token on the client. No lookup needed, but hard to revoke.

**The twist:** The load balancer itself is now a single point of failure. And when you deploy, the servers restart one by one and users still lose their carts. Introduce health checks, redundant load balancers, and the rule of the season: **app servers must be stateless.**

**Verdict:** Separate the database, put 3 stateless app servers behind a load balancer, keep sessions in Redis.

**Diagram adds:** Load Balancer · App Servers ×3 · Redis (sessions) · Database.

**Cliffhanger:** "Three app servers, one database. Every page still runs 40 queries. What happens when traffic grows 20×?"

**Study:** vertical vs horizontal scaling, load balancing algorithms, health checks, stateless vs stateful services, sessions vs tokens.

**Prep**
- **Day 1:** Learn load balancing algorithms, sticky sessions, session stores and JWT basics.
- **Day 2:** Draw a user's requests bouncing between two servers and show exactly why the cart disappears.
- **Day 3:** Calculate requests/sec for the viral spike and how many servers you need. Design the redundant load balancer.
- **Day 4:** Script with the "Facebook share → server dies" hook. Draw the first version of the Architecture Diagram you'll reuse all season.

**Title:** Our App Went Viral and Crashed. Here's How We Scaled It (Load Balancers Explained)

**Description:**
Launch week, one server, and suddenly 200,000 visitors. In this first episode of the Haat series, we take a brand-new app from a single server to a scalable setup on the whiteboard: vertical vs horizontal scaling, load balancers, and the sneaky bug where users get logged out the moment you add a second server. This is the foundation for every episode that follows.

---

## Episode 2: The TV ad that melted our database

**Status Board:** Month 3 · 150K users · 3K orders/day · 3 engineers
**Callback:** "Last time we fixed our crash by adding app servers behind a load balancer. That worked, but every one of those servers still talks to the same single database."
**Prep time:** 4 days.

**Scenario:** Haat's homepage runs 40 database queries per visit: categories, banners, trending products, prices, seller ratings. Marketing books a TV ad during a cricket match. Traffic jumps 20×. You can add app servers in minutes, but they all hammer the same database, which hits 100% CPU. The site goes down during your biggest marketing moment.

**Numbers:** 20× traffic means ~3,000 requests/sec at 40 queries each, so 120,000 queries/sec against a database that handles maybe 5,000. What cache hit rate do you need to survive?

**Approaches to compare**
- **Cache-aside (lazy loading):** the app checks Redis first and loads from the database on a miss. Simple, but stale data and slow first requests.
- **Write-through:** update the cache on every write. Always fresh, but slower writes and cached data nobody reads.
- **Write-back:** write to the cache and sync to the database later. Very fast, but you can lose data.
- **CDN / full-page caching:** cache whole pages and images at the edge. Massive scale, but hard for personalized content like the cart.
- **Which data gets which strategy?** Category list vs product price vs cart vs stock count. This decision is the real lesson.

**The twist:** A cache stampede. The cached homepage expires at 9:00:00 PM, right as the ad airs, and 50,000 requests miss at once and hit the database together. Solutions: a lock so only one request rebuilds, early random refresh, and serving stale data while refreshing.

**Verdict:** CDN for images and static pages, cache-aside in Redis for product data with TTLs, never cache the cart or stock blindly, plus stampede protection.

**Diagram adds:** CDN · Redis now also holds the cache.

**Cliffhanger:** "The site is fast now. So fast that thousands of sellers want in, and every one of them wants to upload photos and videos."

**Study:** caching patterns, TTL and eviction (LRU, LFU), cache invalidation, thundering herd / cache stampede, CDNs.

**Prep**
- **Day 1:** Learn the caching patterns and eviction policies.
- **Day 2:** On paper, decide which strategy each piece of homepage data gets: banners, prices, stock, cart.
- **Day 3:** Study the stampede and its three fixes. Calculate hit rates and the load the cache removes.
- **Day 4:** Script with the "TV ad crash" hook.

**Title:** The TV Ad That Crashed Our Website: Caching Strategies Explained

**Description:**
Traffic jumps 20× and the database melts. Caching is the obvious answer, but which kind? In this episode of the Haat series we compare cache-aside, write-through, write-back and CDN caching, decide which data should use which strategy, and then fix the sneaky failure that takes down even well-cached systems: the cache stampede.

---

## Episode 3: Why every seller's upload failed at 95%

**Status Board:** Month 5 · 400K users · 5K sellers · 10K orders/day · 4 engineers
**Callback:** "Last time caching made Haat fast. Fast enough that 5,000 new sellers signed up in a month, and every one of them is uploading photos and videos through our app servers."
**Prep time:** 4 days.

**Scenario:** Sellers upload 10 photos and a 1-minute product video per listing, often over shaky mobile data from their shops. Uploads go through your app servers, which now spend most of their bandwidth and memory moving files instead of serving customers. A seller's 300 MB video fails at 95% and starts over from zero. Three times. They post an angry Facebook review.

**Numbers:** 5,000 sellers × 20 new listings/month × (10 photos × 3 MB + 300 MB video) = how many TB per month? How much of your app servers' bandwidth is that?

**Approaches to compare**
- **Upload through the app server:** simplest; fails on big files and steals capacity from customers.
- **Direct upload to object storage via pre-signed URL:** the file skips your servers entirely; the server only hands out a temporary permission.
- **Chunked, resumable uploads:** split the file into 5 MB pieces, upload them in parallel, and resume from the last successful chunk.
- **Background processing pipeline:** after upload, a message on a **queue** triggers thumbnails, image compression and video transcoding. **The queue enters the series here.**

**The twist:** Hundreds of resellers upload the *same* supplier photo for the same product. Should you store it 300 times? Introduce deduplication with content hashing, and its traps: the privacy questions, and what happens when one seller deletes a shared file.

**Verdict:** Pre-signed URLs + resumable chunked uploads to object storage, a queue-driven processing pipeline, CDN delivery, hash-based dedup for images.

**Diagram adds:** Object Storage · Upload/Processing Queue · Media Workers.

**Cliffhanger:** "Two million products in the catalog now. A customer types 'iphone cover', and waits eight seconds."

**Study:** object storage (S3-style), pre-signed URLs, multipart/resumable uploads, content hashing, async processing pipelines, queues (intro).

**Prep**
- **Day 1:** Learn object storage, pre-signed URLs and multipart uploads.
- **Day 2:** Design the full flow on paper: request upload → get URL → upload chunks → confirm → process → serve.
- **Day 3:** Calculate storage per month and bandwidth saved by direct uploads. Design dedup.
- **Day 4:** Script with the "failed at 95%" hook.

**Title:** Why Your Upload Fails at 95%: Designing Resumable Large File Uploads

**Description:**
Haat's sellers are uploading thousands of photos and videos over mobile internet, and our servers are drowning. In this episode we redesign uploads on the whiteboard: why routing files through your servers is a mistake, how pre-signed URLs and chunked uploads make uploads fast and resumable, how a queue handles processing, and how deduplication saves huge amounts of storage.

---

## Episode 4: The search box that took 8 seconds

**Status Board:** Month 7 · 1M users · 2M products · 25K orders/day · 6 engineers
**Callback:** "Last time we fixed uploads, and sellers listed two million products. Now customers can't find any of them."
**Prep time:** 4-5 days.

**Scenario:** Search runs `WHERE name LIKE '%iphone cover%'` on the products table. With 2 million products it takes 8 seconds, and it still misses "iPhone 15 Back Case" and anything typed in Banglish ("mobail kabhar"). Customers who can't find a product can't buy it, and conversion from search drops by half.

**Numbers:** 1M users, ~5M searches a day, peak ~300 searches/sec, plus autocomplete on every keystroke (×5). 2M products × ~1 KB of searchable text.

**Approaches to compare**
- **SQL LIKE:** show why `%...%` can't use an index and scans every row.
- **Database full-text index:** a real improvement and no new infrastructure, but limited ranking, typo handling and scale.
- **Dedicated search engine with an inverted index (Elasticsearch-style):** word → list of products. Fast, ranks by relevance (TF-IDF/BM25 at an intuitive level), handles filters like price and category.
- **Autocomplete:** a trie with precomputed top-k suggestions at each node, or a prefix → suggestions table in Redis, rebuilt periodically.

**The twist:** Two problems. First, the search index is a *copy* of the database: when a seller changes a price, how does the index find out? Compare dual writes vs change data capture (CDC) through the queue from Episode 3. Second, typos and Banglish: "iphon" and "mobail" should still work (fuzzy matching, synonyms). Tease: "Searching by *meaning* instead of words is a Season 2 topic."

**Verdict:** A search engine fed by CDC events, a trie/Redis autocomplete with top-k, fuzzy matching and a synonym list.

**Diagram adds:** Search Cluster · CDC stream from the database · Autocomplete store.

**Cliffhanger:** "Search is fast and public now. Within a week, our logs show someone calling it 3 million times a night. And it's not customers."

**Study:** inverted indexes, full-text search, relevance ranking basics, tries and top-k, fuzzy matching, keeping derived data in sync (CDC).

**Prep**
- **Day 1:** Learn inverted indexes and tries with top-k.
- **Day 2:** Build a tiny inverted index and a tiny trie on paper from 10 product names.
- **Day 3:** Calculate index size and autocomplete QPS. Design index syncing and typo handling.
- **Day 4:** Script with a live typing demo as the hook.

**Title:** Why Our Search Took 8 Seconds, and How Search Engines Really Work

**Description:**
Two million products and a search box that takes 8 seconds and still can't find "iphone cover". In this episode of the Haat series we design product search on the whiteboard: why SQL LIKE fails, how an inverted index works, how autocomplete answers in under 100 ms, and how to keep the search index in sync with your database.

---

## Episode 5: A bot cost us lakhs in OTP SMS overnight

**Status Board:** Month 9 · 1.5M users · 40K orders/day · 7 engineers
**Callback:** "Last time we built fast public search. Bots love fast public APIs. So do people who want to burn your SMS budget."
**Prep time:** 4 days.

**Scenario:** Two attacks in one night. A competitor's scraper hits your search API 3 million times to copy every price. Worse, a bot hits the "Send OTP" login endpoint 2 million times. Every SMS costs money. You wake up to a huge SMS bill, a slow site, and real users who can't log in.

**Numbers:** 2M OTP requests × ৳0.30-0.50 per SMS. (Pick one figure and use it in the title.) What's a reasonable limit for a real human?

**Approaches to compare**
- **Fixed window counter:** "max 5 per minute". Simple, but allows bursts at window edges (5 at 0:59, 5 at 1:00).
- **Sliding window log:** store every request timestamp. Accurate, but memory heavy.
- **Sliding window counter:** a weighted blend of two windows. Good accuracy, low memory.
- **Token bucket / leaky bucket:** allow short bursts but control the average rate.

All of them live in **Redis**, the same Redis that already holds sessions and cache, now counting requests with atomic `INCR`.

**The twist:** You run 6 app servers. If each counts on its own, the attacker gets 6× the limit, so the counter must be shared and atomic. And the attacker rotates phone numbers and IPs. What do you limit on: phone, IP, device, account, or all of them in layers? Add CAPTCHA after N attempts and exponential backoff on OTP resend.

**Verdict:** Token bucket per API key for search, sliding window counters on several keys for OTP, all in Redis, plus CAPTCHA and resend backoff.

**Diagram adds:** Rate Limiter (at the gateway, backed by Redis).

**Cliffhanger:** "Atomic counters in Redis. Remember that. Because on 11.11, we're going to need a counter that never goes below zero, and a million people are going to test it at midnight."

**Study:** rate limiting algorithms, distributed counters with Redis, atomic operations, layered defenses.

**Prep**
- **Day 1:** Learn all four algorithms and draw each one on a timeline.
- **Day 2:** Run the same burst of requests through each algorithm on paper.
- **Day 3:** Design the distributed, multi-key version. Calculate the money saved.
- **Day 4:** Script with the "surprise SMS bill" story as the hook.

**Title:** A Bot Cost Us Lakhs in OTP SMS Overnight: Designing a Rate Limiter

**Description:**
Every OTP costs money, and bots know it. In this episode of the Haat series we design a rate limiter on the whiteboard, comparing fixed windows, sliding windows, token buckets and leaky buckets against the same attack traffic, then make it work across many servers against an attacker who keeps changing phone numbers and IPs.

---

# Act 2: The Big Sale (Episodes 6-9)

*Haat runs its first 11.11 mega sale. Four episodes, one event, and everything that goes wrong around it.*

## Episode 6: We sold 12,000 phones but only had 10,000

**Status Board:** Month 12 · 3M users · 11.11 sale night · 9 engineers
**Callback:** "Last time we used atomic counters in Redis to stop bots. Tonight, 600,000 people are about to tap Buy at the same second."
**Prep time:** 4-5 days.

**Scenario:** Haat's first 11.11 sale: 10,000 phones at 40% off, starting at midnight. 600,000 users are waiting with their thumbs on the button. When the dust settles, 12,000 orders were confirmed. You oversold by 2,000 and now have to cancel 2,000 orders on the biggest night of the year.

**Numbers:** 600K users tapping within ~10 seconds is ~60,000 requests/sec at midnight vs a normal peak of ~500. Stock: 10,000.

**Approaches to compare**
- **Read stock, then update (the bug):** two requests both read "1 left" and both succeed. Draw the race condition on a timeline.
- **Pessimistic locking (SELECT ... FOR UPDATE):** correct, but requests queue on one row and the database chokes.
- **Optimistic locking (version column):** no locks, but under heavy contention almost every request fails and retries.
- **Atomic counter in Redis:** `DECR` in memory and reject below zero, then write orders to the database through the queue.
- **Virtual waiting room:** let users in at a controlled rate with a queue position. Protects everything behind it, but users wait.

**The twist:** A user reserves a phone but never pays. Do you hold stock forever? Design reservations with a 10-minute payment timeout that returns stock to the pool without overselling when it comes back.

**Verdict:** Waiting room in front, atomic Redis counter for stock, orders persisted through the queue, reservations with a timeout.

**Diagram adds:** Waiting Room · Stock Counter (Redis) · Order Queue.

**Cliffhanger:** "The reservation is held. Now the customer has to pay. On a slow network. On a phone that retries automatically."

**Study:** race conditions, database locking, transactions and isolation levels, Redis atomic operations, admission control.

**Prep**
- **Day 1:** Learn race conditions, isolation levels, and pessimistic vs optimistic locking.
- **Day 2:** Draw two requests overselling the last phone, then show how each approach prevents it.
- **Day 3:** Calculate midnight requests/sec. Design the waiting room and reservation timeout.
- **Day 4:** Script with the oversold-orders hook.

**Title:** We Sold 12,000 Phones But Only Had 10,000: Designing a Flash Sale System

**Description:**
600,000 people tap "Buy" at midnight and the stock count goes negative. In this episode of the Haat series we design a flash sale system on the whiteboard, starting with the race condition that causes overselling, then comparing pessimistic locks, optimistic locks, Redis atomic counters and virtual waiting rooms. Finally we handle the customer who reserves a phone and never pays.

---

## Episode 7: The customer who got charged twice

**Status Board:** Month 12 · 3.5M users · 11.11 sale, minute 3 · 9 engineers
**Callback:** "Last time we stopped overselling with reservations. Now 10,000 people need to pay for those reservations, exactly once."
**Prep time:** 4-5 days.

**Scenario:** A customer taps "Pay ৳18,000" for their sale phone. The network is slow, so they tap again. Or the app retries automatically after a timeout. They're charged twice, and they post screenshots everywhere. On sale night, hundreds of customers have the same story.

**Numbers:** 10,000 payments in ~10 minutes. The payment gateway's typical latency, timeout values, and how often mobile networks drop requests in practice.

**Approaches to compare**
- **Disable the button on the frontend:** helps, but doesn't stop app retries, network retries or a second device.
- **Idempotency keys:** the client sends a unique key per payment attempt; the server stores the result by key and returns the same result on repeats.
- **Database unique constraint** on the order's payment: the database itself refuses duplicates.
- **Payment state machine:** CREATED → PROCESSING → SUCCESS / FAILED, with only valid transitions allowed.

**The twist:** You call the bank and the request times out. Did the money move? You can't retry blindly and can't assume failure. Introduce status queries, reconciliation jobs, and why "exactly once" is really "**at least once + idempotency**". Keep that phrase, because it returns in Episodes 8, 9 and 13.

**Verdict:** Idempotency keys + a payment state machine + a unique constraint as the last line of defense + nightly reconciliation.

**Diagram adds:** Payment component with an Idempotency Store · Reconciliation Job · external Payment Gateway.

**Cliffhanger:** "Payments are safe. But next year's sale needs promoting: one push notification to five million phones. Our first version would finish in three days."

**Study:** idempotency, retries, distributed timeouts, payment state machines, reconciliation. Stripe's engineering blog post on idempotency keys is a must-read.

**Prep**
- **Day 1:** Learn idempotency and read the Stripe post.
- **Day 2:** Draw the payment flow and mark every point where a duplicate can happen.
- **Day 3:** Design the timeout scenario and reconciliation. Prepare the "exactly once is a myth" explanation.
- **Day 4:** Script with a relatable mobile-banking "did it go through?" hook.

**Title:** Why You Sometimes Get Charged Twice (and How Payment Systems Prevent It)

**Description:**
A slow network, a second tap, an automatic retry, and a customer pays twice. In this episode of the Haat series we design a payment flow that can't double-charge, comparing frontend fixes, idempotency keys, database constraints and state machines. Then we face the scariest case in payments: a timeout where you don't know whether the money moved.

---

## Episode 8: Sending a push notification to 5 million phones

**Status Board:** Month 15 · 5M users · planning the Eid sale · 12 engineers
**Callback:** "Last time we made payments safe with idempotency. Today we need that idea again, because we're about to send five million messages, and some of them will be sent twice."
**Prep time:** 4 days.

**Scenario:** Marketing wants "Eid sale starts in 10 minutes!" sent to all 5 million users at exactly 8:50 PM. The first attempt is a loop that calls the push provider one user at a time. You do the math: it would finish on Thursday.

**Numbers:** 5M × 50 ms per call = ~70 hours. To finish in 5 minutes you need ~17,000 sends/sec. What are the push provider's rate limits?

**Approaches to compare**
- **Simple loop:** show the math for why it fails.
- **Message queue + worker pool:** split users into batches of 1,000, put the batches on the queue (the one from Episode 3, now doing serious work), and let many workers send in parallel.
- **Fan-out by segments:** precompute segments (region, language, app version, past buyers) and send per segment.
- **Priority queues:** OTPs (Episode 5) and payment confirmations (Episode 7) must never wait behind a marketing blast.

**The twist:** Workers crash mid-batch and retry, so some users get the greeting 3 times while others get nothing (idempotency from Episode 7 fixes it). Then the push lands, 5 million people open Haat at the same moment, and you DDoS yourself. Fix it with staggered sends and the waiting room from Episode 6.

**Verdict:** Queue + worker pool + separate priority lanes + per-user idempotency + staggered delivery.

**Diagram adds:** Notification Service · Priority Queues · Push Workers.

**Cliffhanger:** "Notifications at the right *moment* turn out to be the easy part. Do you remember those 10-minute reservations from 11.11? The cron job that expires them just ran on two servers at once."

**Study:** message queues (Kafka and RabbitMQ concepts), worker pools, at-least-once delivery, deduplication, priority, push provider limits, staggered sends.

**Prep**
- **Day 1:** Learn message queues and delivery guarantees.
- **Day 2:** Design the pipeline on paper: who reads users, who batches, who sends, who retries.
- **Day 3:** Do the throughput math for 5 minutes. Solve duplicates and the self-inflicted spike.
- **Day 4:** Script with the "the loop finishes Thursday" hook.

**Title:** How to Send 5 Million Notifications in 5 Minutes (Without DDoSing Yourself)

**Description:**
Sending one notification is one API call. Sending five million at the same moment is a distributed systems problem. In this episode of the Haat series we design a notification system: queues, worker pools, segmentation and priority lanes. Then we fix the duplicates that crashed workers cause, and the traffic spike when millions of users open the app at once.

---

## Episode 9: The cron job that oversold us again

**Status Board:** Month 15 · 5M users · Eid sale night · 12 engineers
**Callback:** "In Episode 6 we released unpaid reservations after 10 minutes. We did it with a cron job. We now have six servers. All six run that cron."
**Prep time:** 4 days.

**Scenario:** During the Eid sale, a cron job runs every minute: "find reservations older than 10 minutes and return their stock". It runs on several servers at once, so the same expired reservation returns its phone to stock *twice*, the counter goes up by 2, and Haat oversells again, this time by a slower route. Meanwhile, "Your cart is waiting" reminders arrive at 3 AM because the job queue fell hours behind.

**Numbers:** 500K reservations during the sale, each with its own expiry time, plus 2M scheduled reminders a day. How long does a full table scan every minute take?

**Approaches to compare**
- **Cron job scanning the database every minute:** simple, but slow scans, and multiple servers mean duplicates. "Only run cron on one server" is a single point of failure.
- **Delay queue:** schedule a message to become visible at an exact time.
- **Redis sorted set scored by due time:** workers poll for items due now. (Sorted sets come back in Episode 15 for geo.)
- **Timing wheel:** an efficient in-memory structure for millions of timers.

**The twist:** A worker claims a job at 8:59:59 and crashes before finishing. Lost or done twice? Design leases (claim a job for N seconds; if not completed, it becomes visible again), idempotent job handlers ("release reservation #123" is safe to run twice), and a distributed lock only where truly needed.

**Verdict:** Delay queue or sorted set for scheduling, leases for crash safety, idempotent handlers so duplicates are harmless.

**Diagram adds:** Scheduler (sorted set / delay queue) · Scheduler Workers.

**Cliffhanger:** "Act 2 is over. We survived two sales. But the orders table now has 300 million rows, and the database is at 100% CPU on a normal Tuesday."

**Study:** job scheduling, cron limitations, delay queues, sorted sets, timing wheels, leases, distributed locks, at-least-once processing.

**Prep**
- **Day 1:** Learn cron limitations and delay queues. Read about timing wheels.
- **Day 2:** Design each approach on paper with 10 sample reservations.
- **Day 3:** Walk through crash scenarios at each step. Design leases and idempotent handlers.
- **Day 4:** Script with the "counter went UP by 2" reveal as the hook.

**Title:** The Cron Job That Oversold Our Sale: Designing a Distributed Job Scheduler

**Description:**
Running a task at an exact time sounds trivial, until you have millions of tasks and servers that crash. In this episode of the Haat series a humble cron job breaks our sale, and we redesign it on the whiteboard, comparing cron, delay queues, Redis sorted sets and timing wheels, then use leases and idempotency so jobs are never lost and never harmful when repeated.

---

# Act 3: Growing Up (Episodes 10-14)

*Haat is a serious company now. The database, the codebase and the team all outgrow their first versions.*

## Episode 10: The database can't keep up anymore

**Status Board:** Month 20 · 8M users · 300M order rows · 150K orders/day · 25 engineers
**Callback:** "For nine episodes, every feature we built (sessions aside) has leaned on one database. Today the CTO asks: what do we do, and in what order?"
**Prep time:** 5-6 days.

**Scenario:** Haat's PostgreSQL primary is at 100% CPU and 90% disk. Users are growing 15% a month. The orders table has 300 million rows. The CTO wants a plan with costs, in order.

**Numbers:** Reads vs writes ratio (~20:1), data growth per month, disk runway at the current rate.

**Approaches to compare (as a staircase, cheapest first)**
- **Vertical scaling:** a bigger machine. Fast and easy, but has a ceiling and a big bill.
- **Query optimization and indexes:** often the biggest win for the least cost.
- **Caching:** callback to Episode 2; what's already offloaded and what isn't.
- **Read replicas:** send reads to copies. Great for read-heavy apps, but replicas lag. (Plant this; it explodes in Episode 12.)
- **Sharding:** split orders across many databases by range, by hash or by a lookup directory. Each has different trade-offs.

**The twist:** You shard by `user_id % 4`. A year later you need 5 shards, and ~80% of rows must move. Introduce consistent hashing. Then two more problems: one mega-seller's shard is 10× hotter than the rest (hot shard), and the seller dashboard needs orders across *all* shards (cross-shard queries).

**Verdict:** Climb the staircase in order; shard orders by user with consistent hashing; keep a separate path for seller analytics.

**Diagram adds:** Read Replicas · Orders Shards ×4 · Shard Router.

**Cliffhanger:** "Each shard has its own auto-increment ID. Shard 1 just created order #100234. So did shard 3."

**Study:** vertical vs horizontal scaling, indexing, replication, sharding strategies, consistent hashing, hot spots, cross-shard queries.

**Prep**
- **Day 1:** Learn replication and sharding strategies (DDIA chapters 5-6).
- **Day 2:** Shard 12 sample users on paper with each strategy; add a 5th shard and count how many move.
- **Day 3:** Learn consistent hashing and redo the exercise. Prepare the staircase diagram.
- **Day 4:** Script with the "CTO's question" hook.

**Title:** Your Database Is at 100%: What to Do First (Scaling Explained Step by Step)

**Description:**
Haat is growing fast and the database can't keep up. Do you buy a bigger server, add replicas or shard everything? In this episode we walk the scaling staircase on the whiteboard, from the cheapest fixes to sharding, compare range, hash and directory-based sharding, and see why consistent hashing saves you when it's time to add more machines.

---

## Episode 11: Two customers got the same order number

**Status Board:** Month 21 · 8M users · 4 order shards · 160K orders/day · 25 engineers
**Callback:** "Last time we sharded the orders table. We forgot one tiny thing: every shard was still counting 1, 2, 3 on its own."
**Prep time:** 4 days.

**Scenario:** Two customers both receive "Order #100234 confirmed". One calls support, the agent looks up #100234 and reads the *other* customer's address and phone number aloud. Now it's a privacy incident. Separately, sellers have been sharing product links like `haat.com.bd/product/48213`, and a competitor is reading your daily sales volume from how fast the IDs grow.

**Numbers:** 160K orders/day today, 10× in 3 years. How many IDs per second at peak sales? How long must IDs stay unique (forever)?

**Approaches to compare**
- **Central ticket server:** one database hands out IDs. Simple and ordered, but a bottleneck and single point of failure.
- **Auto-increment with offsets:** shard k starts at k and steps by N. Works until you change N.
- **UUID:** no coordination, but 128 bits, unordered (bad for indexes), and ugly for humans.
- **Snowflake-style IDs:** timestamp + machine ID + sequence. Roughly time-ordered, no coordination, 64 bits.
- **Pre-generated ID ranges:** each server grabs a block of 1,000 IDs at a time.

**The twist:** Short, shareable links. Sellers want `haat.com.bd/p/x7Kq2` to paste on Facebook, plus custom names like `haat.com.bd/rahim-fashion`. Introduce Base62 encoding, why sequential short codes leak your business numbers (the German tank problem), and how custom aliases coexist with generated codes. (This is where the original URL shortener topic lives now.)

**Verdict:** Snowflake IDs internally; non-sequential Base62 codes for anything public; a reserved namespace for custom shop links.

**Diagram adds:** ID Generator (library in each service) · Short Link Service.

**Cliffhanger:** "Orders are unique now. But a customer just placed one, opened My Orders, and it's empty. They placed it again."

**Study:** ID generation strategies, Snowflake, UUID versions, Base62, information leakage through IDs, 301 vs 302 redirects.

**Prep**
- **Day 1:** Learn Snowflake's bit layout, UUIDs and Base62.
- **Day 2:** On paper, list every approach and the exact way it fails.
- **Day 3:** Calculate: Snowflake capacity per millisecond, how many 7-character Base62 codes exist, and how fast you'd run out.
- **Day 4:** Script with the "support read out the wrong customer's address" hook.

**Title:** Two Customers Got the Same Order Number: Designing Unique IDs at Scale

**Description:**
After sharding our database, Haat started handing out the same order number twice. In this episode we compare ticket servers, offset auto-increment, UUIDs, Snowflake IDs and pre-allocated ranges on the whiteboard, then design short shareable links and see how sequential IDs can leak your sales numbers to competitors.

---

## Episode 12: "I placed an order but My Orders is empty"

**Status Board:** Month 22 · 10M users · 200K orders/day · 30 engineers
**Callback:** "In Episode 10 we added read replicas to take load off the primary. Replicas are copies, and copies are always a little bit behind."
**Prep time:** 5-6 days.

**Scenario:** A customer pays, sees "Order confirmed", taps My Orders, and the list is empty. They panic and order again (and thanks to Episode 7, that's a *different* order, not a retry). Support tickets pile up. A seller changes a price, but checkout charges the old price for 30 seconds. The cause: reads go to replicas and caches that haven't caught up.

**Numbers:** Typical replication lag (100 ms - 2 s), spiking to 30 s+ during sales. Which reads must be fresh, and how many reads/sec would hit the primary if they all were?

**Approaches to compare**
- **Strong consistency: all reads from the primary.** Always correct, but you lose the scaling benefit of Episode 10.
- **Eventual consistency:** accept short staleness. Fine for review counts, wrong for My Orders.
- **Read-your-own-writes:** after a user writes, route *their* reads to the primary for a few seconds, or until the replica has caught up to their write's position.
- **Monotonic reads / sticky replica:** keep a user on the same replica so they never go back in time.
- **Cache invalidation on write and versioned keys:** so caches can't serve old prices.

**The twist:** Sort Haat's data into tiers: what must be strongly consistent (stock, payment status, price at checkout), what needs read-your-writes (my orders, my cart, my profile), and what can be eventual (reviews, view counts, recommendations). Then introduce the idea that will return in the finale: **you can't have perfect freshness and perfect speed everywhere at once.**

**Verdict:** Read-your-own-writes for the user's own data, primary reads for money and stock, eventual consistency everywhere else.

**Diagram adds:** Consistency-aware read routing in the data layer.

**Cliffhanger:** "Thirty engineers, one codebase, one deploy pipeline. Every release is a fight. The CTO says the words every engineer has heard: 'Let's split it into microservices.'"

**Study:** replication lag, consistency models, read-your-writes, monotonic reads, cache invalidation, CAP/PACELC intuition (light).

**Prep**
- **Day 1:** Learn replication lag and consistency models (DDIA chapter 5).
- **Day 2:** Draw the empty-My-Orders bug as a timeline across primary, replica and cache.
- **Day 3:** Build the data-tier table for Haat. Calculate primary load for each routing strategy.
- **Day 4:** Script with the "order vanished, so they ordered again" hook.

**Title:** "I Placed an Order But It's Not There": Consistency in Distributed Databases

**Description:**
You buy something, open your orders, and it's not there. It's not a bug in the code; it's replication lag. In this episode of the Haat series we recreate the bug on the whiteboard and compare the fixes: strong consistency, eventual consistency, read-your-own-writes and monotonic reads, then decide which of Haat's data needs which guarantee.

---

## Episode 13: Paid, but never shipped

**Status Board:** Month 26 · 12M users · 300K orders/day · 80 engineers · 8 services
**Callback:** "Last time, our engineering team outgrew the monolith. So we split Haat into services: Order, Payment, Inventory, Seller, Shipping. Each with its own database. What could go wrong?"
**Prep time:** 5-6 days.

**Scenario:** A customer pays ৳3,200. The Payment service succeeds. The Inventory service then crashes before reserving the item. The money is taken, nothing is reserved, nothing ships, and there's no single database transaction that can roll it all back. Hundreds of these orders sit in limbo every week.

**Numbers:** 300K orders/day across 4-5 service hops each. Even a 0.1% failure rate per hop means hundreds of broken orders per day.

**Approaches to compare**
- **One big database transaction:** impossible across separate services.
- **Two-phase commit (2PC):** a coordinator asks everyone to prepare, then commit. Correct, but slow, and everyone blocks if the coordinator dies.
- **Saga with choreography:** each service reacts to events on the queue and publishes its own; on failure, services run compensating actions (refund, release stock). Decoupled, but the flow is hard to follow.
- **Saga with orchestration:** one Order Orchestrator tells each service what to do and handles failures. Easier to reason about, but it's a central piece.

**The twist:** The Order service saves the order and crashes *before* publishing "OrderCreated". The database and the event stream now disagree. Introduce the **transactional outbox**. Then: what if the refund itself fails? Compensations must be retried, idempotent (Episode 7 again) and eventually escalated to a human.

**Verdict:** Orchestrated saga for checkout, outbox for every event, idempotent compensations.

**Diagram adds:** The monolith box becomes Order / Payment / Inventory / Seller / Shipping services · Order Orchestrator · Outbox tables · Event Bus.

**Cliffhanger:** "Our checkout now calls five services in a row. One of them belongs to our courier partner. It's about to get very, very slow."

**Study:** distributed transactions, 2PC, saga pattern (choreography vs orchestration), compensating transactions, outbox pattern, event-driven architecture.

**Prep**
- **Day 1:** Learn 2PC and the saga pattern.
- **Day 2:** Draw Haat's checkout as both a choreography saga and an orchestration saga, including compensation.
- **Day 3:** Learn the outbox pattern. Walk through every crash point in the flow.
- **Day 4:** Script with the "paid but nothing happened" hook.

**Title:** Paid But Never Shipped: Handling Transactions Across Microservices (Saga Pattern)

**Description:**
In a monolith, one database transaction keeps everything consistent. In microservices, a payment can succeed while the next step fails. In this episode of the Haat series we design checkout across services and compare two-phase commit, choreography sagas and orchestration sagas, then fix the subtle bug where your database and your events disagree using the transactional outbox.

---

## Episode 14: One slow courier API took down all of Haat

**Status Board:** Month 28 · 14M users · Pohela Boishakh sale · 90 engineers
**Callback:** "Last time we made checkout a saga across services. One of those calls goes to our courier partner for a delivery quote. Today, that partner is having a bad day."
**Prep time:** 5-6 days.

**Scenario:** Mid-sale, the courier partner's API starts answering in 30 seconds instead of 300 milliseconds. Checkout threads wait. The thread pools fill. Then search, product pages and even the homepage, none of which need the courier, stop responding. Within six minutes all of Haat is down, because of someone else's server.

**Numbers:** 200 threads per instance ÷ 30 s per call means each instance handles ~7 requests/sec instead of thousands. Draw it filling up minute by minute.

**Approaches to compare**
- **Timeouts:** stop waiting forever. But what's the right value, and why is "30 seconds" almost always wrong?
- **Retries with exponential backoff and jitter:** retry smartly, because naive retries multiply the load on a struggling service.
- **Circuit breaker:** after repeated failures, stop calling for a while and fail fast. Closed → Open → Half-open.
- **Bulkheads:** separate thread/connection pools per dependency, so the courier can only exhaust *its* pool.
- **Graceful degradation:** show a flat estimated delivery fee and confirm later instead of failing checkout.

**The twist:** The courier recovers, and every client retries at the same instant (a retry storm) and knocks it down again. Why jitter matters, and how load shedding (Episode 5's rate limiting, now protecting a service) gives it room to recover.

**Verdict:** Tight timeouts + backoff with jitter + circuit breaker + bulkheads + a degraded fallback for every non-critical dependency.

**Diagram adds:** Circuit breakers and bulkheads around every external call.

**Cliffhanger:** "After the incident, the CEO makes a decision: 'We will never depend on someone else's riders again.' Haat is launching its own delivery fleet. Fifty thousand riders. Who gets which order?"

**Study:** cascading failures, thread/connection pool exhaustion, timeouts, retry strategies, circuit breakers, bulkheads, load shedding. Read about Netflix's resilience engineering and *Release It!*.

**Prep**
- **Day 1:** Learn how cascading failures happen through resource exhaustion.
- **Day 2:** Draw the failure spreading through Haat's services minute by minute, then add each defense.
- **Day 3:** Study retry storms and jitter. Draw the circuit breaker state machine.
- **Day 4:** Script it as a minute-by-minute incident report.

**Title:** How One Slow API Took Down Our Entire App (and the Patterns That Prevent It)

**Description:**
A single slow dependency can bring down a whole system, even the parts that don't use it. In this episode of the Haat series we replay a cascading failure minute by minute on the whiteboard, then add the defenses one by one: timeouts, retries with backoff and jitter, circuit breakers, bulkheads and graceful degradation.

---

# Act 4: Haat Comes Alive (Episodes 15-18)

*Haat stops being just a website. It gets its own riders, real-time tracking, live shopping and a social feed.*

## Episode 15: Finding the nearest rider among 50,000

**Status Board:** Month 31 · 16M users · 50K riders · 400K orders/day · 110 engineers
**Callback:** "After Episode 14's outage, Haat launched its own delivery fleet. Fifty thousand riders across Dhaka and Chattogram. Our first rider-assignment code just took 4 seconds per order."
**Prep time:** 4-5 days.

**Scenario:** When a seller marks an order ready, Haat must find the nearest available rider in under a second. Riders send their location every 4 seconds. Version 1 calculates the distance to all 50,000 riders for every order. At lunch rush it falls minutes behind, and food-grade orders (a new category) arrive cold.

**Numbers:** 50K riders ÷ 4 s = 12,500 location writes/sec. 400K orders/day with a lunch peak of ~30 assignments/sec, each needing a "nearest" query.

**Approaches to compare**
- **Brute force distance:** show the math for why it collapses.
- **Fixed grid:** split the city into squares and search your square and its neighbors. Simple, but Gulshan at rush hour and an empty field get the same square size.
- **Quadtree:** split busy areas into smaller squares. Adapts to density, but costly to update as riders move.
- **Geohash:** encode location as a string where nearby places share a prefix. Easy to store in Redis or any database (Redis has built-in geo commands on sorted sets, a callback to Episode 9).
- **Hexagon grid (Uber's H3):** all neighbors are equally distant, which helps for zones and surge.

**The twist:** Two problems. First, write-heavy: 12,500 updates/sec. Indexes that are fast to read are often slow to update, so which survives constant movement? Second, straight-line nearest isn't road nearest. A rider 400 m away across the Buriganga or a one-way road can be 25 minutes away. Introduce candidate filtering by distance, then ranking by ETA.

**Verdict:** In-memory geohash/H3 index for candidates, ETA-based ranking on the top 20, location writes kept out of the main database.

**Diagram adds:** Rider Location Service (in-memory geo index) · Assignment Service.

**Cliffhanger:** "We know where every rider is, every four seconds. Customers want to know too: 'Where is my rider?' Right now the app refreshes the page."

**Study:** spatial indexing, grids, quadtrees, geohash and its edge problem, H3, in-memory location stores.

**Prep**
- **Day 1:** Learn grids, quadtrees and geohash. Read Uber's engineering post introducing H3.
- **Day 2:** Encode a real Dhaka location as a geohash by hand. Draw a quadtree for 15 riders.
- **Day 3:** Calculate update volume. Compare structures for read speed vs update cost.
- **Day 4:** Use a real Dhaka map as the whiteboard background. Script and rehearse.

**Title:** How Delivery Apps Find the Nearest Rider in Milliseconds

**Description:**
Finding the nearest rider sounds easy, until there are 50,000 of them moving every few seconds. In this episode of the Haat series we compare brute force, fixed grids, quadtrees, geohash and hexagon grids on a real Dhaka map, and discover why the best structure for searching isn't always the best for constant updates, and why "nearest" isn't the same as "fastest".

---

## Episode 16: "Where is my rider?"

**Status Board:** Month 34 · 18M users · 500K orders/day · 130 engineers
**Callback:** "Last time we built a live map of 50,000 riders inside our servers. Now customers want to watch their rider move, and send a message: 'I'm at the blue gate.'"
**Prep time:** 5-6 days.

**Scenario:** Customers refresh the tracking page every few seconds. The rider dot jumps around or freezes. Messages between customer and rider arrive late and out of order ("ok coming" before "which gate?"), and the "delivered" tick lies. Ratings drop.

**Numbers:** 150K orders in delivery at peak, each with a watching customer, so 150K open connections receiving a location every 4 s. Chat volume per order. Connections per server.

**Approaches to compare**
- **Short polling:** "any update?" every 2 seconds. Simple, but wastes battery and server capacity.
- **Long polling:** the server holds the request until there's news. Better, but still clumsy.
- **Server-Sent Events:** one-way stream from server to client. Perfect for location updates.
- **WebSockets:** one persistent two-way connection. Real-time chat, but now you track hundreds of thousands of open connections across servers.
- **Push notifications** (Episode 8) for users who closed the app.

**The twist:** The customer is connected to gateway server 3; the rider's messages arrive at server 7. How does a message cross between them? Introduce pub/sub and a connection registry. Then: message ordering with per-conversation sequence numbers, and the meaning of the ticks (sent → delivered → read) as acks.

**Verdict:** WebSockets for chat, SSE or WebSocket for tracking, pub/sub to route between gateway servers, sequence numbers for order, push for offline.

**Diagram adds:** Real-time Gateway (WebSocket servers) · Pub/Sub · Connection Registry · Chat Store.

**Cliffhanger:** "We built a real-time pipe to every customer's phone. A seller looked at it and asked: 'Can I go *live*?'"

**Study:** polling vs long polling vs SSE vs WebSockets, connection management, pub/sub, message ordering, delivery receipts, offline delivery.

**Prep**
- **Day 1:** Learn the four connection methods and draw the request flow for each.
- **Day 2:** Design the message flow: send → gateway → store → route → deliver → ack → ticks.
- **Day 3:** Solve cross-server routing and ordering. Estimate connections per server and servers needed.
- **Day 4:** Script with the "frozen rider dot + one tick" frustration as the hook.

**Title:** How Apps Show Your Rider Moving Live (and What Those Message Ticks Really Mean)

**Description:**
The dot moving on the map and the ticks next to your message hide a serious distributed system. In this episode of the Haat series we design live tracking and chat on the whiteboard, comparing polling, long polling, Server-Sent Events and WebSockets, then solve the hard parts: routing messages between servers, keeping them in order, and making the ticks tell the truth.

---

## Episode 17: 100,000 likes per second on one live stream

**Status Board:** Month 37 · 22M users · launch of Haat Live · 150 engineers
**Callback:** "Last time we built a real-time pipe to every phone. Sellers turned it into live shopping, and tonight a celebrity seller goes live to half a million people."
**Prep time:** 4-5 days.

**Scenario:** A famous actor's clothing brand goes live on Haat. 500,000 viewers. Hearts fly at 100,000 per second, and every one runs `UPDATE streams SET likes = likes + 1`. That single row becomes the hottest spot in the system and slows down checkout for everyone, right as the actor says "first 1,000 orders get 50% off".

**Numbers:** 100K likes/sec on one row; a single row lock manages maybe a few thousand updates/sec. 500K viewers × comments fan-out.

**Approaches to compare**
- **Single row update:** correct but serialized; it can't go faster than one row's lock.
- **Sharded counters:** split into 100 sub-counters, increment a random one, sum on read.
- **Write buffering / batching:** aggregate likes in memory per gateway server and flush totals every second.
- **Approximate counts:** show "1.2M", because nobody needs the exact number live.

**The twist:** Two parts. First, a viewer taps the heart and sees the count unchanged (batch not flushed), so they tap again. Make users see their own like instantly (read-your-own-writes from Episode 12, client side). Second, the "first 1,000 orders" announcement is Episode 6's flash sale *inside* a live stream, so callback to waiting rooms and atomic stock.

**Verdict:** Client-side optimistic count + per-server batching into sharded counters + approximate display; reuse the flash-sale stock system for live deals.

**Diagram adds:** Live Service · Counter Aggregators.

**Cliffhanger:** "Millions of people now *follow* shops so they don't miss the next live. The actor's shop has 5 million followers. They just posted a new product."

**Study:** write contention, hot keys, sharded counters, batching, approximate counting, eventual consistency, read-your-writes.

**Prep**
- **Day 1:** Learn why hot rows serialize writes, and the sharded counter pattern.
- **Day 2:** Design all four approaches, including how like/unlike works in each.
- **Day 3:** Calculate throughput for each approach. Solve the "my like disappeared" twist.
- **Day 4:** Script with the live-stream hearts as the hook.

**Title:** 100,000 Likes Per Second: How Live Streams Count Without Melting the Database

**Description:**
Counting likes sounds like adding 1. At 100,000 likes per second on one live stream, that "+1" can bring down your database. In this episode of the Haat series we compare single-row updates, sharded counters, batching and approximate counts, then fix the subtle bug where viewers can't see their own like.

---

## Episode 18: The shop with 5 million followers broke the feed

**Status Board:** Month 40 · 25M users · 600K orders/day · 170 engineers
**Callback:** "Live shopping made following shops the most popular feature on Haat. Every user now has a home feed of new products from shops they follow. It worked fine, until the biggest shop posted."
**Prep time:** 5-6 days.

**Scenario:** The actor's shop (5M followers) posts a new collection. Haat tries to insert that post into 5 million feeds immediately. The feed system falls an hour behind, and smaller sellers' posts don't appear at all. Small sellers, the heart of Haat, are furious.

**Numbers:** Average user follows 80 shops; average shop has 300 followers, but the top 100 shops have 1M+. Posts/day, feed opens/day, writes per celebrity post in push mode.

**Approaches to compare**
- **Fan-out on read (pull):** on open, fetch recent posts from every followed shop and merge. Cheap writes, slow reads.
- **Fan-out on write (push):** on post, write into every follower's precomputed feed (through Episode 8's queue). Instant reads, terrible for big shops.
- **Hybrid:** push for normal shops, pull for big shops, merge at read time. What large social networks actually do.

**The twist:** Ranking. Users don't want newest-first; they want "most interesting first", and small sellers need a fair chance against big ones. Where does ranking run in each approach, and how do you avoid recomputing it on every refresh? End on the open question: *"How do we decide what's interesting to each user? That's not a rule you can write by hand. That's machine learning, and that's Season 2."*

**Verdict:** Hybrid fan-out, cached feeds, a simple scoring function for ranking now, and a note that it should be learned, not hand-written.

**Diagram adds:** Feed Service · Feed Cache · Ranking (placeholder box labelled "Season 2").

**Cliffhanger:** "Act 4 is done. Haat runs 40 services and deploys dozens of times a day. Last Thursday, a one-line database change broke checkout for everyone."

**Study:** fan-out patterns, precomputed feeds, timeline caching, the celebrity/hot-user problem, feed ranking signals.

**Prep**
- **Day 1:** Learn push vs pull vs hybrid feed designs.
- **Day 2:** Simulate on paper: 5 users, 1 celebrity shop, a follow graph; trace post and open.
- **Day 3:** Calculate write volume for a 5M-follower post in push mode. Design the hybrid merge and where ranking runs.
- **Day 4:** Script with the celebrity-post hook and the Season 2 teaser.

**Title:** Why One Celebrity Post Can Break a Feed: Designing a News Feed

**Description:**
When a shop with 5 million followers posts, where does that post go? In this episode of the Haat series we design a home feed on the whiteboard, comparing fan-out on read, fan-out on write and the hybrid approach large social networks use. Then we add ranking, and hit the question that leads into machine learning.

---

# Act 5: Never Go Down (Episodes 19-20)

*Haat is too big to fail. These episodes are about changing it safely and surviving disasters.*

## Episode 19: The one-line change that broke checkout

**Status Board:** Month 44 · 28M users · 40 services · 60 deploys/day · 200 engineers
**Callback:** "Since Episode 13 we've been splitting Haat into services. There are 40 now, and they deploy 60 times a day. Every one of those deploys is a chance to break something."
**Prep time:** 4-5 days.

**Scenario:** A developer renames the column `delivery_addr` to `delivery_address`. During the rolling deploy, half the Order service instances run the old code, which looks for a column that no longer exists. Checkout fails for 40% of customers for 18 minutes. The team's response is to ban deploys after 4 PM and on Fridays. Velocity collapses.

**Numbers:** 60 deploys/day × 40 services. Error budget: what does 18 minutes of failed checkout cost at 600K+ orders/day?

**Approaches to compare**
- **Rolling deployment:** replace instances one at a time. No downtime, but old and new versions run together.
- **Blue-green deployment:** run a full copy with the new version and switch all traffic at once. Instant rollback, but double the infrastructure.
- **Canary release:** send 1% of traffic to the new version, watch error rates, then expand, with automatic rollback.
- **Feature flags:** ship code turned off, then enable it per user group. Separates deploying from releasing.

**The twist:** Back to the column rename. In any strategy, old and new code coexist, so the database must work for both. Introduce **expand and contract**: add the new column → write to both → backfill → switch reads → stop writing old → drop old. Six safe steps instead of one dangerous one.

**Verdict:** Canary + automatic rollback for services, feature flags for risky features, expand-and-contract for every schema change. Friday deploys are allowed again.

**Diagram adds:** Deployment pipeline · Feature Flag Service · Monitoring and auto-rollback.

**Cliffhanger:** "We can change Haat safely now. But every one of those 40 services, every database and every shard still runs in one building in Dhaka. On Eid night, that building loses power."

**Study:** deployment strategies, load balancer traffic shifting, health checks, monitoring and automatic rollback, feature flags, backward-compatible migrations.

**Prep**
- **Day 1:** Learn the four strategies and draw traffic flow for each.
- **Day 2:** Compare them on paper: downtime, rollback speed, cost, risk.
- **Day 3:** Work through the column rename with expand-and-contract, step by step.
- **Day 4:** Script with the "18 minutes of broken checkout" incident as the hook.

**Title:** How Big Companies Deploy 100 Times a Day With Zero Downtime

**Description:**
Deploying shouldn't mean downtime, 3 AM releases or crossed fingers. In this episode of the Haat series we compare rolling deployments, blue-green, canary releases and feature flags on the whiteboard, then solve the problem that breaks many zero-downtime deploys: database schema changes, with the expand-and-contract pattern.

---

## Episode 20 (Season Finale): The night the data center went dark

**Status Board:** Year 4 · 30M users · Eid sale night · 1.2M orders/day · 220 engineers
**Callback:** A 60-second montage of the Architecture Diagram growing from Episode 1's single box to today. Then: "Everything we built this season lives in one data center. Tonight, at 9:12 PM, during the biggest Eid sale in Haat's history, that data center loses power."
**Prep time:** 6-7 days (includes the season recap).

**Scenario:** The generator fails to start. Haat is completely offline at the peak of the Eid sale. Backups exist, but restoring them would take 9 hours and lose the last 6 hours of orders. Every minute down costs lakhs and trust.

**Numbers:** Define **RPO** (how much data you can afford to lose) and **RTO** (how long you can afford to be down). Revenue per minute on sale night. Cost of each approach as a multiple of the current infrastructure bill.

**Approaches to compare**
- **Backup and restore:** cheapest; RPO of hours, RTO of hours.
- **Pilot light / warm standby:** a second region with replicated data and minimal servers, scaled up on failover. RPO of seconds, RTO of minutes.
- **Active-passive:** a full second region ready to take traffic, with DNS/global load balancer failover.
- **Active-active:** both regions serve users all the time. Best availability, hardest consistency.

**The twist:** The network link between the regions fails, but both regions are still up. Each thinks the other is dead and starts accepting orders (split brain). The last 50 sale phones get sold in *both* regions, which is Episode 6's oversell at planetary scale. Now explain the **CAP theorem** in plain words: during a network partition, you must choose between staying correct (refuse some requests) and staying available (accept them and reconcile later). Then show Haat's answer per data type, straight from Episode 12's tiers: stock and payments choose correctness; carts, feeds and likes choose availability.

**Verdict:** Active-passive with fast failover for the core, active-active only for data that tolerates conflict, and regular failure drills ("game days").

**Diagram adds:** The whole diagram is duplicated into a second region, with replication arrows and a global load balancer. This is the final season diagram.

**Season recap (2-3 min):** Walk the final diagram box by box: "This box is Episode 2. This one stopped the bots in Episode 5. This one is why nobody gets charged twice..."

**Season 2 teaser:** "Haat has four years of data: what people searched, clicked, bought and watched. In Season 2 we'll use it. How does Haat know what to show you? How can search understand 'something to wear to a wedding'? How does ChatGPT work? See you there."

**Study:** RPO/RTO, backup strategies, replication across regions, failover, split brain and quorum, CAP and PACELC, chaos engineering.

**Prep**
- **Day 1:** Learn RPO/RTO and the four DR strategies.
- **Day 2:** On paper, draw each strategy's failover sequence for Haat's diagram.
- **Day 3:** Design the split-brain scenario and Haat's per-data CAP choices. Prepare the recap.
- **Day 4-5:** Script the finale, build the diagram montage, record the recap and teaser.

**Title:** The Night Our Data Center Went Dark: Disaster Recovery and the CAP Theorem

**Description:**
At the peak of our biggest sale, the entire data center loses power. In the Season 1 finale of the Haat series we design for disaster: backups, warm standby, active-passive and active-active regions, and the split-brain problem that forces every system to choose between being correct and being available. Plus a full recap of everything we built this season, and a look at Season 2.

---

# Season 2 preview: Haat Gets Smart (AI)

Season 2 continues the same story; it doesn't start a new one. Natural first episodes, each growing out of Season 1:

- **The feed ranking box from Episode 18:** how recommendation systems work ("customers also bought").
- **Search from Episode 4:** semantic search, embeddings and vector databases ("something to wear to a wedding").
- **Payments from Episode 7:** fraud detection, plus how a neural network learns.
- **Chat from Episode 16:** a customer support chatbot, and finally how ChatGPT works.

We'll plan this phase in detail when Season 1 is recorded.

---

# Continuity Bible

Check this before scripting any episode. The numbers must never go backwards.

### Growth timeline

| Ep | Time since launch | Users | Orders/day | Engineers | Key infrastructure added |
|---|---|---|---|---|---|
| 1 | Week 1 | 20K | ~500 | 2 | Load balancer, 3 app servers, Redis sessions |
| 2 | Month 3 | 150K | 3K | 3 | CDN, Redis cache |
| 3 | Month 5 | 400K | 10K | 4 | Object storage, first queue, media workers |
| 4 | Month 7 | 1M | 25K | 6 | Search cluster, CDC |
| 5 | Month 9 | 1.5M | 40K | 7 | Rate limiter |
| 6 | Month 12 | 3M | sale night | 9 | Waiting room, Redis stock counter |
| 7 | Month 12 | 3.5M | sale night | 9 | Payment idempotency, reconciliation |
| 8 | Month 15 | 5M | — | 12 | Notification service, priority queues |
| 9 | Month 15 | 5M | sale night | 12 | Scheduler with leases |
| 10 | Month 20 | 8M | 150K | 25 | Read replicas, 4 order shards |
| 11 | Month 21 | 8M | 160K | 25 | Snowflake IDs, short links |
| 12 | Month 22 | 10M | 200K | 30 | Consistency-aware reads |
| 13 | Month 26 | 12M | 300K | 80 | Microservices, saga orchestrator, outbox |
| 14 | Month 28 | 14M | sale | 90 | Circuit breakers, bulkheads |
| 15 | Month 31 | 16M | 400K | 110 | Rider location + assignment services |
| 16 | Month 34 | 18M | 500K | 130 | Real-time gateway, pub/sub |
| 17 | Month 37 | 22M | — | 150 | Live service, counter aggregators |
| 18 | Month 40 | 25M | 600K | 170 | Feed service, feed cache |
| 19 | Month 44 | 28M | 650K+ | 200 | Deploy pipeline, feature flags |
| 20 | Year 4 | 30M | 1.2M (sale) | 220 | Second region |

### Recurring building blocks (use these for callbacks)

- **Redis:** sessions (1) → cache (2) → rate limit counters (5) → stock counter (6) → idempotency store (7) → scheduler sorted sets (9) → rider geo index (15) → like counters (17).
- **The queue:** media processing (3) → order writes (6) → notifications (8) → scheduled jobs (9) → saga events (13) → feed fan-out (18).
- **Idempotency / "at least once + idempotency":** payments (7) → notifications (8) → scheduled jobs (9) → compensations (13).
- **Read-your-own-writes:** My Orders (12) → my own like (17).
- **Oversell:** race condition (6) → double-released stock (9) → live deals (17) → split brain across regions (20).
- **Consistency tiers:** defined in (12), reused for the CAP decision in (20).

---

# Checklist for every whiteboard video

Copy this for each episode and tick it off before publishing.

**Before recording**
- [ ] Status Board numbers checked against the Continuity Bible
- [ ] Callback line written and connects to the previous episode's cliffhanger
- [ ] One-page explanation written in your own words
- [ ] Tiny worked example solved on paper, arithmetic double-checked
- [ ] Hook ready: open with the disaster, not the definition
- [ ] Architecture Diagram updated; new boxes highlighted
- [ ] Whiteboard planned panel by panel (or slides in Excalidraw if you record digitally)
- [ ] Answers ready for the 3 questions viewers are most likely to ask
- [ ] Cliffhanger for the next episode written

**While recording**
- [ ] Show the naive solution failing before the clever one
- [ ] Pause and ask viewers "how would you solve this?" before revealing the answer
- [ ] Use real-world numbers to show why scale matters
- [ ] Use local examples (Dhaka traffic, bKash, Banglish search, Eid and 11.11 sales) where they fit

**After recording**
- [ ] Aim for 12-20 minutes; cut anything that doesn't move the story
- [ ] Add chapters (timestamps) to the description
- [ ] Add "Previous episode" and "Next episode" links, plus the season playlist
- [ ] Thumbnail: one striking drawing plus 3-5 words, readable on a phone
- [ ] Link the sources you studied in the description
- [ ] Cut a 30-60 second Short from the best "aha" moment
