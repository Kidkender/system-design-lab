import { apiClient } from '@/lib/api/client'
import type {
  Session,
  SessionSummary,
  SessionMode,
  SubmitChoiceRequest,
  SubmitChoiceResponse,
  StartSessionRequest,
} from '@/types/api'

export async function createSession(
  userId: string,
  scenarioId: string,
  mode?: SessionMode,
): Promise<Session> {
  const body: StartSessionRequest = { userId, scenarioId, mode }
  const res = await apiClient.post<Session>('/sessions', body)
  return res.data
}

export async function getSession(sessionId: string): Promise<Session> {
  const res = await apiClient.get<Session>(`/sessions/${sessionId}`)
  return res.data
}

export async function submitChoice(
  sessionId: string,
  body: SubmitChoiceRequest,
): Promise<SubmitChoiceResponse> {
  const res = await apiClient.post<SubmitChoiceResponse>(`/sessions/${sessionId}/submit`, body)
  return res.data
}

export async function getSessionSummary(sessionId: string): Promise<SessionSummary> {
  const res = await apiClient.get<SessionSummary>(`/sessions/${sessionId}/summary`)
  return res.data
}

export async function abandonSession(sessionId: string): Promise<void> {
  await apiClient.post(`/sessions/${sessionId}/abandon`)
}

export async function restartSession(sessionId: string): Promise<Session> {
  const res = await apiClient.post<Session>(`/sessions/${sessionId}/restart`)
  return res.data
}
