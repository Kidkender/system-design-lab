export const PATHS = {
  home: '/',
  quests: '/quests',
  questBegin: (scenarioId: string) => `/quests/${scenarioId}/begin`,
  leaderboard: (scenarioId: string) => `/quests/${scenarioId}/leaderboard`,
  play: (sessionId: string) => `/play/${sessionId}`,
  summary: (sessionId: string) => `/play/${sessionId}/summary`,
  progress: '/profile/progress',
} as const
