# "How Would You Do It?": Engineering Scenarios from Real Case Studies

*Real engineering problems across the whole stack, how real teams solved them, and how to do it safely · Nov 12, 2026 · @Lion*

Each topic is a situation real teams face: **"We have problem X. What's going on, and how do we fix it without breaking anything?"** Most of the video is the solution.

**Part A** has 29 topics built on **published case studies** from company engineering blogs, papers and conference talks, spread across 10 areas so the channel isn't only about databases. **Part B** has everyday scenarios from normal jobs that don't have one famous case study.

> **How this list was researched:** Part A topics and facts come from web searches of engineering blogs, papers, conference talks and newsletters (Nov 2026). I worked from search excerpts and summaries, **not the full articles**. Before recording, read the original post (the first link in each topic) and take every number from it, not from this doc.

### Topic balance

| Area | Part A topics |
|---|---|
| 1. Performance | 3 |
| 2. Reliability and incidents | 3 |
| 3. Infrastructure and cost | 2 |
| 4. Security and authentication | 2 |
| 5. Frontend, mobile and client apps | 3 |
| 6. Real-time and messaging | 3 |
| 7. APIs | 3 |
| 8. Architecture | 2 |
| 9. Codebase and team | 4 |
| 10. Data and databases | 4 |

---

## The 1-day prep routine

| Time | Task |
|---|---|
| Morning (2-3 h) | Read the original case study fully. Note: the problem, the constraints, the steps they took, what went wrong, what they'd do differently. |
| Midday (1 h) | Generalize it: draw the "before" and "after", and write the solution as steps your viewer could apply to their own system. |
| Afternoon (2 h) | Whiteboard panels, hook, bullet-point script. |
| Evening (1-2 h) | Record and do a light edit. |

## The video shape (8-12 minutes)

1. **The situation (1 min):** "Imagine your team has this problem ..." Pose it as *your* problem first.
2. **The naive fix (1-2 min):** what most people would try first, and why it doesn't work.
3. **What's really going on (2 min):** the concept behind the problem. Ask viewers "what would you check first?"
4. **The solution (4-5 min):** step by step, including how to roll it out safely. The main part.
5. **How [Company] actually did it (1-2 min):** the real story and numbers. This builds credibility.
6. **Checklist (30 s):** 3-4 rules to remember. This also makes a good LinkedIn post.

---

## The 6 patterns behind almost every safe change

Many topics below, from database moves to API changes to language upgrades, reuse the same few ideas. Make **one early video about these patterns**, then refer back to them ("this is pattern #5 again").

1. **Expand and contract (parallel change):** add the new thing next to the old, move everyone over, then remove the old. ([Martin Fowler: Parallel Change](https://www.martinfowler.com/bliki/ParallelChange.html))
2. **Dual write → backfill → compare → switch reads → stop old writes:** the standard way to move data. (Stripe, topic 10.1)
3. **Strangler fig:** replace an old system piece by piece, routing one feature at a time to the new one. ([Wikipedia: Strangler fig pattern](https://en.wikipedia.org/wiki/Strangler_fig_pattern))
4. **Branch by abstraction:** put an interface in front of the thing you're replacing, so the swap happens behind it.
5. **Shadow mode / side-by-side comparison:** run old and new together, compare results, keep serving the old one until they match. (Khan Academy, Netflix, Stripe)
6. **Gradual rollout with an escape hatch:** staged groups, feature flags, canaries, and a way to go back instantly. (GitHub 2FA, Dropbox, Cloudflare's lesson)

---

# Part A: Topics based on real case studies

## 1. Performance

**1.1 "Our service has latency spikes every 2 minutes" (Discord: Go → Rust)**
- **Situation:** a fast service shows regular latency spikes, even though the code barely allocates memory.
- **Real case:** Discord's "Read States" service (tracks which messages each user has read, so it's on the hot path) had spikes that looked like garbage collection. Reading Go's source, they found Go forced a GC run at least every 2 minutes. They rewrote the service in Rust, where evicted cache entries are freed immediately.
- **What you teach:** how garbage collectors work, how to find the cause of periodic spikes, and when changing language is (and isn't) justified.
- **Twist:** commenters noted the migration started on Go 1.9, and Go's GC improved a lot afterwards. Was rewriting right, or would upgrading Go have been enough? A good debate for comments.
- **Read:** [Discord: Why Discord is switching from Go to Rust](https://discord.com/blog/why-discord-is-switching-from-go-to-rust)

**1.2 "We turned off garbage collection and got 10% more capacity" (Instagram)**
- **Situation:** your Python web servers use more memory than they should, so each machine runs fewer processes.
- **Real case:** Instagram disabled Python's garbage collector on its web fleet and reported ~10% more efficiency. The cause: worker processes share memory with their parent through copy-on-write, and the GC touching objects forced memory to be copied. Later they found fully disabling GC wasn't a great idea, and added `gc.freeze()` to Python so GC could be mostly turned back on.
- **What you teach:** process forking, copy-on-write memory, how an "odd" fix comes from understanding the runtime.
- **Twist:** the follow-up: why the first fix had to be reversed, and the better fix that went into Python itself.
- **Read:** [Instagram: Dismissing Python Garbage Collection at Instagram](https://instagram-engineering.com/dismissing-python-garbage-collection-at-instagram-4dca40b29172) · [Instagram: Copy-on-write friendly Python garbage collection](https://instagram-engineering.com/copy-on-write-friendly-python-garbage-collection-ad6ed5233ddf) · [PyCon 2018 talk: disable and re-enable GC at Instagram](https://pycon-archive.python.org/2018/schedule/presentation/72/)

**1.3 "A popular cache key expires and the database melts" (Facebook memcache)**
- **Situation:** a heavily read cache key is deleted or expires, and thousands of requests miss at once and hit the database (thundering herd). Also: a slow request writes old data into the cache after a newer one (stale set).
- **Real case:** Facebook's memcache paper introduced **leases**: on a miss, the cache gives one client a token to fetch from the database; others wait or retry. Tokens are handed out at most once every 10 seconds per key, and a delete invalidates the lease so stale writes are rejected. One summary reports peak database queries dropping from 17K/s to 1.3K/s.
- **What you teach:** cache-aside, thundering herd, stale sets, and how one small mechanism fixes both.
- **Read:** "Scaling Memcache at Facebook" (NSDI 2013; search the title for the PDF) · [MIT 6.824 FAQ on the paper](https://pdos.csail.mit.edu/6.824/papers/memcache-faq.txt) · [Engineer's Codex summary](https://read.engineerscodex.com/p/how-facebook-scaled-memcached)

## 2. Reliability and incidents

**2.1 "One firewall rule took down the whole network in minutes" (Cloudflare, 2019)**
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

**3.2 "Move half an exabyte off S3 without losing a file" (Dropbox Magic Pocket)**
- **Situation:** you store user files in a cloud storage service and decide to build your own.
- **Real case:** Dropbox kept metadata in its own data centers but file data in S3. Project "Magic Pocket" moved data to custom storage across three regions; the 90% target was hit in October 2015. During the move, data lived in both places but reads came only from Magic Pocket, so they could switch back if something went wrong.
- **What you teach:** large-scale data migration, verifying data, the "keep a way back" principle.
- **Read:** [InfoQ: Dropbox moves off AWS](https://www.infoq.com/news/2016/03/Dropbox-AWS/) · [InfoQ podcast with Preslav Le](https://www.infoq.com/podcasts/preslav-le-dropbox) · Dropbox's original 2016 announcement (search "Scaling to exabytes and beyond Dropbox")

## 4. Security and authentication

**4.1 "Make 2FA mandatory for millions of users without locking them out" (GitHub)**
- **Situation:** security wants 2FA required for everyone. Support fears thousands of locked-out users.
- **Real case:** GitHub required 2FA for all code contributors, starting in March 2023 with small groups and expanding over the year, after a pilot on npm and months of interviews. Users got 45 days' notice. Results: nearly 95% opt-in, and a one-third drop in 2FA-related support tickets, credited to UX work.
- **What you teach:** rolling out a security requirement like a product launch: cohorts, notice, recovery flows, measuring support load.
- **Read:** [GitHub: Raising the bar: next steps for 2FA](https://github.blog/news-insights/product-news/raising-the-bar-for-software-security-next-steps-for-github-com-2fa/) · [GitHub: Securing millions of developers through 2FA](https://github.blog/security/supply-chain-security/securing-millions-of-developers-through-2fa/)

**4.2 "Our password hashing is weak. Fix it without forcing everyone to reset" (Dropbox + OWASP)**
- **Situation:** you discover passwords are stored with a fast hash.
- **Real case:** Dropbox's 2012 breach exposed credentials, some hashed with SHA-1 and others with bcrypt. Its later design is layered: SHA-512 first (to avoid bcrypt's 72-byte input limit), then bcrypt with a per-user salt, then AES-256 encryption with a separately stored key.
- **What you teach:** fast vs slow hashes, salt and pepper, and the migration trick: wrap old hashes now (`bcrypt(old_hash)`), re-hash to the new scheme at each user's next login.
- **Read:** [Dropbox: How Dropbox securely stores your passwords](https://dropbox.tech/security/how-dropbox-securely-stores-your-passwords) · OWASP Password Storage Cheat Sheet (search the title)

## 5. Frontend, mobile and client apps

**5.1 "Why Airbnb stopped using React Native"**
- **Situation:** a debate video. Should your company build mobile apps with a cross-platform framework?
- **Real case:** in 2018 Airbnb announced it was sunsetting React Native, citing technical and organizational issues; teams had very different experiences, and much of the pain came from the **hybrid** approach (mixing native and React Native screens). Others argued Airbnb's situation was unusual, not proof that cross-platform is bad.
- **What you teach:** cross-platform vs native trade-offs, the hidden cost of hybrid setups, how team structure affects technology choices.
- **Read:** [Airbnb: Sunsetting React Native](https://medium.com/airbnb-engineering/1868ba28e30a) · [Simon Willison's note](https://www.simonwillison.net/2018/Jun/19/sunsetting-react-native-airbnb) · [Expo: Should we use React Native?](https://blog.expo.dev/should-we-use-react-native-1465d8b607ac)

**5.2 "Our mobile website is too slow for users on cheap phones" (Pinterest PWA)**
- **Situation:** growth in countries with slow networks and low-end phones, but the mobile web is heavy.
- **Real case:** Pinterest rebuilt its mobile web as a Progressive Web App: lean JavaScript bundles, service workers for network resilience, dominant-color placeholders for images, app shell, push notifications and caching. A later post reports a large year-over-year increase in mobile web weekly active users in India.
- **What you teach:** web performance basics (bundle size, code splitting, caching), service workers, measuring performance as a business metric.
- **Read:** [Pinterest: A one year PWA retrospective](https://medium.com/pinterest-engineering/a-one-year-pwa-retrospective-f4a2f4129e05) · Addy Osmani: "A Pinterest Progressive Web App Performance Case Study" (search the title) · [Pinterest: Driving user growth with performance improvements](https://medium.com/pinterest-engineering/driving-user-growth-with-performance-improvements-cfc50dafadd7)

**5.3 "Upgrade the language inside an app running on millions of computers" (Dropbox Python 3)**
- **Situation:** you can't redeploy software installed on users' machines, so a bad release can't simply be rolled back on a server.
- **Real case:** Dropbox moved its desktop app from Python 2 to 3, starting at a 2015 hack week. A system called **Hydra** let the app choose the Python 2 or 3 interpreter at startup, so users could be switched back quickly. It rolled out to internal users, then Beta, then everyone.
- **What you teach:** client-side releases, staged rollouts to user groups, remote kill switches, an escape hatch (pattern 6).
- **Read:** [Dropbox: How we rolled out one of the largest Python 3 migrations ever](https://dropbox.tech/application/how-we-rolled-out-one-of-the-largest-python-3-migrations-ever) · [Dropbox: Incrementally migrating over one million lines of code](https://dropbox.tech/application/incrementally-migrating-over-one-million-lines-of-code-from-python-2-to-python-3)

## 6. Real-time and messaging

**6.1 "Our job queue ran out of memory, and everything stopped" (Slack)**
- **Situation:** your async job queue (Redis) is a single point of failure; when it fills up, every feature that enqueues jobs fails.
- **Real case:** Slack's Redis job queue ran out of memory, so new jobs couldn't be enqueued. Slack put **Kafka in front of Redis** as a durable buffer, with a relay service (JQRelay) controlling how fast jobs move from Kafka into Redis, leaving the app's queue logic unchanged. One summary reports over 1.4 billion jobs on busy days, peaking at 33,000/s.
- **What you teach:** queues, back-pressure, durability vs speed, adding a buffer without rewriting the system.
- **Read:** [Slack: Scaling Slack's job queue](https://slack.engineering/scaling-slacks-job-queue/)

**6.2 "Five million people connected at once" (Discord + Elixir)**
- **Situation:** a real-time chat platform where some groups have tens of thousands of users online in the same room.
- **Real case:** Discord built its real-time system on Elixir/Erlang VM. It reached nearly 5 million concurrent users; big servers (e.g. one with ~30,000 concurrent users) made processes fall behind on their message queues at peak, and bursts of reconnecting users were another problem. Later they used Rust (via Rustler) to scale further.
- **What you teach:** fan-out of messages to many connections, actor/process model, handling reconnect storms.
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
- **Read:** [Stripe: APIs as infrastructure: future-proofing Stripe with versioning](https://stripe.com/ae/blog/api-versioning) · [Stripe docs: API versioning](https://docs.stripe.com/api/versioning?lang=curl)

**7.2 "Some customers are flooding our API, and real payments are failing" (Stripe rate limiters)**
- **Situation:** one client's script hammers your API, and important traffic suffers.
- **Real case:** Stripe runs four types of limiters in production: two rate limiters and two **load shedders**. One load shedder reserves part of the fleet for the most important requests; another limits damage during incidents already happening. Advice includes clear 429/503 errors for users and switches to turn limiters off.
- **What you teach:** rate limiting vs load shedding, prioritizing critical traffic, designing what clients see when you say "no".
- **Read:** [Stripe: Scaling your API with rate limiters](https://stripe.com/blog/rate-limiters) · [Stripe docs: rate limits](https://docs.stripe.com/rate-limits)

**7.3 "Replace the API behind a mobile app with no downtime" (Netflix → GraphQL)**
- **Situation:** you want to replace the API layer your mobile app depends on, with millions of active users.
- **Real case:** Netflix moved its mobile apps from its internal Falcor API to GraphQL with zero downtime. Phase 1: a GraphQL shim on top of Falcor, validated with A/B tests. Phase 2: a new implementation, validated with **replay testing** (compare outputs of old and new for read-only queries) and **sticky canaries** (route some users persistently to the new cluster, which also works for writes).
- **What you teach:** pattern 5 (shadow comparison) for APIs, canaries, why read-only and write traffic need different testing.
- **Read:** [The New Stack: Netflix's testing strategies for migrating to GraphQL](https://thenewstack.io/netflixs-testing-strategies-for-migrating-to-graphql/) · Netflix Tech Blog: "Migrating Netflix to GraphQL Safely" (search the title)

## 8. Architecture

**8.1 "Rewrite the backend in a new language, without a big-bang rewrite" (Khan Academy)**
- **Situation:** your language version is dying (Python 2 end of life, Jan 2020), and you decide to move to new services.
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
- **(2 days)**: or split it into "why teams split" and "why teams merge back".

## 9. Codebase and team

**9.1 "Upgrade a framework three major versions behind, without stopping feature work" (GitHub Rails)**
- **Situation:** you're stuck on an old framework version, and the business won't pause features.
- **Real case:** GitHub took a year and a half to go from Rails 3.2 to 5.2. The key technique was **dual booting**: the app could run on either version and CI tested both. They shipped once production ran 30 minutes at peak traffic with no visible impact. Their own fork of Rails had made it harder.
- **What you teach:** compatibility with both versions, small steps on the main branch, staying close to upstream.
- **Read:** [GitHub: Upgrading GitHub from Rails 3.2 to 5.2](https://github.blog/engineering/upgrading-github-from-rails-3-2-to-5-2/)

**9.2 "Big bang or step by step? Moving millions of lines to TypeScript" (Stripe vs Airbnb)**
- **Situation:** a debate video. Your large JavaScript codebase needs a type system.
- **Real cases:** **Stripe** converted 3.7M+ lines from Flow to TypeScript **in a single pull request** (merged March 2022) using a codemod. **Airbnb** built **ts-migrate** to convert projects automatically, reporting 86% of a ~6M-line frontend converted.
- **What you teach:** codemods, when big bang is *safer* than incremental, and what preparation makes it possible.
- **Twist:** Pinterest also reported a 3.7M-line migration. Don't mix up the numbers.
- **Read:** [Stripe: Migrating millions of lines of code to TypeScript](https://stripe.com/blog/migrating-to-typescript) · [airbnb/ts-migrate](https://github.com/airbnb/ts-migrate) · [InfoQ: Airbnb TypeScript migration](https://www.infoq.com/news/2020/08/airbnb-typescript-migration)

**9.3 "One repo or many?" (Google's monorepo)**
- **Situation:** a debate video. Your company has 40 repositories and cross-repo changes are painful. Should you merge them?
- **Real case:** Google's 2016 paper describes one repository with about two billion lines of code in nine million files (2015 numbers), about 15 million lines changed per week in 2014, hierarchical code ownership, and a custom version control system (Piper). The paper discusses both benefits and costs.
- **What you teach:** monorepo vs polyrepo trade-offs: atomic changes, shared tooling, code ownership, and the tooling investment a monorepo needs.
- **Read:** [Google Research: Why Google stores billions of lines of code in a single repository](https://research.google/pubs/why-google-stores-billions-of-lines-of-code-in-a-single-repository/) · [ACM full text](https://cacmb4.acm.org/magazines/2016/7/204032-why-google-stores-billions-of-lines-of-code-in-a-single-repository/fulltext)

**9.4 "We have 2,000 old feature flags nobody dares to remove" (Uber Piranha)**
- **Situation:** feature flags made releases safe, but now the code is full of dead branches.
- **Real case:** Uber built Piranha, which analyzes code to find flags that are fully rolled out, generates a code change removing them, and sends it to the flag's author for review. It runs periodically on Uber's Android and iOS code and removed around two thousand stale flags. Hard parts: deciding what's truly stale, owners who left, and flags intentionally kept as kill switches.
- **What you teach:** feature flag lifecycle, technical debt, automated refactoring.
- **Read:** [Uber: Introducing Piranha](https://www.uber.com/blog/piranha/)

## 10. Data and databases

**10.1 "Split a huge table into a new one while the app is live" (Stripe)**
- **Real case:** Stripe moved hundreds of millions of Subscriptions out of Customer records with a four-step dual-writing pattern, using the Scientist library to compare reads from both.
- **What you teach:** pattern 2 in full.
- **Read:** [Stripe: Online migrations at scale](https://stripe.com/se/blog/online-migrations) · [Gusto: splitting a large table](https://engineering.gusto.com/old-write)

**10.2 "Our database is too big for one machine. Shard it without a rewrite" (Figma, Notion)**
- **Real cases:** **Figma** went from one Postgres on AWS's largest instance to sharding, doing **logical sharding first** (routing as if sharded while data stays in one database) and physical sharding later, through a custom proxy. **Notion** sharded by workspace ID, the block table first, using Postgres logical replication.
- **What you teach:** shard keys, why "change the routing before moving the data" reduces risk.
- **Read:** [Figma: How Figma's databases team lived to tell the scale](https://www.figma.com/en-jp/blog/how-figmas-databases-team-lived-to-tell-the-scale/) · [Notion: Herding elephants](https://notion.com/blog/sharding-postgres-at-notion)

**10.3 "Move trillions of messages to a new database with zero downtime" (Discord)**
- **Real case:** Discord went MongoDB → Cassandra → ScyllaDB (177 Cassandra nodes became 72 ScyllaDB nodes). ScyllaDB speaks Cassandra's protocol, so app logic didn't need rewriting. They used dual writes plus a bulk migrator.
- **What you teach:** hot partitions (a data-modeling problem), dual writes + bulk copy.
- **Read:** [Discord: How Discord stores trillions of messages](https://discord.com/blog/how-discord-stores-trillions-of-messages)

**10.4 "Change the schema of a huge table without locking it" (GitHub gh-ost)**
- **Real case:** GitHub's gh-ost copies data to a "ghost" table and follows ongoing changes through MySQL's binary log instead of triggers; it can pause and throttle on load and replication lag.
- **What you teach:** shadow table + catch up + atomic swap.
- **Read:** [GitHub: gh-ost announcement](https://github.blog/2016-08-01-gh-ost-github-s-online-migration-tool-for-mysql) · [gh-ost README](https://github.com/github/gh-ost)

**More database case studies, if you want extra:** [Uber: Postgres to MySQL](https://www.uber.com/blog/postgres-to-mysql-migration/) · [Slack: Vitess](https://slack.engineering/scaling-datastores-at-slack-with-vitess/) · [GitHub: MySQL 8.0 upgrade](https://github.blog/engineering/infrastructure/upgrading-github-com-to-mysql-8-0/) · [Shopify: shard balancing](https://shopify.engineering/blogs/engineering/mysql-database-shard-balancing-terabyte-scale) · [Wix: DB migration service](https://www.wix.engineering/post/how-we-built-a-zero-downtime-database-migration-service-at-wix) · [InfoQ: Netflix data migrations](https://www.infoq.com/news/2018/11/netflix-data-migrations)

---

# Part B: Everyday scenarios (no single famous case study)

Common tasks in normal jobs. No one big write-up, but the solution is standard and documented, so they're still 1-day topics. Use the reference to check your plan.

| # | Area | Scenario | The solution in one line | Reference to check |
|---|---|---|---|---|
| B1 | Performance | **An endpoint got 10× slower after a release** | Compare before/after → profile → flame graph → find the hot path → fix and add a performance test | Brendan Gregg's website (flame graphs, USE method) |
| B2 | Performance | **Memory keeps growing until the server restarts** | Watch memory over time → heap dumps at two moments → compare → find what keeps growing (caches, listeners, globals) | Your language's profiler docs |
| B3 | Reliability | **On-call gets 200 alerts a week** | Alert on user-facing symptoms, not causes → define SLOs → delete or downgrade noisy alerts | Google SRE book (sre.google/books): "Monitoring Distributed Systems", "Service Level Objectives" |
| B4 | Frontend | **The web page takes 8 seconds on a phone** | Measure (Core Web Vitals) → shrink the bundle → code splitting / lazy loading → optimize images → cache | web.dev |
| B5 | Mobile | **A crash shipped to the app store and you can't roll back** | Staged rollout → remote config / kill switch → hotfix release → crash monitoring | Google Play staged rollouts docs; Apple phased release docs |
| B6 | Security | **Add "Login with Google" / SSO to an existing app** | OpenID Connect → link to existing accounts by verified email → keep password login during transition | OpenID Connect docs; OWASP cheat sheets |
| B7 | Security | **An API key was pushed to GitHub** | Rotate first → check logs → clean history → secret manager + secret scanning | GitHub docs: "Removing sensitive data from a repository" |
| B8 | Security / data | **"Delete my data" across 30 services** | Data inventory → `UserDeleted` event → per-service tracking → crypto-shredding for backups | Search "crypto shredding" |
| B9 | Architecture | **The report button times out** | Async job + job ID + polling/notification + file in object storage | Microsoft Azure Architecture Center: "Asynchronous Request-Reply pattern" |
| B10 | Architecture | **Switching payment providers** | Your own `PaymentProvider` interface → new provider behind a flag → old payments finish on the old one | Martin Fowler: "BranchByAbstraction" |
| B11 | Codebase | **Changing code that has no tests** | Characterization tests → small refactors → shadow-run old vs new | Michael Feathers, *Working Effectively with Legacy Code*; GitHub's Scientist library |
| B12 | Codebase | **CI takes 45 minutes** | Measure → cache → parallelize → run affected tests only → quarantine flaky tests | Google Testing Blog posts on flaky tests |
| B13 | Data | **Money stored as float** | Integer minor units (paisa) or `DECIMAL`; dual write; documented rounding; reconcile totals | Martin Fowler: "Money" pattern |
| B14 | Data | **All timestamps stored in local time** | Classify columns (instant vs date vs wall time) → UTC for instants → convert only in the UI | Jon Skeet: "Storing UTC is not a silver bullet" |
| B15 | Data / performance | **A page runs 1,200 queries (N+1)** | Measure queries per request → eager/batch loading → test that caps query count | Your ORM's eager-loading docs |

---

## Where to find more topics

**Collections and newsletters (to discover topics)**
- [Pragmatic Engineer: Real-World Engineering Challenges](https://newsletter.pragmaticengineer.com/p/real-world-4): real engineering stories across many areas. Issue #6 covers migrations; #8 covers Khan Academy. Some content is paid.
- [kilimchoi/engineering-blogs](https://github.com/kilimchoi/engineering-blogs): 400+ company engineering blogs. Not updated recently, so some links may be dead.
- [Simon Willison's link blog](https://simonwillison.net/): links and summarizes many of the best engineering posts.
- **InfoQ** (infoq.com): conference talks and news summaries across all areas.
- **Hacker News** (search old discussions at hn.algolia.com): the comments add counter-arguments and twists.
- **Papers:** classic industry papers (Facebook memcache, Google monorepo, Amazon Dynamo) are short enough for a 2-day prep and make strong videos.

**Where to search by area**
- **Performance:** Discord, Instagram, Netflix, Brendan Gregg's blog.
- **Reliability:** Cloudflare blog (detailed postmortems), Slack Engineering, Google SRE book, Surfing Complexity (incident analysis).
- **Frontend and mobile:** Airbnb, Pinterest, Shopify, web.dev case studies.
- **Real-time:** Discord, Slack, Figma.
- **APIs:** Stripe blog and docs.
- **Code and team:** GitHub, Uber, Google, Stripe.

**Patterns and books**
- [Martin Fowler: Parallel Change](https://www.martinfowler.com/bliki/ParallelChange.html), Strangler Fig, Branch by Abstraction, Feature Toggles.
- *Designing Data-Intensive Applications* (Kleppmann), *Release It!* (Nygard), *Site Reliability Engineering* (Google, free online), *Working Effectively with Legacy Code* (Feathers).

**Summaries vs originals:** ByteByteGo, Medium, dev.to and YouTube summaries are good for *finding* topics, but some are AI-generated or wrong. Take your facts from the company's original post.

**Your best source is still your own experience:** every performance bug, outage or tricky change from your past jobs is a video only you can make, and an interview answer.

---

## Sources used for this list

- **Performance:** [Discord: Go to Rust](https://discord.com/blog/why-discord-is-switching-from-go-to-rust) · [Instagram: dismissing Python GC](https://instagram-engineering.com/dismissing-python-garbage-collection-at-instagram-4dca40b29172) · [Instagram: copy-on-write friendly GC](https://instagram-engineering.com/copy-on-write-friendly-python-garbage-collection-ad6ed5233ddf) · [PyCon 2018 talk](https://pycon-archive.python.org/2018/schedule/presentation/72/) · [MIT 6.824 memcache FAQ](https://pdos.csail.mit.edu/6.824/papers/memcache-faq.txt) · [Engineer's Codex: Facebook memcached](https://read.engineerscodex.com/p/how-facebook-scaled-memcached)
- **Reliability:** [InfoQ: Cloudflare WAF outage](https://www.infoq.com/news/2019/07/waf-causes-outage/) · [Surfing Complexity: Cloudflare](https://surfingcomplexity.blog/2019/08/02/contributors-mitigators-risks-cloudflare-2019-07-02-outage/) · [Slack: Jan 4 2021 outage](https://slack.engineering/slacks-outage-on-january-4th-2021/) · [The Register: Slack outage](https://www.theregister.com/2021/02/02/slack_outage_aws_autoscaling/) · [Surfing Complexity: Slack](https://surfingcomplexity.blog/2021/02/08/slacks-jan-2021-outage-a-tale-of-saturation/) · [Gremlin: origin of Chaos Monkey](https://gremlin.com/chaos-monkey/the-origin-of-chaos-monkey) · [IEEE Spectrum: chaos engineering](https://spectrum.ieee.org/chaos-engineering-saved-your-netflix)
- **Infrastructure:** [DHH: cloud-exit savings](https://world.hey.com/dhh/our-cloud-exit-savings-will-now-top-ten-million-over-five-years-c7d9b5bd) · [The Register (2024)](https://www.theregister.com/2024/10/21/37signals_aws_savings) · [The Register (2025)](https://www.theregister.com/2025/05/09/37signals_cloud_repatriation_storage_savings) · [InfoQ: Dropbox off AWS](https://www.infoq.com/news/2016/03/Dropbox-AWS/) · [InfoQ podcast: Preslav Le](https://www.infoq.com/podcasts/preslav-le-dropbox)
- **Security:** [GitHub: 2FA next steps](https://github.blog/news-insights/product-news/raising-the-bar-for-software-security-next-steps-for-github-com-2fa/) · [GitHub: securing millions of developers through 2FA](https://github.blog/security/supply-chain-security/securing-millions-of-developers-through-2fa/) · [Dropbox: password storage](https://dropbox.tech/security/how-dropbox-securely-stores-your-passwords)
- **Frontend and mobile:** [Airbnb: Sunsetting React Native](https://medium.com/airbnb-engineering/1868ba28e30a) · [Simon Willison on Airbnb RN](https://www.simonwillison.net/2018/Jun/19/sunsetting-react-native-airbnb) · [Expo: Should we use React Native?](https://blog.expo.dev/should-we-use-react-native-1465d8b607ac) · [Pinterest: one year PWA retrospective](https://medium.com/pinterest-engineering/a-one-year-pwa-retrospective-f4a2f4129e05) · [Pinterest: performance and growth](https://medium.com/pinterest-engineering/driving-user-growth-with-performance-improvements-cfc50dafadd7) · [Dropbox: Python 3 rollout](https://dropbox.tech/application/how-we-rolled-out-one-of-the-largest-python-3-migrations-ever) · [Dropbox: incremental Python 3](https://dropbox.tech/application/incrementally-migrating-over-one-million-lines-of-code-from-python-2-to-python-3)
- **Real-time:** [Slack: job queue](https://slack.engineering/scaling-slacks-job-queue/) · [Discord: Elixir 5M users](https://discord.com/blog/how-discord-scaled-elixir-to-5-000-000-concurrent-users) · [Elixir blog: Discord](https://elixir-lang.org/blog/2020/10/08/real-time-communication-at-scale-with-elixir-at-discord/) · [Figma multiplayer (archived)](https://madebyevan.com/figma/how-figmas-multiplayer-technology-works/) · [Liveblocks: sync engines](https://www.liveblocks.io/blog/understanding-sync-engines-how-figma-linear-and-google-docs-work)
- **APIs:** [Stripe: API versioning](https://stripe.com/ae/blog/api-versioning) · [Stripe docs: versioning](https://docs.stripe.com/api/versioning?lang=curl) · [Stripe: rate limiters](https://stripe.com/blog/rate-limiters) · [Stripe docs: rate limits](https://docs.stripe.com/rate-limits) · [The New Stack: Netflix GraphQL](https://thenewstack.io/netflixs-testing-strategies-for-migrating-to-graphql/)
- **Architecture:** [Khan Academy: Goliath](https://blog.khanacademy.org/go-services-one-goliath-project) · [Incremental rewrites with GraphQL](https://blog.khanacademy.org/incremental-rewrites-with-graphql) · [Beating the odds](https://blog.khanacademy.org/beating-the-odds-khan-academys-successful-monolith%e2%86%92services-rewrite) · [Pragmatic Engineer #8](https://newsletter.pragmaticengineer.com/p/real-world-eng-8) · [InfoQ: Airbnb SOA](https://www.infoq.com/news/2019/02/airbnb-monolith-migration-soa/) · [Segment: Goodbye Microservices](https://twilio.com/en-us/blog/developers/best-practices/goodbye-microservices) · [InfoQ: Prime Video](https://infoq.com/news/2023/05/prime-ec2-ecs-saves-costs) · [Shopify: Deconstructing the Monolith](https://shopify.engineering/deconstructing-monolith-designing-software-maximizes-developer-productivity)
- **Codebase and team:** [GitHub: Rails upgrade](https://github.blog/engineering/upgrading-github-from-rails-3-2-to-5-2/) · [Stripe: TypeScript](https://stripe.com/blog/migrating-to-typescript) · [airbnb/ts-migrate](https://github.com/airbnb/ts-migrate) · [InfoQ: Airbnb TypeScript](https://www.infoq.com/news/2020/08/airbnb-typescript-migration) · [Google: monorepo paper](https://research.google/pubs/why-google-stores-billions-of-lines-of-code-in-a-single-repository/) · [Uber: Piranha](https://www.uber.com/blog/piranha/)
- **Data:** [Stripe: online migrations](https://stripe.com/se/blog/online-migrations) · [Gusto: old write](https://engineering.gusto.com/old-write) · [Figma: databases team](https://www.figma.com/en-jp/blog/how-figmas-databases-team-lived-to-tell-the-scale/) · [Notion: Herding elephants](https://notion.com/blog/sharding-postgres-at-notion) · [Discord: trillions of messages](https://discord.com/blog/how-discord-stores-trillions-of-messages) · [GitHub: gh-ost](https://github.blog/2016-08-01-gh-ost-github-s-online-migration-tool-for-mysql) · [gh-ost repo](https://github.com/github/gh-ost)
- **General:** [Pragmatic Engineer: Real-World Engineering Challenges](https://newsletter.pragmaticengineer.com/p/real-world-4) · [Martin Fowler: Parallel Change](https://www.martinfowler.com/bliki/ParallelChange.html) · [Wikipedia: Strangler fig pattern](https://en.wikipedia.org/wiki/Strangler_fig_pattern) · [kilimchoi/engineering-blogs](https://github.com/kilimchoi/engineering-blogs)
