import { apiClient } from '@/lib/api/client'
import type { UserProgressItem, UserSessionListItem } from '@/types/api'

export async function getUserProgress(userId: string): Promise<UserProgressItem[]> {
  const res = await apiClient.get<UserProgressItem[]>(`/users/${userId}/progress`)
  return res.data
}

export async function getUserSessions(userId: string): Promise<UserSessionListItem[]> {
  const res = await apiClient.get<UserSessionListItem[]>(`/users/${userId}/sessions`)
  return res.data
}
