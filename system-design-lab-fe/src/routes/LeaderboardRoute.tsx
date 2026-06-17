import { useParams, Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { getLeaderboard } from '@/features/scenarios/leaderboard'
import { PixelHeading } from '@/components/pixel/PixelHeading'
import { PixelPanel } from '@/components/pixel/PixelPanel'
import { PixelButton } from '@/components/pixel/PixelButton'
import { PATHS } from '@/lib/router/paths'
import type { LeaderboardEntry } from '@/types/api'

const rankColors: Record<number, string> = {
  1: 'var(--gold)',
  2: 'var(--parchment)',
  3: '#cd7f32',
}

function LeaderboardRow({ entry }: { entry: LeaderboardEntry }) {
  const color = rankColors[entry.rank] ?? 'var(--parchment-dim)'
  return (
    <PixelPanel variant="dark" className="flex items-center gap-4">
      <span
        className="font-['Press_Start_2P'] text-[10px] w-8 shrink-0 text-center"
        style={{ color }}
      >
        #{entry.rank}
      </span>
      <div className="flex-1 min-w-0">
        <p className="font-['VT323'] text-lg text-[var(--parchment)] truncate">{entry.username}</p>
        <p className="font-['Press_Start_2P'] text-[6px] text-[var(--parchment-dim)]">
          {entry.correctChoices}/{entry.totalChoices} correct
        </p>
      </div>
      <div className="text-right shrink-0">
        <p className="font-['Press_Start_2P'] text-[10px]" style={{ color }}>
          {Math.round(entry.score)}%
        </p>
      </div>
    </PixelPanel>
  )
}

export function LeaderboardRoute() {
  const { scenarioId } = useParams<{ scenarioId: string }>()

  const { data, isLoading, isError } = useQuery({
    queryKey: ['leaderboard', scenarioId],
    queryFn: () => getLeaderboard(scenarioId!, 10),
    enabled: !!scenarioId,
    staleTime: 30_000,
  })

  return (
    <div className="flex flex-col gap-8 max-w-2xl mx-auto">
      <div className="text-center">
        <PixelHeading level={1} className="text-xl mb-4">🏆 LEADERBOARD</PixelHeading>
        <p className="text-[var(--parchment-dim)] font-['VT323'] text-lg">
          Top architects who conquered this challenge
        </p>
      </div>

      {isLoading && (
        <div className="text-center">
          <p className="font-['Press_Start_2P'] text-[10px] text-[var(--parchment-dim)] blink">
            LOADING...
          </p>
        </div>
      )}

      {isError && (
        <PixelPanel variant="blood" className="text-center">
          <p className="font-['Press_Start_2P'] text-[8px] text-[var(--blood)]">⚠ FAILED TO LOAD</p>
        </PixelPanel>
      )}

      {data && data.length === 0 && (
        <PixelPanel variant="dark" className="text-center">
          <p className="font-['VT323'] text-xl text-[var(--parchment-dim)]">
            No completed runs yet. Be the first!
          </p>
        </PixelPanel>
      )}

      {data && data.length > 0 && (
        <div className="flex flex-col gap-3">
          {data.map((entry) => (
            <LeaderboardRow key={entry.sessionId} entry={entry} />
          ))}
        </div>
      )}

      <div className="flex gap-4 justify-center">
        <Link to={PATHS.quests}>
          <PixelButton variant="default">◀ QUEST BOARD</PixelButton>
        </Link>
        {scenarioId && (
          <Link to={PATHS.questBegin(scenarioId)}>
            <PixelButton variant="gold">▶ ACCEPT QUEST</PixelButton>
          </Link>
        )}
      </div>
    </div>
  )
}
