import { apiClient } from '@/lib/api/client'
import type { LeaderboardEntry } from '@/types/api'

export async function getLeaderboard(
  scenarioId: string,
  topN = 10,
): Promise<LeaderboardEntry[]> {
  const res = await apiClient.get<LeaderboardEntry[]>(
    `/scenarios/${scenarioId}/leaderboard`,
    { params: { top_n: topN } },
  )
  return res.data
}
