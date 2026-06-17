# System Design Lab

A platform for learning system design through real decisions. Instead of drawing diagrams, you choose — and every choice has consequences.

## The Problem

Junior developers don't know where to start when designing systems. Existing tools only offer free-form diagramming with no feedback and no explanation of why a decision is right or wrong.

## The Solution

System Design Lab places you in real-world scenarios (Chat App, URL Shortener, E-commerce...) and guides you through a structured decision flow:

```
Question → Choice → Feedback + Metrics → Next Question
```

Every decision directly impacts system metrics. Wrong choice → system degrades. Right choice → metrics improve.

## Features

**Gameplay**
- Scenario-based learning: each challenge is a real system with 5–7 decision steps
- Metrics update in real-time with each choice (latency, cost, scalability)
- Instant feedback after every step
- Explanations at 3 levels: Beginner / Intermediate / Advanced
- Hints available for extra context when needed

**Sessions**
- Normal mode: learn at your own pace
- Interview mode: time-limited, simulates a real system design interview
- Abandon or retry any quest at any time

**Progress & Leaderboard**
- Track progress across all scenarios
- Top 10 leaderboard per scenario
- Score history and completion stats

## Tech Stack

| Layer | Stack |
|-------|-------|
| Backend | Go, PostgreSQL, sqlc, pgx/v5 |
| Frontend | React, TypeScript, Vite, TanStack Query |

## Repository Structure

```
system-design-lab/
├── system-design-lab-api/   # Go backend
└── system-design-lab-fe/    # React frontend
```

See setup instructions in each directory:
- [Backend →](./system-design-lab-api/README.md)
- [Frontend →](./system-design-lab-fe/README.md)

## Roadmap

- [ ] AI Coach — asks follow-up questions and explains decisions in depth
- [ ] Visual Architecture — diagram that updates with each decision
- [ ] Multiplayer — team learning and decision comparison
- [ ] Advanced Interview Mode — timed pressure scenarios
