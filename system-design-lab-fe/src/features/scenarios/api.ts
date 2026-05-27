import { apiClient } from '@/lib/api/client'
import type { PaginatedResponse, ScenarioListItem, ScenarioDetail } from '@/types/api'

export async function listScenarios(page = 1, limit = 9): Promise<PaginatedResponse<ScenarioListItem>> {
  const res = await apiClient.get<PaginatedResponse<ScenarioListItem>>('/scenarios', {
    params: { page, limit },
  })
  return res.data
}

export async function getScenario(id: string): Promise<ScenarioDetail> {
  const res = await apiClient.get<ScenarioDetail>(`/scenarios/${id}`)
  return res.data
}
