# Study Resources: Episodes 1-43

*What to read, watch and build to prepare each episode in depth · Nov 13, 2026 · @Lion*

Companion to [engineering-scenarios-quick-topics.md](engineering-scenarios-quick-topics.md). Episode numbers and titles match the episode order there.

---

## How to use this

Each episode has four kinds of resources:

| Tier | What it is | When |
|---|---|---|
| **Core** | The resource that teaches the concept properly. Read this first, fully. | Always (morning of prep day) |
| **Case** | The real-world story you'll tell in the video. | Always |
| **Deeper** | Extra depth for the comments, follow-up questions and your own understanding. | For your best preparation; skim if short on time |
| **Practice** | A small hands-on exercise. Real numbers from your own demo make the video stronger. | Strongly recommended, especially from Level 2 onwards |

**Time:** Core + Case fits the 1-day routine. Adding Deeper and Practice makes it 1.5-2 days, which is what "highest preparation" really takes. A good rhythm: full depth for every second episode, Core + Case for the rest.

**About these links:** most are well-known primary sources (official docs, books, company engineering blogs, original papers). Some links were found through web searches; where I wasn't sure of the exact URL, the entry says "search the title" instead. If a link is dead, search the title.

---

## Core library (used across many episodes)

You don't need to buy everything at once. The two starred ones cover the most episodes.

**Books**
- ⭐ **Designing Data-Intensive Applications** (Martin Kleppmann; 2nd edition with Chris Riccomini, 2026). The best single book for Levels 2-4. Chapter numbers differ between editions, so this guide refers to chapters by **topic** (Replication, Sharding/Partitioning, Transactions, Encoding and Evolution, Stream Processing). [dataintensive.net](https://dataintensive.net/)
- ⭐ **Site Reliability Engineering** and **The Site Reliability Workbook** (Google), **free online** at [sre.google/books](https://sre.google/books/). Used in the Level 3 reliability episodes.
- **Release It!** (Michael Nygard, 2nd ed.): stability patterns (timeouts, circuit breakers, bulkheads, load shedding, back-pressure).
- **System Design Interview Vol. 1** (Alex Xu): short chapters on rate limiters, consistent hashing, unique IDs and chat systems. Good for a quick overview, not depth.
- **Software Engineering at Google**, **free online** at [abseil.io/resources/swe-book](https://abseil.io/resources/swe-book): code, testing, dependencies and large-scale changes.
- **Working Effectively with Legacy Code** (Michael Feathers): changing code without tests.

**Free sites**
- [Use The Index, Luke](https://use-the-index-luke.com/): SQL indexing and query performance.
- [Martin Fowler's website](https://martinfowler.com/): patterns for safe change (Parallel Change, Strangler Fig, Branch by Abstraction, Feature Toggles).
- [AWS Builders' Library](https://aws.amazon.com/builders-library/): Amazon engineers on timeouts, retries, load shedding, queues and caching. Excellent and free.
- [OWASP Cheat Sheet Series](https://cheatsheetseries.owasp.org/): security topics.
- [PortSwigger Web Security Academy](https://portswigger.net/web-security): free security labs.
- [web.dev Learn Performance](https://web.dev/learn/performance): web performance.

**Free university lectures (for Levels 3-4)**
- **CMU 15-445 Database Systems** (Andy Pavlo), on YouTube: how databases work inside.
- **MIT 6.5840 Distributed Systems** (formerly 6.824), lectures and papers online.
- **Martin Kleppmann's Distributed Systems lecture series** (Cambridge), on YouTube.

---

# Level 1: "The bugs every backend developer meets"

### Ep 1. A page runs 1,200 queries (N+1)
- **Core:** [Use The Index, Luke: Nested Loops (the N+1 problem)](https://use-the-index-luke.com/sql/join/nested-loops-join-n1-problem)
- **Core:** the docs for your data layer on loading related data, e.g. [GORM Preload](https://gorm.io/docs/preload.html) (Go), Mongoose `populate`, Prisma's query optimization docs
- **Case:** [Shopify: Solving the N+1 problem for GraphQL through batching](https://shopify.engineering/solving-the-n-1-problem-for-graphql-through-batching)
- **Deeper:** [graphql/dataloader README](https://github.com/graphql/dataloader): the batching idea in its simplest form
- **Practice:** build one endpoint with N+1 and one with batching (`WHERE id IN (...)` or `$in`) against about 1,000 rows. Log query count and response time for both.

### Ep 2. The database is almost idle, but the API keeps timing out
- **Core:** [HikariCP: About Pool Sizing](https://github.com/brettwooldridge/HikariCP/wiki/About-Pool-Sizing): why a *smaller* pool is often faster (Java library, but the reasoning is universal)
- **Core:** [Go docs: Managing connections](https://go.dev/doc/database/manage-connections), or your driver's pool docs
- **Case:** [Layers: Using multiple PgBouncers to isolate application and worker connections](https://www.uselayers.com/articles/using-multiple-pgbouncers-to-isolate-application-and-worker-connections) · [Decagon: Debugging a latent PgBouncer bug](https://decagon.ai/blog/debugging-a-latent-pgbouncer-bug-across-four-layers)
- **Deeper:** *Release It!*, the Bulkheads pattern · [PgBouncer docs](https://www.pgbouncer.org/)
- **Practice:** set your pool to 5 connections, add one slow endpoint (`pg_sleep(5)`), load-test both endpoints together and watch fast requests time out. Then give each a separate pool.

### Ep 3. A popular cache key expires and the database melts
- **Core:** [AWS Builders' Library: Caching challenges and strategies](https://aws.amazon.com/builders-library/caching-challenges-and-strategies/)
- **Core:** [Wikipedia: Cache stampede](https://en.wikipedia.org/wiki/Cache_stampede): short summary of the three standard fixes (locking, early recomputation, probabilistic early expiration)
- **Case:** "Scaling Memcache at Facebook" (NSDI 2013; search the title for the PDF), read sections on leases only · [MIT 6.824 FAQ on the paper](https://pdos.csail.mit.edu/6.824/papers/memcache-faq.txt) · [Shopify: storefront rendering with caching](https://shopify.engineering/simplify-batch-cache-optimized-server-side-storefront-rendering)
- **Deeper:** "Optimal Probabilistic Cache Stampede Prevention" (Vattani et al., VLDB 2015; search the title): the math behind early expiration
- **Practice:** cache one hot key in Redis with a short TTL, send 500 concurrent requests as it expires, count database hits. Add a lock (`SET NX`) so only one request refills it, and count again.

### Ep 4. Your average response time is lying to you
- **Core:** [The Morning Paper: The Tail at Scale](https://blog.acolyer.org/2015/01/15/the-tail-at-scale/), then the original paper (Dean & Barroso, CACM 2013; search the title)
- **Core:** Gil Tene's talk **"How NOT to Measure Latency"** (search the title on YouTube): percentiles and "coordinated omission"
- **Case:** the hedged-requests results in "The Tail at Scale"
- **Deeper:** *SRE book*, [Monitoring Distributed Systems](https://sre.google/sre-book/monitoring-distributed-systems/) (the "four golden signals")
- **Practice:** load-test your Ep 1 demo with a tool like k6 and show average vs p50 vs p99 side by side. Add one random slow dependency and watch p99 change while the average barely moves.

### Ep 5. An endpoint got 10× slower after a release
- **Core:** [Brendan Gregg: Flame Graphs](https://www.brendangregg.com/flamegraphs.html) · [The USE Method](https://www.brendangregg.com/usemethod.html)
- **Core:** profiling in your language: [Go: Profiling Go Programs (pprof)](https://go.dev/blog/pprof), or Node.js CPU profiling with Chrome DevTools
- **Case:** no single famous case. Use a regression you create in your demo (see Practice).
- **Deeper:** *SRE Workbook*, "Canarying Releases" chapter (at [sre.google/books](https://sre.google/books/)): catching regressions before everyone gets them
- **Practice:** add an accidental O(n²) loop or a missing index to your demo, measure the slowdown, find it with a flame graph, fix it, measure again.

### Ep 6. Memory keeps growing until the server restarts
- **Core:** [Node.js docs: Using Heap Snapshot](https://nodejs.org/en/learn/diagnostics/memory/using-heap-snapshot) · [Using Heap Profiler](https://nodejs.org/id/learn/diagnostics/memory/using-heap-profiler)
- **Core:** [Chrome DevTools: Fix memory problems](https://developer.chrome.com/docs/devtools/memory-problems)
- **Case:** no single famous case. Use the common causes: unbounded in-memory caches, event listeners never removed, global arrays, closures holding large objects.
- **Deeper:** Go heap profiling with [pprof](https://go.dev/blog/pprof), to show the idea is the same in any language
- **Practice:** write a leak on purpose (an in-memory cache that never evicts), take two heap snapshots, find it by comparing them, fix it with a size limit (LRU).

### Ep 7. The report button times out
- **Core:** [Microsoft Azure: Asynchronous Request-Reply pattern](https://learn.microsoft.com/azure/architecture/patterns/async-request-reply)
- **Core:** [Microsoft Azure: Queue-Based Load Leveling pattern](https://learn.microsoft.com/azure/architecture/patterns/queue-based-load-leveling)
- **Case:** no single famous case; this is a standard pattern.
- **Deeper:** [AWS Builders' Library: Avoiding insurmountable queue backlogs](https://aws.amazon.com/builders-library/avoiding-insurmountable-queue-backlogs)
- **Practice:** turn a slow endpoint into "create job → return job ID → worker processes from a Redis queue → client polls status".

### Ep 8. A customer paid ৳1 for a ৳500 order: never trust the client
- **Core:** [PortSwigger: Business logic vulnerabilities](https://portswigger.net/web-security/logic-flaws) · [Examples](https://portswigger.net/web-security/logic-flaws/examples)
- **Case:** [PortSwigger lab: Excessive trust in client-side controls](https://portswigger.net/web-security/logic-flaws/examples/lab-logic-flaws-excessive-trust-in-client-side-controls). Do the lab yourself (free account) and show it on screen.
- **Deeper:** [OWASP Top 10: A04 Insecure Design](https://owasp.org/Top10/A04_2021-Insecure_Design/) · OWASP Web Security Testing Guide, "Business Logic Testing" section (search the title)
- **Practice:** do 3-4 of PortSwigger's business logic labs (prices, quantities, coupons). They're free and make great video footage.

### Ep 9. Money stored as float
- **Core:** [The Floating-Point Guide](https://floating-point-gui.de/): short and clear
- **Core:** [Stripe docs: Currencies](https://docs.stripe.com/currencies): amounts in the smallest currency unit, and zero-decimal currencies
- **Case:** Martin Fowler's **Money** pattern (*Patterns of Enterprise Application Architecture*; also search "Martin Fowler Money pattern")
- **Deeper:** "What Every Computer Scientist Should Know About Floating-Point Arithmetic" (David Goldberg; search the title). Heavy; read the first sections only.
- **Practice:** show `0.1 + 0.2` in JavaScript and Go, then split ৳100 three ways and show where the leftover paisa goes with integer math.

### Ep 10. All timestamps stored in local time
- **Core:** [PostgreSQL docs: Date/Time types](https://www.postgresql.org/docs/current/datatype-datetime.html), especially `timestamp` vs `timestamptz`
- **Core:** Jon Skeet, **"Storing UTC is not a silver bullet"** (search the title)
- **Case:** "Falsehoods programmers believe about time" (Noah Sussman; search the title), for the hook
- **Deeper:** the IANA time zone database (search "IANA tz database"): why time zone *rules* change
- **Practice:** store a reservation as UTC, as local time without a zone, and as a "wall time + zone", then show which one breaks when a user's time zone changes.

### Ep 11. An API key was pushed to GitHub
- **Core:** [GitHub docs: Removing sensitive data from a repository](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/removing-sensitive-data-from-a-repository)
- **Core:** [OWASP: Secrets Management Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Secrets_Management_Cheat_Sheet.html)
- **Case:** GitGuardian's yearly **"State of Secrets Sprawl"** report (search the title) for real numbers on leaked secrets
- **Deeper:** GitHub secret scanning and push protection docs (search "GitHub secret scanning push protection") · tools: gitleaks, trufflehog
- **Practice:** in a throwaway repo, commit a fake key, delete it in a new commit, then show it's still in history. Run gitleaks to find it.

### Ep 12. One regex took down Cloudflare's network
- **Core:** [Cloudflare: Details of the Cloudflare outage on July 2, 2019](https://blog.cloudflare.com/details-of-the-cloudflare-outage-on-july-2-2019/). Read it fully; it's one of the best postmortems ever written.
- **Core:** [Russ Cox: Regular Expression Matching Can Be Simple And Fast](https://swtch.com/~rsc/regexp/regexp1.html): why backtracking explodes (fits your algorithm background)
- **Case:** the Cloudflare post above · [OWASP: ReDoS](https://owasp.org/www-community/attacks/Regular_expression_Denial_of_Service_-_ReDoS)
- **Deeper:** Stack Overflow's July 20, 2016 outage postmortem (search the title): same root cause, different company
- **Practice:** time a backtracking-prone regex on growing inputs in JavaScript, then run the same pattern in Go's `regexp` (which doesn't backtrack) and compare.

### Ep 13. Does competitive programming help in a real job? (debate)
- **Core:** Erik Bernhardsson, ["Norvig's claim that programming competitions correlate negatively with being good on the job"](https://erikbern.com/2015/04/07/norvigs-claim-that-programming-competitions-correlate-negatively-with-being-good-on-the-job): the famous claim, explained with selection bias (Berkson's paradox)
- **Case:** Peter Norvig's remark (Q&A of a 2015 lecture; see the post above) · the counter-argument: ["No, programming competitions don't produce bad engineers"](https://ruberik.blogspot.com/2017/07/no-programming-competitions-dont.html)
- **Deeper:** the [Codeforces discussion of the claim](https://codeforces.com/blog/entry/17302) and its Hacker News threads, for both sides
- **Practice:** list 3 backend problems from Episodes 1-12 where an algorithm idea mattered (e.g. Ep 12 regex backtracking, Ep 3 LRU eviction) and 3 where it didn't. That list is your script.

---

# Level 2: "Changing things without breaking them"

### Ep 14. How big companies change things with zero downtime (the 6 patterns)
- **Core:** Martin Fowler: [Parallel Change](https://martinfowler.com/bliki/ParallelChange.html) · [Strangler Fig Application](https://martinfowler.com/bliki/StranglerFigApplication.html) · [Branch By Abstraction](https://martinfowler.com/bliki/BranchByAbstraction.html) · [Feature Toggles](https://martinfowler.com/articles/feature-toggles.html)
- **Case:** [Stripe: Online migrations at scale](https://stripe.com/blog/online-migrations)
- **Deeper:** [Pragmatic Engineer: Migrations done well](https://blog.pragmaticengineer.com/typical-migration-approaches/) · [GitHub's Scientist library](https://github.com/github/scientist) (shadow comparison)
- **Practice:** draw each pattern as a 3-panel before/during/after diagram. These drawings become your recurring visuals for Level 2.

### Ep 15. Switching payment providers
- **Core:** [Martin Fowler: Branch By Abstraction](https://martinfowler.com/bliki/BranchByAbstraction.html)
- **Core:** [Stripe: Designing robust and predictable APIs with idempotency](https://stripe.com/blog/idempotency): payments must be safe to retry, whichever provider
- **Case:** no single famous case. Compare how two real gateways model a payment (e.g. Stripe vs a local gateway's docs) to show why an interface is needed.
- **Deeper:** payment provider docs on webhooks and reconciliation (search "Stripe webhooks best practices")
- **Practice:** write a `PaymentProvider` interface in Go with two fake implementations and a flag that routes a percentage of payments to the new one.

### Ep 16. Change your API without breaking anyone
- **Core:** [Stripe: APIs as infrastructure: future-proofing Stripe with versioning](https://stripe.com/blog/api-versioning)
- **Core:** [Google AIP-180: Backwards compatibility](https://google.aip.dev/180): a precise list of what counts as a breaking change
- **Case:** Stripe's versioning (above) · [Stripe docs: API versioning](https://docs.stripe.com/api/versioning)
- **Deeper:** [Microsoft REST API Guidelines](https://github.com/microsoft/api-guidelines) (versioning section)
- **Practice:** take one endpoint, make a breaking change the safe way: add the new field, keep the old one, track who still uses it, then remove it.

### Ep 17. Split a huge table into a new one while the app is live
- **Core:** [Stripe: Online migrations at scale](https://stripe.com/blog/online-migrations)
- **Case:** Stripe (above) · [Gusto: splitting a large table](https://engineering.gusto.com/old-write)
- **Deeper:** Martin Kleppmann, **"Using logs to build a solid data infrastructure (or: why dual writes are a bad idea)"** (search the title) · DDIA, the Stream Processing chapter (change data capture)
- **Practice:** in PostgreSQL, move a column into a new table with dual write → backfill in batches → compare → switch reads, while a script keeps writing.

### Ep 18. Our password hashing is weak. Fix it without forcing resets
- **Core:** [OWASP: Password Storage Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html)
- **Core:** [Coda Hale: How to safely store a password](https://codahale.com/how-to-safely-store-a-password/): fast vs slow hashes, in plain words
- **Case:** [Dropbox: How Dropbox securely stores your passwords](https://dropbox.tech/security/how-dropbox-securely-stores-your-passwords)
- **Deeper:** [NIST SP 800-63B](https://pages.nist.gov/800-63-3/sp800-63b.html), the password guidance section
- **Practice:** time MD5 vs bcrypt in Go for 1,000 hashes, then implement "wrap old hash now, re-hash at next login".

### Ep 19. Add "Login with Google" to an existing app
- **Core:** [OAuth 2.0 Simplified (Aaron Parecki)](https://www.oauth.com/) · OpenID Connect basics at [openid.net](https://openid.net/) (search "How OpenID Connect works")
- **Core:** [Firebase: Link multiple auth providers to an account](https://firebase.google.com/docs/auth/web/account-linking)
- **Case:** no single famous case; the account-linking problem (same email, two login methods) is the story
- **Deeper:** [OWASP: Authentication Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Authentication_Cheat_Sheet.html)
- **Practice:** add Google login to a small app that already has email/password, and handle "this email already has an account".

### Ep 20. The web page takes 8 seconds on a phone
- **Core:** [web.dev: Learn Performance](https://web.dev/learn/performance) · [Web Vitals](https://web.dev/articles/vitals)
- **Core:** [Next.js docs: Image component](https://nextjs.org/docs/app/api-reference/components/image)
- **Case:** [Pinterest: A one year PWA retrospective](https://medium.com/pinterest-engineering/a-one-year-pwa-retrospective-f4a2f4129e05) · Addy Osmani, "A Pinterest Progressive Web App Performance Case Study" (search the title)
- **Deeper:** Chrome DevTools Performance panel and Lighthouse docs
- **Practice:** run Lighthouse on a page with mobile throttling, fix images, bundle size and caching, and show the before/after scores.

### Ep 21. Changing code that has no tests
- **Core:** *Working Effectively with Legacy Code* (Feathers): chapters on characterization tests and dependency-breaking
- **Core:** [Wikipedia: Characterization test](https://en.wikipedia.org/wiki/Characterization_test)
- **Case:** [GitHub's Scientist library](https://github.com/github/scientist): GitHub's tool for running old and new code side by side in production
- **Deeper:** Martin Fowler, *Refactoring* (2nd ed.), the first chapter's worked example
- **Practice:** take a messy 100-line function, write characterization tests that record its current output, then refactor with tests green.

### Ep 22. Upgrade a core dependency years out of date
- **Core:** [Software Engineering at Google: Dependency Management](https://abseil.io/resources/swe-book/html/ch21.html)
- **Case:** [GitHub: Upgrading GitHub from Rails 3.2 to 5.2](https://github.blog/engineering/upgrading-github-from-rails-3-2-to-5-2/)
- **Deeper:** [GitHub: Upgrading GitHub.com to MySQL 8.0](https://github.blog/engineering/infrastructure/upgrading-github-com-to-mysql-8-0/), the same discipline for infrastructure
- **Practice:** read a Next.js major version upgrade guide and list which changes could be made *before* the upgrade (compatible with both versions).

### Ep 23. Big bang or step by step? Changing millions of lines (debate)
- **Core:** [Software Engineering at Google: Large-Scale Changes](https://abseil.io/resources/swe-book/html/ch22.html)
- **Case:** [Stripe: Migrating millions of lines of code to TypeScript](https://stripe.com/blog/migrating-to-typescript) · [airbnb/ts-migrate](https://github.com/airbnb/ts-migrate)
- **Deeper:** [jscodeshift](https://github.com/facebook/jscodeshift): how codemods work
- **Practice:** write a tiny codemod that renames a function across a sample project.

### Ep 24. CI takes 45 minutes
- **Core:** [Martin Fowler: Continuous Integration](https://martinfowler.com/articles/continuousIntegration.html)
- **Core:** [Google Testing Blog: Flaky tests at Google and how we mitigate them](https://testing.googleblog.com/2016/05/flaky-tests-at-google-and-how-we.html)
- **Case:** Google's flaky-test data (above)
- **Deeper:** [Martin Fowler: Eradicating non-determinism in tests](https://martinfowler.com/articles/nonDeterminism.html) · your CI provider's docs on caching and test splitting
- **Practice:** take a slow GitHub Actions workflow, add dependency caching and parallel jobs, and compare run times.

### Ep 25. Writing specs so AI coding agents don't break your code
- **Core:** [GitHub: Spec-driven development with AI](https://github.blog/ai-and-ml/generative-ai/spec-driven-development-with-ai-get-started-with-a-new-open-source-toolkit/)
- **Core:** [Claude Code docs: Best practices](https://code.claude.com/docs/en/best-practices)
- **Case:** GitHub's Spec Kit project (github.com/github/spec-kit)
- **Deeper:** Anthropic, **"Building effective agents"** (search the title)
- **Practice:** give an AI agent the same feature twice, once with a vague prompt and once with a proper spec (goal, constraints, files not to touch, how to verify), and compare the results on screen.

---

# Level 3: "Systems under pressure"

### Ep 26. Our job queue ran out of memory, and everything stopped
- **Core:** [Fred Hebert: Queues Don't Fix Overload](https://ferd.ca/queues-don-t-fix-overload.html): short and memorable
- **Core:** [AWS Builders' Library: Avoiding insurmountable queue backlogs](https://aws.amazon.com/builders-library/avoiding-insurmountable-queue-backlogs)
- **Case:** [Slack: Scaling Slack's job queue](https://slack.engineering/scaling-slacks-job-queue/)
- **Deeper:** DDIA, the Stream Processing chapter (logs vs traditional queues) · Kafka's introduction docs (search "Apache Kafka introduction")
- **Practice:** fill a Redis queue faster than workers drain it, watch memory, then add back-pressure (reject or slow producers when the queue is too long).

### Ep 27. Some customers are flooding our API
- **Core:** [Stripe: Scaling your API with rate limiters](https://stripe.com/blog/rate-limiters)
- **Core:** [AWS Builders' Library: Using load shedding to avoid overload](https://aws.amazon.com/builders-library/using-load-shedding-to-avoid-overload/)
- **Case:** Stripe (above)
- **Deeper:** *SRE book*, [Handling Overload](https://sre.google/sre-book/handling-overload/) · *System Design Interview Vol. 1*, rate limiter chapter (algorithms)
- **Practice:** build a token-bucket rate limiter in Go backed by Redis, and a load shedder that rejects low-priority requests above a CPU or concurrency limit.

### Ep 28. Everyone came back from holiday and autoscaling made it worse
- **Core:** *SRE book*, [Addressing Cascading Failures](https://sre.google/sre-book/addressing-cascading-failures/)
- **Core:** [AWS Builders' Library: Timeouts, retries, and backoff with jitter](https://aws.amazon.com/builders-library/timeouts-retries-and-backoff-with-jitter)
- **Case:** [Slack: Slack's outage on January 4th 2021](https://slack.engineering/slacks-outage-on-january-4th-2021/)
- **Deeper:** [AWS Architecture Blog: Exponential backoff and jitter](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/) · [Surfing Complexity: Slack's Jan 2021 outage](https://surfingcomplexity.blog/2021/02/08/slacks-jan-2021-outage-a-tale-of-saturation/)
- **Practice:** simulate 1,000 clients retrying at once vs with exponential backoff + jitter, and chart the load on the server.

### Ep 29. The dashboard is green, but users are angry (SLOs and alerting)
- **Core:** *SRE book*, [Service Level Objectives](https://sre.google/sre-book/service-level-objectives/) and [Monitoring Distributed Systems](https://sre.google/sre-book/monitoring-distributed-systems/)
- **Core:** *SRE Workbook*, "Implementing SLOs" and "Alerting on SLOs" chapters (at [sre.google/books](https://sre.google/books/))
- **Case:** Rob Ewaschuk, **"My Philosophy on Alerting"** (search the title): written by a Google SRE, the classic on alerting on symptoms
- **Deeper:** *SRE book*, [Postmortem Culture](https://sre.google/sre-book/postmortem-culture/)
- **Practice:** define 2 SLOs for your demo API (e.g. 99% of requests under 300 ms), calculate the error budget, and write one alert rule that fires on budget burn.

### Ep 30. Break your own servers on purpose (Chaos Monkey)
- **Core:** [Principles of Chaos Engineering](https://principlesofchaos.org/)
- **Case:** [Gremlin: The origin of Chaos Monkey](https://gremlin.com/chaos-monkey/the-origin-of-chaos-monkey) · Netflix Tech Blog posts on Chaos Monkey and the Simian Army (search the titles)
- **Deeper:** [IEEE Spectrum: Chaos engineering saved your Netflix](https://spectrum.ieee.org/chaos-engineering-saved-your-netflix)
- **Practice:** run 3 copies of your demo behind a load balancer, kill one during a load test, and show what users see. Then fix it (health checks, retries).

### Ep 31. Replace an old system piece by piece
- **Core:** [Martin Fowler: Strangler Fig Application](https://martinfowler.com/bliki/StranglerFigApplication.html)
- **Core:** Sam Newman, *Monolith to Microservices*: the chapters on splitting the monolith
- **Case:** [Khan Academy: Go + Services = One Goliath Project](https://blog.khanacademy.org/go-services-one-goliath-project) · [Incremental rewrites with GraphQL](https://blog.khanacademy.org/incremental-rewrites-with-graphql) · [Beating the odds](https://blog.khanacademy.org/beating-the-odds-khan-academys-successful-monolith%e2%86%92services-rewrite)
- **Deeper:** [Pragmatic Engineer: Breaking up a monolith](https://newsletter.pragmaticengineer.com/p/real-world-eng-8)
- **Practice:** put a small proxy in front of an old endpoint that calls both old and new implementations, compares results and logs differences.

### Ep 32. Replace the API behind a mobile app with no downtime
- **Core:** Martin Fowler's [Feature Toggles](https://martinfowler.com/articles/feature-toggles.html) (sections on canary releases)
- **Case:** [The New Stack: Netflix's testing strategies for migrating to GraphQL](https://thenewstack.io/netflixs-testing-strategies-for-migrating-to-graphql/) · Netflix Tech Blog, "Migrating Netflix to GraphQL Safely" (search the title)
- **Deeper:** *SRE Workbook*, "Canarying Releases" chapter
- **Practice:** record real requests from your demo, replay them against old and new versions, and diff the responses (replay testing).

### Ep 33. Ship a risky change to an app on millions of devices
- **Core:** [Firebase Remote Config docs](https://firebase.google.com/docs/remote-config) (kill switches and staged feature rollout)
- **Core:** Google Play "staged rollouts" and Apple "phased release" docs (search the titles)
- **Case:** [Dropbox: How we rolled out one of the largest Python 3 migrations ever](https://dropbox.tech/application/how-we-rolled-out-one-of-the-largest-python-3-migrations-ever)
- **Deeper:** [Dropbox: Incrementally migrating over one million lines of code](https://dropbox.tech/application/incrementally-migrating-over-one-million-lines-of-code-from-python-2-to-python-3)
- **Practice:** add a remote kill switch to a web or demo app that turns a new feature off without a deploy.

### Ep 34. One repo or many? (debate)
- **Core:** [Google: Why Google stores billions of lines of code in a single repository](https://research.google/pubs/why-google-stores-billions-of-lines-of-code-in-a-single-repository/)
- **Core:** [monorepo.tools](https://monorepo.tools/): a neutral overview of monorepo features and tools
- **Case:** the Google paper (above)
- **Deeper:** Hacker News discussions of monorepo vs polyrepo (search "monorepo vs polyrepo hacker news"), for both sides
- **Practice:** put two small services and a shared library in one repo with a workspace tool, and show an atomic change across all three.

### Ep 35. Microservices: in, and back out again (debate, 2 days)
- **Core:** Martin Fowler: [Monolith First](https://martinfowler.com/bliki/MonolithFirst.html) · [Microservice Premium](https://martinfowler.com/bliki/MicroservicePremium.html)
- **Core:** Sam Newman, *Building Microservices* (2nd ed.), first chapters
- **Case:** [Segment: Goodbye Microservices](https://twilio.com/en-us/blog/developers/best-practices/goodbye-microservices) · [InfoQ: Prime Video](https://infoq.com/news/2023/05/prime-ec2-ecs-saves-costs) · [InfoQ: Airbnb monolith → SOA](https://www.infoq.com/news/2019/02/airbnb-monolith-migration-soa/) · [Shopify: Deconstructing the Monolith](https://shopify.engineering/deconstructing-monolith-designing-software-maximizes-developer-productivity)
- **Deeper:** Segment's QCon talk on InfoQ (search "Segment microservices QCon InfoQ")
- **Practice:** write a one-page decision table: team size, deploy frequency, scaling needs, operational cost. Use it as the video's main visual.

---

# Level 4: "Scale"

### Ep 36. One customer could see another customer's data (multi-tenant isolation)
- **Core:** [Microsoft Azure: Architect multitenant solutions](https://learn.microsoft.com/azure/architecture/guide/multitenant/overview)
- **Core:** [PostgreSQL docs: Row Security Policies](https://www.postgresql.org/docs/current/ddl-rowsecurity.html)
- **Case:** AWS whitepaper **"SaaS Tenant Isolation Strategies"** (search the title)
- **Deeper:** [LoginRadius: Multi-tenant authorization patterns](https://www.loginradius.com/blog/identity/what-is-multi-tenant-authorization)
- **Practice:** create a shared `orders` table, write a query that forgets the tenant filter and leaks data, then turn on row-level security and show the same query is now safe.

### Ep 37. One chat room with 30,000 people online
- **Core:** *System Design Interview Vol. 1*, chat system chapter (WebSockets, connection servers, fan-out)
- **Core:** [AWS Architecture Blog: Exponential backoff and jitter](https://aws.amazon.com/blogs/architecture/exponential-backoff-and-jitter/) (for reconnect storms)
- **Case:** [Discord: How Discord scaled Elixir to 5,000,000 concurrent users](https://discord.com/blog/how-discord-scaled-elixir-to-5-000-000-concurrent-users) · [Elixir blog: Real-time communication at scale at Discord](https://elixir-lang.org/blog/2020/10/08/real-time-communication-at-scale-with-elixir-at-discord/)
- **Deeper:** WhatsApp's 2012 post on reaching 2 million connections on one server (search "WhatsApp 1 million is so 2011")
- **Practice:** a Go WebSocket server that broadcasts to 10,000 simulated clients. Measure message delay as clients increase, then make all clients reconnect at once with and without jitter.

### Ep 38. Change the schema of a huge table without locking it
- **Core:** [ankane/strong_migrations README](https://github.com/ankane/strong_migrations): a clear list of dangerous schema operations and safe alternatives (Rails gem, but the advice applies to any app)
- **Core:** [PostgreSQL docs: ALTER TABLE](https://www.postgresql.org/docs/current/sql-altertable.html) (which operations take which locks)
- **Case:** [GitHub: gh-ost announcement](https://github.blog/2016-08-01-gh-ost-github-s-online-migration-tool-for-mysql) · [gh-ost README](https://github.com/github/gh-ost)
- **Deeper:** DDIA, the Encoding and Evolution chapter
- **Practice:** on a 10-million-row PostgreSQL table, run a locking change while a script writes continuously, measure blocked time, then do it the safe way.

### Ep 39. Move trillions of messages to a new database with zero downtime
- **Core:** DDIA, the Sharding/Partitioning chapter (hot spots and partition keys)
- **Case:** [Discord: How Discord stores trillions of messages](https://discord.com/blog/how-discord-stores-trillions-of-messages)
- **Deeper:** Discord's earlier post "How Discord stores billions of messages" (search the title): why they chose Cassandra in the first place
- **Practice:** model chat messages with a partition key of `channel_id` vs `(channel_id, time bucket)` and show how one busy channel creates a hot partition.

### Ep 40. Our database is too big for one machine. Shard it (2 days)
- **Core:** DDIA, the Sharding/Partitioning chapter (fully)
- **Core:** *System Design Interview Vol. 1*, consistent hashing chapter
- **Case:** [Figma: How Figma's databases team lived to tell the scale](https://www.figma.com/blog/how-figmas-databases-team-lived-to-tell-the-scale/) · [Notion: Herding elephants](https://notion.com/blog/sharding-postgres-at-notion) · [Notion: The Great Re-shard](https://notion.com/blog/the-great-re-shard)
- **Deeper:** CMU 15-445 lectures on distributed databases (YouTube)
- **Practice:** shard 1,000 users across 3 PostgreSQL databases by `user_id`, then add a 4th shard and count how many users must move with modulo vs consistent hashing.

### Ep 41. "Delete my data" across 30 services
- **Core:** [GDPR Article 17: Right to erasure](https://gdpr-info.eu/art-17-gdpr/)
- **Core:** [Wikipedia: Crypto-shredding](https://en.wikipedia.org/wiki/Crypto-shredding)
- **Case:** no single famous case; build the story around a data inventory and a `UserDeleted` event
- **Deeper:** DDIA, the Stream Processing chapter (events, logs, and deleting data in event-sourced systems)
- **Practice:** publish a `UserDeleted` event that three small services handle idempotently, and track which ones have finished.

### Ep 42. Two people edit the same design at the same time (2 days)
- **Core:** [crdt.tech](https://crdt.tech/): introduction and resources
- **Core:** Martin Kleppmann's talk **"CRDTs: The Hard Parts"** (search the title on YouTube)
- **Case:** [Evan Wallace: How Figma's multiplayer technology works](https://madebyevan.com/figma/how-figmas-multiplayer-technology-works/)
- **Deeper:** [Liveblocks: Understanding sync engines](https://www.liveblocks.io/blog/understanding-sync-engines-how-figma-linear-and-google-docs-work)
- **Practice:** build a tiny shared whiteboard where two browser tabs move shapes, using last-writer-wins per property through a WebSocket server, like Figma.

### Ep 43. Our cloud bill is too high. Should we leave the cloud? (finale debate)
- **Core:** a16z, [The Cost of Cloud, a Trillion Dollar Paradox](https://a16z.com/the-cost-of-cloud-a-trillion-dollar-paradox/)
- **Case:** [DHH: Our cloud-exit savings will now top ten million over five years](https://world.hey.com/dhh/our-cloud-exit-savings-will-now-top-ten-million-over-five-years-c7d9b5bd) · [InfoQ: Dropbox moves off AWS](https://www.infoq.com/news/2016/03/Dropbox-AWS/)
- **Deeper:** [The Register: 37signals completes its move](https://www.theregister.com/2025/05/09/37signals_cloud_repatriation_storage_savings) (including what the savings figures leave out)
- **Practice:** build a simple cost model for one app at 3 growth stages (cloud vs own servers), including staff and hardware refresh. This is the kind of analysis you've done for Firebase, and it makes the debate concrete.

---

## Sources checked while building this guide

- [O'Reilly: Designing Data-Intensive Applications, 2nd edition](https://oreilly.com/library/view/designing-data-intensive-applications/9781098119058/)
- [AWS Builders' Library: Timeouts, retries, and backoff with jitter](https://aws.amazon.com/builders-library/timeouts-retries-and-backoff-with-jitter) · [Avoiding insurmountable queue backlogs](https://aws.amazon.com/builders-library/avoiding-insurmountable-queue-backlogs)
- [Shopify: Solving the N+1 problem for GraphQL through batching](https://shopify.engineering/solving-the-n-1-problem-for-graphql-through-batching)
- [Erik Bernhardsson on Norvig's claim](https://erikbern.com/2015/04/07/norvigs-claim-that-programming-competitions-correlate-negatively-with-being-good-on-the-job) · [ruberik: No, programming competitions don't produce bad engineers](https://ruberik.blogspot.com/2017/07/no-programming-competitions-dont.html) · [Codeforces discussion](https://codeforces.com/blog/entry/17302)
- [Software Engineering at Google: Dependency Management](https://abseil.io/resources/swe-book/html/ch21.html) · [Large-Scale Changes](https://abseil.io/resources/swe-book/html/ch22.html)
- [Node.js: Using Heap Snapshot](https://nodejs.org/en/learn/diagnostics/memory/using-heap-snapshot) · [Using Heap Profiler](https://nodejs.org/id/learn/diagnostics/memory/using-heap-profiler)
- Case-study links are the same ones listed and sourced in [engineering-scenarios-quick-topics.md](engineering-scenarios-quick-topics.md).
