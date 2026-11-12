# Engineering Stories: Quick-Prep Video Topics

*Real software problems, bugs and outages that are fun to explain and fast to prepare · Nov 12, 2026 · @Lion*

Each topic is a **true story** with a clear "aha" moment and one core concept. Most take **1 day** to prepare because a detailed write-up (a postmortem, paper or well-known article) already exists. Topics marked **(2 days)** need more background reading.

> **Rule:** Every fact, number and date in your script must come from the primary source listed (the company's own postmortem or official report), not from memory, this list or a YouTube video. The list tells you where to look; the source tells you what to say.

---

## The 1-day prep routine

| Time | Task |
|---|---|
| Morning (2-3 h) | Read the primary source fully. Then read one explainer (a blog post or HN discussion) for the bits you didn't understand. |
| Midday (1 h) | Answer 4 questions on paper: **What happened? Why did it happen (the concept)? How was it fixed? Where could this happen in my code?** |
| Afternoon (2 h) | Draw 3-5 whiteboard panels. Write the hook and a bullet-point script (not word for word). |
| Evening (1-2 h) | Record. Light edit: cut mistakes and long pauses, nothing fancy. |

## The video shape (6-10 minutes)

1. **Hook (20 s):** the most surprising fact. *"A music video broke YouTube's math."*
2. **The story (1-2 min):** what happened, when, and the damage.
3. **The concept (3-4 min):** the engineering idea behind it, on the whiteboard. This is the main part.
4. **The fix (1 min):** what the engineers did.
5. **The lesson for you (1 min):** where this exact bug can hide in everyday code, or how it shows up in interviews.
6. **Question for comments (10 s):** *"Have you ever hit a bug like this?"*

Each video also gives you **2 LinkedIn posts**: the story as a question ("Why did a 2014 music video force YouTube to change its database?") and your explanation, linking the video.

---

## 1. Numbers that break things

**1. Gangnam Style broke YouTube's view counter (2014)**
- Hook: a K-pop video passed 2,147,483,647 views, the maximum of a 32-bit signed integer.
- Concept: integer types, overflow, why 2³¹ − 1.
- Source: YouTube's official Google+ announcement (Dec 2014), widely quoted in news coverage; the Wikipedia article on the video covers the incident.

**2. Planes and PCs that must be rebooted (Boeing 787, Windows 95)**
- Hook: a Boeing 787 had to be power-cycled at least every 248 days; Windows 95 hung after 49.7 days.
- Concept: counters overflowing. 2³¹ hundredths of a second ≈ 248 days; 2³² milliseconds ≈ 49.7 days. Pair it with episode 1 or make it a short.
- Source: FAA Airworthiness Directive 2015-09-07 (787); Microsoft support article KB216641 (Windows 95/98).

**3. The rocket that exploded over one number conversion (Ariane 5, 1996)**
- Hook: a 64-bit floating point number was converted to a 16-bit integer, and a rocket was destroyed about 37 seconds after launch.
- Concept: type conversion, overflow, reusing code in a new context.
- Source: *Ariane 5 Flight 501 Failure: Report by the Inquiry Board* (ESA, 1996), short and readable.

**4. Why 0.1 + 0.2 ≠ 0.3, and the missile clock that drifted (Patriot, 1991)**
- Hook: 0.1 can't be stored exactly in binary. After 100 hours of uptime, that tiny error had grown to about a third of a second.
- Concept: floating point representation, accumulated error, why money is never stored as float.
- Source: US GAO report IMTEC-92-26, *Patriot Missile Defense: Software Problem Led to System Failure at Dhahran*. Treat it respectfully: 28 soldiers died.

**5. The Pentium chip that couldn't divide (1994)**
- Hook: Intel's processor gave wrong answers for some divisions, and the recall cost about $475 million.
- Concept: lookup tables, how a professor found it, testing edge cases.
- Source: Thomas Nicely's original reports; Intel's statements; the Wikipedia article "Pentium FDIV bug" links to both.

**6. The Year 2038 problem**
- Hook: on 19 January 2038, 32-bit Unix time runs out.
- Concept: Unix timestamps, signed 32-bit integers, why Y2K was not a one-off.
- Source: the Wikipedia article "Year 2038 problem"; Linux kernel discussions on 64-bit time_t.

## 2. Time is hard

**7. The leap second that crashed servers (2012) and Cloudflare's DNS (2017)**
- Hook: one extra second in a minute made Linux servers spin at 100% CPU, and years later produced a negative duration in Cloudflare's code.
- Concept: wall clock vs monotonic clock; never measure durations with the wall clock.
- Source: Cloudflare blog "How and why the leap second affected Cloudflare DNS" (Jan 2017).

**8. Falsehoods programmers believe about time and names**
- Hook: "every day has 24 hours", "everyone has a first and last name". Both are false.
- Concept: time zones, DST, names, addresses: assumptions that break real software. Easy to make Bangladesh-specific (single names, Bangla script).
- Source: Patrick McKenzie, "Falsehoods Programmers Believe About Names" (2010); GitHub repo `kdeldycke/awesome-falsehood` collects the whole series.

## 3. One small change took down the internet

**9. Knight Capital lost $440 million in 45 minutes (2012)**
- Hook: a deployment missed 1 of 8 servers, a reused feature flag woke up old dead code, and the company was nearly bankrupt by lunch.
- Concept: deployment consistency, dead code, feature flag hygiene.
- Source: US SEC order against Knight Capital Americas (Oct 2013), which describes the technical failure in detail.

**10. A regular expression took down Cloudflare (2019)**
- Hook: one regex in a firewall rule pushed CPUs to 100% across Cloudflare's whole network.
- Concept: regex backtracking, ReDoS, staged rollouts.
- Source: Cloudflare blog "Details of the Cloudflare outage on July 2, 2019", one of the best postmortems ever written.
- Bonus: Stack Overflow's 2016 outage had the same cause (a whitespace-trimming regex). Source: Stack Overflow's own postmortem (July 20, 2016).

**11. A typo took down Amazon S3 (2017)**
- Hook: an engineer typed a command to remove a few servers and removed far more, taking down a huge part of the web.
- Concept: blast radius, safe tooling, guardrails.
- Source: AWS "Summary of the Amazon S3 Service Disruption in the Northern Virginia (US-EAST-1) Region" (Feb 28, 2017).

**12. Facebook disappeared from the internet for 6 hours (2021)**
- Hook: a configuration change withdrew Facebook's routes from the internet, and engineers reportedly had trouble even getting into the buildings.
- Concept: BGP and DNS basics: how the internet finds a website.
- Source: Meta engineering blog "More details about the October 4 outage"; Cloudflare's blog "Understanding how Facebook disappeared from the Internet". **(2 days if BGP is new to you)**

**13. GitLab deleted its production database, and its backups didn't work (2017)**
- Hook: an engineer ran a delete on the wrong server, then the team found that its backup methods had been failing. They livestreamed the recovery.
- Concept: untested backups are not backups; replication isn't backup.
- Source: GitLab's own postmortem "Postmortem of database outage of January 31".

**14. 43 seconds of network trouble, 24 hours of degraded GitHub (2018)**
- Hook: a short network cut between data centers caused databases to fail over across the US, and the cleanup took a full day.
- Concept: failover, network partitions, split brain (a preview of CAP).
- Source: GitHub blog "October 21 post-incident analysis". **(2 days)**

**15. The update that crashed 8.5 million Windows PCs (CrowdStrike, 2024)**
- Hook: airports, banks and hospitals showed blue screens worldwide because of one content update.
- Concept: kernel-level software, out-of-bounds memory reads, why staged rollouts matter.
- Source: CrowdStrike's Preliminary Post Incident Review and its Root Cause Analysis (2024).

**16. 11 lines of code broke the JavaScript world (left-pad, 2016)**
- Hook: a developer unpublished a tiny package, and builds broke across the ecosystem, including big projects.
- Concept: dependencies, package registries, lock files and vendoring.
- Source: the author's blog post "I've Just Liberated My Modules" and npm's blog post on the incident.

## 4. Security stories

**17. The xz backdoor caught because SSH was half a second slow (2024)**
- Hook: an engineer noticed SSH logins were ~500 ms slower and uncovered a backdoor that someone had spent about two years planting in an open-source project.
- Concept: supply-chain attacks, trust in open source, noticing small anomalies.
- Source: Andres Freund's original message to the oss-security mailing list (March 29, 2024).

**18. Heartbleed: asking a server for 64 KB it shouldn't give you (2014)**
- Hook: "Send me the 5-letter word 'hello', which is 500 letters long", and the server sent back secrets from its memory.
- Concept: buffer over-read, bounds checking, trusting length fields.
- Source: heartbleed.com; the xkcd comic #1354 explains it in one panel (great for the video).

**19. Log4Shell: a log message that runs code (2021)**
- Hook: typing a special string into a chat box or username could take over a server.
- Concept: input that gets interpreted, JNDI lookups, why logging untrusted input is dangerous.
- Source: Apache Log4j security page (CVE-2021-44228); LunaSec's original write-up.

**20. Hash flooding: crashing web servers with a carefully chosen form (2011)**
- Hook: one HTTP request with thousands of crafted parameters could keep a CPU busy for a long time.
- Concept: hash tables, collisions, O(1) average vs O(n) worst case, randomized hashing.
- Source: Klink & Wälde's talk at 28C3, "Efficient Denial of Service Attacks on Web Application Platforms" (2011). Great for interview-style data structures content.

## 5. Concurrency and systems

**21. The bug that kept resetting a robot on Mars (Pathfinder, 1997)**
- Hook: the rover kept rebooting on Mars, and engineers fixed it from Earth by switching on one setting.
- Concept: priority inversion, mutexes, priority inheritance.
- Source: Glenn Reeves's account "What really happened on Mars?" (JPL); Mike Jones's summary page of the same title.

**22. Therac-25: a race condition in a medical machine (1985-87)**
- Hook: typing too fast on a radiation therapy machine's keyboard could trigger a massive overdose.
- Concept: race conditions, removing hardware safety locks, why safety-critical software is different.
- Source: Nancy Leveson & Clark Turner, "An Investigation of the Therac-25 Accidents" (IEEE Computer, 1993). Treat it with care; people died. **(2 days)**

**23. Mars Climate Orbiter: pounds vs newtons (1999)**
- Hook: one team used imperial units, another used metric, and a spacecraft was lost.
- Concept: API contracts, units in types, integration testing.
- Source: NASA's *Mars Climate Orbiter Mishap Investigation Board Phase I Report* (1999).

## 6. Clever solutions and fun stories

**24. The 500-mile email**
- Hook: "We can't send email more than 500 miles." It was true, and the reason is the speed of light.
- Concept: timeouts, network latency, debugging by measurement. Pure fun and very shareable.
- Source: Trey Harris's original post (search "the case of the 500-mile email").

**25. The fast inverse square root (Quake III, 0x5f3759df)**
- Hook: a mysterious hex constant made 3D games fast in 1999.
- Concept: floating point bit layout, Newton's method, performance tricks.
- Source: Quake III Arena source code; Chris Lomont's paper "Fast Inverse Square Root". **(2 days for the math)**

**26. SHAttered: two different PDFs with the same SHA-1 (2017)**
- Hook: Google produced two different files with the same "fingerprint".
- Concept: hash functions, collisions, why Git and browsers moved away from SHA-1.
- Source: shattered.io (Google & CWI Amsterdam).

**27. The text message that crashed iPhones ("Effective Power", 2015)**
- Hook: receiving one specific string of text crashed a phone.
- Concept: Unicode, text rendering, why strings are not simple.
- Source: news coverage and Apple's iOS 8.4 security notes; pair it with Joel Spolsky's "The Absolute Minimum Every Software Developer Must Know About Unicode".

## 7. Architecture decisions (great for debates)

**28. Prime Video moved from microservices back to a monolith and cut costs by 90% (2023)**
- Hook: Amazon's own team said "we went back to a monolith".
- Concept: microservices vs monolith, when each makes sense. Comments will argue, which is good for reach.
- Source: Prime Video Tech blog "Scaling up the Prime Video audio/video monitoring service and reducing costs by 90%" (2023).

**29. Why Discord switched a service from Go to Rust (2020)**
- Hook: performance spikes every two minutes, caused by garbage collection.
- Concept: garbage collection, latency spikes, choosing languages for a workload.
- Source: Discord blog "Why Discord is switching from Go to Rust" (Feb 2020).

**30. How Stack Overflow serves millions of users with a handful of servers**
- Hook: one of the biggest sites in the world runs on far fewer servers than you'd guess.
- Concept: vertical scaling, caching, "you probably don't need microservices".
- Source: Nick Craver's blog series "Stack Overflow: The Architecture" (2016 edition).

---

## Where these topics come from, and how to find more

**Where this list came from:** I didn't run a live search for it. These are well-known stories from my training knowledge, the ones engineers retell and link to most. They keep showing up in the places below, and that's also where *you* will find new ones. Before recording, always check details against the primary source.

**Primary sources (where you get your facts)**
- **Company engineering blogs and postmortems:** Cloudflare blog (the best for detailed outage write-ups), GitHub blog, GitLab, Meta Engineering, Discord, Netflix Tech Blog, Uber Engineering, Stripe, Dropbox, Slack.
- **Cloud provider incident reports:** AWS Post-Event Summaries page; Google Cloud status page incident reports; Azure status history.
- **Official investigation reports:** NASA, ESA, GAO and FAA reports for the historic bugs. Long but authoritative; read the summary and the "root cause" section.

**Collections (where you find new topics)**
- **GitHub `danluu/post-mortems`:** a big categorized list of public postmortems. Probably the best single topic mine.
- **k8s.af:** Kubernetes failure stories.
- **GitHub `kdeldycke/awesome-falsehood`:** the "falsehoods programmers believe" series.
- **Wikipedia "List of software bugs":** historic bugs with links to sources.
- **The RISKS Digest** (catless.ncl.ac.uk/Risks): decades of computer failure stories.

**Ongoing feeds (to keep your topic list growing)**
- **Hacker News** (and hn.algolia.com to search old discussions; the comments often have insiders explaining the details).
- **SRE Weekly** newsletter: a weekly list of incidents and postmortems.
- **The Pragmatic Engineer** newsletter by Gergely Orosz: industry stories and deep dives.
- **Reddit:** r/programming, r/ExperiencedDevs, r/sysadmin (sysadmin war stories).

**Books with many short stories**
- *Release It!* by Michael Nygard: each chapter opens with a real production failure.
- *Site Reliability Engineering* (Google, free online at sre.google/books): incident management and lessons.

**Watch, but don't copy**
- **Kevin Fang** on YouTube makes outage-story videos in this style. Watch a few to see what works, then make yours different: your own whiteboard explanations, the "where this bug hides in your code / in interviews" angle, and local examples for your audience.

**Keep a topic inbox:** whenever you read about an interesting bug or outage, add one line to a note: title, link, one-line hook. Pick from it each week, so you never start from zero.
