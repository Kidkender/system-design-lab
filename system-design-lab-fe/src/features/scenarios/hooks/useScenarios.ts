import { useQuery } from '@tanstack/react-query'
import { listScenarios } from '../api'

export function useScenarios(page: number) {
  return useQuery({
    queryKey: ['scenarios', { page }],
    queryFn: () => listScenarios(page),
    staleTime: 60_000,
  })
}
