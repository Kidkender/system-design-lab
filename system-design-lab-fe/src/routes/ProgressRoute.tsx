import { Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { getUserProgress } from '@/features/user/progress'
import { PixelHeading } from '@/components/pixel/PixelHeading'
import { PixelPanel } from '@/components/pixel/PixelPanel'
import { PixelButton } from '@/components/pixel/PixelButton'
import { PATHS } from '@/lib/router/paths'
import type { UserProgressItem } from '@/types/api'
import { cn } from '@/lib/utils'

const difficultyColor: Record<string, string> = {
  easy: 'var(--poison)',
  medium: 'var(--gold)',
  hard: 'var(--blood)',
}

function ScoreBar({ score }: { score: number }) {
  const pct = Math.min(100, Math.round(score))
  const color = pct >= 80 ? 'var(--poison)' : pct >= 50 ? 'var(--gold)' : 'var(--blood)'
  return (
    <div className="flex items-center gap-3">
      <div className="flex-1 h-2 bg-[var(--dungeon-deep)] border border-[var(--dungeon-border)]">
        <div
          className="h-full transition-none"
          style={{ width: `${pct}%`, backgroundColor: color }}
        />
      </div>
      <span className="font-['Press_Start_2P'] text-[8px] w-10 text-right" style={{ color }}>
        {pct}%
      </span>
    </div>
  )
}

function ProgressCard({ item }: { item: UserProgressItem }) {
  const isCompleted = item.completions > 0
  return (
    <PixelPanel className={cn('flex flex-col gap-3', !isCompleted && 'opacity-75')}>
      <div className="flex items-start justify-between gap-2">
        <PixelHeading level={3} className="text-xs leading-loose flex-1">
          {item.title}
        </PixelHeading>
        <span
          className="font-['Press_Start_2P'] text-[6px] shrink-0 mt-1"
          style={{ color: difficultyColor[item.difficulty] ?? 'var(--parchment-dim)' }}
        >
          {item.difficulty.toUpperCase()}
        </span>
      </div>

      {isCompleted ? (
        <>
          <ScoreBar score={item.bestScore} />
          <p className="font-['Press_Start_2P'] text-[6px] text-[var(--parchment-dim)]">
            {item.completions}/{item.attempts} runs completed
          </p>
        </>
      ) : item.attempts > 0 ? (
        <p className="font-['VT323'] text-lg text-[var(--blood)]">
          {item.attempts} attempt{item.attempts > 1 ? 's' : ''} — not yet completed
        </p>
      ) : (
        <p className="font-['VT323'] text-lg text-[var(--dungeon-border)]">Not started</p>
      )}

      <Link to={PATHS.leaderboard(item.scenarioId)}>
        <PixelButton size="sm" variant="default" className="text-[6px]">
          🏆 LEADERBOARD
        </PixelButton>
      </Link>
    </PixelPanel>
  )
}

export function ProgressRoute() {
  const userId = localStorage.getItem('userId')

  const { data, isLoading, isError } = useQuery({
    queryKey: ['progress', userId],
    queryFn: () => getUserProgress(userId!),
    enabled: !!userId,
    staleTime: 30_000,
  })

  if (!userId) {
    return (
      <div className="flex flex-col items-center gap-6 max-w-md mx-auto mt-16 text-center">
        <PixelHeading level={1} className="text-lg">📜 YOUR PROGRESS</PixelHeading>
        <PixelPanel variant="dark">
          <p className="font-['VT323'] text-xl text-[var(--parchment-dim)]">
            Start a quest first to track your progress.
          </p>
        </PixelPanel>
        <Link to={PATHS.quests}>
          <PixelButton variant="gold">▶ QUEST BOARD</PixelButton>
        </Link>
      </div>
    )
  }

  return (
    <div className="flex flex-col gap-8 max-w-5xl mx-auto">
      <div className="text-center">
        <PixelHeading level={1} className="text-xl mb-4">📜 YOUR PROGRESS</PixelHeading>
        <p className="text-[var(--parchment-dim)] font-['VT323'] text-lg">
          Track your journey through every system design challenge
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

      {data && (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {data.map((item) => (
            <ProgressCard key={item.scenarioId} item={item} />
          ))}
        </div>
      )}
    </div>
  )
}
