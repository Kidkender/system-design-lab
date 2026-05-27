export const PATHS = {
  home: '/',
  quests: '/quests',
  questBegin: (scenarioId: string) => `/quests/${scenarioId}/begin`,
  play: (sessionId: string) => `/play/${sessionId}`,
  summary: (sessionId: string) => `/play/${sessionId}/summary`,
} as const
