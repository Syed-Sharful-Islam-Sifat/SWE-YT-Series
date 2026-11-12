# "How Would You Do It?": Engineering Scenarios from Real Case Studies

*Real engineering changes, how real teams did them, and how to do them without breaking anything · Nov 12, 2026 · @Lion*

Each topic is a situation real teams face: **"We need to change X. What can go wrong, and how do we do it safely?"** Most of the video is the solution.

**Part A** is built on **published case studies** from company engineering blogs and conference talks. Each one gives you a real story, real numbers and a real solution to explain, which is what makes 1-day prep possible. **Part B** has everyday scenarios that don't have one famous case study but are just as useful.

> **How this list was researched:** the Part A topics and facts come from web searches of engineering blogs, conference talks and newsletters (Nov 2026). I worked from search excerpts and summaries, **not the full articles**. Before recording, read the original post (the first link in each topic) and take every number from it, not from this doc.

---

## The 1-day prep routine

| Time | Task |
|---|---|
| Morning (2-3 h) | Read the original case study fully. Note: the problem, the constraints, the steps they took, what went wrong, what they'd do differently. |
| Midday (1 h) | Generalize it: draw the "before" and "after", and write the safe plan as numbered steps your viewer could apply to their own system. |
| Afternoon (2 h) | Whiteboard panels, hook, bullet-point script. |
| Evening (1-2 h) | Record and do a light edit. |

## The video shape (8-12 minutes)

1. **The situation (1 min):** "Imagine your team has to ..." Pose it as *your* problem first.
2. **The naive way (1-2 min):** what most people would do, and how it breaks.
3. **The challenges (2 min):** list them. Ask viewers "what else could go wrong?"
4. **The safe plan (4-5 min):** step by step, with rollback at every step. The main part.
5. **How [Company] actually did it (1-2 min):** the real story and numbers. This builds credibility.
6. **Checklist (30 s):** 3-4 rules to remember. This also makes a good LinkedIn post.

---

## The 6 patterns behind almost every safe change

The same few ideas solve very different problems, and every case study below uses them. Make **one early video about these patterns**, then refer back to them ("this is pattern #2 again"). That links your videos together.

1. **Expand and contract (parallel change):** add the new thing next to the old, move everyone over, then remove the old. ([Martin Fowler: Parallel Change](https://www.martinfowler.com/bliki/ParallelChange.html))
2. **Dual write → backfill → compare → switch reads → stop old writes:** the standard way to move data. (Stripe's four-step migration, topic A1)
3. **Strangler fig:** replace an old system piece by piece, routing one feature at a time to the new one. ([Wikipedia: Strangler fig pattern](https://en.wikipedia.org/wiki/Strangler_fig_pattern), [Microsoft Azure: Strangler Fig](https://learn.microsoft.com/azure/architecture/patterns/strangler-fig))
4. **Branch by abstraction:** put an interface in front of the thing you're replacing, so the swap happens behind it.
5. **Shadow mode / side-by-side comparison:** run old and new together, compare results, keep serving the old one until they match. (Stripe, Khan Academy, Netflix all did this.)
6. **Gradual rollout with an escape hatch:** feature flags or traffic percentages, and a way to go back instantly. (Dropbox's Hydra, GitHub's rollback planning.)

---

# Part A: Topics based on real case studies

## Moving and reshaping data

**A1. "Split a huge table into a new one while the app is live" (Stripe)**
- **Situation:** a big object (customers) has a nested piece (subscriptions) that now needs its own table. Hundreds of millions of records, payments running 24/7.
- **Real case:** Stripe moved hundreds of millions of Subscriptions out of Customer records using a four-step dual-writing pattern, and used the Scientist library to read from both and flag mismatches.
- **What you teach:** pattern 2 in full: dual write → change reads → change writes → remove old data, with a comparison step before trusting the new table.
- **Twist:** a write succeeds in the old table but fails in the new one. How you detect and repair the difference.
- **Read:** [Stripe: Online migrations at scale](https://stripe.com/se/blog/online-migrations) · [Simon Willison's note on it](https://simonwillison.net/2023/Nov/5/online-migrations-at-scale/) · smaller real example: [Gusto: splitting a large table](https://engineering.gusto.com/old-write)

**A2. "Our database is too big for one machine. Shard it, without a rewrite" (Figma, Notion)**
- **Situation:** one Postgres server, already on the largest instance; the product grows ~100× in a few years.
- **Real cases:**
  - **Figma** went from a single Postgres on AWS's largest instance (2020) to vertical partitioning, then horizontal sharding. Their key idea: **logical sharding first** (queries routed as if sharded while data still sits in one database), **then physical sharding**, with an easy rollback, through a custom Go proxy (DBProxy). They evaluated CockroachDB, TiDB, Spanner and Vitess, then stayed on Postgres for team expertise and an incremental path.
  - **Notion** sharded application-side by workspace ID, the block table first, using Postgres logical replication to copy history and stream new changes.
- **What you teach:** choosing a shard key, why "change the router before you move the data" reduces risk, why not just switch databases.
- **Twist:** queries that join across shards; re-sharding later (Notion wrote a follow-up, "The Great Re-shard").
- **Read:** [Figma: How Figma's databases team lived to tell the scale](https://www.figma.com/en-jp/blog/how-figmas-databases-team-lived-to-tell-the-scale/) · [Notion: Herding elephants: sharding Postgres](https://notion.com/blog/sharding-postgres-at-notion) · [Notion: The Great Re-shard](https://notion.com/blog/the-great-re-shard) · [pganalyze: 5 mins of Postgres on Figma's DBProxy](https://pganalyze.com/blog/5mins-postgres-figma-dbproxy-sharding-postgres)
- **(2 days)**: there's a lot of material. Cover one company in depth and mention the other.

**A3. "Move trillions of messages to a new database with zero downtime" (Discord)**
- **Situation:** your NoSQL cluster has hot partitions and latency spikes; you want a different database without stopping the app.
- **Real case:** Discord went MongoDB → Cassandra → ScyllaDB. The Cassandra cluster grew from 12 nodes (2017) to 177; the ScyllaDB cluster runs on 72. Because ScyllaDB speaks Cassandra's protocol, app logic didn't need rewriting. Requirements were zero downtime and speed. They planned dual writes plus a Spark-based migrator, and also put intermediary data services in front of the database to absorb hot partitions.
- **What you teach:** why the database was the wrong fix alone (hot partitions are a data-modeling problem), dual writes + bulk copy, why a compatible protocol is a huge advantage.
- **Twist:** the bulk copy would take months, so how do you make it faster? (The original post tells this story.)
- **Read:** [Discord: How Discord stores trillions of messages](https://discord.com/blog/how-discord-stores-trillions-of-messages) · [InfoQ summary](https://www.infoq.com/news/2023/06/discord-cassandra-scylladb/)

**A4. "We're switching from Postgres to MySQL." Wait, why? (Uber)**
- **Situation:** a debate video. When is changing your relational database actually worth it?
- **Real case:** in 2016 Uber explained its move from Postgres to MySQL: Postgres updated every index when a row moved, its low-level WAL made cross-data-center replication expensive, and it had issues with replica MVCC and upgrades. Uber then built Schemaless, a sharding layer on MySQL. Uber's own post notes the analysis was based on the old Postgres 9.2.
- **What you teach:** how MVCC and replication work at a simple level; how to evaluate "the database is the problem" claims; why context and version matter.
- **Twist:** the post caused a big debate, and Postgres has changed a lot since. Show both sides.
- **Read:** [Uber: Why Uber Engineering switched from Postgres to MySQL](https://www.uber.com/blog/postgres-to-mysql-migration/) · [LWN discussion](https://lwn.net/Articles/695999/)

**A5. "Change the schema of a huge MySQL table without locking it" (GitHub, Shopify)**
- **Situation:** you need to alter a table with billions of rows that is written to constantly.
- **Real cases:**
  - **GitHub** built gh-ost, a triggerless online schema migration tool. Instead of database triggers, it reads the binary log to copy changes to a "ghost" table, and it can be paused and throttled based on load and replication lag.
  - **Shopify** applies a developer's migration asynchronously across 100+ shards without downtime.
- **What you teach:** shadow table + copy + catch up + atomic swap; why triggers add load; throttling on replication lag.
- **Twist:** the final cut-over moment: how to swap tables with near-zero lock time.
- **Read:** [GitHub: gh-ost announcement](https://github.blog/2016-08-01-gh-ost-github-s-online-migration-tool-for-mysql) · [gh-ost README](https://github.com/github/gh-ost) · [Rails at Scale: swappable migration backends](https://railsatscale.com/2025-12-08-swappable-migration-backends-in-rails/)

**A6. "One shard is overloaded. Move a customer to another database while they're using it" (Shopify)**
- **Situation:** with sharding, a few big customers make one shard hot.
- **Real case:** Shopify moves a shop's data between pods with its tool *ghostferry* to rebalance load, aiming for no downtime visible to merchants.
- **What you teach:** copy + tail changes + short cut-over for a single tenant; the routing table that says which shard a shop lives on.
- **Twist:** what about writes that arrive during the few seconds of cut-over?
- **Read:** [Shopify: MySQL database shard balancing at terabyte scale](https://shopify.engineering/blogs/engineering/mysql-database-shard-balancing-terabyte-scale) · [Arpit Bhayani's video summary](https://arpitbhayani.me/videos/shopify-zero-downtime-data-migration-hot-shards)

**A7. "Move to a new database cluster, and change nothing in the app" (Wix, Netflix)**
- **Situation:** many teams must move data between clusters; you don't want each team to write migration code.
- **Real cases:**
  - **Wix** built DB Mover, an internal service that moves data between clusters at scale without application code changes, with a "reroute" step as the key to zero downtime.
  - **Netflix** moved billing from Oracle to MySQL with "perceived zero downtime": table by table, writes shadowed between the stores, results compared and reconciled. It took over a year with no customer impact.
- **What you teach:** making migration a platform capability instead of a one-off project; shadowing and reconciliation.
- **Read:** [Wix: How we built a zero-downtime database migration service](https://www.wix.engineering/post/how-we-built-a-zero-downtime-database-migration-service-at-wix) · [InfoQ: Netflix data migrations](https://www.infoq.com/news/2018/11/netflix-data-migrations) · [Tines: Zero-downtime database migrations, lessons from moving a live production database](https://www.tines.com/blog/zero-downtime-database-migrations-lessons-from-moving-a-live-production/)

**A8. "Migrate databases with feature flags" (LaunchDarkly)**
- **Situation:** you want every migration step to be gradual and reversible.
- **Real case:** LaunchDarkly describes moving production databases (including its own move from Mongo and PostgreSQL to CockroachDB) with gradual, reversible, verifiable steps controlled by feature flags.
- **What you teach:** pattern 6: a flag per step (dual write on/off, read from new for X%), and what "reversible" means at each step.
- **Read:** [LaunchDarkly: database migration using LaunchDarkly](https://launchdarkly.com/blog/database-migration-using-launchdarkly/) · [3 best practices for zero-downtime database migrations](https://launchdarkly.com/blog/3-best-practices-for-zero-downtime-database-migrations/)

**A9. "Move to a sharding system and don't change three things at once" (Slack)**
- **Situation:** your MySQL setup can't grow any further; you adopt a sharding layer.
- **Real case:** Slack started moving to Vitess in 2017; by December 2020 Vitess served 99% of query load. They kept MySQL and the relational model on purpose. Their honest lesson: they changed the storage layout, the per-row data format and the migration path at the same time, "three goals with one migration... three times the complexity".
- **What you teach:** why one change at a time is safer, even though it looks slower.
- **Read:** [Slack: Scaling datastores at Slack with Vitess](https://slack.engineering/scaling-datastores-at-slack-with-vitess/) · [CNCF case study](https://www.cncf.io/case-studies/slack/)

## Upgrading what's already running

**A10. "Upgrade the database version on 1,200 servers" (GitHub)**
- **Situation:** your MySQL version is reaching end of life; hundreds of hosts, no downtime allowed.
- **Real case:** GitHub upgraded 1,200+ hosts from MySQL 5.7 to 8.0 over more than a year. It ran 5.7 and 8.0 side by side in CI for every application, gave developers 8.0 test containers, and planned rollback carefully: 5.7 hosts were blocked as failover candidates so the topology manager couldn't promote a host that would make going back impossible.
- **What you teach:** an upgrade is a migration; test both versions in CI; plan the rollback *before* you start.
- **Twist:** once a replica is on the new version, can you still fail back? (Replication usually only works old → new.)
- **Read:** [GitHub: Upgrading GitHub.com to MySQL 8.0](https://github.blog/engineering/infrastructure/upgrading-github-com-to-mysql-8-0/)

**A11. "Upgrade a framework three major versions behind" (GitHub's Rails upgrade)**
- **Situation:** you're stuck on an old framework version, and feature work can't stop for the upgrade.
- **Real case:** GitHub took a year and a half to go from Rails 3.2 to 5.2 without pausing features. The key technique was **dual booting**: the app could run on either version, and CI tested both. They shipped once they could run 30 minutes at peak traffic in production with no visible impact. Their own fork of Rails had made it harder, so a lesson was to stay close to upstream.
- **What you teach:** dual booting / compatibility with both versions; small steps on the main branch, not a giant branch.
- **Read:** [GitHub: Upgrading GitHub from Rails 3.2 to 5.2](https://github.blog/engineering/upgrading-github-from-rails-3-2-to-5-2/) · [Planet Argon: 6 Rails upgrade tales](https://blog.planetargon.com/blog/entries/6-rails-upgrade-tales-to-learn-from)

**A12. "Move a desktop app used by millions from Python 2 to 3" (Dropbox)**
- **Situation:** a language version upgrade, but the code runs on millions of users' computers you can't log into.
- **Real case:** Dropbox started during a 2015 hack week. It built a hybrid setup first, and a system called **Hydra** let the app choose the Python 2 or 3 interpreter at startup, so users could be switched back quickly if something broke. It rolled out to internal users, then Beta, then everyone.
- **What you teach:** the escape hatch (pattern 6) when you can't redeploy servers; staged rollout to user groups.
- **Read:** [Dropbox: How we rolled out one of the largest Python 3 migrations ever](https://dropbox.tech/application/how-we-rolled-out-one-of-the-largest-python-3-migrations-ever) · [Dropbox: Incrementally migrating over one million lines of code from Python 2 to Python 3](https://dropbox.tech/application/incrementally-migrating-over-one-million-lines-of-code-from-python-2-to-python-3)

**A13. "Big bang or step by step? Moving millions of lines of code to TypeScript" (Stripe vs Airbnb)**
- **Situation:** a debate video. Your large JavaScript codebase needs a type system. Do you convert it all at once or gradually?
- **Real cases:**
  - **Stripe** converted 3.7M+ lines from Flow to TypeScript **in a single pull request** (merged in March 2022), using a codemod, so engineers never had to deal with two type systems.
  - **Airbnb** built **ts-migrate**, which turns a JS project into a compiling TS project with automated codemod plugins; it reported 86% of a ~6M-line frontend converted.
- **What you teach:** codemods (automated code rewriting), when big bang is *safer* than incremental, and what preparation makes a one-PR migration possible.
- **Twist:** Pinterest also reported a 3.7M-line migration. Don't mix up the numbers; check the sources.
- **Read:** [Stripe: Migrating millions of lines of code to TypeScript](https://stripe.com/blog/migrating-to-typescript) · [airbnb/ts-migrate on GitHub](https://github.com/airbnb/ts-migrate) · [InfoQ: Airbnb TypeScript migration](https://www.infoq.com/news/2020/08/airbnb-typescript-migration)

## Changing the architecture

**A14. "Rewrite the backend in a new language, without a big-bang rewrite" (Khan Academy)**
- **Situation:** your language version is dying (Python 2 end of life, Jan 2020) and you decide to move to new services.
- **Real case:** Khan Academy's project "Goliath" moved a ~1M-line Python 2 monolith to Go services with a GraphQL API. Incremental method: when Go could handle a request, the gateway called **both** Python and Go, compared the results, logged differences and still returned Python's answer; then traffic moved to Go with Python as fallback; then Python was deleted. By August 2021, the new backend handled 95% of traffic. A later talk reports about 3.5 years and ~40 person-years.
- **What you teach:** patterns 3 and 5 together, the cleanest real example of each.
- **Read:** [Khan Academy: Go + Services = One Goliath Project](https://blog.khanacademy.org/go-services-one-goliath-project) · [Incremental rewrites with GraphQL](https://blog.khanacademy.org/incremental-rewrites-with-graphql) · [Beating the odds: Khan Academy's successful monolith→services rewrite](https://blog.khanacademy.org/beating-the-odds-khan-academys-successful-monolith%e2%86%92services-rewrite) · [Pragmatic Engineer: Breaking up a monolith](https://newsletter.pragmaticengineer.com/p/real-world-eng-8)

**A15. "Microservices: in, and back out again" (Airbnb, Segment, Prime Video, Shopify)**
- **Situation:** a debate video, perfect for comments. Should your team split the monolith?
- **Real cases:**
  - **Airbnb** moved from its Rails monolith ("Monorail") to services when deploys started backing up at over 100 engineers; it had ~500 engineers when the migration began in 2016 and later 1,000+ services. Payments had to run two parallel orchestration systems through most of the migration.
  - **Segment** went the other way: one service per destination became 100+ services, and developer productivity dropped; they merged them back into one service ("Centrifuge") and accepted the trade-offs.
  - **Prime Video** moved one monitoring tool from serverless pieces (Step Functions, Lambda, S3) into a single process and cut that tool's cost by 90%. The debate that followed was about this specific workload, not microservices in general.
  - **Shopify** chose a third path: a **modular monolith**, one codebase with enforced boundaries between components.
- **What you teach:** the real trade-offs (team size, deploy independence, operational cost), and why "it depends on your stage" is the honest answer.
- **Read:** [InfoQ: Airbnb monolith → SOA](https://www.infoq.com/news/2019/02/airbnb-monolith-migration-soa/) · [Segment: Goodbye Microservices](https://twilio.com/en-us/blog/developers/best-practices/goodbye-microservices) · [InfoQ: Prime Video monolith](https://infoq.com/news/2023/05/prime-ec2-ecs-saves-costs) · [Shopify: Deconstructing the Monolith](https://shopify.engineering/deconstructing-monolith-designing-software-maximizes-developer-productivity) · [Shopify: follow-up post on the state of its monolith](https://shopify.engineering/shopify-monolith)
- **(2 days)**: or split it into two videos: "why teams split" and "why teams merge back".

## APIs and security

**A16. "Change your API without breaking anyone who uses it" (Stripe)**
- **Situation:** thousands of businesses call your API; any change can break their code, and you can't make them all update.
- **Real case:** Stripe has kept compatibility with every API version since 2011. Each account is pinned to the version it first used, and changes are released as dated versions. Stripe still tries hard to get the design right the first time so it rarely needs a new version.
- **What you teach:** additive changes, version pinning, translating responses between versions inside the server, deprecation.
- **Twist:** your mobile app is a client of your own API, and old app versions stay installed for years.
- **Read:** [Stripe: APIs as infrastructure: future-proofing Stripe with versioning](https://stripe.com/ae/blog/api-versioning) · [Stripe docs: API versioning](https://docs.stripe.com/api/versioning?lang=curl)

**A17. "Our password hashing is weak. Fix it without forcing everyone to reset" (Dropbox + OWASP)**
- **Situation:** you discover passwords are stored with a fast hash.
- **Real case:** Dropbox's 2012 breach exposed credentials, some hashed with SHA-1 and others with bcrypt. Its later design is layered: SHA-512 first (to avoid bcrypt's 72-byte input limit), then bcrypt with a per-user salt, then AES-256 encryption with a key stored separately.
- **What you teach:** fast vs slow hashes, salt and pepper, and the migration trick: wrap old hashes immediately (`bcrypt(old_hash)`), then re-hash to the new scheme when each user next logs in.
- **Read:** [Dropbox: How Dropbox securely stores your passwords](https://dropbox.tech/security/how-dropbox-securely-stores-your-passwords) · OWASP Password Storage Cheat Sheet (search the title)

---

# Part B: Everyday scenarios (no single famous case study)

These are common tasks in normal jobs. There's no one big write-up, but the solution is standard and documented, so they're still 1-day topics. Use the reference to check your plan.

| # | Scenario | The safe plan in one line | Reference to check |
|---|---|---|---|
| B1 | **ID column about to overflow 2.1 billion** | New `BIGINT` column → sync trigger → batch backfill → swap; watch out for JavaScript's 2⁵³ limit | PostgreSQL docs on `ALTER TABLE` locking |
| B2 | **Money stored as float** | Integer minor units (paisa) or `DECIMAL`; dual write; documented rounding rule; reconcile totals | Martin Fowler, "Money" pattern |
| B3 | **All timestamps stored in local time** | Classify columns (instant vs date vs wall time) → UTC for instants → convert only in the UI | Jon Skeet, "Storing UTC is not a silver bullet" |
| B4 | **Page 5,000 takes 12 seconds** | Replace OFFSET with keyset/cursor pagination | Markus Winand, *Use The Index, Luke*, "no offset" |
| B5 | **A page runs 1,200 queries (N+1)** | Measure queries per request → eager/batch loading → test that caps query count | Your ORM's eager-loading docs; GraphQL DataLoader |
| B6 | **The report button times out** | Async job + job ID + polling/notification + file in object storage | Microsoft Azure Architecture Center, "Asynchronous Request-Reply pattern" |
| B7 | **Switching payment providers** | Your own `PaymentProvider` interface → new provider behind a flag → old payments finish on the old one | Martin Fowler, "BranchByAbstraction" |
| B8 | **A big customer wants isolated data** | Tenant-aware data layer / row-level security → routing table → move one tenant | PostgreSQL docs, Row Level Security |
| B9 | **Changing an event 12 services consume** | Additive changes only; schema registry; versioned event for breaking changes | DDIA chapter 4; Confluent Schema Registry docs |
| B10 | **A 1-million-row CSV import fails halfway** | Upload → validate → chunked async processing → idempotent upsert → error report | Your database's bulk-load docs |
| B11 | **An API key was pushed to GitHub** | Rotate first → check logs → clean history → secret manager + scanning | GitHub docs, "Removing sensitive data from a repository" |
| B12 | **"Delete my data" across 30 services** | Data inventory → `UserDeleted` event → per-service tracking → crypto-shredding for backups | Search "crypto shredding" |
| B13 | **Changing code that has no tests** | Characterization tests → small refactors → shadow-run old vs new | Michael Feathers, *Working Effectively with Legacy Code*; GitHub's Scientist library |
| B14 | **CI takes 45 minutes** | Measure → cache → parallelize → run affected tests only → quarantine flaky tests | Google Testing Blog on flaky tests |
| B15 | **50 million files from local disk to cloud storage** | New uploads to object storage → read with fallback → resumable checksummed copy → remove fallback | Your cloud provider's bulk transfer docs |

---

## Where to find more topics

**Collections and newsletters (to discover topics)**
- [Pragmatic Engineer: Real-World Engineering Challenges](https://newsletter.pragmaticengineer.com/p/real-world-4): a series summarizing real engineering stories. Issue **#6 is about migrations** (Box, Stripe, Pinterest, DoorDash, LinkedIn); **#8 is Khan Academy**. Also [Migrations Done Well: typical migration approaches](https://blog.pragmaticengineer.com/typical-migration-approaches/). Some content is paid.
- [kilimchoi/engineering-blogs](https://github.com/kilimchoi/engineering-blogs): a list of 400+ company engineering blogs. Not updated recently, so some links may be dead.
- [Simon Willison's link blog](https://simonwillison.net/): he links and summarizes many of the best engineering posts (Stripe, Figma, Discord, Slack all appeared there). Search his site for "migration".
- **InfoQ** (infoq.com): conference talk recordings and news summaries, e.g. Airbnb's SOA migration, Segment's microservices talk, Netflix's data migrations.
- **Hacker News** (search old discussions at hn.algolia.com): almost every post above was discussed there, and the comments add counter-arguments and twists for your videos.

**Company engineering blogs that publish migration stories**
- Stripe, GitHub, Shopify, Dropbox, Discord, Figma, Notion, Slack, Uber, Netflix, Airbnb, Khan Academy, Wix, Gusto. Search: `<company> engineering blog migration` or `<company> zero downtime`.

**Patterns and books (for the "how" part)**
- [Martin Fowler: Parallel Change](https://www.martinfowler.com/bliki/ParallelChange.html), Strangler Fig Application, Branch by Abstraction.
- *Designing Data-Intensive Applications* (Kleppmann), *Monolith to Microservices* (Sam Newman), *Working Effectively with Legacy Code* (Feathers), *Refactoring Databases* (Ambler & Sadalage).

**Summaries vs originals:** sites like ByteByteGo, Medium, dev.to and YouTube summaries are good for *finding* topics, but some are AI-generated or contain mistakes. Always take your facts from the company's original post.

**Your best source is still your own experience:** list every migration or tricky change from your previous jobs. Those are videos only you can make, and they're your answers to "tell me about a time you changed a system safely" in interviews.

---

## Sources used for this list

- [Stripe: Online migrations at scale](https://stripe.com/se/blog/online-migrations) · [Simon Willison on it](https://simonwillison.net/2023/Nov/5/online-migrations-at-scale/)
- [Gusto: old write](https://engineering.gusto.com/old-write) · [Wix: zero-downtime DB migration service](https://www.wix.engineering/post/how-we-built-a-zero-downtime-database-migration-service-at-wix) · [Tines: zero downtime database migrations](https://www.tines.com/blog/zero-downtime-database-migrations-lessons-from-moving-a-live-production/) · [LaunchDarkly: database migration](https://launchdarkly.com/blog/database-migration-using-launchdarkly/) · [LaunchDarkly: 3 best practices](https://launchdarkly.com/blog/3-best-practices-for-zero-downtime-database-migrations/)
- [Figma: databases team lived to tell the scale](https://www.figma.com/en-jp/blog/how-figmas-databases-team-lived-to-tell-the-scale/) · [Simon Willison on Figma](https://simonwillison.net/2024/Mar/14/how-figmas-databases-team-lived-to-tell-the-scale) · [pganalyze on DBProxy](https://pganalyze.com/blog/5mins-postgres-figma-dbproxy-sharding-postgres)
- [Notion: Herding elephants](https://notion.com/blog/sharding-postgres-at-notion) · [Notion: The Great Re-shard](https://notion.com/blog/the-great-re-shard)
- [Discord: How Discord stores trillions of messages](https://discord.com/blog/how-discord-stores-trillions-of-messages) · [InfoQ on Discord](https://www.infoq.com/news/2023/06/discord-cassandra-scylladb/)
- [Uber: Postgres to MySQL](https://www.uber.com/blog/postgres-to-mysql-migration/) · [LWN](https://lwn.net/Articles/695999/)
- [GitHub: gh-ost](https://github.blog/2016-08-01-gh-ost-github-s-online-migration-tool-for-mysql) · [gh-ost repo](https://github.com/github/gh-ost) · [GitHub: MySQL 8.0 upgrade](https://github.blog/engineering/infrastructure/upgrading-github-com-to-mysql-8-0/) · [GitHub: Rails 3.2 → 5.2](https://github.blog/engineering/upgrading-github-from-rails-3-2-to-5-2/)
- [Shopify: shard balancing](https://shopify.engineering/blogs/engineering/mysql-database-shard-balancing-terabyte-scale) · [Arpit Bhayani on Shopify hot shards](https://arpitbhayani.me/videos/shopify-zero-downtime-data-migration-hot-shards) · [Rails at Scale: migration backends](https://railsatscale.com/2025-12-08-swappable-migration-backends-in-rails/) · [Shopify: Deconstructing the Monolith](https://shopify.engineering/deconstructing-monolith-designing-software-maximizes-developer-productivity) · [Shopify monolith update](https://shopify.engineering/shopify-monolith)
- [InfoQ: Netflix data migrations](https://www.infoq.com/news/2018/11/netflix-data-migrations)
- [Slack: Vitess](https://slack.engineering/scaling-datastores-at-slack-with-vitess/) · [CNCF: Slack case study](https://www.cncf.io/case-studies/slack/)
- [Dropbox: Python 3 rollout](https://dropbox.tech/application/how-we-rolled-out-one-of-the-largest-python-3-migrations-ever) · [Dropbox: incremental Python 3 migration](https://dropbox.tech/application/incrementally-migrating-over-one-million-lines-of-code-from-python-2-to-python-3) · [Dropbox: password storage](https://dropbox.tech/security/how-dropbox-securely-stores-your-passwords)
- [Stripe: TypeScript migration](https://stripe.com/blog/migrating-to-typescript) · [airbnb/ts-migrate](https://github.com/airbnb/ts-migrate) · [InfoQ: Airbnb TypeScript](https://www.infoq.com/news/2020/08/airbnb-typescript-migration)
- [Khan Academy: Goliath](https://blog.khanacademy.org/go-services-one-goliath-project) · [Incremental rewrites with GraphQL](https://blog.khanacademy.org/incremental-rewrites-with-graphql) · [Beating the odds](https://blog.khanacademy.org/beating-the-odds-khan-academys-successful-monolith%e2%86%92services-rewrite) · [Pragmatic Engineer #8](https://newsletter.pragmaticengineer.com/p/real-world-eng-8)
- [InfoQ: Airbnb SOA](https://www.infoq.com/news/2019/02/airbnb-monolith-migration-soa/) · [Segment: Goodbye Microservices](https://twilio.com/en-us/blog/developers/best-practices/goodbye-microservices) · [InfoQ: Segment](https://www.infoq.com/news/2018/07/segment-microservices) · [InfoQ: Prime Video](https://infoq.com/news/2023/05/prime-ec2-ecs-saves-costs) · [The New Stack: Prime Video](https://thenewstack.io/return-of-the-monolith-amazon-dumps-microservices-for-video-monitoring/)
- [Stripe: API versioning](https://stripe.com/ae/blog/api-versioning) · [Stripe docs: versioning](https://docs.stripe.com/api/versioning?lang=curl)
- [Pragmatic Engineer: Real-World Engineering Challenges #4](https://newsletter.pragmaticengineer.com/p/real-world-4) · [Migrations Done Well](https://blog.pragmaticengineer.com/typical-migration-approaches/)
- [Martin Fowler: Parallel Change](https://www.martinfowler.com/bliki/ParallelChange.html) · [Wikipedia: Strangler fig pattern](https://en.wikipedia.org/wiki/Strangler_fig_pattern) · [Microsoft: Strangler Fig](https://learn.microsoft.com/azure/architecture/patterns/strangler-fig)
- [kilimchoi/engineering-blogs](https://github.com/kilimchoi/engineering-blogs)
