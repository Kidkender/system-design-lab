# system-design-lab-fe

React frontend for [System Design Lab](../README.md).

## Requirements

- Node.js 18+
- npm or pnpm
- Backend running at `http://localhost:8080`

## Quick Start

```bash
npm install
npm run dev
```

App runs at `http://localhost:5173`.

The Vite proxy forwards `/api` to `http://localhost:8080` automatically — no extra config needed for local development.

## Environment Variables

Create a `.env.local` file to override the default API URL:

```env
VITE_API_BASE_URL=/api/v1
```

The default works out of the box via the Vite proxy. Only change this for production deployments.

## Commands

```bash
# Dev server with hot reload
npm run dev

# Production build
npm run build

# Preview production build
npm run preview
```

## Pages

| Path | Description |
|------|-------------|
| `/` | Landing page |
| `/quests` | Scenario list |
| `/quests/:id/begin` | Start a quest |
| `/quests/:id/leaderboard` | Scenario leaderboard |
| `/play/:sessionId` | Gameplay — read, choose, submit |
| `/play/:sessionId/summary` | Session summary and score |
| `/profile/progress` | Personal progress across all scenarios |
