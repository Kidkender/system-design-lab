import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { getSession, submitChoice } from '../api'
import type { SubmitChoiceRequest } from '@/types/api'

export function useSession(sessionId: string) {
  return useQuery({
    queryKey: ['session', sessionId],
    queryFn: () => getSession(sessionId),
    staleTime: 0,
    retry: 1,
  })
}

export function useSubmitChoice(sessionId: string) {
  const queryClient = useQueryClient()
  return useMutation({
    mutationFn: (body: SubmitChoiceRequest) => submitChoice(sessionId, body),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['session', sessionId] })
    },
    retry: 0,
  })
}
