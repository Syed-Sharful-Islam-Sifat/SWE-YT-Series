# SWE YouTube Series, Season 1: Building Haat

*System design through one app that keeps growing · Updated Nov 3, 2026 · @Lion*

---

## How the series works

The first 21 videos are scenario-based system design. Something goes wrong ("we sold 12,000 phones but only had 10,000"), and you solve it on the whiteboard by comparing several approaches. There's no coding on screen, just reasoning, diagrams and back-of-envelope numbers. AI topics start in Season 2, after Episode 21.

**Season 0, The Toolbox, comes first.** It's a separate, standalone series: one hands-on deep-dive video per tool (Nginx, Postgres, Redis, queues and so on), not connected to the Haat story. It's how you build real, practised knowledge of every component before designing with them here. The tool list, depth levels, practice exercises and schedule are in [season-0-toolbox.md](season-0-toolbox.md).

## One app, one story

**Haat starts as a Facebook page.** A clothing seller in Dhaka has run their business on Facebook for years: about 80,000 followers, orders through Messenger comments and inbox, an Excel sheet for orders, cash on delivery and a courier service. It works, but it's falling apart: orders get lost in the inbox, the same dress gets sold twice, every post needs a paid boost to be seen, and one page ban could end the business overnight.

The owner isn't technical, but they hear that "anyone can build an app with AI now". They find **a junior who just graduated from university** and offer a small fee: *"Build me an e-commerce app in 14 days."*

**You are that junior.** The viewer sees every decision through your eyes, from the first sketch on paper to a nationwide platform with 30 million users, its own delivery riders and live shopping. The owner is a recurring character who brings business demands ("Eid sale next week!", "Other sellers want to join!"), and those demands drive the story.

**How Haat grows, in one line per act:**
- **Act 1:** one shop's website → it survives launch and Eid → other Facebook sellers ask to join, and Haat becomes a marketplace.
- **Act 2:** the first big 11.11 sale.
- **Act 3:** the database, the codebase and the team outgrow their first versions.
- **Act 4:** own riders, live tracking, live shopping (Facebook sellers already sell through Facebook Live), shop feeds.
- **Act 5:** changing a huge system safely and surviving a disaster.

Haat **never stops being what it is**: an online store for sellers who started on Facebook. Every new feature is something a real marketplace adds as it grows. Viewers always know which app they're looking at.

### The two rules of the season

1. **Every episode starts from the previous one.** The new problem is caused either by last episode's fix (the side effect), by the growth that fix made possible (the success), or by a reasonable shortcut taken in Episode 1 that has finally hit its limit. Never by "and now we also do X".
2. **The numbers only go up.** Users, orders, products and team size grow every episode, and they never contradict each other. The Continuity Bible at the end is the source of truth.

### Why the problems feel natural: planted shortcuts

In Episode 1 the junior makes simple choices that are **correct for 50 orders a day**. Each one quietly waits for the day the business outgrows it. When it breaks, the viewer remembers Episode 1 and thinks "that choice was right at the time". That's what makes the story feel real instead of forced.

| Shortcut taken in Episode 1 (right at the time) | Breaks in |
|---|---|
| One server; login sessions stored in that server's memory | Ep 2: launch night crash, then logouts after adding servers |
| Product images saved in a folder on the server's disk | Ep 2 (a quick hack) and Ep 4 (sellers' uploads choke everything) |
| Each page section queries the database directly | Ep 3: the Eid campaign melts the database |
| Search is `WHERE name LIKE '%...%'` | Ep 5: 2M products, 8-second search |
| OTP login sends an SMS to any number, no limits | Ep 6: bots run up the SMS bill |
| Checkout reads the stock, then saves the order (two steps) | Ep 7: the flash sale oversells |
| Cash on delivery only | Ep 8: prepaid payments, double charges |
| A cron job cancels unconfirmed orders | Ep 10: the cron runs on six servers at once |
| One database for everything | Ep 11: 300M rows, 100% CPU |
| Auto-increment order IDs | Ep 12: two customers get the same order number |
| One deploy of one app, by one person | Ep 20: 40 services, 60 deploys a day |
| Everything in one data center, daily backups | Ep 21: the data center goes dark |

### Two recurring on-screen devices

- **The Status Board** (top corner, first 30 seconds): time since launch, users, orders per day, engineers. Viewers feel the company growing.
- **The Architecture Diagram**: one diagram for the whole season. Episode 1 draws the first version: a single box. Each episode adds or changes one or two boxes, highlighted in a new colour. Open on last episode's diagram and end on the new one. By Episode 21 it's the full system, and that final diagram is your season thumbnail.

### The opening callback (30 seconds, every episode from Episode 2)

> "Last time, [problem] almost killed Haat, and we fixed it with [solution]. That worked. Maybe too well. Because today, [solution / its success / a shortcut from Episode 1] is exactly what causes our next disaster..."

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

---

# The confidence playbook: how to prepare any episode

Every episode below has its own **Learning guide** (what to learn, in what order, from where, and how to know you're ready) and its own **Prep plan**. This section is the method behind all of them. Read it once before Episode 1 and come back whenever you feel unsure. If you haven't started the series yet, begin with **Before Episode 1: the pre-season plan** below.

## Where confidence on camera comes from

Confidence isn't a personality trait here. It comes from four things you can prepare:
1. **You can explain every concept without notes**, in plain words, in under two minutes.
2. **You solved the problem yourself** before reading anyone's solution, so the approaches feel like *your* reasoning, not a recited list.
3. **Every number on screen is checked**, and you know where it came from.
4. **You've already answered the hard questions** viewers will ask, so nothing in the comments surprises you.

The prep plan for each episode is built to give you exactly these four things.

## The 4-day routine (about 3-4 hours a day)

**Day 1: Learn (concepts)**
- Read the episode's *Read first* resources. Skip the *Go deeper* ones unless something is still unclear.
- For each concept in *Learn in this order*, write **one paragraph in your own words** in a notebook. If you can't, re-read that part.
- **The 2-minute test:** explain each concept out loud, without notes, as if to a friend who isn't an engineer. Record it on your phone if you can. Anything you stumble on goes on a "still unclear" list.
- End the day by clearing the "still unclear" list with a second source (a different book, article or video).

**Day 2: Solve it yourself (approaches)**
- Close every resource. On paper, write the scenario and the numbers, then write **every approach you can think of**, even bad ones.
- For each approach, draw the request flow and mark exactly where it breaks.
- Only then open the resources and compare. Add anything you missed. The approaches you found yourself are the ones you'll explain best.
- Fill in a pros / cons / "when I'd choose this" table. This table becomes the middle of your video.

**Day 3: Numbers, twist and questions**
- Do the capacity math in the episode's *Numbers* section with a calculator, then do it again a second way (for example, per second and then per day). Both must agree.
- Check every real-world number (prices, latencies, limits) against a source, and note the source for the video description. If you can't find a source, say "roughly" on camera.
- Work through the twist on paper until you can draw it from memory.
- Write short answers to the *Questions viewers will ask*.
- Go through *You're ready to record when you can*. Anything you can't tick goes back to Day 1 material.

**Day 4: Script and rehearse**
- Plan the whiteboard **panel by panel** (a panel is one screen-worth of drawing). Typical episode: 8-12 panels.
- Write the **hook** (first 15 seconds), the **callback**, the **verdict** and the **cliffhanger** word for word. Everything else is bullet points. Reading a full script sounds robotic; bullets make you sound like you're thinking, which is the point of the series.
- Update the Architecture Diagram and highlight the new boxes.
- **Dry run 1:** talk through the whole episode with the panels, no recording. Note where you got lost.
- **Dry run 2:** record it on your phone, watch it at 1.5× speed, and fix the two or three most confusing moments.

**Recording day:** record in sections (hook, scenario, each approach, twist, verdict). Re-recording one section is much easier than re-recording 15 minutes.

## Rules for numbers on screen

- **Say where a number comes from.** "We assumed...", "the provider's docs say...", "a rough industry figure is...". Viewers trust you more when you label estimates honestly.
- **Round aggressively.** 86,400 seconds a day ≈ 100,000. 1,157 requests/sec ≈ 1,000. Exact numbers distract from the reasoning.
- **Show the formula, then the answer.** "200,000 views ÷ 900 seconds ≈ 220 per second." Viewers learn the method, not just the result.
- **Peak is not average.** Real traffic has peaks 2-10× the daily average (sales and launches more). Always say which one you mean.
- **Check numbers against the Continuity Bible** at the end of this document before every recording.

## Back-of-envelope cheat sheet

Keep this next to you while preparing.

| Fact | Value to use |
|---|---|
| Seconds in a day | 86,400 ≈ **100K** (use 100K for quick math) |
| Requests/day → requests/sec | divide by 100K. 1M/day ≈ 10/sec average |
| Peak factor | normal day 2-3×; launches, sales, viral posts 10× or more |
| Reading from memory (RAM) | ~100 nanoseconds |
| Reading 1 MB from memory | ~a few microseconds to ~0.25 ms |
| SSD random read | ~0.1 ms (100 µs) |
| Round trip inside one data center | ~0.5 ms |
| Round trip Dhaka ↔ Singapore | ~50-80 ms |
| Round trip Dhaka ↔ US | ~200-250 ms |
| One simple database query (indexed, warm cache) | ~1 ms |
| One modest app server, simple dynamic pages | hundreds of requests/sec (very app-dependent) |
| One Redis instance, simple operations | ~100K operations/sec |
| 1 KB × 1M | 1 GB |
| 1 MB × 1M | 1 TB |

Source for the latency rows: Jeff Dean's "Latency numbers every programmer should know" (search the title; several updated versions exist). Treat them as orders of magnitude, not exact values.

## When you're unsure on camera

- It's fine to say **"it depends, and here's what it depends on"**. That's the most senior sentence in engineering.
- If a viewer corrects you in the comments and they're right, **pin their comment and thank them**. It builds trust and the channel's community.
- Never invent a company's internal detail ("Daraz uses X"). Say "a common approach is..." unless you have a public source, and link the source.

## Resource kit for the whole season

You don't need to read these cover to cover. Each episode's learning guide points to the exact chapter or article.

**Books (the core four)**
- **Designing Data-Intensive Applications (DDIA)** by Martin Kleppmann: the deep "why" behind most episodes. Chapter numbers in this document refer to the **1st edition (2017)**. If you have the 2nd edition, find the chapter with the same title.
- **System Design Interview, Vol. 1 and Vol. 2** by Alex Xu (Vol. 2 with Sahn Lam): many of these scenarios in interview form, with clear diagrams.
- **Release It! (2nd edition)** by Michael Nygard: for the reliability episodes (15, 20, 21).
- **Site Reliability Engineering** (the "SRE book") by Google, free to read online at sre.google/books.

**Free online references**
- **The System Design Primer** (GitHub: donnemartin/system-design-primer): a free overview of almost every topic in the season.
- **AWS Builders' Library** (aws.amazon.com/builders-library): short, practical articles by Amazon engineers on timeouts, retries, caching, load shedding and more.
- **martinfowler.com**: clear articles on architecture patterns (monolith first, circuit breaker, deployment patterns, feature toggles).
- **microservices.io** by Chris Richardson: the saga and outbox patterns.
- **Engineering blogs**: Stripe, Uber, Discord, Netflix, Shopify, Meta, Instagram, Notion, Figma, Cloudflare. Use them for real-world stories to quote.

**YouTube (to see how others explain it)**
- **ByteByteGo** (Alex Xu's channel): short animated explanations.
- **Hussein Nasser**: backend fundamentals (connections, proxies, databases, WebSockets).
- **Martin Kleppmann's distributed systems lecture series** (University of Cambridge, 2020, free on YouTube): for Episodes 11-14 and 21.
- **Arpit Bhayani**: database internals and system design from real engineering experience.

Watch a video *after* you've done your own Day 2 solving, not before, so your explanation stays yours.

**Tools**
- **Excalidraw** (excalidraw.com): free, hand-drawn-style whiteboard. Good for both planning and recording.
- **OBS Studio**: free screen and camera recording.
- A basic USB or lavalier microphone matters more than the camera. Viewers forgive average video, not bad audio.

---

## Before Episode 1: the pre-season plan (4 weeks)

> **Updated plan:** Season 0 ([season-0-toolbox.md](season-0-toolbox.md)) now covers Weeks 1-3 in far more depth, with hands-on practice for every tool. If you've finished Season 0, treat these weeks as a quick self-check. Week 4's production rehearsal happens before Toolbox #1 instead.

**How much to cover before you start:** go deep on the foundations that come back in every act, sketch a rough map of the rest of the season, and fully prepare only Episodes 1-3. Don't try to master all 21 topics first. You'll learn Episode 14 best in the weeks before Episode 14, and over-preparing is the most common way a series never starts.

Budget: about 2-3 hours a day, roughly 60-70 hours in total.

### Week 1: Core foundations (go deep)
These come up in almost every episode. If you're shaky on them, every episode will feel shaky.
- **How a request travels:** DNS, TCP/TLS at headline level, HTTP request/response, web server → app → database. Draw it from memory.
- **Databases:** tables and indexes (why `LIKE '%dress%'` is slow), transactions, what "isolation level" means, `SELECT ... FOR UPDATE`.
- **Estimation:** daily → per second, peak vs average, storage per year. Do 10 practice estimates using the cheat sheet above.
- *Read:* Alex Xu Vol. 1, Chapters 1-3 (all of Chapter 1 this time); DDIA Chapter 1, Chapter 3 (skim, for indexes) and Chapter 7 "Transactions" (the first half).
- **Weekly test:** explain each topic out loud in 2 minutes, without notes.

### Week 2: The recurring building blocks (go deep, hands-on)
These are the boxes the whole season keeps reusing (see *Recurring building blocks* in the Continuity Bible).
- **Redis:** strings, `INCR`, `SET NX EX`, sorted sets, TTLs. These cover sessions, cache, rate limits, stock counters and leases.
- **Caching:** cache-aside, TTL, invalidation, what a stampede is.
- **Queues:** producer/consumer, at-least-once delivery, retries, dead-letter queues.
- **Idempotency:** why retries cause duplicates and how an idempotency key stops them.
- **Replication basics:** primary/replica and replication lag (just enough to plant it; Episode 11 goes deep).
- *Read:* DDIA Chapter 5 "Replication" (the first section), Chapter 11 (the messaging part only); the Redis docs on data types; the AWS Builders' Library article "Making retries safe with idempotent APIs".
- **Labs (do these; they're where most of the confidence comes from):**
  1. In Postgres, open two terminals and make two "checkouts" oversell one item. Then fix it with `FOR UPDATE` and again with an atomic `UPDATE ... WHERE stock > 0`. That's Episode 7, done by hand.
  2. In Redis, build a 10-line rate limiter with `INCR` + `EXPIRE`.
  3. Run a queue (Redis lists, or RabbitMQ in Docker), kill the consumer halfway through a job, and watch the message come back.

### Week 3: A map of the whole season (skim only)
The goal is to know where the story goes, so that the shortcuts, callbacks and cliffhangers you plant in Act 1 point the right way. It's not to master Acts 2-5.
- Skim about 30 minutes per chapter (diagrams and summary only): Alex Xu Vol. 1 Chapters 4 (rate limiter), 5 (consistent hashing), 7 (unique IDs), 10 (notifications), 11 (news feed), 12 (chat), 13 (autocomplete); Vol. 2 "Proximity service", "Hotel reservation" and "Payment system".
- Watch the first 2-3 lectures of Martin Kleppmann's distributed systems series.
- Read the Continuity Bible and the *Season at a glance* table until you can tell the whole Haat story in 3 minutes.
- **Weekly test:** for each of the 21 episodes, write one sentence: what breaks, and the core idea of the fix. If you can write all 21, your map is good enough.

### Week 4: Production rehearsal and the launch buffer
- **Day 1-2:** set up the microphone, OBS and the Excalidraw template (Episode 1's Day 4 covers the visual style). Record a throwaway 5-minute pilot ("What is a cache?"), watch it back, and fix your audio and pacing. Nobody ever sees this one.
- **Day 3 onward:** start Episode 1's prep plan (6 days), then Episode 2's.

### The launch rule
**Don't publish Episode 1 until Episodes 1-3 are recorded and Episode 4 is in prep.** A two-episode buffer means a bad week never breaks your upload schedule, and you can still go back and add a callback in Episode 1 if Episode 3's prep reveals one.

### Ready-to-start checklist
- [ ] I can draw the request flow and explain an index, a transaction and a race condition without notes.
- [ ] I've done all three Week 2 labs myself.
- [ ] I can write the 21 one-sentence summaries from memory.
- [ ] I've watched my pilot recording and fixed the audio.
- [ ] Episodes 1-3 are recorded.

### After launch: keep a pipeline going
- **Rhythm:** while you edit Episode N, you're learning Episode N+2. Use each episode's 4-day routine.
- **Plan for the heavy episodes:** 7, 8, 11, 13, 14 and 21 each need an extra day or two of prep. Schedule a lighter week before them.
- **Refresh before each act** (one weekend each):
  - *Before Act 2 (Ep 7):* DDIA Chapter 7, the whole chapter.
  - *Before Act 3 (Ep 11):* DDIA Chapters 5, 6 and 9, plus Kleppmann lectures on replication and consistency.
  - *Before Act 4 (Ep 16):* Alex Xu Vol. 2 "Proximity service" and "Nearby friends"; Hussein Nasser on WebSockets.
  - *Before Act 5 (Ep 20):* *Release It!* Part I, and the SRE book chapters "Embracing Risk" and "Managing Incidents".

---

## Season at a glance

| # | What breaks | Topic | Caused by |
|---|---|---|---|
| 1 | A Facebook seller needs an app in 14 days, on a tiny budget | Requirements, estimation, monolith-first design, backups | The Facebook business outgrew Messenger and Excel |
| 2 | Launch on Facebook Live kills the server; adding a 2nd server logs users out | Load balancing, stateless servers | Ep 1's single server meets the page's 80K followers |
| 3 | Eid boosted campaign, 20× traffic, the database melts | Caching, cache stampede | Ep 2's servers all hit one database |
| 4 | Other sellers join; their uploads choke the servers and fail at 95% | Object storage, resumable uploads | Ep 3's Eid success made other Facebook sellers ask to join |
| 5 | 2M products and search takes 8 seconds | Search engines, autocomplete | Ep 4 filled the catalog; Ep 1's `LIKE` search |
| 6 | Bots scrape search and spam OTP login, leaving a huge SMS bill | Rate limiting | Ep 5's public search API; Ep 1's unlimited OTP |
| 7 | 11.11 flash sale: 12,000 phones sold, only 10,000 in stock | Race conditions, locking, atomic counters | Ep 1's check-then-save checkout; Ep 6's Redis counters |
| 8 | A customer is charged twice for a reservation | Idempotency, payment state machine | Ep 7's sale items must be prepaid, not COD |
| 9 | "Sale starts in 10 min" push would take 3 days, then self-DDoS | Message queues, fan-out | Ep 7's sale needs promoting |
| 10 | Expired reservations released twice, so oversold again | Job scheduling, leases | Ep 1's cron job, cloned onto 6 servers |
| 11 | The orders table hits 300M rows and the database is at 100% | Replication, sharding, consistent hashing | Ep 1's single database, ten episodes later |
| 12 | Two customers get the same order number | Distributed ID generation, Base62 links | Ep 11's shards each auto-increment |
| 13 | "I placed an order but My Orders is empty" | Replication lag, consistency models | Ep 11's read replicas |
| 14 | Microservices: order paid, never shipped | Sagas, transactional outbox | Team outgrew the monolith |
| 15 | The courier partner's API slows down and takes down the whole app | Cascading failures, circuit breakers | Ep 14's chain of services |
| 16 | Haat launches its own riders: who is nearest? | Geospatial indexing | Ep 15: the courier can't be trusted |
| 17 | "Where is my rider?" Live tracking + chat | WebSockets, pub/sub, delivery receipts | Ep 16's riders stream locations |
| 18 | Celebrity seller goes live: 500K viewers, 100K likes/sec | Hot keys, sharded counters | Ep 17's WebSockets; sellers' Facebook Live habit |
| 19 | A 5M-follower shop posts and feeds lag for an hour | News feed fan-out, ranking | Ep 18 made following shops popular |
| 20 | 40 services deploy daily; a column rename breaks checkout | Zero-downtime deployment | Ep 14 onward: many services, many deploys |
| 21 | **Finale:** the Dhaka data center goes dark during the Eid sale | Multi-region, disaster recovery, CAP | Everything we built |

---

# Act 1: From Facebook Page to Its Own App (Episodes 1-6)

*A Facebook shop gets its own website, survives launch night and Eid, and turns into a marketplace. One database, a few servers, and every lesson learned the hard way.*

## Episode 1: "Build me an e-commerce app in 14 days"

**Status Board:** Day 0 · Facebook page with 80K followers · ~50 orders/day (via Messenger) · 300 products · 1 engineer (you, a fresh graduate)
**Callback:** None; this is the origin story. Open on the owner's phone: a Messenger inbox with 400 unread "price?" messages, and an Excel sheet with rows highlighted red for "sold twice".
**Prep time:** 5-6 days (first episode: add time to set up your recording workflow and draw the base diagram style you'll reuse all season).

**Scenario:** The owner explains the business to you over tea. Orders arrive as Facebook comments and Messenger messages. Two staff copy them into Excel by hand. Some orders are lost in the inbox; some dresses are sold twice because the sheet wasn't updated in time. Facebook reach drops every month, so every post needs a paid boost. A friend's page was hacked last year and their business vanished overnight. The owner says: "Everyone is making apps with AI now. Make me one in 14 days. I can't pay much."

You could open an AI coding tool and start typing prompts. Instead, you take out a sheet of paper. **The AI can write the code fast. It can't decide what the system should be. That's your job.** This episode is the junior's first instinct done right: understand the business, size the problem, and choose the simplest design that works.

**Step 1: Requirements (what the owner actually needs)**
- *Must have (functional):* product catalog with photos and sizes; cart; checkout with name, phone, address and **cash on delivery**; login with phone number + OTP (customers don't want passwords; their phone number is their identity on Facebook commerce too); order confirmation SMS; an admin panel for staff to add products, update stock, see and update orders, and print courier slips.
- *The owner's real pain:* "Stock must be correct. No more selling the same dress twice."
- *Must have (non-functional):* works well on mid-range Android phones and slow mobile data, because customers arrive from Facebook links on their phones; cheap to run; built in 14 days; one person can run it; **never lose an order**.
- *Out of scope for now:* online payment (bKash/cards), a native mobile app, other sellers, advanced search, own delivery.
- *Website, not a mobile app:* customers tap a link inside Facebook and land in the in-app browser. Nobody installs an app from a Facebook post. A fast, mobile-first website is the right "app" for version 1.

**Step 2: Numbers (back-of-envelope)**
- 80K followers, ~50 orders/day today. Estimate ~5,000 visits/day once the site is live, ~8 pages per visit = 40,000 page views/day.
- 40,000 ÷ 100,000 seconds ≈ **0.4 page views/sec on average**. A boosted post might cause a 20-50× spike: **~10-20 page views/sec**.
- Data: 300 products × 5 photos × ~1 MB = ~1.5 GB of images. Orders: 50/day × 365 × ~2 KB ≈ 36 MB a year. Tiny.
- Conclusion: **one small server (~$10-20/month) handles this comfortably.** The numbers decide the architecture, not fashion.

**Step 3: Data model (one panel)**
Users (phone, name, address) · Products (name, price, description) · Product images (file path) · Stock per size · Orders (customer, status, total, created time) · Order items. Order statuses: PENDING → CONFIRMED (phone call) → SHIPPED → DELIVERED / RETURNED. IDs are auto-increment numbers (1, 2, 3...).

**Approaches to compare**
- **A. Don't build: use an off-the-shelf store (Shopify, WooCommerce and similar).** Ready in days, cheap, proven. For many small shops this is honestly the right answer. Limits for this owner: local flows (COD phone confirmation, local courier booking, Messenger order import) need plugins or workarounds, monthly fees add up, and the owner wants a platform they can grow and own. Say clearly on camera: "If you're a small shop, start here."
- **B. The "YouTube architecture": microservices, Kubernetes, Kafka, separate databases, because big companies use them.** Show what it would take: a dozen moving parts, weeks of setup, a monthly bill many times larger, and one junior to run it all. It solves problems Haat doesn't have, and would never ship in 14 days.
- **C. One simple app on one server (a "monolith").** One web application (store + admin in one codebase), one database, product images in a folder on the same server, a web server (like Nginx) in front, a scheduled job (cron) that cancels orders nobody confirmed within 24 hours. Fast to build, cheap, easy to understand, easy for AI-written code to be reviewed. Its limits are real, but they're far away at 50 orders a day.

**The twist:** Everything is on one machine, including **every order the business has ever taken**. If the disk fails, the server is hacked, or someone runs the wrong command, the business loses its order history and customer list. That's worse than the old Excel sheet. Two cheap fixes, done right away:
1. **Automatic daily database backups stored somewhere else** (different provider or object storage), following the 3-2-1 idea: 3 copies, 2 different kinds of storage, 1 off-site.
2. **A test restore.** A backup you've never restored is a hope, not a backup. Tell the GitLab 2017 story: during an outage they found their backup methods weren't working as expected.

Plus three habits that cost nothing: secrets (database password, SMS API key) in environment variables, not in the code; HTTPS from day one; the server's operating system updated automatically.

**Verdict:** Ship approach C plus backups. For 50 orders a day, simple is the correct engineering decision, not the lazy one. Be honest that approach A is right for many shops, and that approach B is what you grow into only when the problems actually arrive.

**Diagram v1:** Customer's phone (from a Facebook link) → domain (DNS) → **one VPS** [Nginx → App (store, cart, checkout, admin) → Database · /uploads folder · cron job] → SMS gateway · daily backups to a separate location.

**Planted for later:** Show the "planted shortcuts" table from the top of this document as a quick, light montage: "Every one of these is the right choice today. Remember them."

**Cliffhanger:** The owner loves it: "I'll announce it on Facebook Live this Friday at 9 PM. All 80,000 followers." You look at the single $10 server...

### Learning guide

**Learn in this order**
1. **How a web request works end to end:** phone → DNS → server → web server → app → database → back. You must be able to draw this from memory.
2. **Requirements gathering:** functional vs non-functional requirements, and why "out of scope" is as important as "in scope".
3. **Back-of-envelope estimation:** daily → per second, average vs peak, storage per year.
4. **Monolith vs microservices:** what each is, the real costs of microservices, and why "monolith first" is common advice.
5. **Vertical structure of a simple web app:** web server (Nginx), app, database, file storage, scheduled jobs.
6. **Basic data modelling:** tables, primary keys, auto-increment IDs, order status as a simple state flow.
7. **Backups and restore:** full backups vs point-in-time recovery, the 3-2-1 rule, why restore tests matter.
8. **Security basics for a first app:** HTTPS, secrets in environment variables, OTP basics, the OWASP Top 10 at a headline level.
9. **What "vibe coding" is**, where it helps (speed, boilerplate) and where it doesn't (architecture decisions, security review, data model).

**Resources**
- *Read first:*
  - Alex Xu, *System Design Interview Vol. 1*: Chapter 1 "Scale from zero to millions of users" (read only the single-server part for now), Chapter 2 "Back-of-the-envelope estimation", Chapter 3 "A framework for system design interviews" (requirements first).
  - Martin Fowler, "MonolithFirst" (martinfowler.com, 2015) and "Microservice Premium" (2015).
  - DDIA Chapter 1 "Reliable, Scalable, and Maintainable Applications", the section on describing load.
- *Go deeper:*
  - DHH, "The Majestic Monolith" (Signal v. Noise / Basecamp, 2016).
  - Nick Craver, "Stack Overflow: The Architecture - 2016 Edition": a huge site on surprisingly few servers.
  - The Twelve-Factor App (12factor.net): read "III. Config" and "VI. Processes".
  - PostgreSQL documentation, chapter "Backup and Restore" (or the MySQL equivalent if you use MySQL).
  - OWASP Top 10 (owasp.org): read the one-line summary of each item.
- *Real-world stories to quote:*
  - GitLab, "Postmortem of database outage of January 31" (2017): backups that didn't work when needed.
  - Shopify's posts on keeping a "modular monolith" (search "Shopify deconstructing the monolith").
  - Andrej Karpathy coined "vibe coding" in early 2025; one quote is enough context.
- *Watch:* any ByteByteGo or Hussein Nasser video on "what happens when you type a URL" or "monolith vs microservices", after your Day 2 solving.

**Try it yourself (optional, 2-3 hours, builds real confidence)**
- Rent the cheapest VPS for a day (or use a local virtual machine). Install Nginx and a database, deploy a "hello world" app, take a database backup, delete a table, and restore it. Having done it once, you'll talk about it like someone who has.

**You're ready to record when you can**
- [ ] Draw the request flow from a Facebook link to the database and back, from memory.
- [ ] Turn "5,000 visits a day" into requests per second in your head and explain peak vs average.
- [ ] Give three honest reasons a small shop should use Shopify/WooCommerce, and three reasons this owner might not.
- [ ] Explain why microservices are expensive for a team of one, without saying "microservices are bad".
- [ ] Explain the 3-2-1 backup rule and why a restore test matters.
- [ ] Name every planted shortcut and say why each one is fine today.

**Questions viewers will ask (prepare answers)**
- *"Why not just use Firebase / Supabase / serverless?"* Good options too; they trade control and cost predictability for speed. The principle (start simple, one place for data) is the same. Mention you chose a VPS so viewers can *see* every box.
- *"Which language or framework?"* Deliberately not the point of the series. Any mainstream framework (Laravel, Django, Rails, Node, Spring) works for this design.
- *"Isn't one server a single point of failure?"* Yes, and you accept it knowingly at this stage, with backups to limit the damage. Episode 2 deals with it.
- *"Can AI really build this in 14 days?"* It can write most of the code. It can't decide requirements, data model, security or backups. That's the engineer's job, and it's what this series teaches.

**Don't get this wrong on camera**
- Don't present the monolith as "the beginner choice". Present it as the **right choice for these numbers**, with a clear point where it stops being right.
- Don't mock off-the-shelf platforms; many viewers run shops on them.
- Don't quote exact server prices as facts; say "around $10-20 a month for a small VPS" and keep it approximate.

### Prep plan
- **Day 1 (learn, ~4 h):** Alex Xu chapters 1-3; Fowler's two articles; DDIA chapter 1 "describing load". Write one paragraph each on: request flow, requirements, estimation, monolith vs microservices, backups.
- **Day 2 (solve, ~3 h):** Without resources, write the owner's requirements as if you were interviewing them. Then design approaches A, B and C on paper. For B, list every component and estimate setup time and monthly cost. Then compare with Alex Xu chapter 1.
- **Day 3 (numbers and twist, ~3 h):** Do the traffic, storage and cost math twice. Read the GitLab postmortem summary. Write the backup plan in three lines. Prepare answers to the four viewer questions.
- **Day 4 (design your visual style, ~3 h):** This is the first episode, so decide the season's look: colours for each box type (servers, databases, queues, external services), the Status Board template, the "new this episode" highlight colour. Draw Diagram v1 cleanly in Excalidraw and save it as the season template.
- **Day 5 (script and rehearse, ~4 h):** Write the hook word for word: open on the Messenger inbox with 400 unread messages and the "sold twice" Excel rows. Plan 10-12 panels. Two dry runs, the second recorded.
- **Day 6 (record):** Record in sections. Add a short channel intro at the end: what the series is, how episodes connect, and that Episode 2 is launch night.

**Title:** A Facebook Seller Asked Me to Build Their E-commerce App in 14 Days. Here's How I Designed It

**Description:**
A clothing business running on Facebook, Messenger and Excel wants its own online store, built in 14 days by a fresh graduate. Before writing a single prompt for an AI coding tool, we design it on the whiteboard: gathering real requirements, doing the back-of-envelope math, comparing an off-the-shelf store, a "big tech" microservices setup and a simple monolith, and adding the one thing nobody remembers until it's too late: backups. This is Episode 1 of the Haat series, where one app grows from a single server to a nationwide platform.

---

## Episode 2: The Facebook Live that killed our server

**Status Board:** Launch night · 80K page followers → 60K visitors in 2 hours · ~800 orders on launch day · 1 engineer (you)
**Callback:** "Last time, a Facebook seller hired us to build their store in 14 days. We built one simple server, the right call for 50 orders a day. Tonight at 9 PM, the owner goes live on Facebook and tells 80,000 followers to tap the link."
**Prep time:** 4-5 days.

**Scenario:** The owner goes live at 9 PM. 8,000 people are watching; the post gets shared. 25,000 people tap the link in the first 15 minutes. The single server, running the app, the database and serving every product photo, hits 100% CPU and stops responding. The owner is still live, saying "the link is in the comments!" while the comments fill with "site not working". You rush to separate the database and add two more app servers behind a load balancer, and the site comes back. Then the support messages start: "I keep getting logged out", "my cart is empty again".

**Numbers:** 25,000 visitors in 15 minutes × ~8 page views = 200,000 page views ÷ 900 s ≈ **220 page views/sec**. Each page view is about 1 HTML page + 3 API calls + ~10 product images, all served by the same small server: roughly **3,000 HTTP requests/sec** on a machine sized for 20. How many page views/sec can one small app server handle? Where's the first bottleneck: CPU, memory, the database, or bandwidth for images?

**Approaches to compare**
- **Vertical scaling:** buy a bigger server. Fastest fix (minutes on most cloud providers), but it has a ceiling, a bigger bill and is still a single point of failure.
- **Separate the database from the app server:** the first and cheapest split; each can now scale on its own, and the database isn't fighting the app for CPU.
- **Let Nginx serve images directly** instead of the app: a quick win, because static files shouldn't go through application code.
- **Horizontal scaling + load balancer:** many app servers behind one entry point. Compare round robin, least connections and IP hash.
- **Fixing the logouts:**
  - *Sticky sessions:* the load balancer always sends a user to the same server. Easy, but if that server dies, the user's session dies too.
  - *Shared session store (Redis):* every server reads sessions from one fast store. **Redis enters the series here.**
  - *Stateless tokens (JWT):* the session lives in a signed token on the client. No lookup needed, but hard to revoke.

**The twist:** Three problems appear after the fix:
1. The load balancer itself is now a single point of failure. Introduce health checks and redundant load balancers (or a managed cloud load balancer).
2. When you deploy, servers restart one by one and users still lose their carts. The rule of the season arrives: **app servers must be stateless.**
3. Staff upload a new product photo through the admin panel. It's saved to server 2's disk. Customers routed to servers 1 and 3 see a broken image. Your quick fix at 1 AM: one shared folder on a small file server (NFS) that all app servers use. Say on camera: "This is a hack. Remember it." (It breaks in Episode 4.) Also, the cron job from Episode 1 now exists on three servers, so you switch it off on two of them and keep it only on server 1. (That comes back in Episode 10.)

**Verdict:** Separate the database, put 3 stateless app servers behind a load balancer with health checks, keep sessions in Redis, let Nginx serve images, and accept the shared image folder as a temporary hack.

**Diagram adds:** Load Balancer · App Servers ×3 · Redis (sessions) · Database on its own server · Shared file server for images (marked "temporary").

**Cliffhanger:** "Three app servers, one database. Every page still runs 40 database queries, written in a hurry in Episode 1. And the owner just booked a boosted Facebook campaign for Eid."

### Learning guide

**Learn in this order**
1. **Vertical vs horizontal scaling:** what each means, cost and limits.
2. **What a load balancer does:** single entry point, distributing traffic, health checks, removing dead servers.
3. **Load balancing algorithms:** round robin, weighted round robin, least connections, IP hash. When each is useful.
4. **Layer 4 vs Layer 7 load balancing** (at a high level: does it look at connections or at HTTP requests?).
5. **Sessions and state:** what a session is, where it's stored, why storing it in server memory breaks with multiple servers.
6. **Sticky sessions vs shared session store vs JWT:** trade-offs, especially revocation (logging someone out, a stolen token).
7. **Stateless servers:** what "stateless" means in practice (no local sessions, no local files, no local scheduled jobs).
8. **Removing the load balancer as a single point of failure:** redundant pairs, floating IPs, managed cloud load balancers.

**Resources**
- *Read first:*
  - Alex Xu Vol. 1, Chapter 1 "Scale from zero to millions of users": the sections on separating the database, load balancers and stateless web tier.
  - Sam Rose, "Load Balancing" (samwho.dev, interactive visual article): the best visual intro to the algorithms.
  - OWASP "Session Management Cheat Sheet" (cheatsheetseries.owasp.org): sessions done safely.
- *Go deeper:*
  - NGINX documentation, "HTTP Load Balancing" (methods, health checks, session persistence).
  - jwt.io "Introduction to JSON Web Tokens", then search "Stop using JWT for sessions" (joepie91, 2016) for the strongest counter-argument. Present both sides.
  - The Twelve-Factor App, "VI. Processes" (stateless processes).
- *Watch:* Hussein Nasser's videos on load balancers and Layer 4 vs Layer 7 proxies.

**Try it yourself (optional, 2 hours)**
- Run two copies of a tiny web app on your laptop on different ports, each printing its own name, with Nginx round-robin in front. Store a session in memory and watch it "disappear" when Nginx switches servers. Then move it to Redis. You'll explain the logout bug far better after seeing it.

**You're ready to record when you can**
- [ ] Calculate the launch-night requests/sec and explain why images multiply the load.
- [ ] Draw a user's three requests bouncing across servers and show exactly why the cart disappears.
- [ ] Compare sticky sessions, Redis sessions and JWT in one sentence each, including one weakness.
- [ ] Explain why the load balancer is now the single point of failure, and two ways to fix it.
- [ ] Define "stateless server" and list three kinds of hidden state (sessions, local files, local cron).

**Questions viewers will ask (prepare answers)**
- *"Why not autoscaling?"* Autoscaling adds servers automatically, but new servers take minutes to start; a launch spike happens in seconds. And it doesn't help if servers aren't stateless. (Autoscaling arrives later in the season.)
- *"JWT is better, everyone uses it."* JWT avoids a lookup but makes logout and token theft harder to handle. Many teams use short-lived JWTs plus a refresh token in a store. "It depends" is the honest answer.
- *"Why not a CDN for the images?"* A great idea, and it's coming in Episode 3. Tonight, the fastest fix was letting Nginx serve files.
- *"Is NFS a good idea?"* Not long-term, and you say so on camera. It's a deliberate, labelled hack.

**Don't get this wrong on camera**
- Don't say "a load balancer makes your app infinitely scalable". The database is still single, and Episode 3 is about exactly that.
- Don't confuse "stateless server" with "stateless app". The state still exists; it just lives in Redis and the database, not on the app server.

### Prep plan
- **Day 1 (learn):** Alex Xu chapter 1 (the sections above), the Sam Rose article and the OWASP cheat sheet. Paragraph per concept; 2-minute out-loud test on sessions vs JWT.
- **Day 2 (solve):** Draw a user's requests bouncing between two servers and show exactly why the cart disappears. Write all three fixes and their failure cases. Then add the image-folder and cron problems on your own before reading anything.
- **Day 3 (numbers and twist):** Do the launch-night math twice (per second and per 15 minutes). Estimate how many app servers you need if one handles ~80 page views/sec. Design the redundant load balancer. Do the optional Nginx experiment.
- **Day 4 (script and rehearse):** Hook: the owner live on Facebook saying "link in the comments" while the comments say "site not working". Update Diagram v1 → v2. Two dry runs.

**Title:** Our App Crashed on Launch Night During a Facebook Live. Here's How We Scaled It (Load Balancers Explained)

**Description:**
Launch night. The owner goes live on Facebook, 25,000 people tap the link in 15 minutes, and our single server dies. In this episode of the Haat series we take the app from one server to a scalable setup on the whiteboard: vertical vs horizontal scaling, load balancers and their algorithms, and the sneaky bug where users get logged out the moment you add a second server. We end with the rule that shapes the rest of the season: app servers must be stateless.

---

## Episode 3: The Eid campaign that melted our database

**Status Board:** Month 4 · 150K users · 2K orders/day · 3 engineers (the owner hired two more juniors)
**Callback:** "Last time we survived launch night by adding app servers behind a load balancer. That worked, but every one of those servers still talks to the same single database."
**Prep time:** 4 days.

**Scenario:** Haat's homepage runs 40 database queries per visit: categories, banners, new arrivals, prices, stock, reviews. That's how the pages were written in Episode 1, one query per section, and it was fine at 50 orders a day. For Eid, the owner books a boosted Facebook campaign and a popular influencer's post, both at 9 PM. Traffic jumps 20×. You can add app servers in minutes, but they all hammer the same database, which hits 100% CPU. The site goes down during the most important sales week of the year.

**Numbers:** Normal evening peak is ~150 requests/sec. 20× means ~3,000 requests/sec at 40 queries each, so **120,000 queries/sec** against a database that handles maybe 5,000. What cache hit rate do you need to survive? (To get down to ~5,000 queries/sec you need roughly 96% of queries served from cache.)

**Approaches to compare**
- **Cache-aside (lazy loading):** the app checks Redis first and loads from the database on a miss. Simple, but stale data and slow first requests.
- **Write-through:** update the cache on every write. Always fresh, but slower writes and cached data nobody reads.
- **Write-back:** write to the cache and sync to the database later. Very fast, but you can lose data.
- **CDN / full-page caching:** cache whole pages and images at the edge, close to the user. Massive scale, but hard for personalized content like the cart.
- **Which data gets which strategy?** Category list vs product price vs cart vs stock count. This decision is the real lesson.

**The twist:** A cache stampede. The cached homepage expires at 9:00:00 PM, right as the influencer's post goes up, and 50,000 requests miss at once and hit the database together. Solutions: a lock so only one request rebuilds, early random refresh, and serving stale data while refreshing.

**Verdict:** CDN for images and static pages, cache-aside in Redis for product data with TTLs, never cache the cart or stock blindly, plus stampede protection.

**Diagram adds:** CDN · Redis now also holds the cache.

**Cliffhanger:** "Eid was a record. Other Facebook sellers noticed. Fifty of them messaged the owner the same week: 'Can we sell on Haat too?'"

### Learning guide

**Learn in this order**
1. **Why caching works:** the 80/20 pattern (a small set of data gets most reads), memory vs disk speed.
2. **Cache hit rate** and how to calculate the load it removes.
3. **Caching patterns:** cache-aside, read-through, write-through, write-back (write-behind).
4. **TTL (time to live) and eviction policies:** LRU, LFU, and what Redis does when memory is full.
5. **Cache invalidation:** why it's famously hard; delete-on-write vs update-on-write.
6. **CDNs:** edge locations, what can and can't be cached, Cache-Control headers at a basic level.
7. **Cache stampede / thundering herd:** what causes it and the three standard fixes (locking/request coalescing, early probabilistic refresh, stale-while-revalidate).
8. **What not to cache:** carts, stock for checkout, anything where a wrong answer costs money.

**Resources**
- *Read first:*
  - Alex Xu Vol. 1, Chapter 1: the sections on cache and CDN.
  - AWS whitepaper "Database Caching Strategies Using Redis".
  - AWS Builders' Library, "Caching challenges and strategies".
- *Go deeper:*
  - "Scaling Memcache at Facebook" (NSDI 2013 paper): read the parts about leases and thundering herds. Excellent real-world story.
  - "Optimal Probabilistic Cache Stampede Prevention" (Vattani et al., VLDB 2015): the math behind early random refresh. Read the abstract and the core idea only.
  - MDN Web Docs, "HTTP caching" (Cache-Control, stale-while-revalidate).
  - Redis documentation on key eviction policies.
  - Cloudflare Learning Center, "What is a CDN?".
- *Watch:* ByteByteGo videos on caching strategies and cache problems (search "ByteByteGo cache").

**Try it yourself (optional, 1-2 hours)**
- Use Redis locally: set a key with `SET key value EX 10`, read it, wait for it to expire. Write a tiny script that simulates 100 simultaneous misses and count how many "database calls" happen with and without a lock.

**You're ready to record when you can**
- [ ] Calculate the required cache hit rate from the numbers.
- [ ] Explain cache-aside step by step on a diagram, including the miss path.
- [ ] Explain why write-back can lose data and when it's still acceptable.
- [ ] Decide, for each piece of homepage data (banners, categories, prices, stock, cart), a strategy and a TTL, with a reason.
- [ ] Draw the stampede as a timeline and show each fix.

**Questions viewers will ask (prepare answers)**
- *"Why not just make the database bigger?"* That's vertical scaling: it buys time, but 120K queries/sec of mostly identical reads is exactly what caches are for, and they cost far less.
- *"Redis or Memcached?"* Both work for caching. Haat already has Redis for sessions, so one less system to run. Mention Memcached is simpler and great for pure caching.
- *"What about stale prices?"* Short TTLs plus delete-on-write for price changes, and checkout always reads the real price from the database.

**Don't get this wrong on camera**
- Don't say "just cache everything". The lesson is choosing what to cache.
- Be careful with "write-through is always consistent": a failure between the cache write and the database write can still leave them different. Say "much fresher", not "perfectly consistent".

### Prep plan
- **Day 1 (learn):** Alex Xu sections, the AWS whitepaper and the Builders' Library article. Paragraph per pattern.
- **Day 2 (solve):** On paper, decide which strategy each piece of homepage data gets: banners, categories, prices, stock, cart. Justify each TTL.
- **Day 3 (numbers and twist):** Calculate the hit rate needed. Read the Facebook Memcache leases section and draw the stampede with its three fixes.
- **Day 4 (script):** Hook: the influencer's post goes live at 9:00 PM and the database graph goes vertical. Update the diagram. Two dry runs.

**Title:** The Eid Campaign That Crashed Our Website: Caching Strategies Explained

**Description:**
A boosted Facebook campaign and an influencer post send 20× traffic, and the database melts. Caching is the obvious answer, but which kind? In this episode of the Haat series we compare cache-aside, write-through, write-back and CDN caching, decide which data should use which strategy, and then fix the sneaky failure that takes down even well-cached systems: the cache stampede.

---

## Episode 4: Why every seller's upload failed at 95%

**Status Board:** Month 6 · 400K users · 5K sellers · 10K orders/day · 4 engineers
**Callback:** "Last time caching got us through Eid. Other Facebook sellers saw it, and the owner made the biggest decision in Haat's life: open Haat to other sellers. Five thousand joined in two months. Every one of their photos and videos goes into the shared folder we hacked together in Episode 2."
**Prep time:** 4 days.

**Scenario:** Haat is now a marketplace. Sellers upload 10 photos and a 1-minute product video per listing, often over shaky mobile data from their shops. Uploads go through your app servers into the shared file server, which is running out of disk. The app servers spend most of their bandwidth and memory moving files instead of serving customers. A seller's 300 MB video fails at 95% and starts over from zero. Three times. They post an angry review in a Facebook sellers' group.

**Numbers:** 5,000 sellers × 20 new listings/month × (10 photos × 3 MB + 300 MB video) = how many TB per month? (≈ 33 TB/month.) How much of your app servers' bandwidth is that? How long does the shared file server's disk last?

**Approaches to compare**
- **Upload through the app server:** simplest; fails on big files and steals capacity from customers.
- **Direct upload to object storage via pre-signed URL:** the file skips your servers entirely; the server only hands out a temporary permission.
- **Chunked, resumable uploads:** split the file into 5 MB pieces, upload them in parallel, and resume from the last successful chunk.
- **Background processing pipeline:** after upload, a message on a **queue** triggers thumbnails, image compression and video transcoding. **The queue enters the series here.**

**The twist:** Hundreds of resellers upload the *same* supplier photo for the same product (very common for Facebook resellers). Should you store it 300 times? Introduce deduplication with content hashing, and its traps: the privacy questions, and what happens when one seller deletes a shared file.

**Verdict:** Pre-signed URLs + resumable chunked uploads to object storage, a queue-driven processing pipeline, CDN delivery, hash-based dedup for images. Retire the shared folder hack from Episode 2.

**Diagram adds:** Object Storage · Upload/Processing Queue · Media Workers. The "temporary" shared file server is crossed out.

**Cliffhanger:** "Two million products in the catalog now. A customer types 'iphone cover', and waits eight seconds. The search is still the one line of SQL we wrote in Episode 1."

### Learning guide

**Learn in this order**
1. **Object storage vs file systems vs block storage:** what S3-style storage is (buckets, keys, objects), why it's cheap and durable.
2. **Why uploads shouldn't pass through app servers:** bandwidth, memory, slow clients tying up connections.
3. **Pre-signed URLs:** how a server grants a temporary, limited permission without exposing credentials.
4. **Multipart / chunked uploads and resumability:** how chunks, part numbers and a final "complete" call work.
5. **Asynchronous processing with a queue:** producer, queue, consumer (worker); why thumbnails and transcoding happen after upload.
6. **Image and video processing basics:** resizing, compression formats (WebP/AVIF), video transcoding to multiple qualities (HLS at a high level).
7. **Content hashing and deduplication:** SHA-256 as a fingerprint, reference counting for deletes.
8. **CDN for media delivery** (connects to Episode 3).

**Resources**
- *Read first:*
  - Alex Xu Vol. 1, Chapter 14 "Design YouTube": the upload flow, pre-signed URLs and transcoding pipeline.
  - Amazon S3 documentation: "Uploading objects with presigned URLs" and "Uploading and copying objects using multipart upload".
- *Go deeper:*
  - Alex Xu Vol. 1, Chapter 15 "Design Google Drive": block-level uploads and deduplication.
  - Alex Xu Vol. 2, Chapter 9 "S3-like Object Storage": how object storage works inside.
  - The tus protocol (tus.io): an open standard for resumable uploads; read the "Protocol" overview.
  - *Pro Git*, Chapter 10 "Git Internals - Git Objects" (free online): the clearest intro to content-addressed storage by hash.
- *Real-world stories:* engineering posts on image pipelines from Instagram, Pinterest or Shopify (search "<company> image processing pipeline").

**Try it yourself (optional, 1-2 hours)**
- Create a free-tier bucket (Cloudflare R2, Backblaze B2, AWS S3 or a local MinIO), generate a pre-signed upload URL, and upload a file with `curl`. Then run `sha256sum` on two copies of the same image to see identical hashes.

**You're ready to record when you can**
- [ ] Draw the full flow: request upload → get URL → upload chunks → confirm → queue → process → serve via CDN.
- [ ] Explain why a pre-signed URL is safe (expiry, one specific object, one action).
- [ ] Explain exactly what happens when chunk 57 of 60 fails, and how resume works.
- [ ] Calculate storage per month and the bandwidth saved on app servers.
- [ ] Explain dedup with reference counting, and why you might *not* dedup private files.

**Questions viewers will ask (prepare answers)**
- *"What if someone uploads a 10 GB file or malware?"* The pre-signed URL can limit size and type; a worker scans and validates after upload before the product goes live.
- *"Which object storage is cheapest?"* Prices change; compare storage *and* download (egress) costs. Don't name a winner on camera.
- *"Why not compress on the phone?"* Good idea and complementary: the app can resize before upload, which helps on slow mobile data.

**Don't get this wrong on camera**
- Don't say object storage is "a hard drive in the cloud". It has no folders (only key prefixes) and you replace whole objects rather than editing them.
- Don't promise dedup saves "90%"; say "it can save a lot when many sellers share supplier photos".

### Prep plan
- **Day 1 (learn):** Alex Xu chapter 14 upload section, S3 pre-signed URL and multipart docs. Paragraph per concept.
- **Day 2 (solve):** Design the full flow on paper: request upload → get URL → upload chunks → confirm → process → serve. Mark every failure point.
- **Day 3 (numbers and twist):** Calculate storage per month and bandwidth saved. Design dedup with reference counts. Do the optional bucket experiment.
- **Day 4 (script):** Hook: a seller's progress bar at 95%... and back to 0%. Update the diagram (cross out the Episode 2 hack). Two dry runs.

**Title:** Why Your Upload Fails at 95%: Designing Resumable Large File Uploads

**Description:**
Haat just opened to other Facebook sellers, and thousands of them are uploading photos and videos over mobile internet. Our servers are drowning. In this episode we redesign uploads on the whiteboard: why routing files through your servers is a mistake, how pre-signed URLs and chunked uploads make uploads fast and resumable, how a queue handles processing, and how deduplication saves huge amounts of storage.

---

## Episode 5: The search box that took 8 seconds

**Status Board:** Month 8 · 1M users · 2M products · 25K orders/day · 6 engineers
**Callback:** "Last time we fixed uploads, and sellers listed two million products. Now customers can't find any of them, because search is still the one line of SQL from Episode 1."
**Prep time:** 4-5 days.

**Scenario:** Search runs `WHERE name LIKE '%iphone cover%'` on the products table. With 2 million products it takes 8 seconds, and it still misses "iPhone 15 Back Case" and anything typed in Banglish ("mobail kabhar"). Customers who can't find a product can't buy it, and conversion from search drops by half.

**Numbers:** 1M users, ~5M searches a day, peak ~300 searches/sec, plus autocomplete on every keystroke (×5). 2M products × ~1 KB of searchable text = ~2 GB of text to index.

**Approaches to compare**
- **SQL LIKE:** show why `%...%` can't use a normal index and scans every row.
- **Database full-text index:** a real improvement and no new infrastructure, but limited ranking, typo handling and scale.
- **Dedicated search engine with an inverted index (Elasticsearch/OpenSearch-style):** word → list of products. Fast, ranks by relevance (TF-IDF/BM25 at an intuitive level), handles filters like price and category.
- **Autocomplete:** a trie with precomputed top-k suggestions at each node, or a prefix → suggestions table in Redis, rebuilt periodically.

**The twist:** Two problems. First, the search index is a *copy* of the database: when a seller changes a price, how does the index find out? Compare dual writes vs change data capture (CDC) through the queue from Episode 4. Second, typos and Banglish: "iphon" and "mobail" should still work (fuzzy matching, synonyms). Tease: "Searching by *meaning* instead of words is a Season 2 topic."

**Verdict:** A search engine fed by CDC events, a trie/Redis autocomplete with top-k, fuzzy matching and a synonym list.

**Diagram adds:** Search Cluster · CDC stream from the database · Autocomplete store.

**Cliffhanger:** "Search is fast and public now. Within a week, our logs show someone calling it 3 million times a night. And it's not customers."

### Learning guide

**Learn in this order**
1. **How a database index works (B-tree) and why `LIKE '%x%'` can't use it.**
2. **Inverted index:** tokenizing text, term → posting list, intersecting lists for multi-word queries.
3. **Text analysis:** lowercasing, stemming, stop words, tokenizers.
4. **Relevance ranking:** TF-IDF intuition, BM25 as the modern default (intuition only, no formulas needed on camera).
5. **Database full-text search** (PostgreSQL `tsvector`, MySQL FULLTEXT) and its limits.
6. **Search engines:** Elasticsearch/OpenSearch basics (index, document, shard, replica, query vs filter).
7. **Autocomplete:** tries, top-k at each node, prefix tables; why it must answer in under ~100 ms.
8. **Fuzzy matching and synonyms:** edit distance (Levenshtein), synonym lists for Banglish.
9. **Keeping derived data in sync:** dual writes and their race conditions vs change data capture (CDC).

**Resources**
- *Read first:*
  - *Introduction to Information Retrieval* by Manning, Raghavan & Schütze (free online at the Stanford NLP site): Chapter 1 "Boolean retrieval" (inverted index). Skim Chapter 6 for TF-IDF.
  - Alex Xu Vol. 1, Chapter 13 "Design a search autocomplete system".
  - "Use The Index, Luke" (use-the-index-luke.com) by Markus Winand: the pages on how B-tree indexes work and on LIKE filters.
- *Go deeper:*
  - *Introduction to Information Retrieval*, Chapter 3 "Dictionaries and tolerant retrieval" (wildcards and spelling correction).
  - PostgreSQL documentation, Chapter "Full Text Search", and the `pg_trgm` extension.
  - Elasticsearch documentation: full-text queries, fuzzy queries, synonyms, the completion suggester.
  - DDIA Chapter 3 "Storage and Retrieval": the section on full-text search and fuzzy indexes; Chapter 11 on change data capture.
  - Debezium documentation (introduction to CDC).
  - Martin Kleppmann, "Turning the database inside-out" (talk, 2015): why derived data should be fed from a change log.

**Try it yourself (optional, 2 hours)**
- In PostgreSQL, create a table with 1M generated product names, time a `LIKE '%cover%'` query, then add a full-text or trigram index and time it again. Real numbers from your own laptop make a great on-screen moment.

**You're ready to record when you can**
- [ ] Explain on a diagram why `LIKE '%cover%'` scans every row.
- [ ] Build an inverted index by hand for 10 product names and answer a two-word query with it.
- [ ] Explain BM25 in one sentence ("rare words matter more, repeated words matter less and less, long text is normalised").
- [ ] Draw a trie for 8 words with top-3 suggestions at each node.
- [ ] Explain why dual writes can leave the index wrong, and how CDC fixes it.
- [ ] Explain how "mobail kabhar" can match "mobile cover" (synonyms + fuzzy).

**Questions viewers will ask (prepare answers)**
- *"Why not just use PostgreSQL full-text search?"* At this size it can be enough for many teams. The search engine wins on relevance tuning, typo tolerance, facets and scaling search separately from the database.
- *"Elasticsearch vs OpenSearch vs Meilisearch/Typesense?"* All use inverted indexes. Pick by team skills, scale and hosting. Not the point of the episode.
- *"How does Google do it?"* Same core idea (inverted index) at enormous scale, plus much more ranking. Keep it brief.

**Don't get this wrong on camera**
- Don't claim a search engine is "always better". It's another system to run and keep in sync.
- Don't show BM25 formulas unless you can explain every symbol. Intuition is enough.

### Prep plan
- **Day 1 (learn):** IR book chapter 1, Use The Index Luke pages, Alex Xu chapter 13. Paragraph per concept.
- **Day 2 (solve):** Build a tiny inverted index and a tiny trie on paper from 10 product names (include one Banglish name). Design each approach.
- **Day 3 (numbers and twist):** Calculate index size and autocomplete QPS (searches × 5 keystrokes). Design index syncing (dual write vs CDC) and typo handling. Do the optional PostgreSQL timing experiment.
- **Day 4 (script):** Hook: a live typing demo, "iphone cover"... spinner... 8 seconds... no results. Two dry runs.

**Title:** Why Our Search Took 8 Seconds, and How Search Engines Really Work

**Description:**
Two million products and a search box that takes 8 seconds and still can't find "iphone cover". In this episode of the Haat series we design product search on the whiteboard: why SQL LIKE fails, how an inverted index works, how autocomplete answers in under 100 ms, how to handle typos and Banglish, and how to keep the search index in sync with your database.

---

## Episode 6: A bot cost us lakhs in OTP SMS overnight

**Status Board:** Month 10 · 1.5M users · 40K orders/day · 7 engineers
**Callback:** "Last time we built fast public search. Bots love fast public APIs. And since Episode 1, our login sends an OTP SMS to any phone number anyone types, with no limit."
**Prep time:** 4 days.

**Scenario:** Two attacks in one night. A competitor's scraper hits your search API 3 million times to copy every price. Worse, a bot hits the "Send OTP" login endpoint 2 million times. Every SMS costs money. You wake up to a huge SMS bill, a slow site, and real users who can't log in.

**Numbers:** 2M OTP requests × ৳0.30-0.50 per SMS. (Check your SMS provider's real price and pick one figure for the title.) What's a reasonable limit for a real human? (A real person might request 3 OTPs in 10 minutes, not 300.)

**Approaches to compare**
- **Fixed window counter:** "max 5 per minute". Simple, but allows bursts at window edges (5 at 0:59, 5 at 1:00).
- **Sliding window log:** store every request timestamp. Accurate, but memory heavy.
- **Sliding window counter:** a weighted blend of two windows. Good accuracy, low memory.
- **Token bucket / leaky bucket:** allow short bursts but control the average rate.

All of them live in **Redis**, the same Redis that already holds sessions (Episode 2) and cache (Episode 3), now counting requests with atomic `INCR`.

**The twist:** You run 6 app servers. If each counts on its own, the attacker gets 6× the limit, so the counter must be shared and atomic. And the attacker rotates phone numbers and IPs. What do you limit on: phone, IP, device, account, or all of them in layers? Add CAPTCHA after N attempts and exponential backoff on OTP resend.

**Verdict:** Token bucket per client for search, sliding window counters on several keys for OTP, all in Redis, plus CAPTCHA and resend backoff.

**Diagram adds:** Rate Limiter (at the gateway, backed by Redis).

**Cliffhanger:** "Atomic counters in Redis. Remember that. Because on 11.11, we're going to need a counter that never goes below zero, and a million people are going to test it at midnight. And our checkout still uses the stock check we wrote in Episode 1."

### Learning guide

**Learn in this order**
1. **Why rate limiting exists:** abuse, cost protection, fairness, protecting your own servers.
2. **Where it runs:** client, API gateway, application, per-service; HTTP 429 and the `Retry-After` header.
3. **The four algorithms:** fixed window, sliding window log, sliding window counter, token bucket (and leaky bucket as its cousin). Draw each on a timeline.
4. **Distributed rate limiting:** why per-server counters fail; atomic operations in Redis (`INCR` + `EXPIRE`, or a Lua script).
5. **Choosing the key:** per IP, per phone, per device, per account, per API key; layered limits.
6. **OTP-specific defenses:** resend backoff, CAPTCHA, limits per phone *and* per IP, SMS pumping fraud.
7. **Race conditions in counters** (a preview of Episode 7): why "read, then write" isn't atomic.

**Resources**
- *Read first:*
  - Alex Xu Vol. 1, Chapter 4 "Design a rate limiter". The core resource for this episode.
  - Stripe engineering blog, "Scaling your API with rate limiters" (2017).
  - Cloudflare blog, "How we built rate limiting capable of scaling to millions of domains" (2017): the sliding window counter in practice.
- *Go deeper:*
  - Redis documentation for `INCR` (it includes a rate limiter pattern) and for Lua scripting (atomic multi-step operations).
  - Twilio's articles on SMS pumping / SMS toll fraud: a real-world explanation of why OTP endpoints are attacked.
  - OWASP "Automated Threats to Web Applications" (scraping, credential stuffing).
  - Google SRE book, chapter "Handling Overload" (for the server-protection side).
- *Watch:* ByteByteGo's rate limiter video, after your Day 2 solving.

**Try it yourself (optional, 1 hour)**
- In Redis, implement a fixed window with `INCR key` + `EXPIRE key 60`, then fire 20 requests in a loop and watch the 6th get rejected. Then try to break it at the window edge.

**You're ready to record when you can**
- [ ] Run the same burst of 12 requests through all four algorithms on a timeline and show which get rejected.
- [ ] Explain the fixed window edge problem with a concrete example.
- [ ] Explain why 6 servers with local counters give the attacker 6× the limit.
- [ ] Design layered OTP limits (per phone, per IP, per device) and explain what each layer stops.
- [ ] Calculate the cost of the attack and the money saved.

**Questions viewers will ask (prepare answers)**
- *"Can't attackers just use thousands of IPs?"* Yes, which is why you layer limits (phone, device, account), add CAPTCHA and watch for patterns. No single limit stops everything.
- *"Won't a strict limit block real users behind one shared IP (like a mobile carrier's NAT)?"* Yes. That's why IP limits are loose and phone/account limits are tighter.
- *"Why not use the cloud provider's or Cloudflare's rate limiting?"* Often a great choice at the edge. You still need business-aware limits (per phone) in the app.

**Don't get this wrong on camera**
- Don't call the leaky bucket and token bucket "the same thing". They're related: token bucket allows bursts up to bucket size; leaky bucket smooths output to a constant rate.
- Don't present CAPTCHA as a complete solution; bots can solve many CAPTCHAs.

### Prep plan
- **Day 1 (learn):** Alex Xu chapter 4, the Stripe and Cloudflare posts. Draw each algorithm on a timeline.
- **Day 2 (solve):** Run the same burst of requests through each algorithm on paper. Decide limits for search and OTP yourself before comparing.
- **Day 3 (numbers and twist):** Design the distributed, multi-key version. Calculate the money lost and saved. Do the optional Redis experiment.
- **Day 4 (script):** Hook: the morning SMS bill notification. Update the diagram. Two dry runs.

**Title:** A Bot Cost Us Lakhs in OTP SMS Overnight: Designing a Rate Limiter

**Description:**
Every OTP costs money, and bots know it. In this episode of the Haat series we design a rate limiter on the whiteboard, comparing fixed windows, sliding windows, token buckets and leaky buckets against the same attack traffic, then make it work across many servers against an attacker who keeps changing phone numbers and IPs.

---

# Act 2: The Big Sale (Episodes 7-10)

*Haat runs its first 11.11 mega sale. Four episodes, one event, and everything that goes wrong around it.*

## Episode 7: We sold 12,000 phones but only had 10,000

**Status Board:** Month 12 · 3M users · 11.11 sale night · 9 engineers
**Callback:** "Last time we used atomic counters in Redis to stop bots. Tonight, 600,000 people are about to tap Buy at the same second, and our checkout still does what it did in Episode 1: read the stock, then save the order."
**Prep time:** 4-5 days.

**Scenario:** Haat's first 11.11 sale: 10,000 phones from an electronics seller at 40% off, starting at midnight. 600,000 users are waiting with their thumbs on the button. When the dust settles, 12,000 orders were confirmed. You oversold by 2,000 and now have to cancel 2,000 orders on the biggest night of the year. The Facebook sellers' groups are on fire.

**Numbers:** 600K users tapping within ~10 seconds is ~60,000 requests/sec at midnight vs a normal peak of ~500. Stock: 10,000.

**Approaches to compare**
- **Read stock, then update (the Episode 1 code):** two requests both read "1 left" and both succeed. Draw the race condition on a timeline. At 50 orders a day this almost never happened; at 60,000 requests a second it happens thousands of times.
- **Atomic conditional update in the database:** `UPDATE stock SET qty = qty - 1 WHERE id = ? AND qty > 0`. Correct and simple, but every request still queues on one hot row.
- **Pessimistic locking (`SELECT ... FOR UPDATE`):** correct, but requests queue on one row and the database chokes.
- **Optimistic locking (version column):** no locks, but under heavy contention almost every request fails and retries.
- **Atomic counter in Redis:** `DECR` in memory and reject below zero, then write orders to the database through the queue (from Episode 4).
- **Virtual waiting room:** let users in at a controlled rate with a queue position. Protects everything behind it, but users wait.

**The twist:** A user reserves a phone but never pays. Do you hold stock forever? Design reservations with a 10-minute payment timeout that returns stock to the pool without overselling when it comes back.

**Verdict:** Waiting room in front, atomic Redis counter for stock, orders persisted through the queue, reservations with a timeout. Keep the database constraint as a last safety net.

**Diagram adds:** Waiting Room · Stock Counter (Redis) · Order Queue.

**Cliffhanger:** "The reservation is held. Now the customer has to pay. Until tonight, Haat was cash on delivery only. On a slow network. On a phone that retries automatically."

### Learning guide

**Learn in this order**
1. **Race conditions:** what they are, drawn as two timelines interleaving. The "lost update" problem.
2. **Transactions and ACID** at a practical level; what isolation means.
3. **Isolation levels:** read committed, repeatable read, serializable, and which anomalies each allows (focus on lost update and write skew).
4. **Atomic conditional updates** (`UPDATE ... WHERE qty > 0`) and checking affected rows.
5. **Pessimistic locking** (`SELECT ... FOR UPDATE`) and **optimistic locking** (version column, compare-and-set).
6. **Hot rows / contention:** why one popular row limits throughput no matter how big the database is.
7. **Redis atomic operations:** `DECR`, and Lua scripts for "check and decrement" in one step.
8. **Admission control / virtual waiting rooms:** letting traffic in at a rate the system can handle.
9. **Reservations with timeouts:** holding, expiring and returning stock.

**Resources**
- *Read first:*
  - DDIA Chapter 7 "Transactions": the sections on weak isolation levels, preventing lost updates (atomic writes, explicit locking, compare-and-set) and write skew.
  - Alex Xu Vol. 2, Chapter 7 "Hotel Reservation System": the concurrency section compares pessimistic locking, optimistic locking and database constraints. Very close to this episode.
- *Go deeper:*
  - PostgreSQL documentation: "Transaction Isolation" and "Explicit Locking" chapters.
  - Redis documentation: `DECR`, and "Scripting with Lua" (atomic multi-step logic).
  - Cloudflare blog on how they built their Waiting Room product (2021): a real virtual waiting room design.
  - Martin Kleppmann's Cambridge lecture series on concurrency (for a visual intro to races).
- *Watch:* Arpit Bhayani or Hussein Nasser on "pessimistic vs optimistic locking", after your Day 2 solving.

**Try it yourself (optional, 2 hours)**
- In PostgreSQL, open two terminal sessions. In both, `BEGIN; SELECT qty FROM stock WHERE id=1;`, then both update `qty = <value read> - 1` and commit. Watch the lost update happen. Then repeat with `FOR UPDATE` and watch the second session wait. Seeing this once makes the episode easy to explain.

**You're ready to record when you can**
- [ ] Draw two requests overselling the last phone on a timeline, step by step.
- [ ] Explain how each approach prevents it, and its cost under 60,000 requests/sec.
- [ ] Explain why optimistic locking performs badly under heavy contention.
- [ ] Explain why a single Redis `DECR` is atomic, and why "GET then SET" in Redis is not.
- [ ] Design the reservation timeout and explain how stock returns without overselling.
- [ ] Explain what the waiting room protects and what it costs the user.

**Questions viewers will ask (prepare answers)**
- *"What if Redis crashes during the sale?"* You lose in-memory counts unless persisted or replicated. Mitigations: Redis replication/persistence, and the database constraint as a final check. Be honest that this is a real risk.
- *"Why not just use SERIALIZABLE isolation?"* Correct but it aborts many transactions under contention; you still need retries, and throughput drops.
- *"How do Daraz / Alibaba / ticket sites do it?"* Don't claim internal details. Say "large flash-sale systems commonly combine waiting rooms, in-memory stock and asynchronous order creation".

**Don't get this wrong on camera**
- Don't say "transactions prevent all race conditions". Under the default isolation level of many databases, the read-then-write pattern can still lose updates.
- Don't forget to show the conditional `UPDATE ... WHERE qty > 0` option; it's the simplest correct fix and viewers will ask why it's missing.

### Prep plan
- **Day 1 (learn):** DDIA chapter 7 (the sections listed) and Alex Xu Vol. 2 chapter 7 concurrency section. Paragraph per concept.
- **Day 2 (solve):** Draw two requests overselling the last phone, then show how each approach prevents it. Rank the approaches yourself before reading comparisons.
- **Day 3 (numbers and twist):** Calculate midnight requests/sec. Design the waiting room admission rate and the reservation timeout. Do the two-terminal experiment.
- **Day 4 (script):** Hook: "12,000 orders confirmed. 10,000 phones in the warehouse." Two dry runs.

**Title:** We Sold 12,000 Phones But Only Had 10,000: Designing a Flash Sale System

**Description:**
600,000 people tap "Buy" at midnight and the stock count goes negative. In this episode of the Haat series we design a flash sale system on the whiteboard, starting with the race condition that causes overselling, then comparing conditional updates, pessimistic locks, optimistic locks, Redis atomic counters and virtual waiting rooms. Finally we handle the customer who reserves a phone and never pays.

---

## Episode 8: The customer who got charged twice

**Status Board:** Month 12 · 3.5M users · 11.11 sale, minute 3 · 9 engineers
**Callback:** "Last time we stopped overselling with reservations. Since Episode 1, Haat has been cash on delivery only. But for flash-sale phones, the owner insists on prepayment, because fake COD orders would lock up stock. Now 10,000 people need to pay through bKash or a card, exactly once."
**Prep time:** 4-5 days.

**Scenario:** A customer taps "Pay ৳18,000" for their sale phone. The network is slow, so they tap again. Or the app retries automatically after a timeout. They're charged twice, and they post screenshots everywhere. On sale night, hundreds of customers have the same story.

**Numbers:** 10,000 payments in ~10 minutes (~17/sec, with bursts). The payment gateway's typical latency, timeout values, and how often mobile networks drop requests in practice.

**Approaches to compare**
- **Disable the button on the frontend:** helps, but doesn't stop app retries, network retries or a second device.
- **Idempotency keys:** the client sends a unique key per payment attempt; the server stores the result by key and returns the same result on repeats.
- **Database unique constraint** on the order's payment: the database itself refuses duplicates.
- **Payment state machine:** CREATED → PROCESSING → SUCCESS / FAILED, with only valid transitions allowed.

**The twist:** You call the payment gateway and the request times out. Did the money move? You can't retry blindly and can't assume failure. Introduce status queries (most gateways offer a "query payment status" API), callbacks/webhooks from the gateway, reconciliation jobs, and why "exactly once" is really "**at least once + idempotency**". Keep that phrase, because it returns in Episodes 9, 10 and 14.

**Verdict:** Idempotency keys + a payment state machine + a unique constraint as the last line of defense + gateway status queries + nightly reconciliation.

**Diagram adds:** Payment component with an Idempotency Store · Reconciliation Job · external Payment Gateway.

**Cliffhanger:** "Payments are safe. But next sale needs promoting: one push notification to five million phones. Our first version would finish in three days."

### Learning guide

**Learn in this order**
1. **Why requests get duplicated:** double taps, client retries, network timeouts, load balancer retries.
2. **Idempotency:** definition with everyday examples (pressing a lift button twice), idempotent vs non-idempotent HTTP methods.
3. **Idempotency keys:** who generates them, where they're stored, how long, what's returned on a repeat, what happens if two identical requests arrive at the same moment.
4. **Database unique constraints** as a final safety net.
5. **State machines for payments:** states, allowed transitions, why "update status where status = expected" matters.
6. **Timeouts with unknown outcome:** the hardest case; status queries and webhooks/callbacks (IPN).
7. **Reconciliation:** comparing your records with the gateway's settlement report.
8. **"Exactly once" vs "at least once + idempotency".**
9. **How local gateways work at a high level:** redirect/hosted checkout, callback URL, status query (bKash, SSLCommerz and similar).

**Resources**
- *Read first:*
  - Stripe engineering blog, "Designing robust and predictable APIs with idempotency" (Brandur Leach, 2017). Must-read.
  - Alex Xu Vol. 2, Chapter 11 "Payment System": exactly-once, retries, idempotency, reconciliation.
  - AWS Builders' Library, "Making retries safe with idempotent APIs".
- *Go deeper:*
  - Brandur Leach, "Implementing Stripe-like Idempotency Keys in Postgres": a detailed, practical design.
  - Airbnb engineering, "Avoiding Double Payments in a Distributed Payments System" (2019).
  - The developer documentation for bKash and SSLCommerz (or the gateway you know best): find their payment status query API and their callback/IPN mechanism. Using real local examples on screen is powerful.
  - Alex Xu Vol. 2, Chapter 12 "Digital Wallet" (optional, for event sourcing ideas).

**Try it yourself (optional, 1-2 hours)**
- Sketch (or code) a tiny endpoint that stores `idempotency_key → response` in a table with a unique constraint. Send the same request twice and confirm the second returns the stored result. Then send two at the exact same moment and think about what happens.

**You're ready to record when you can**
- [ ] Draw the payment flow and mark every point where a duplicate can happen.
- [ ] Explain idempotency keys in under a minute, including where the key is created and stored.
- [ ] Draw the payment state machine and name one invalid transition it prevents.
- [ ] Walk through the timeout case: what you do at 0 s, 30 s, 5 minutes and the next morning.
- [ ] Explain why "exactly once" over a network is a myth, and what you do instead.

**Questions viewers will ask (prepare answers)**
- *"Who generates the idempotency key, client or server?"* The client, per payment attempt, so its retries share the key. The server can derive one from the order ID as a second layer.
- *"How long do you keep keys?"* Long enough to cover any retry window (Stripe-style systems use about a day); payment records themselves are kept permanently.
- *"What does the customer see during an unknown outcome?"* "Payment processing, we'll confirm shortly", never "failed, try again".

**Don't get this wrong on camera**
- Don't say "use a UUID and you're done". The hard parts are storage, concurrent identical requests and unknown outcomes.
- Don't describe any specific gateway's internal behaviour without checking their docs.

### Prep plan
- **Day 1 (learn):** Stripe post, Alex Xu Vol. 2 chapter 11, the Builders' Library article.
- **Day 2 (solve):** Draw the payment flow and mark every duplicate point. Design your fixes before reading Brandur's implementation post.
- **Day 3 (numbers and twist):** Design the timeout scenario and reconciliation. Read one local gateway's status/callback docs. Prepare the "exactly once is a myth" explanation.
- **Day 4 (script):** Hook: the relatable mobile-banking "did it go through?" moment with a spinning loader. Two dry runs.

**Title:** Why You Sometimes Get Charged Twice (and How Payment Systems Prevent It)

**Description:**
A slow network, a second tap, an automatic retry, and a customer pays twice. In this episode of the Haat series we design a payment flow that can't double-charge, comparing frontend fixes, idempotency keys, database constraints and state machines. Then we face the scariest case in payments: a timeout where you don't know whether the money moved.

---

## Episode 9: Sending a push notification to 5 million phones

**Status Board:** Month 15 · 5M users · planning the Eid sale · 12 engineers
**Callback:** "Last time we made payments safe with idempotency. Today we need that idea again, because we're about to send five million messages, and some of them will be sent twice."
**Prep time:** 4 days.

**Scenario:** Haat now has a mobile app alongside the website, and push notifications work far better than paying for Facebook boosts. Marketing wants "Eid sale starts in 10 minutes!" sent to all 5 million users at exactly 8:50 PM. The first attempt is a loop that calls the push provider one user at a time. You do the math: it would finish on Thursday.

**Numbers:** 5M × 50 ms per call = ~70 hours. To finish in 5 minutes you need ~17,000 sends/sec. What are the push provider's rate limits and batch sizes?

**Approaches to compare**
- **Simple loop:** show the math for why it fails.
- **Message queue + worker pool:** split users into batches of 1,000, put the batches on the queue (the one from Episode 4, now doing serious work), and let many workers send in parallel.
- **Fan-out by segments:** precompute segments (region, language, app version, past buyers) and send per segment, or use the provider's topic feature.
- **Priority queues:** OTPs (Episode 6) and payment confirmations (Episode 8) must never wait behind a marketing blast.

**The twist:** Workers crash mid-batch and retry, so some users get the greeting 3 times while others get nothing (idempotency from Episode 8 fixes it). Then the push lands, 5 million people open Haat at the same moment, and you DDoS yourself. Fix it with staggered sends and the waiting room from Episode 7.

**Verdict:** Queue + worker pool + separate priority lanes + per-user idempotency + staggered delivery.

**Diagram adds:** Notification Service · Priority Queues · Push Workers.

**Cliffhanger:** "Notifications at the right *moment* turn out to be the easy part. Remember the 10-minute reservations from 11.11? The cron job that expires them, the same cron from Episode 1, just ran on six servers at once."

### Learning guide

**Learn in this order**
1. **Why sequential work doesn't scale:** latency × count; parallelism as the answer.
2. **Message queues:** producer, broker, consumer; acknowledgements; why queues decouple and absorb spikes.
3. **Two families of brokers:** work queues (RabbitMQ, SQS) vs logs (Kafka). Know the difference at a high level: messages deleted after processing vs kept and replayed by offset.
4. **Delivery guarantees:** at-most-once, at-least-once, "exactly once" (and why at-least-once + idempotency is the practical answer).
5. **Worker pools and batching:** choosing batch sizes, respecting provider rate limits.
6. **Priority:** separate queues/lanes for critical vs bulk messages.
7. **Dead-letter queues:** what happens to messages that keep failing.
8. **How mobile push works:** your server → FCM (Android) / APNs (iOS) → device; device tokens; topics.
9. **Thundering herd from your own notifications:** staggered/jittered sends.

**Resources**
- *Read first:*
  - Alex Xu Vol. 1, Chapter 10 "Design a notification system".
  - RabbitMQ tutorials 1-3 ("Hello World", "Work Queues", "Publish/Subscribe"): short and clear even if you never use RabbitMQ.
  - DDIA Chapter 11 "Stream Processing": the section on transmitting event streams (messaging systems, acknowledgements, logs vs queues).
- *Go deeper:*
  - Jay Kreps, "The Log: What every software engineer should know about real-time data's unifying abstraction" (2013).
  - Alex Xu Vol. 2, Chapter 4 "Distributed Message Queue".
  - Firebase Cloud Messaging documentation: the pages on topics, throttling and sending at scale.
  - AWS Builders' Library, "Avoiding insurmountable queue backlogs".
  - Amazon SQS documentation on visibility timeout and dead-letter queues (you'll reuse this in Episode 10).

**Try it yourself (optional, 1-2 hours)**
- Run RabbitMQ (or Redis lists) locally, push 10,000 fake "send" messages, start 1 worker then 10 workers, and compare total time. Kill a worker mid-run and see which messages are redelivered.

**You're ready to record when you can**
- [ ] Do the 70-hour math and the 17,000/sec math from memory.
- [ ] Draw the pipeline: who reads users, who batches, who sends, who retries, where failures go.
- [ ] Explain at-least-once delivery and why it causes duplicates.
- [ ] Explain how per-user idempotency stops triple greetings.
- [ ] Explain why OTPs need their own priority lane.
- [ ] Explain how staggering prevents the self-inflicted spike, with numbers.

**Questions viewers will ask (prepare answers)**
- *"Kafka or RabbitMQ?"* For a notification work queue either works; Kafka shines when many consumers need to replay the same stream. Don't turn the episode into a product comparison.
- *"Can't FCM send to everyone with one call using topics?"* Topics help a lot for broadcast messages. Personalised messages (names, recommendations) still need per-user sends, which is where the pipeline matters.
- *"What about SMS and email?"* Same pipeline, different provider workers, each with its own rate limits.

**Don't get this wrong on camera**
- Don't claim any queue gives "exactly once" end to end without explaining the conditions. Use the season phrase: at least once + idempotency.
- Don't quote FCM/APNs rate limits from memory; check current docs or say "the provider sets limits".

### Prep plan
- **Day 1 (learn):** Alex Xu chapter 10, RabbitMQ tutorials 1-3, DDIA chapter 11 messaging section.
- **Day 2 (solve):** Design the pipeline on paper: who reads users, who batches, who sends, who retries. Add priority lanes yourself.
- **Day 3 (numbers and twist):** Throughput math for 5 minutes; number of workers needed. Solve duplicates and the self-inflicted spike. Optional local queue experiment.
- **Day 4 (script):** Hook: "the loop finishes Thursday". Two dry runs.

**Title:** How to Send 5 Million Notifications in 5 Minutes (Without DDoSing Yourself)

**Description:**
Sending one notification is one API call. Sending five million at the same moment is a distributed systems problem. In this episode of the Haat series we design a notification system: queues, worker pools, segmentation and priority lanes. Then we fix the duplicates that crashed workers cause, and the traffic spike when millions of users open the app at once.

---

## Episode 10: The cron job that oversold us again

**Status Board:** Month 15 · 5M users · Eid sale night · 12 engineers
**Callback:** "In Episode 7 we released unpaid reservations after 10 minutes, using the cron job from Episode 1. In Episode 2 we kept that cron on server 1 only. But last month we turned on autoscaling, and every new server is a copy of server 1, crontab included. Tonight there are six servers. All six run that cron."
**Prep time:** 4 days.

**Scenario:** During the Eid sale, a cron job runs every minute: "find reservations older than 10 minutes and return their stock". It runs on several servers at once, so the same expired reservation returns its phone to stock *twice*, the counter goes up by 2, and Haat oversells again, this time by a slower route. Meanwhile, "Your cart is waiting" reminders arrive at 3 AM because the job queue fell hours behind.

**Numbers:** 500K reservations during the sale, each with its own expiry time, plus 2M scheduled reminders a day. How long does a full table scan every minute take as the table grows?

**Approaches to compare**
- **Cron job scanning the database every minute:** simple, but slow scans, and multiple servers mean duplicates. "Only run cron on one server" (what we did in Episode 2) is a single point of failure, and autoscaling just broke it.
- **Delay queue:** schedule a message to become visible at an exact time.
- **Redis sorted set scored by due time:** workers poll for items due now. (Sorted sets come back in Episode 16 for geo.)
- **Timing wheel:** an efficient in-memory structure for millions of timers.

**The twist:** A worker claims a job at 8:59:59 and crashes before finishing. Lost or done twice? Design leases (claim a job for N seconds; if not completed, it becomes visible again), idempotent job handlers ("release reservation #123" is safe to run twice because it checks the reservation's state first), and a distributed lock only where truly needed (with fencing tokens).

**Verdict:** Delay queue or sorted set for scheduling, leases for crash safety, idempotent handlers so duplicates are harmless.

**Diagram adds:** Scheduler (sorted set / delay queue) · Scheduler Workers. The cron box on server 1 is removed.

**Cliffhanger:** "Act 2 is over. We survived two sales. But the orders table, the same single database from Episode 1, now has 300 million rows, and it's at 100% CPU on a normal Tuesday."

### Learning guide

**Learn in this order**
1. **How cron works** and its limits in a multi-server world (duplicates, single point of failure, no retries, no visibility).
2. **Delay queues and visibility timeouts:** how SQS-style queues hide a message until a time, and re-show it if not acknowledged.
3. **Redis sorted sets:** score = due timestamp; `ZADD`, range by score, and atomically claiming an item (`ZREM` succeeding for only one worker, or a Lua script).
4. **Timing wheels:** the idea (a clock with buckets), why it's efficient for millions of timers.
5. **Leases:** time-limited ownership; what happens on crash and on slow workers.
6. **Idempotent job handlers:** check-state-then-act inside one atomic step.
7. **Distributed locks and their pitfalls:** process pauses, clock issues, and **fencing tokens**.
8. **At-least-once processing** (the season phrase again).

**Resources**
- *Read first:*
  - Google SRE book, chapter "Distributed Periodic Scheduling with Cron". Exactly this problem, at Google scale.
  - Redis documentation on sorted sets (`ZADD`, `ZRANGE ... BYSCORE`, `ZREM`).
  - Amazon SQS documentation: "visibility timeout" and "delay queues". Visibility timeout *is* a lease; great to explain it with a real product.
- *Go deeper:*
  - Martin Kleppmann, "How to do distributed locking" (2016): why locks need fencing tokens. Essential before saying "just use a lock".
  - DDIA Chapter 8 "The Trouble with Distributed Systems": process pauses, leases and fencing.
  - Varghese & Lauck, "Hashed and Hierarchical Timing Wheels" (1987 paper): read the introduction and the core idea only.
  - Slack engineering, "Scaling Slack's Job Queue" (2020): a real-world job queue story.

**Try it yourself (optional, 1 hour)**
- In Redis, `ZADD jobs <timestamp> job1` for several jobs, then fetch due ones with a score range. Run two "workers" that both try `ZREM` on the same job and see that only one succeeds. That's the core of safe claiming.

**You're ready to record when you can**
- [ ] Explain exactly how the counter went UP by 2, with a timeline across two servers.
- [ ] Explain why "run cron on one server" is fragile, using Haat's own history (Episode 2 → autoscaling).
- [ ] Design each approach with 10 sample reservations on paper.
- [ ] Explain leases and what happens when a worker crashes or is just slow.
- [ ] Make "release reservation #123" idempotent and explain how.
- [ ] Explain fencing tokens in one simple example.

**Questions viewers will ask (prepare answers)**
- *"Why not Kubernetes CronJob / a managed scheduler?"* They solve "run once per schedule" reasonably well, but not millions of individually-timed jobs, and you still need idempotent handlers.
- *"Is Redlock safe?"* It's debated. Present Kleppmann's argument and say: design handlers to be safe even if two workers run, and use fencing where correctness matters.
- *"Why not just a database query with `SKIP LOCKED`?"* A valid, popular approach for job queues in PostgreSQL/MySQL; mention it as another good option at moderate scale.

**Don't get this wrong on camera**
- Don't say "a distributed lock makes it safe". Locks can expire while the holder still thinks it owns them. Idempotency is the real safety net.
- Don't blame cron itself; the problem is running it on many servers without coordination.

### Prep plan
- **Day 1 (learn):** SRE book cron chapter, Redis sorted sets, SQS visibility timeout docs. Read about timing wheels.
- **Day 2 (solve):** Design each approach on paper with 10 sample reservations. Find the duplicate yourself before reading.
- **Day 3 (numbers and twist):** Walk through crash scenarios at each step. Design leases and idempotent handlers. Read Kleppmann's locking post.
- **Day 4 (script):** Hook: the "stock counter went UP by 2" reveal. Two dry runs.

**Title:** The Cron Job That Oversold Our Sale: Designing a Distributed Job Scheduler

**Description:**
Running a task at an exact time sounds trivial, until you have millions of tasks and servers that crash. In this episode of the Haat series a humble cron job from our very first version breaks our sale, and we redesign it on the whiteboard, comparing cron, delay queues, Redis sorted sets and timing wheels, then use leases and idempotency so jobs are never lost and never harmful when repeated.

---

# Act 3: Growing Up (Episodes 11-15)

*Haat is a serious company now. The database, the codebase and the team all outgrow their first versions.*

## Episode 11: The database can't keep up anymore

**Status Board:** Month 20 · 8M users · 300M order rows · 150K orders/day · 25 engineers
**Callback:** "For ten episodes, every feature we built (sessions aside) has leaned on the one database from Episode 1. Today the CTO (yes, Haat has a CTO now) asks: what do we do, and in what order?"
**Prep time:** 5-6 days.

**Scenario:** Haat's PostgreSQL primary is at 100% CPU and 90% disk. Users are growing 15% a month. The orders table has 300 million rows. The CTO wants a plan with costs, in order.

**Numbers:** Reads vs writes ratio (~20:1), data growth per month, disk runway at the current rate. (Example: 150K orders/day × ~2 KB per order with items ≈ 300 MB/day ≈ 9 GB/month for orders alone, before indexes.)

**Approaches to compare (as a staircase, cheapest first)**
- **Vertical scaling:** a bigger machine. Fast and easy, but has a ceiling and a big bill.
- **Query optimization and indexes:** often the biggest win for the least cost. Many queries from Episode 1 still have no proper indexes.
- **Caching:** callback to Episode 3; what's already offloaded and what isn't.
- **Archiving old data:** move orders older than a year to cheaper storage.
- **Read replicas:** send reads to copies. Great for read-heavy apps, but replicas lag. (Plant this; it explodes in Episode 13.)
- **Sharding:** split orders across many databases by range, by hash or by a lookup directory. Each has different trade-offs.

**The twist:** You shard by `user_id % 4`. A year later you need 5 shards, and ~80% of rows must move. Introduce consistent hashing. Then two more problems: one mega-seller's shard is 10× hotter than the rest (hot shard), and the seller dashboard needs orders across *all* shards (cross-shard queries).

**Verdict:** Climb the staircase in order; shard orders by user with consistent hashing; keep a separate path for seller analytics.

**Diagram adds:** Read Replicas · Orders Shards ×4 · Shard Router.

**Cliffhanger:** "Each shard has its own auto-increment ID, just like the orders table we designed in Episode 1. Shard 1 just created order #100234. So did shard 3."

### Learning guide

**Learn in this order**
1. **How to find what's slow:** query plans (`EXPLAIN ANALYZE`), slow query logs, missing indexes.
2. **Indexes in depth:** B-trees, composite indexes, covering indexes, the cost of too many indexes on writes.
3. **Vertical scaling limits** and cost curves.
4. **Replication:** leader/follower, synchronous vs asynchronous, replication lag (just the idea; Episode 13 goes deep).
5. **Partitioning vs sharding:** splitting a table inside one database vs across databases.
6. **Sharding strategies:** range, hash, directory/lookup; choosing a shard key (user_id vs order_id vs seller_id).
7. **Resharding and consistent hashing:** the ring, virtual nodes, how few keys move when you add a node.
8. **Hot shards / skew** and ways to relieve them.
9. **Cross-shard queries:** scatter-gather, and moving analytics to a separate system.

**Resources**
- *Read first:*
  - DDIA Chapter 5 "Replication" (leaders and followers section) and Chapter 6 "Partitioning" (whole chapter, ~25 pages).
  - Alex Xu Vol. 1, Chapter 1 (the database scaling section) and Chapter 5 "Design consistent hashing".
- *Go deeper:*
  - "Use The Index, Luke" (use-the-index-luke.com): indexes explained for developers.
  - Notion engineering, "Herding elephants: lessons learned from sharding Postgres at Notion" (2021).
  - Figma engineering, "How Figma's databases team lived to tell the scale" (2024).
  - Instagram engineering, "Sharding & IDs at Instagram" (also used in Episode 12).
  - The Amazon Dynamo paper (2007), section on consistent hashing and virtual nodes.
  - Vitess or Citus documentation (how real tools shard MySQL/PostgreSQL).
- *Watch:* Martin Kleppmann's Cambridge lectures on replication; Arpit Bhayani on sharding and consistent hashing.

**Try it yourself (optional, 2 hours)**
- On paper or in a spreadsheet, assign 12 user IDs to 4 shards with `id % 4`, then to 5 shards with `id % 5`, and count how many move. Then place the same users on a hash ring with 4 and 5 nodes and count again. This is the core demo of the episode.

**You're ready to record when you can**
- [ ] Present the staircase and explain why each step comes before the next.
- [ ] Read a simple `EXPLAIN` output and spot a missing index.
- [ ] Compare range, hash and directory sharding with one advantage and one problem each.
- [ ] Explain why `% N` moves most keys when N changes, and how consistent hashing moves only about 1/N.
- [ ] Explain virtual nodes in one sentence.
- [ ] Explain the hot shard problem and the cross-shard query problem with Haat examples.

**Questions viewers will ask (prepare answers)**
- *"Why not just switch to MongoDB / Cassandra / a NewSQL database?"* They shard automatically but come with their own trade-offs and a full migration. Distributed SQL databases (CockroachDB, YugabyteDB, TiDB, Spanner-style) are a legitimate option; mention them honestly.
- *"Why shard by user and not by seller?"* Most queries are "my orders"; sharding by user keeps them on one shard. Seller views go to a separate analytics path.
- *"When should I shard?"* As late as possible, after the cheaper steps. That's the whole point of the staircase.

**Don't get this wrong on camera**
- Don't jump straight to sharding; the staircase order is the lesson.
- Don't say consistent hashing "moves no data"; it moves roughly 1/N of it instead of most of it.

### Prep plan
- **Day 1 (learn):** DDIA chapters 5 (first part) and 6. Paragraph per concept.
- **Day 2 (solve):** Shard 12 sample users on paper with each strategy; add a 5th shard and count how many move.
- **Day 3 (numbers and twist):** Learn consistent hashing and redo the exercise. Calculate disk runway. Read the Notion or Figma story. Prepare the staircase diagram.
- **Day 4 (polish):** Build the staircase as one strong visual. Decide which real company story you'll tell.
- **Day 5 (script):** Hook: the CTO's question. Two dry runs; this episode is dense, so rehearse transitions between staircase steps.

**Title:** Your Database Is at 100%: What to Do First (Scaling Explained Step by Step)

**Description:**
Haat is growing fast and the database we created in Episode 1 can't keep up. Do you buy a bigger server, add replicas or shard everything? In this episode we walk the scaling staircase on the whiteboard, from the cheapest fixes to sharding, compare range, hash and directory-based sharding, and see why consistent hashing saves you when it's time to add more machines.

---

## Episode 12: Two customers got the same order number

**Status Board:** Month 21 · 8M users · 4 order shards · 160K orders/day · 25 engineers
**Callback:** "Last time we sharded the orders table. We forgot one tiny thing: every shard was still counting 1, 2, 3 on its own, exactly the way we set up IDs in Episode 1."
**Prep time:** 4 days.

**Scenario:** Two customers both receive "Order #100234 confirmed". One calls support, the agent looks up #100234 and reads the *other* customer's address and phone number aloud. Now it's a privacy incident. Separately, sellers have been sharing product links like `haat.com.bd/product/48213` on their Facebook pages, and a competitor is reading your daily sales volume from how fast the IDs grow.

**Numbers:** 160K orders/day today, 10× in 3 years. How many IDs per second at peak sales? How long must IDs stay unique (forever)?

**Approaches to compare**
- **Central ticket server:** one database hands out IDs. Simple and ordered, but a bottleneck and single point of failure.
- **Auto-increment with offsets:** shard k starts at k and steps by N. Works until you change N.
- **UUID (v4):** no coordination, but 128 bits, random order (bad for index locality), and ugly for humans. Mention UUIDv7, which is time-ordered.
- **Snowflake-style IDs:** timestamp + machine ID + sequence. Roughly time-ordered, no coordination, 64 bits.
- **Pre-generated ID ranges:** each server grabs a block of 1,000 IDs at a time.

**The twist:** Short, shareable links. Sellers want `haat.com.bd/p/x7Kq2` to paste on Facebook, plus custom names like `haat.com.bd/rahim-fashion`. Introduce Base62 encoding, why sequential short codes leak your business numbers (the German tank problem), and how custom aliases coexist with generated codes.

**Verdict:** Snowflake IDs internally; non-sequential Base62 codes for anything public; a reserved namespace for custom shop links.

**Diagram adds:** ID Generator (library in each service) · Short Link Service.

**Cliffhanger:** "Orders are unique now. But a customer just placed one, opened My Orders, and it's empty. They placed it again."

### Learning guide

**Learn in this order**
1. **Why auto-increment breaks across shards** (and why it leaks information even on one database).
2. **Requirements for IDs:** unique, fits in 64 bits?, sortable by time?, generated without coordination?, safe to expose?
3. **Ticket servers** (Flickr's approach) and their trade-offs.
4. **UUIDs:** v4 (random) vs v7 (time-ordered); why random keys hurt B-tree indexes.
5. **Snowflake layout:** sign bit, timestamp, machine/worker ID, sequence; how many IDs per millisecond per machine.
6. **Clock problems:** what happens if a server's clock goes backwards.
7. **Base62 encoding:** how a number becomes `x7Kq2`; how many codes a given length gives.
8. **Information leakage:** the German tank problem; why public IDs should be non-sequential.
9. **HTTP redirects:** 301 vs 302 for short links (caching vs analytics).

**Resources**
- *Read first:*
  - Alex Xu Vol. 1, Chapter 7 "Design a unique ID generator in distributed systems" and Chapter 8 "Design a URL shortener".
  - Instagram engineering, "Sharding & IDs at Instagram".
  - Twitter engineering blog, "Announcing Snowflake" (2010).
- *Go deeper:*
  - Flickr engineering, "Ticket Servers: Distributed Unique Primary Keys on the Cheap" (2010).
  - RFC 9562 (2024), the UUID specification including UUIDv7: read the motivation section.
  - Wikipedia, "German tank problem": a great story for the episode.
  - The Sqids library (formerly Hashids): one way to turn numbers into short, non-sequential-looking codes.
  - MDN Web Docs: "301 Moved Permanently" and "302 Found".

**Try it yourself (optional, 1 hour)**
- Write the Snowflake bit layout on paper and build one ID by hand from a timestamp, machine ID 5 and sequence 12. Encode the number 125,000 in Base62 by hand.

**You're ready to record when you can**
- [ ] List every approach and the exact way it fails.
- [ ] Draw the Snowflake bit layout and calculate IDs per millisecond per machine and the number of years it lasts.
- [ ] Explain why random UUIDs slow down inserts into a B-tree index.
- [ ] Calculate how many 7-character Base62 codes exist (62⁷ ≈ 3.5 trillion) and how long they'd last.
- [ ] Tell the German tank problem in 30 seconds and connect it to Haat's product IDs.
- [ ] Explain 301 vs 302 and which one Haat uses for shareable links and why.

**Questions viewers will ask (prepare answers)**
- *"Why not just UUID everything?"* Many teams do, and it works. The costs are size, index locality (less so with UUIDv7) and ugliness for humans. Snowflake fits in 64 bits and sorts by time.
- *"What if two machines get the same machine ID?"* You'll get duplicates. Machine IDs must be assigned centrally or from a coordination service when a machine starts.
- *"Isn't hiding IDs just security through obscurity?"* Non-sequential IDs aren't for security (you still check permissions); they stop accidental information leaks like sales volume.

**Don't get this wrong on camera**
- Don't say Snowflake IDs are "perfectly ordered"; they're roughly time-ordered across machines.
- Don't mix up Base62 and Base64 (Base64 includes `+` and `/`, which are awkward in URLs).

### Prep plan
- **Day 1 (learn):** Alex Xu chapters 7 and 8, the Instagram post and the Snowflake announcement.
- **Day 2 (solve):** On paper, list every approach and the exact way it fails. Design the short link scheme yourself.
- **Day 3 (numbers and twist):** Snowflake capacity per millisecond, years before the timestamp runs out, number of 7-character Base62 codes. Prepare the German tank story.
- **Day 4 (script):** Hook: the support agent reading the wrong customer's address. Two dry runs.

**Title:** Two Customers Got the Same Order Number: Designing Unique IDs at Scale

**Description:**
After sharding our database, Haat started handing out the same order number twice. In this episode we compare ticket servers, offset auto-increment, UUIDs, Snowflake IDs and pre-allocated ranges on the whiteboard, then design short shareable links for sellers' Facebook posts and see how sequential IDs can leak your sales numbers to competitors.

---

## Episode 13: "I placed an order but My Orders is empty"

**Status Board:** Month 22 · 10M users · 200K orders/day · 30 engineers
**Callback:** "In Episode 11 we added read replicas to take load off the primary. Replicas are copies, and copies are always a little bit behind."
**Prep time:** 5-6 days.

**Scenario:** A customer pays, sees "Order confirmed", taps My Orders, and the list is empty. They panic and order again (and thanks to Episode 8, that's a *different* order, not a retry). Support tickets pile up. A seller changes a price, but checkout charges the old price for 30 seconds. The cause: reads go to replicas and caches that haven't caught up.

**Numbers:** Typical replication lag (100 ms - 2 s), spiking to 30 s+ during sales. Which reads must be fresh, and how many reads/sec would hit the primary if they all were?

**Approaches to compare**
- **Strong consistency: all reads from the primary.** Always correct, but you lose the scaling benefit of Episode 11.
- **Eventual consistency:** accept short staleness. Fine for review counts, wrong for My Orders.
- **Read-your-own-writes:** after a user writes, route *their* reads to the primary for a few seconds, or until the replica has caught up to their write's position.
- **Monotonic reads / sticky replica:** keep a user on the same replica so they never go back in time.
- **Cache invalidation on write and versioned keys:** so caches can't serve old prices.

**The twist:** Sort Haat's data into tiers: what must be strongly consistent (stock, payment status, price at checkout), what needs read-your-writes (my orders, my cart, my profile), and what can be eventual (reviews, view counts, recommendations). Then introduce the idea that will return in the finale: **you can't have perfect freshness and perfect speed everywhere at once.**

**Verdict:** Read-your-own-writes for the user's own data, primary reads for money and stock, eventual consistency everywhere else.

**Diagram adds:** Consistency-aware read routing in the data layer.

**Cliffhanger:** "Thirty engineers, one codebase, one deploy pipeline, the same app we started in Episode 1. Every release is a fight. The CTO says the words every engineer has heard: 'Let's split it into microservices.'"

### Learning guide

**Learn in this order**
1. **Replication lag:** why asynchronous replicas fall behind and what makes lag spike (heavy writes, long queries, network).
2. **Read-after-write (read-your-own-writes) consistency.**
3. **Monotonic reads** (never going back in time) and **consistent prefix reads** (seeing causes before effects).
4. **Strong vs eventual consistency** in plain words; linearizability as "behaves like a single copy".
5. **Implementation techniques:** reading from the primary for a short window, tracking the write's log position (LSN) and waiting for the replica, sticky sessions to a replica.
6. **Caches as another source of staleness:** invalidation on write, versioned keys.
7. **Consistency tiers:** classifying data by how fresh it must be.
8. **CAP and PACELC intuition (light):** even without failures, there's a trade-off between latency and consistency.

**Resources**
- *Read first:*
  - DDIA Chapter 5, section "Problems with Replication Lag" (reading your own writes, monotonic reads, consistent prefix reads). The core resource; about 10 pages.
  - Werner Vogels, "Eventually Consistent" (All Things Distributed / ACM Queue, 2008).
- *Go deeper:*
  - Jepsen, "Consistency Models" (jepsen.io/consistency): a clickable map of consistency models.
  - DDIA Chapter 9 "Consistency and Consensus", the linearizability section.
  - Daniel Abadi, "Consistency Tradeoffs in Modern Distributed Database System Design" (2012): the PACELC paper; read the introduction.
  - "Scaling Memcache at Facebook" (NSDI 2013): the cross-region consistency section (remote markers) is a real read-your-writes story.
  - PostgreSQL documentation on streaming replication and monitoring replication lag.
- *Watch:* Martin Kleppmann's Cambridge lectures on consistency models and linearizability.

**Try it yourself (optional, 1 hour)**
- On paper, draw three timelines (primary, replica, cache) and place the customer's write and reads on them with lag. Show each fix as a change in which timeline the read goes to. This drawing becomes your main panel.

**You're ready to record when you can**
- [ ] Draw the empty-My-Orders bug as a timeline across primary, replica and cache.
- [ ] Explain read-your-own-writes, monotonic reads and consistent prefix with one Haat example each.
- [ ] Explain two ways to implement read-your-own-writes and their costs.
- [ ] Build the data-tier table for Haat and defend each row.
- [ ] Calculate primary load if all reads were strongly consistent vs with tiers.

**Questions viewers will ask (prepare answers)**
- *"Why not use synchronous replication?"* Then every write waits for replicas; one slow replica slows or blocks writes. Some systems use one synchronous replica as a middle ground.
- *"Isn't this what CAP is about?"* Partly. CAP is specifically about network partitions; replication lag is the everyday latency vs consistency trade-off (PACELC). Save full CAP for the finale.
- *"How long should 'read from primary' last after a write?"* Longer than typical lag (a few seconds), or better, until the replica has caught up to the write's position.

**Don't get this wrong on camera**
- Don't define "eventual consistency" as "it becomes consistent in a few seconds"; it only promises it becomes consistent *eventually*, with no time bound.
- Don't present CAP as "pick 2 of 3"; that framing is misleading. You'll handle CAP properly in Episode 21.

### Prep plan
- **Day 1 (learn):** DDIA chapter 5 replication lag section and Vogels' article.
- **Day 2 (solve):** Draw the empty-My-Orders bug as a timeline across primary, replica and cache. Design fixes before reading further.
- **Day 3 (numbers):** Build the data-tier table for Haat. Calculate primary load for each routing strategy.
- **Day 4 (depth):** Jepsen consistency map and a Kleppmann lecture, to make sure every term you use on screen is precise.
- **Day 5 (script):** Hook: "order vanished, so they ordered again". Two dry runs.

**Title:** "I Placed an Order But It's Not There": Consistency in Distributed Databases

**Description:**
You buy something, open your orders, and it's not there. It's not a bug in the code; it's replication lag. In this episode of the Haat series we recreate the bug on the whiteboard and compare the fixes: strong consistency, eventual consistency, read-your-own-writes and monotonic reads, then decide which of Haat's data needs which guarantee.

---

## Episode 14: Paid, but never shipped

**Status Board:** Month 26 · 12M users · 300K orders/day · 80 engineers · 8 services
**Callback:** "Last time, our engineering team outgrew the monolith we started in Episode 1. So we split Haat into services: Order, Payment, Inventory, Seller, Shipping. Each with its own database. What could go wrong?"
**Prep time:** 5-6 days.

**Scenario:** A customer pays ৳3,200. The Payment service succeeds. The Inventory service then crashes before reserving the item. The money is taken, nothing is reserved, nothing ships, and there's no single database transaction that can roll it all back. Hundreds of these orders sit in limbo every week.

**Numbers:** 300K orders/day across 4-5 service hops each. Even a 0.1% failure rate per hop means hundreds of broken orders per day.

**Approaches to compare**
- **One big database transaction:** impossible across separate services with separate databases.
- **Two-phase commit (2PC):** a coordinator asks everyone to prepare, then commit. Correct, but slow, and everyone blocks if the coordinator dies.
- **Saga with choreography:** each service reacts to events on the queue and publishes its own; on failure, services run compensating actions (refund, release stock). Decoupled, but the flow is hard to follow.
- **Saga with orchestration:** one Order Orchestrator tells each service what to do and handles failures. Easier to reason about, but it's a central piece.

**The twist:** The Order service saves the order and crashes *before* publishing "OrderCreated". The database and the event stream now disagree. Introduce the **transactional outbox**. Then: what if the refund itself fails? Compensations must be retried, idempotent (Episode 8 again) and eventually escalated to a human.

**Verdict:** Orchestrated saga for checkout, outbox for every event, idempotent compensations. Add an honest note: splitting into services was a team decision, not a traffic decision, and a well-structured monolith can go a long way.

**Diagram adds:** The monolith box becomes Order / Payment / Inventory / Seller / Shipping services · Order Orchestrator · Outbox tables · Event Bus.

**Cliffhanger:** "Our checkout now calls five services in a row. One of them belongs to our courier partner, the same courier the owner used to book by phone back in Episode 1. It's about to get very, very slow."

### Learning guide

**Learn in this order**
1. **Why teams split monoliths:** team independence, deploy conflicts, ownership. And the costs: network calls, partial failure, data spread across databases.
2. **Database per service** and why it removes cross-service transactions.
3. **Two-phase commit:** prepare and commit phases, the blocking problem when the coordinator fails.
4. **The saga pattern:** a sequence of local transactions, each with a compensating action.
5. **Choreography vs orchestration:** event-driven vs central coordinator; draw both for Haat's checkout.
6. **Compensating transactions:** semantic undo (refund, not delete); compensations that fail.
7. **The dual-write problem** between a database and a message broker.
8. **Transactional outbox:** write the event into an outbox table in the same transaction; a relay (polling or CDC) publishes it.
9. **Idempotent consumers** (again): events may arrive more than once.

**Resources**
- *Read first:*
  - microservices.io (Chris Richardson): the "Saga" and "Transactional outbox" pattern pages.
  - DDIA Chapter 9, the section on atomic commit and two-phase commit.
  - Martin Fowler, "MonolithFirst" again, plus Shopify's modular monolith posts, so your verdict is balanced.
- *Go deeper:*
  - Chris Richardson, *Microservices Patterns*: Chapter 3 (messaging, transactional outbox) and Chapter 4 "Managing transactions with sagas".
  - Sam Newman, *Building Microservices* (2nd edition): the chapter on workflow and sagas.
  - Caitie McCaffrey, "Distributed Sagas: A Protocol for Coordinating Microservices" (conference talk, 2017).
  - Garcia-Molina & Salem, "Sagas" (1987 paper): the original idea; the introduction is enough.
  - Debezium documentation, "Outbox Event Router".
  - Temporal documentation (an example of a workflow/orchestration engine).
- *Watch:* Martin Kleppmann's Cambridge lecture on two-phase commit.

**Try it yourself (optional, 1-2 hours)**
- On paper, list every step of Haat's checkout as a saga with its compensation in a two-column table. Then mark each arrow where a crash could happen and write what the system does. This table is half your episode.

**You're ready to record when you can**
- [ ] Explain why a single transaction is impossible across services.
- [ ] Draw 2PC and show exactly where it blocks.
- [ ] Draw Haat's checkout as a choreography saga and an orchestration saga, including compensation.
- [ ] Explain the dual-write problem and how the outbox fixes it.
- [ ] Walk through every crash point in the flow and say what happens.
- [ ] Explain what happens when a refund (compensation) fails.

**Questions viewers will ask (prepare answers)**
- *"Should we even have split into microservices?"* Fair question, and you can say so. Haat did it for team reasons at 80 engineers. A modular monolith is a strong alternative.
- *"Choreography or orchestration?"* Choreography for simple flows with few steps; orchestration for business-critical, multi-step flows like checkout where you need visibility.
- *"Isn't the outbox just a queue in the database?"* Essentially yes, and that's the trick: it's written in the same transaction as your data.

**Don't get this wrong on camera**
- Don't present sagas as giving the same guarantees as ACID transactions. Sagas lack isolation; other requests can see intermediate states.
- Don't call 2PC "wrong"; it's used in practice, it's just costly across independent services.

### Prep plan
- **Day 1 (learn):** microservices.io saga and outbox pages; DDIA 2PC section.
- **Day 2 (solve):** Draw Haat's checkout as both a choreography saga and an orchestration saga, including compensation.
- **Day 3 (twist):** Learn the outbox pattern. Walk through every crash point in the flow.
- **Day 4 (balance):** Read Fowler's MonolithFirst and a Shopify modular monolith post so your verdict is honest.
- **Day 5 (script):** Hook: "paid but nothing happened". Two dry runs.

**Title:** Paid But Never Shipped: Handling Transactions Across Microservices (Saga Pattern)

**Description:**
In a monolith, one database transaction keeps everything consistent. In microservices, a payment can succeed while the next step fails. In this episode of the Haat series we design checkout across services and compare two-phase commit, choreography sagas and orchestration sagas, then fix the subtle bug where your database and your events disagree using the transactional outbox.

---

## Episode 15: One slow courier API took down all of Haat

**Status Board:** Month 28 · 14M users · Pohela Boishakh sale · 90 engineers
**Callback:** "Last time we made checkout a saga across services. One of those calls goes to our courier partner for a delivery quote. Today, that partner is having a bad day."
**Prep time:** 5-6 days.

**Scenario:** Mid-sale, the courier partner's API starts answering in 30 seconds instead of 300 milliseconds. Checkout threads wait. The thread pools fill. Then search, product pages and even the homepage, none of which need the courier, stop responding. Within six minutes all of Haat is down, because of someone else's server.

**Numbers:** 200 threads per instance ÷ 30 s per call means each instance handles ~7 requests/sec instead of hundreds. Draw it filling up minute by minute.

**Approaches to compare**
- **Timeouts:** stop waiting forever. But what's the right value, and why is "30 seconds" almost always wrong?
- **Retries with exponential backoff and jitter:** retry smartly, because naive retries multiply the load on a struggling service.
- **Circuit breaker:** after repeated failures, stop calling for a while and fail fast. Closed → Open → Half-open.
- **Bulkheads:** separate thread/connection pools per dependency, so the courier can only exhaust *its* pool.
- **Graceful degradation:** show a flat estimated delivery fee and confirm later instead of failing checkout.

**The twist:** The courier recovers, and every client retries at the same instant (a retry storm) and knocks it down again. Why jitter matters, and how load shedding (Episode 6's rate limiting, now protecting a service) gives it room to recover.

**Verdict:** Tight timeouts + backoff with jitter + circuit breaker + bulkheads + a degraded fallback for every non-critical dependency.

**Diagram adds:** Circuit breakers and bulkheads around every external call.

**Cliffhanger:** "After the incident, the owner, now the CEO, makes a decision: 'We will never depend on someone else's riders again.' Haat is launching its own delivery fleet. Fifty thousand riders. Who gets which order?"

### Learning guide

**Learn in this order**
1. **Resource exhaustion:** threads, connections and memory as finite pools; how slow calls fill them.
2. **Cascading failure:** how one slow dependency spreads to unrelated features (shared pools, shared servers).
3. **Timeouts:** connect vs read timeouts; choosing values from latency percentiles (p99), not guesses.
4. **Retries:** when to retry (idempotent operations only), retry budgets, exponential backoff, **jitter**.
5. **Circuit breaker:** closed, open, half-open; thresholds and probe requests.
6. **Bulkheads:** isolating pools per dependency (the ship compartment analogy).
7. **Fail fast and graceful degradation:** fallbacks that keep the core flow working.
8. **Load shedding:** rejecting work early to protect a struggling service.
9. **Retry storms and metastable failures** at an intuitive level.

**Resources**
- *Read first:*
  - Michael Nygard, *Release It!* (2nd edition): the chapters on stability antipatterns (integration points, cascading failures, blocked threads) and stability patterns (timeouts, circuit breaker, bulkheads, fail fast, shed load).
  - Martin Fowler, "CircuitBreaker" (martinfowler.com).
  - Marc Brooker, "Exponential Backoff And Jitter" (AWS Architecture Blog, 2015): the charts are perfect for your whiteboard.
- *Go deeper:*
  - AWS Builders' Library: "Timeouts, retries, and backoff with jitter" and "Using load shedding to avoid overload".
  - Google SRE book: chapters "Addressing Cascading Failures" and "Handling Overload".
  - Netflix Hystrix wiki, "How it Works" (the library is retired, but the explanation of bulkheads and circuit breakers is excellent).
  - Resilience4j documentation (a current library implementing these patterns).

**Try it yourself (optional, 1 hour)**
- Draw a timeline of one app instance with a 200-thread pool, minute by minute, as courier calls take 30 s. Then redraw it with a 1 s timeout, then with a bulkhead of 20 threads for the courier. These three drawings are the backbone of the episode.

**You're ready to record when you can**
- [ ] Draw the failure spreading through Haat minute by minute, and explain why the homepage died.
- [ ] Calculate requests/sec per instance with 30 s calls vs 300 ms calls.
- [ ] Explain how to choose a timeout from p99 latency.
- [ ] Draw the circuit breaker state machine and explain half-open.
- [ ] Explain why retries without jitter cause a retry storm.
- [ ] Describe Haat's fallback for the courier quote and why it's acceptable to the business.

**Questions viewers will ask (prepare answers)**
- *"Should we retry at all?"* Only idempotent operations, with a limit and backoff. Sometimes the right answer is to not retry and use the fallback.
- *"Doesn't a service mesh (Istio, Linkerd) do this?"* It can provide timeouts, retries and circuit breaking at the network layer. You still need to choose the values and design the fallbacks.
- *"What's a good circuit breaker threshold?"* It depends on traffic and normal error rates; start from observed data, e.g. open after 50% failures over a window with a minimum number of calls.

**Don't get this wrong on camera**
- Don't say "add retries" as a fix on its own; retries without limits make outages worse.
- Don't confuse a circuit breaker with a rate limiter: one stops calls to a failing dependency, the other limits incoming traffic.

### Prep plan
- **Day 1 (learn):** *Release It!* stability chapters; Fowler's CircuitBreaker; Brooker's jitter post.
- **Day 2 (solve):** Draw the failure spreading through Haat's services minute by minute, then add each defense.
- **Day 3 (twist):** Study retry storms and jitter. Draw the circuit breaker state machine.
- **Day 4 (depth):** SRE book cascading failures chapter, to collect one or two strong real-world quotes.
- **Day 5 (script):** Script it as a minute-by-minute incident report. Two dry runs.

**Title:** How One Slow API Took Down Our Entire App (and the Patterns That Prevent It)

**Description:**
A single slow dependency can bring down a whole system, even the parts that don't use it. In this episode of the Haat series we replay a cascading failure minute by minute on the whiteboard, then add the defenses one by one: timeouts, retries with backoff and jitter, circuit breakers, bulkheads and graceful degradation.

---

# Act 4: Haat Comes Alive (Episodes 16-19)

*Haat stops being just a website. It gets its own riders, real-time tracking, live shopping and a social feed.*

## Episode 16: Finding the nearest rider among 50,000

**Status Board:** Month 31 · 16M users · 50K riders · 400K orders/day · 110 engineers
**Callback:** "After Episode 15's outage, Haat launched its own delivery fleet. Fifty thousand riders across Dhaka and Chattogram. Our first rider-assignment code just took 4 seconds per order."
**Prep time:** 4-5 days.

**Scenario:** When a seller marks an order ready, Haat must find the nearest available rider in under a second. Riders send their location every 4 seconds. Version 1 calculates the distance to all 50,000 riders for every order. At lunch rush it falls minutes behind, and food orders (a new category) arrive cold.

**Numbers:** 50K riders ÷ 4 s = 12,500 location writes/sec. 400K orders/day with a lunch peak of ~30 assignments/sec, each needing a "nearest" query.

**Approaches to compare**
- **Brute force distance:** show the math for why it collapses.
- **Fixed grid:** split the city into squares and search your square and its neighbors. Simple, but Gulshan at rush hour and an empty field get the same square size.
- **Quadtree:** split busy areas into smaller squares. Adapts to density, but costly to update as riders move.
- **Geohash:** encode location as a string where nearby places share a prefix. Easy to store in Redis or any database (Redis has built-in geo commands on sorted sets, a callback to Episode 10).
- **Hexagon grid (Uber's H3):** all neighbors are equally distant, which helps for zones and surge.

**The twist:** Two problems. First, write-heavy: 12,500 updates/sec. Indexes that are fast to read are often slow to update, so which survives constant movement? Second, straight-line nearest isn't road nearest. A rider 400 m away across the Buriganga or a one-way road can be 25 minutes away. Introduce candidate filtering by distance, then ranking by ETA.

**Verdict:** In-memory geohash/H3 index for candidates, ETA-based ranking on the top 20, location writes kept out of the main database.

**Diagram adds:** Rider Location Service (in-memory geo index) · Assignment Service.

**Cliffhanger:** "We know where every rider is, every four seconds. Customers want to know too: 'Where is my rider?' Right now the app refreshes the page."

### Learning guide

**Learn in this order**
1. **Why "nearest" is hard at scale:** brute force cost = riders × queries.
2. **Distance basics:** latitude/longitude, the haversine formula (just know it exists and what it does).
3. **Fixed grids:** cell lookups and neighbor cells; the density problem.
4. **Quadtrees:** recursive splitting by density; building, querying, updating.
5. **Geohash:** interleaving bits, Base32, shared prefixes; the **edge problem** (close points with different prefixes) and why you search neighbor cells.
6. **Hexagonal grids (H3)** and why equal neighbor distance matters; S2 as Google's alternative.
7. **Read-heavy vs write-heavy indexes:** what survives 12,500 updates/sec.
8. **Candidate generation then ranking:** filter by distance, rank by ETA (road network, traffic).
9. **Redis GEO commands** as a practical implementation.

**Resources**
- *Read first:*
  - Alex Xu Vol. 2, Chapter 1 "Proximity Service": geohash, quadtree and S2 compared clearly. The core resource.
  - Uber engineering blog, "H3: Uber's Hexagonal Hierarchical Spatial Index" (2018).
- *Go deeper:*
  - Alex Xu Vol. 2, Chapter 2 "Nearby Friends": handling constantly-moving locations.
  - h3geo.org documentation (with interactive examples you can use on screen).
  - Redis documentation: `GEOADD` and `GEOSEARCH`.
  - Wikipedia, "Geohash" (good diagrams of the prefix idea).
  - s2geometry.io overview (Google's S2).
- *Real-world stories:* Uber and Grab engineering posts on dispatch and ETA prediction (search "<company> dispatch system" or "ETA prediction").

**Try it yourself (optional, 1-2 hours)**
- Encode a real Dhaka location (for example, a well-known landmark in Gulshan) as a geohash by hand to 5 characters, then check it with an online geohash tool. Use the H3 website to view the hexagon over the same place. Both make strong on-screen visuals.

**You're ready to record when you can**
- [ ] Do the brute-force math and the 12,500 writes/sec math.
- [ ] Encode a location as a geohash by hand and explain shared prefixes.
- [ ] Explain the geohash edge problem and the neighbor-cell fix.
- [ ] Draw a quadtree for 15 riders and explain why updates are costly.
- [ ] Explain why hexagons are better for zones than squares.
- [ ] Explain "candidates by distance, rank by ETA" with the Buriganga example.

**Questions viewers will ask (prepare answers)**
- *"Why not PostGIS?"* PostGIS is excellent for geographic queries; the challenge here is 12,500 writes/sec of constantly-moving points, which is why the hot path lives in memory. PostGIS fits zones and history.
- *"How do you get ETA?"* A routing engine on the road network (OSRM and similar), adjusted with historical traffic data. In Season 2 this becomes an ML prediction.
- *"What about riders who are 'available' but not moving?"* Rider state (available, busy, offline) is stored with the location; only available riders are candidates.

**Don't get this wrong on camera**
- Don't say geohash prefixes guarantee closeness both ways; nearby points can have different prefixes (the edge problem).
- Don't claim to know how Pathao, Foodpanda or Uber assign riders internally; describe the common pattern.

### Prep plan
- **Day 1 (learn):** Alex Xu Vol. 2 chapter 1; Uber's H3 post.
- **Day 2 (solve):** Encode a real Dhaka location as a geohash by hand. Draw a quadtree for 15 riders.
- **Day 3 (numbers and twist):** Calculate update volume. Compare structures for read speed vs update cost. Design candidate filtering + ETA ranking.
- **Day 4 (visuals):** Use a real Dhaka map as the whiteboard background; overlay a grid, geohash cells and hexagons.
- **Day 5 (script):** Hook: a cold biryani arriving 50 minutes late. Two dry runs.

**Title:** How Delivery Apps Find the Nearest Rider in Milliseconds

**Description:**
Finding the nearest rider sounds easy, until there are 50,000 of them moving every few seconds. In this episode of the Haat series we compare brute force, fixed grids, quadtrees, geohash and hexagon grids on a real Dhaka map, and discover why the best structure for searching isn't always the best for constant updates, and why "nearest" isn't the same as "fastest".

---

## Episode 17: "Where is my rider?"

**Status Board:** Month 34 · 18M users · 500K orders/day · 130 engineers
**Callback:** "Last time we built a live map of 50,000 riders inside our servers. Now customers want to watch their rider move, and send a message: 'I'm at the blue gate.'"
**Prep time:** 5-6 days.

**Scenario:** Customers refresh the tracking page every few seconds. The rider dot jumps around or freezes. Messages between customer and rider arrive late and out of order ("ok coming" before "which gate?"), and the "delivered" tick lies. Ratings drop.

**Numbers:** 150K orders in delivery at peak, each with a watching customer, so 150K open connections receiving a location every 4 s (~37,500 messages/sec). Chat volume per order. Connections per server.

**Approaches to compare**
- **Short polling:** "any update?" every 2 seconds. Simple, but wastes battery and server capacity.
- **Long polling:** the server holds the request until there's news. Better, but still clumsy.
- **Server-Sent Events (SSE):** one-way stream from server to client. Perfect for location updates.
- **WebSockets:** one persistent two-way connection. Real-time chat, but now you track hundreds of thousands of open connections across servers.
- **Push notifications** (Episode 9) for users who closed the app.

**The twist:** The customer is connected to gateway server 3; the rider's messages arrive at server 7. How does a message cross between them? Introduce pub/sub and a connection registry. Then: message ordering with per-conversation sequence numbers, and the meaning of the ticks (sent → delivered → read) as acknowledgements.

**Verdict:** WebSockets for chat, SSE or WebSocket for tracking, pub/sub to route between gateway servers, sequence numbers for order, push for offline.

**Diagram adds:** Real-time Gateway (WebSocket servers) · Pub/Sub · Connection Registry · Chat Store.

**Cliffhanger:** "We built a real-time pipe to every customer's phone. A seller who used to sell through Facebook Live looked at it and asked: 'Can I go *live* here?'"

### Learning guide

**Learn in this order**
1. **HTTP request/response limits** for real-time data.
2. **Short polling, long polling, SSE, WebSockets:** how each works on the wire, draw the request flow for each.
3. **Stateful connections:** why WebSocket servers are stateful and what that means for load balancing (sticky routing, connection draining on deploy).
4. **Connections per server:** memory per connection, file descriptor limits, heartbeats.
5. **Pub/sub:** publishers, channels/topics, subscribers; routing messages between gateway servers.
6. **Connection registry:** "which server holds user X's connection?"
7. **Message ordering:** per-conversation sequence numbers; why timestamps alone aren't enough.
8. **Delivery receipts:** sent/delivered/read as acknowledgements at different stages.
9. **Offline delivery:** storing messages and falling back to push.
10. **Reconnection:** resuming from the last received sequence number.

**Resources**
- *Read first:*
  - Alex Xu Vol. 1, Chapter 12 "Design a chat system". The core resource.
  - MDN Web Docs: "The WebSocket API" and "Using server-sent events".
- *Go deeper:*
  - Alex Xu Vol. 2, Chapter 2 "Nearby Friends": location updates over WebSockets with pub/sub.
  - Discord engineering, "How Discord Scaled Elixir to 5,000,000 Concurrent Users" (2017).
  - Slack engineering's posts on their real-time messaging architecture.
  - Redis documentation, "Pub/Sub" (and its limits: messages aren't stored).
  - RFC 6455 (The WebSocket Protocol): read only the introduction and the opening handshake.
- *Watch:* Hussein Nasser's videos on WebSockets, SSE and long polling.

**Try it yourself (optional, 1-2 hours)**
- Use a tiny WebSocket echo server (many frameworks have a 20-line example), open it from two browser tabs, and watch messages flow. Then open your browser's developer tools Network tab and look at the WebSocket frames. Showing the "Upgrade" handshake is a great 20-second on-screen moment.

**You're ready to record when you can**
- [ ] Draw the request flow for all four methods and compare battery, server cost and latency.
- [ ] Calculate connections and messages per second at peak, and servers needed.
- [ ] Design the message flow: send → gateway → store → route → deliver → ack → ticks.
- [ ] Explain cross-server routing with pub/sub and a connection registry.
- [ ] Explain why "ok coming" can arrive before "which gate?" and how sequence numbers fix it.
- [ ] Explain what each tick means and which component sets it.

**Questions viewers will ask (prepare answers)**
- *"Why not WebSockets for everything?"* SSE is simpler for one-way streams, works over normal HTTP and reconnects automatically. Use the simplest tool that fits.
- *"What about MQTT?"* A lightweight protocol often used for mobile and IoT; a valid alternative for rider location streams.
- *"How does WhatsApp do it?"* Don't claim internal details beyond public talks; describe the common pattern.

**Don't get this wrong on camera**
- Don't say WebSockets "don't use HTTP"; they start with an HTTP Upgrade handshake.
- Don't forget the deploy problem: restarting a WebSocket server drops connections, so clients must reconnect gracefully.

### Prep plan
- **Day 1 (learn):** MDN WebSocket and SSE pages; Alex Xu chapter 12. Draw the request flow for each method.
- **Day 2 (solve):** Design the message flow: send → gateway → store → route → deliver → ack → ticks.
- **Day 3 (numbers and twist):** Solve cross-server routing and ordering. Estimate connections per server and servers needed.
- **Day 4 (depth):** Discord's scaling post; the optional WebSocket demo.
- **Day 5 (script):** Hook: the frozen rider dot and the single tick. Two dry runs.

**Title:** How Apps Show Your Rider Moving Live (and What Those Message Ticks Really Mean)

**Description:**
The dot moving on the map and the ticks next to your message hide a serious distributed system. In this episode of the Haat series we design live tracking and chat on the whiteboard, comparing polling, long polling, Server-Sent Events and WebSockets, then solve the hard parts: routing messages between servers, keeping them in order, and making the ticks tell the truth.

---

## Episode 18: 100,000 likes per second on one live stream

**Status Board:** Month 37 · 22M users · launch of Haat Live · 150 engineers
**Callback:** "Last time we built a real-time pipe to every phone. Our sellers came from Facebook, where they already sold through Facebook Live, so they turned it into live shopping. Tonight a celebrity seller goes live to half a million people."
**Prep time:** 4-5 days.

**Scenario:** A famous actor's clothing brand goes live on Haat. 500,000 viewers. Hearts fly at 100,000 per second, and every one runs `UPDATE streams SET likes = likes + 1`. That single row becomes the hottest spot in the system and slows down checkout for everyone, right as the actor says "first 1,000 orders get 50% off".

**Numbers:** 100K likes/sec on one row; a single row lock manages maybe a few thousand updates/sec. 500K viewers × comments fan-out.

**Approaches to compare**
- **Single row update:** correct but serialized; it can't go faster than one row's lock.
- **Sharded counters:** split into 100 sub-counters, increment a random one, sum on read.
- **Write buffering / batching:** aggregate likes in memory per gateway server and flush totals every second.
- **Approximate counts:** show "1.2M", because nobody needs the exact number live.

**The twist:** Two parts. First, a viewer taps the heart and sees the count unchanged (batch not flushed), so they tap again. Make users see their own like instantly (read-your-own-writes from Episode 13, client side). Second, the "first 1,000 orders" announcement is Episode 7's flash sale *inside* a live stream, so callback to waiting rooms and atomic stock.

**Verdict:** Client-side optimistic count + per-server batching into sharded counters + approximate display; reuse the flash-sale stock system for live deals.

**Diagram adds:** Live Service · Counter Aggregators.

**Cliffhanger:** "Millions of people now *follow* shops so they don't miss the next live, just like they followed Facebook pages. The actor's shop has 5 million followers. They just posted a new product."

### Learning guide

**Learn in this order**
1. **Write contention and hot rows:** why one row serializes writes regardless of hardware.
2. **Hot keys** in caches and databases (the same problem in Redis, Episode 11's hot shards).
3. **Sharded counters:** N sub-counters, random increment, sum on read; choosing N.
4. **Write batching / buffering:** aggregating in memory and flushing; what's lost if a server crashes.
5. **Approximate counting:** rounding for display; HyperLogLog for *unique* counts (unique viewers).
6. **Optimistic UI:** showing the user's own action immediately.
7. **Fan-out of comments to viewers:** batching broadcasts instead of one message per comment per viewer.
8. **Reusing the flash-sale design** (Episode 7) for live deals.

**Resources**
- *Read first:*
  - Google Cloud Firestore documentation, "Distributed counters": a short, clear explanation of sharded counters.
  - DDIA Chapter 6, section "Skewed Workloads and Relieving Hot Spots".
- *Go deeper:*
  - Amazon DynamoDB documentation on write sharding to distribute workloads evenly.
  - Redis documentation, "HyperLogLog" (counting unique viewers in a few KB).
  - Discord engineering, "How Discord Stores Trillions of Messages" (2023): hot partitions and request coalescing in practice.
  - Facebook engineering posts on live video at scale (search "Facebook Live engineering broadcasting millions").
- *Watch:* any good explanation of "hot key problem" or "sharded counters" (ByteByteGo, Arpit Bhayani), after Day 2.

**Try it yourself (optional, 1 hour)**
- On paper, simulate 20 likes across 4 sub-counters with dice or random numbers, then sum. Then simulate 3 gateway servers buffering likes for one second each and flushing totals. Count database writes in each approach: 20 vs 20 (spread out) vs 3.

**You're ready to record when you can**
- [ ] Explain why 100K updates/sec on one row can't work, even on a huge database.
- [ ] Design all four approaches, including how like/unlike works in each.
- [ ] Calculate database writes/sec for each approach.
- [ ] Explain what's lost if a gateway server crashes before flushing, and why that's acceptable for likes (but not for orders).
- [ ] Solve the "my like disappeared" twist with optimistic UI.
- [ ] Connect the live deal to Episode 7's flash-sale design.

**Questions viewers will ask (prepare answers)**
- *"Why not just use Redis INCR?"* Much better than a database row, but a single Redis key is still a hot key at this rate; you'd shard or batch there too.
- *"Do apps fake the like counts?"* Don't speculate. Explain that displayed counts are often approximate and eventually consistent, by design.
- *"What about unlikes?"* Track per-user like state separately (a set of users who liked), and treat the counter as a derived, approximate number.

**Don't get this wrong on camera**
- Don't say batching "loses nothing"; a crash loses the unflushed buffer. Be clear that it's an accepted trade-off for likes.
- Don't confuse HyperLogLog (unique counts) with total counts.

### Prep plan
- **Day 1 (learn):** Firestore distributed counters; DDIA hot spots section.
- **Day 2 (solve):** Design all four approaches, including how like/unlike works in each.
- **Day 3 (numbers and twist):** Calculate throughput for each approach. Solve the "my like disappeared" twist. Connect to Episode 7.
- **Day 4 (script):** Hook: the live-stream hearts flying, then checkout freezing. Two dry runs.

**Title:** 100,000 Likes Per Second: How Live Streams Count Without Melting the Database

**Description:**
Counting likes sounds like adding 1. At 100,000 likes per second on one live stream, that "+1" can bring down your database. In this episode of the Haat series we compare single-row updates, sharded counters, batching and approximate counts, then fix the subtle bug where viewers can't see their own like.

---

## Episode 19: The shop with 5 million followers broke the feed

**Status Board:** Month 40 · 25M users · 600K orders/day · 170 engineers
**Callback:** "Live shopping made following shops the most popular feature on Haat. Every user now has a home feed of new products from shops they follow. It worked fine, until the biggest shop posted."
**Prep time:** 5-6 days.

**Scenario:** The actor's shop (5M followers) posts a new collection. Haat tries to insert that post into 5 million feeds immediately. The feed system falls an hour behind, and smaller sellers' posts don't appear at all. Small sellers, the people Haat was built for since Episode 1, are furious.

**Numbers:** Average user follows 80 shops; average shop has 300 followers, but the top 100 shops have 1M+. Posts/day, feed opens/day, writes per celebrity post in push mode.

**Approaches to compare**
- **Fan-out on read (pull):** on open, fetch recent posts from every followed shop and merge. Cheap writes, slow reads.
- **Fan-out on write (push):** on post, write into every follower's precomputed feed (through Episode 9's queue). Instant reads, terrible for big shops.
- **Hybrid:** push for normal shops, pull for big shops, merge at read time. What large social networks describe doing.

**The twist:** Ranking. Users don't want newest-first; they want "most interesting first", and small sellers need a fair chance against big ones. Where does ranking run in each approach, and how do you avoid recomputing it on every refresh? End on the open question: *"How do we decide what's interesting to each user? That's not a rule you can write by hand. That's machine learning, and that's Season 2."*

**Verdict:** Hybrid fan-out, cached feeds, a simple scoring function for ranking now, and a note that it should be learned, not hand-written.

**Diagram adds:** Feed Service · Feed Cache · Ranking (placeholder box labelled "Season 2").

**Cliffhanger:** "Act 4 is done. Haat runs 40 services and deploys dozens of times a day. Last Thursday, a one-line database change broke checkout for everyone."

### Learning guide

**Learn in this order**
1. **The follow graph:** users, shops, follower counts and the long-tail distribution (a few huge shops, many small ones).
2. **Fan-out on read (pull):** the merge at read time and its latency cost.
3. **Fan-out on write (push):** precomputed feeds per user; write amplification.
4. **The celebrity problem** and the **hybrid** model with a follower-count threshold.
5. **Feed storage:** feed caches holding only IDs, then "hydrating" post details.
6. **Pagination for feeds:** cursor-based, not page numbers.
7. **Ranking basics:** signals (recency, engagement, relationship, seller fairness), a hand-written score, and where ranking runs (at write time vs read time).
8. **Fairness for small sellers** as a product requirement.

**Resources**
- *Read first:*
  - Alex Xu Vol. 1, Chapter 11 "Design a news feed system".
  - DDIA Chapter 1: the Twitter home timeline example (fan-out and the celebrity hybrid), only a few pages.
- *Go deeper:*
  - Raffi Krikorian, "Timelines at Scale" (QCon 2013 talk, on InfoQ): how Twitter's timeline fan-out worked.
  - Silberstein et al., "Feeding Frenzy: Selectively Materializing Users' Event Feeds" (SIGMOD 2010): the academic basis for choosing push vs pull per user/producer. Read the introduction.
  - Meta engineering, posts explaining how News Feed ranking works at a high level (search "how machine learning powers Facebook's News Feed ranking").
  - Slack or Pinterest engineering posts on home feed architecture (search "<company> home feed architecture").

**Try it yourself (optional, 1 hour)**
- Simulate on paper: 5 users, 3 normal shops, 1 celebrity shop, a follow graph. Trace one post and one feed open in pull mode, push mode and hybrid mode. Count reads and writes for each. This becomes your main whiteboard panel.

**You're ready to record when you can**
- [ ] Explain pull and push with read/write costs for each.
- [ ] Calculate writes for a 5M-follower post in push mode, and how long it takes at a given write rate.
- [ ] Explain the hybrid model and choose a reasonable follower threshold.
- [ ] Draw the merge at read time in the hybrid model.
- [ ] List 4 ranking signals and explain where ranking runs.
- [ ] Explain why small-seller fairness is a business requirement, not just a technical one.

**Questions viewers will ask (prepare answers)**
- *"What threshold makes a shop 'big'?"* It's tuned from data (cost of push vs cost of pull); a number like 100K followers is a reasonable starting point to discuss.
- *"Do inactive users get pushed feeds?"* A common optimisation: skip push for users who haven't opened the app in a while, and build their feed on demand when they return.
- *"Isn't this exactly how Instagram/Facebook works?"* Similar ideas are publicly described; avoid claiming exact internals.

**Don't get this wrong on camera**
- Don't present push as "wrong"; it's right for most shops. The hybrid is about the long tail.
- Don't over-promise on ranking; keep it a simple, explainable score and hand the hard part to Season 2.

### Prep plan
- **Day 1 (learn):** Alex Xu chapter 11; DDIA chapter 1 Twitter example.
- **Day 2 (solve):** Simulate on paper: 5 users, 1 celebrity shop, a follow graph; trace post and open in each model.
- **Day 3 (numbers and twist):** Calculate write volume for a 5M-follower post in push mode. Design the hybrid merge and where ranking runs.
- **Day 4 (depth):** Watch "Timelines at Scale" and collect one quote.
- **Day 5 (script):** Hook: the celebrity post and the furious small sellers. Prepare the Season 2 teaser. Two dry runs.

**Title:** Why One Celebrity Post Can Break a Feed: Designing a News Feed

**Description:**
When a shop with 5 million followers posts, where does that post go? In this episode of the Haat series we design a home feed on the whiteboard, comparing fan-out on read, fan-out on write and the hybrid approach large social networks describe using. Then we add ranking, and hit the question that leads into machine learning.

---

# Act 5: Never Go Down (Episodes 20-21)

*Haat is too big to fail. These episodes are about changing it safely and surviving disasters.*

## Episode 20: The one-line change that broke checkout

**Status Board:** Month 44 · 28M users · 40 services · 60 deploys/day · 200 engineers
**Callback:** "In Episode 1, deploying meant one junior copying files to one server. Since Episode 14 we've been splitting Haat into services. There are 40 now, and they deploy 60 times a day. Every one of those deploys is a chance to break something."
**Prep time:** 4-5 days.

**Scenario:** A developer renames the column `delivery_addr` to `delivery_address`. During the rolling deploy, half the Order service instances run the old code, which looks for a column that no longer exists. Checkout fails for 40% of customers for 18 minutes. The team's response is to ban deploys after 4 PM and on Fridays. Velocity collapses.

**Numbers:** 60 deploys/day × 40 services. Error budget: what does 18 minutes of failed checkout cost at 600K+ orders/day? (600K ÷ 1,440 minutes ≈ 420 orders/minute; 40% failing for 18 minutes ≈ 3,000 lost orders, more at peak hours.)

**Approaches to compare**
- **Rolling deployment:** replace instances one at a time. No downtime, but old and new versions run together.
- **Blue-green deployment:** run a full copy with the new version and switch all traffic at once. Instant rollback, but double the infrastructure.
- **Canary release:** send 1% of traffic to the new version, watch error rates, then expand, with automatic rollback.
- **Feature flags:** ship code turned off, then enable it per user group. Separates deploying from releasing.

**The twist:** Back to the column rename. In any strategy, old and new code coexist, so the database must work for both. Introduce **expand and contract**: add the new column → write to both → backfill → switch reads → stop writing old → drop old. Six safe steps instead of one dangerous one.

**Verdict:** Canary + automatic rollback for services, feature flags for risky features, expand-and-contract for every schema change. Friday deploys are allowed again.

**Diagram adds:** Deployment pipeline · Feature Flag Service · Monitoring and auto-rollback.

**Cliffhanger:** "We can change Haat safely now. But every one of those 40 services, every database and every shard still runs in one building in Dhaka. On Eid night, that building loses power."

### Learning guide

**Learn in this order**
1. **Why deploys cause outages:** mixed versions, schema changes, config changes, untested paths.
2. **Rolling deployments:** how health checks gate each step (connects to Episode 2).
3. **Blue-green deployments:** switching traffic at the load balancer; rollback; the database is still shared.
4. **Canary releases:** small traffic percentage, comparing metrics (error rate, latency) against the baseline, automatic rollback.
5. **Feature flags:** deploy ≠ release; flag types (release, experiment, ops kill switch); flag debt.
6. **Backward-compatible changes:** API and schema compatibility across versions.
7. **Expand and contract (parallel change)** for schema migrations, step by step.
8. **Online schema migration tools** at a high level (avoiding long table locks).
9. **Error budgets and DORA metrics** at a high level (deployment frequency, change failure rate, time to restore).

**Resources**
- *Read first:*
  - Martin Fowler's site: "BlueGreenDeployment", "CanaryRelease" (Danilo Sato), "Feature Toggles (aka Feature Flags)" (Pete Hodgson) and "ParallelChange" (expand and contract).
  - Stripe engineering, "Online migrations at scale" (2017): a real four-step migration story; excellent for the twist.
- *Go deeper:*
  - Google SRE Workbook, chapter "Canarying Releases".
  - Google SRE book, chapter "Release Engineering".
  - *Accelerate* by Forsgren, Humble & Kim: why frequent, small deploys are safer (DORA research).
  - *Continuous Delivery* by Humble & Farley: the classic on deployment pipelines.
  - GitHub's `gh-ost` documentation (online schema migrations for MySQL), as an example of real tooling.
  - Kubernetes documentation on rolling updates and readiness probes.

**Try it yourself (optional, 1 hour)**
- Write the six expand-and-contract steps for the `delivery_addr` rename on paper. For each step, write which code version (old/new) is running and confirm both work. If any step breaks one version, you've found a bug in your plan.

**You're ready to record when you can**
- [ ] Draw traffic flow for rolling, blue-green and canary.
- [ ] Compare them on downtime, rollback speed, cost and risk in a table.
- [ ] Explain why blue-green doesn't solve database changes.
- [ ] Walk through the column rename with expand-and-contract, step by step, from memory.
- [ ] Explain feature flags and one risk of having too many.
- [ ] Calculate the cost of 18 minutes of broken checkout.

**Questions viewers will ask (prepare answers)**
- *"Isn't banning Friday deploys sensible?"* It's understandable, but it hides the real problem. Safe, small, reversible deploys make any day safe; that's the goal.
- *"Which feature flag tool?"* Many exist (open source and paid); a simple config table works at first. Not the point.
- *"How do you test canaries if 1% traffic is tiny?"* At Haat's scale 1% is thousands of requests a minute, enough for error-rate comparison. Small services may need longer canary windows.

**Don't get this wrong on camera**
- Don't present blue-green as "zero risk"; shared databases and long-running connections still need care.
- Don't skip the backfill step in expand-and-contract; it's where large tables need throttling.

### Prep plan
- **Day 1 (learn):** The four Fowler articles; draw traffic flow for each strategy.
- **Day 2 (solve):** Compare them on paper: downtime, rollback speed, cost, risk.
- **Day 3 (twist):** Work through the column rename with expand-and-contract, step by step. Read Stripe's online migrations post.
- **Day 4 (script):** Script with the "18 minutes of broken checkout" incident as the hook, told as a short postmortem. Two dry runs.

**Title:** How Big Companies Deploy 100 Times a Day With Zero Downtime

**Description:**
Deploying shouldn't mean downtime, 3 AM releases or crossed fingers. In this episode of the Haat series we compare rolling deployments, blue-green, canary releases and feature flags on the whiteboard, then solve the problem that breaks many zero-downtime deploys: database schema changes, with the expand-and-contract pattern.

---

## Episode 21 (Season Finale): The night the data center went dark

**Status Board:** Year 4 · 30M users · Eid sale night · 1.2M orders/day · 220 engineers
**Callback:** A 60-second montage of the Architecture Diagram growing from Episode 1's single box, the one a fresh graduate sketched for a Facebook seller, to today. Then: "Everything we built this season lives in one data center. Tonight, at 9:12 PM, during the biggest Eid sale in Haat's history, that data center loses power."
**Prep time:** 6-7 days (includes the season recap).

**Scenario:** The generator fails to start. Haat is completely offline at the peak of the Eid sale. Backups exist, the same kind of daily backups we set up in Episode 1, now much bigger, but restoring them would take 9 hours and lose the last 6 hours of orders. Every minute down costs lakhs and trust.

**Numbers:** Define **RPO** (how much data you can afford to lose) and **RTO** (how long you can afford to be down). Revenue per minute on sale night. Cost of each approach as a multiple of the current infrastructure bill.

**Approaches to compare**
- **Backup and restore:** cheapest; RPO of hours, RTO of hours.
- **Pilot light / warm standby:** a second region with replicated data and minimal servers, scaled up on failover. RPO of seconds, RTO of minutes.
- **Active-passive:** a full second region ready to take traffic, with DNS/global load balancer failover.
- **Active-active:** both regions serve users all the time. Best availability, hardest consistency.

**The twist:** The network link between the regions fails, but both regions are still up. Each thinks the other is dead and starts accepting orders (split brain). The last 50 sale phones get sold in *both* regions, which is Episode 7's oversell at national scale. Now explain the **CAP theorem** in plain words: during a network partition, you must choose between staying correct (refuse some requests) and staying available (accept them and reconcile later). Then show Haat's answer per data type, straight from Episode 13's tiers: stock and payments choose correctness; carts, feeds and likes choose availability.

**Verdict:** Active-passive with fast failover for the core, active-active only for data that tolerates conflict, and regular failure drills ("game days").

**Diagram adds:** The whole diagram is duplicated into a second region, with replication arrows and a global load balancer. This is the final season diagram.

**Season recap (2-3 min):** Walk the final diagram box by box: "This box is Episode 2's launch night. This one stopped the bots in Episode 6. This one is why nobody gets charged twice..." End on the single box from Episode 1: "And it all started here, with one server, a Facebook page and fourteen days."

**Season 2 teaser:** "Haat has four years of data: what people searched, clicked, bought and watched. In Season 2 we'll use it. How does Haat know what to show you? How can search understand 'something to wear to a wedding'? How does ChatGPT work? See you there."

### Learning guide

**Learn in this order**
1. **RPO and RTO:** definitions, and how the business (not engineers) sets them.
2. **Backup types:** full, incremental, point-in-time recovery; restore time as the hidden cost.
3. **The four DR strategies:** backup and restore, pilot light, warm standby, multi-site active-active; cost vs RPO/RTO for each.
4. **Cross-region replication:** synchronous vs asynchronous across long distances, and the latency cost.
5. **Failover mechanics:** health checks, DNS TTLs, global load balancers, who decides to fail over (human vs automatic).
6. **Split brain and quorum:** why a third location or a majority vote prevents both sides acting as leader.
7. **The CAP theorem, properly:** it applies *during a partition*; the choice is consistency or availability for each operation. Plus PACELC for normal operation.
8. **Conflict resolution for active-active data:** last-write-wins and its dangers, merging carts, CRDTs (name only).
9. **Chaos engineering and game days:** testing failover before you need it.

**Resources**
- *Read first:*
  - AWS whitepaper, "Disaster Recovery of Workloads on AWS: Recovery in the Cloud": the four strategies with clear RPO/RTO diagrams.
  - Eric Brewer, "CAP Twelve Years Later: How the 'Rules' Have Changed" (IEEE Computer, 2012): CAP explained by its author.
  - Martin Kleppmann, "Please stop calling databases CP or AP" (2015): why "pick 2 of 3" is misleading.
- *Go deeper:*
  - DDIA Chapter 5 (multi-leader replication and conflict handling) and Chapter 9 (linearizability and the cost of it; the CAP section).
  - Google SRE book, chapter "Data Integrity: What You Read Is What You Wrote".
  - AWS Builders' Library, "Static stability using Availability Zones".
  - Principles of Chaos Engineering (principlesofchaos.org) and Netflix's posts on Chaos Monkey.
  - Martin Kleppmann's Cambridge lectures on quorums and consensus (for split brain).
- *Real-world stories:*
  - The OVHcloud Strasbourg data center fire (March 2021): customers without off-site backups lost data.
  - GitLab's 2017 outage again: a perfect callback to Episode 1's twist.
  - Public postmortems of major cloud region outages (AWS, Google Cloud and Azure publish summaries).

**Try it yourself (optional, 1-2 hours)**
- Build a table: rows = Haat's data types (stock, payments, orders, carts, feeds, likes, sessions); columns = RPO, RTO, CAP choice during a partition, and why. This table is your verdict panel and your strongest season-recap moment.

**You're ready to record when you can**
- [ ] Define RPO and RTO with a Haat example, and calculate the cost of a 9-hour RTO on sale night.
- [ ] Draw each DR strategy's failover sequence on Haat's diagram.
- [ ] Explain why cross-region synchronous replication is slow (distance and round trips).
- [ ] Explain split brain with the 50 phones example, and how quorum prevents it.
- [ ] Explain CAP correctly (during a partition, per operation) without "pick 2 of 3".
- [ ] Present Haat's per-data CAP choices and connect each to an earlier episode.
- [ ] Walk the whole final diagram box by box, naming the episode for each.

**Questions viewers will ask (prepare answers)**
- *"Why not active-active for everything?"* Conflicts. Stock and payments can't safely accept writes in two places during a partition without risking overselling or double charges.
- *"Is CAP still relevant?"* Yes, as a way to reason about partitions; but real systems make the choice per operation and also face latency trade-offs (PACELC) when there's no partition.
- *"How often should you test failover?"* Regularly (quarterly game days are a common starting point), and always after major architecture changes.

**Don't get this wrong on camera**
- Don't say "pick any two of consistency, availability and partition tolerance". Partitions aren't optional; the choice is what to do when one happens.
- Don't claim a strategy gives "zero downtime" or "zero data loss" without stating the cost.

### Prep plan
- **Day 1 (learn):** The AWS DR whitepaper; RPO/RTO and the four strategies.
- **Day 2 (solve):** On paper, draw each strategy's failover sequence for Haat's diagram. Estimate cost multiples.
- **Day 3 (twist):** Read Brewer and Kleppmann on CAP. Design the split-brain scenario and Haat's per-data CAP table.
- **Day 4 (recap):** Prepare the season recap: list every box in the final diagram with its episode number and a one-line memory.
- **Day 5 (script):** Script the finale; build the diagram montage from your saved per-episode diagrams.
- **Day 6-7 (record):** Record the finale, the recap and the Season 2 teaser separately. Rehearse the recap until it flows; it's the most rewatched part of the season.

**Title:** The Night Our Data Center Went Dark: Disaster Recovery and the CAP Theorem

**Description:**
At the peak of our biggest sale, the entire data center loses power. In the Season 1 finale of the Haat series we design for disaster: backups, warm standby, active-passive and active-active regions, and the split-brain problem that forces every system to choose between being correct and being available. Plus a full recap of everything we built this season, from one server for a Facebook seller to a nationwide platform, and a look at Season 2.

---

# Season 2 preview: Haat Gets Smart (AI)

Season 2 continues the same story; it doesn't start a new one. Natural first episodes, each growing out of Season 1:

- **The feed ranking box from Episode 19:** how recommendation systems work ("customers also bought").
- **Search from Episode 5:** semantic search, embeddings and vector databases ("something to wear to a wedding").
- **Payments from Episode 8:** fraud detection (including fake COD orders), plus how a neural network learns.
- **ETA from Episode 16:** predicting delivery time with machine learning.
- **Chat from Episode 17:** a customer support chatbot, and finally how ChatGPT works.
- **Episode 1's "vibe coding" thread:** what AI coding tools are good and bad at, now that the viewer understands the systems underneath.

We'll plan this phase in detail when Season 1 is recorded.

---

# Continuity Bible

Check this before scripting any episode. The numbers must never go backwards.

### Growth timeline

| Ep | Time since launch | Users | Orders/day | Engineers | Key infrastructure added |
|---|---|---|---|---|---|
| 1 | Day 0 (14-day build) | 80K Facebook followers | ~50 (via Messenger) | 1 | One VPS: app + DB + uploads folder + cron; daily off-site backups |
| 2 | Launch night / week 1 | 60K visitors; ~20K accounts | ~800 on launch day | 1 | Load balancer, 3 app servers, Redis sessions, DB on its own server, shared image folder (hack) |
| 3 | Month 4 | 150K | 2K (Eid: 20× traffic) | 3 | CDN, Redis cache |
| 4 | Month 6 | 400K · 5K sellers | 10K | 4 | Object storage, first queue, media workers (image hack retired) |
| 5 | Month 8 | 1M · 2M products | 25K | 6 | Search cluster, CDC |
| 6 | Month 10 | 1.5M | 40K | 7 | Rate limiter |
| 7 | Month 12 | 3M | 11.11 sale night | 9 | Waiting room, Redis stock counter |
| 8 | Month 12 | 3.5M | 11.11 sale night | 9 | Prepaid payments, idempotency store, reconciliation |
| 9 | Month 15 | 5M | — (planning Eid) | 12 | Notification service, priority queues (mobile app now exists) |
| 10 | Month 15 | 5M | Eid sale night | 12 | Scheduler with leases (Episode 1 cron removed) |
| 11 | Month 20 | 8M | 150K | 25 | Read replicas, 4 order shards |
| 12 | Month 21 | 8M | 160K | 25 | Snowflake IDs, short links |
| 13 | Month 22 | 10M | 200K | 30 | Consistency-aware reads |
| 14 | Month 26 | 12M | 300K | 80 | Microservices, saga orchestrator, outbox |
| 15 | Month 28 | 14M | Pohela Boishakh sale | 90 | Circuit breakers, bulkheads |
| 16 | Month 31 | 16M · 50K riders | 400K | 110 | Rider location + assignment services |
| 17 | Month 34 | 18M | 500K | 130 | Real-time gateway, pub/sub |
| 18 | Month 37 | 22M | — (Haat Live launch) | 150 | Live service, counter aggregators |
| 19 | Month 40 | 25M | 600K | 170 | Feed service, feed cache |
| 20 | Month 44 | 28M | 650K+ | 200 | Deploy pipeline, feature flags |
| 21 | Year 4 | 30M | 1.2M (Eid sale) | 220 | Second region |

### Recurring characters

- **You (the junior):** fresh graduate in Episode 1; Haat's first engineer; by the finale, the person who knows why every box exists.
- **The owner:** non-technical Facebook seller; later the CEO. Brings business demands that trigger episodes (launch Live in Ep 2, Eid campaign in Ep 3, opening to other sellers in Ep 4, prepaid-only flash sales in Ep 8, own riders in Ep 16).
- **The CTO:** appears once the team grows (from Episode 11); asks the "what do we do, in what order?" questions.

### Recurring building blocks (use these for callbacks)

- **Episode 1's planted shortcuts:** memory sessions (2) → local image folder (2, 4) → per-section queries (3) → `LIKE` search (5) → unlimited OTP (6) → check-then-save stock (7) → COD only (8) → cron (2, 10) → single database (11) → auto-increment IDs (12) → single deploy (20) → daily backups and one data center (21).
- **Redis:** sessions (2) → cache (3) → rate limit counters (6) → stock counter (7) → idempotency store (8) → scheduler sorted sets (10) → rider geo index (16) → like counters (18).
- **The queue:** media processing (4) → order writes (7) → notifications (9) → scheduled jobs (10) → saga events (14) → feed fan-out (19).
- **Idempotency / "at least once + idempotency":** payments (8) → notifications (9) → scheduled jobs (10) → compensations (14).
- **Read-your-own-writes:** My Orders (13) → my own like (18).
- **Oversell:** race condition (7) → double-released stock (10) → live deals (18) → split brain across regions (21).
- **Consistency tiers:** defined in (13), reused for the CAP decision in (21).
- **The Facebook thread:** Facebook-only business (1) → Facebook Live launch (2) → boosted Eid campaign (3) → Facebook sellers join (4) → links shared on Facebook (12) → Facebook Live habit becomes Haat Live (18) → following shops like pages (19).

---

# Checklist for every whiteboard video

Copy this for each episode and tick it off before publishing.

**Before recording**
- [ ] Status Board numbers checked against the Continuity Bible
- [ ] Callback line written and connects to the previous episode's cliffhanger (or an Episode 1 shortcut)
- [ ] One-page explanation written in your own words
- [ ] Every box in the episode's *You're ready to record when you can* list ticked
- [ ] Tiny worked example solved on paper, arithmetic double-checked
- [ ] Every real-world number has a source noted (or is labelled "rough estimate")
- [ ] Hook ready: open with the disaster, not the definition
- [ ] Architecture Diagram updated; new boxes highlighted
- [ ] Whiteboard planned panel by panel (or slides in Excalidraw if you record digitally)
- [ ] Answers ready for the questions in *Questions viewers will ask*
- [ ] *Don't get this wrong on camera* list re-read on recording day
- [ ] Cliffhanger for the next episode written

**While recording**
- [ ] Show the naive solution failing before the clever one
- [ ] Pause and ask viewers "how would you solve this?" before revealing the answer
- [ ] Use real-world numbers to show why scale matters
- [ ] Use local examples (Facebook sellers, COD, Dhaka traffic, bKash, Banglish search, Eid, Pohela Boishakh and 11.11 sales) where they fit
- [ ] Say "it depends" honestly, and say what it depends on

**After recording**
- [ ] Aim for 12-20 minutes; cut anything that doesn't move the story
- [ ] Add chapters (timestamps) to the description
- [ ] Add "Previous episode" and "Next episode" links, plus the season playlist
- [ ] Thumbnail: one striking drawing plus 3-5 words, readable on a phone
- [ ] Link the sources you studied in the description
- [ ] Cut a 30-60 second Short from the best "aha" moment
- [ ] Save this episode's final diagram as a separate image (you'll need all of them for the finale montage)
