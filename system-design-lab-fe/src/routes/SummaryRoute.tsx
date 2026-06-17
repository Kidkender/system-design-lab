import { useState } from 'react'
import { useParams, Link, useNavigate } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { getSessionSummary, restartSession } from '@/features/session/api'
import { MetricsPanel } from '@/features/metrics/components/MetricsPanel'
import { PixelHeading } from '@/components/pixel/PixelHeading'
import { PixelPanel } from '@/components/pixel/PixelPanel'
import { PixelButton } from '@/components/pixel/PixelButton'
import { PATHS } from '@/lib/router/paths'
import type { ChoiceLogEntry } from '@/types/api'

function ChoiceLogItem({ entry, index }: { entry: ChoiceLogEntry; index: number }) {
  return (
    <PixelPanel variant={entry.isCorrect ? 'default' : 'blood'} className="flex flex-col gap-2">
      <p className="font-['Press_Start_2P'] text-[6px] text-[var(--parchment-dim)]">STEP {index + 1}</p>
      <p className="text-[var(--parchment)] font-['VT323'] text-lg">{entry.question}</p>
      <div className="flex items-center gap-3">
        <span
          style={{ color: entry.isCorrect ? 'var(--poison)' : 'var(--blood)' }}
          className="font-['Press_Start_2P'] text-[8px]"
        >
          {entry.isCorrect ? '✓' : '✗'}
        </span>
        <span className="text-[var(--parchment-dim)] font-['VT323'] text-lg">{entry.label}</span>
      </div>
      {entry.explanations['beginner'] && (
        <p className="text-[var(--parchment-dim)] text-sm italic">{entry.explanations['beginner']}</p>
      )}
    </PixelPanel>
  )
}

export function SummaryRoute() {
  const { sessionId } = useParams<{ sessionId: string }>()
  const navigate = useNavigate()
  const [isRestarting, setIsRestarting] = useState(false)

  async function handleRestart() {
    if (!sessionId) return
    setIsRestarting(true)
    try {
      const newSession = await restartSession(sessionId)
      navigate(PATHS.play(newSession.id))
    } finally {
      setIsRestarting(false)
    }
  }

  const { data: summary, isLoading, isError } = useQuery({
    queryKey: ['summary', sessionId],
    queryFn: () => getSessionSummary(sessionId!),
    enabled: !!sessionId,
  })

  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-64">
        <p className="font-['Press_Start_2P'] text-[10px] text-[var(--parchment-dim)] blink">
          TALLYING RESULTS...
        </p>
      </div>
    )
  }

  if (isError || !summary) {
    return (
      <PixelPanel variant="blood" className="max-w-md mx-auto mt-16 text-center">
        <p className="font-['Press_Start_2P'] text-[8px] text-[var(--blood)]">⚠ SUMMARY NOT FOUND</p>
      </PixelPanel>
    )
  }

  const correctCount = summary.choices.filter((c) => c.isCorrect).length
  const totalCount = summary.choices.length
  const score = totalCount > 0 ? Math.round((correctCount / totalCount) * 100) : 0

  return (
    <div className="flex flex-col gap-8 max-w-5xl mx-auto">

      <div className="text-center flex flex-col gap-4">
        <PixelHeading level={1} className="text-2xl" color={score >= 60 ? 'gold' : 'blood'}>
          {score >= 80 ? '★ QUEST COMPLETE ★' : score >= 60 ? '✓ QUEST DONE' : '✗ QUEST FAILED'}
        </PixelHeading>
        <p className="font-['Press_Start_2P'] text-[10px] text-[var(--parchment-dim)]">
          SCORE: {correctCount}/{totalCount} CORRECT ({score}%)
        </p>
      </div>

      <div className="flex flex-col lg:flex-row gap-8">
        <div className="flex-1 flex flex-col gap-4">
          <PixelHeading level={2} className="text-sm">DECISION LOG</PixelHeading>
          {summary.choices.map((entry, i) => (
            <ChoiceLogItem key={entry.choiceId} entry={entry} index={i} />
          ))}
        </div>

        <aside className="lg:w-56 flex flex-col gap-4">
          <PixelHeading level={2} className="text-sm">FINAL STATS</PixelHeading>
          <MetricsPanel metrics={summary.metrics} />
          <PixelButton
            variant="default"
            className="w-full justify-center"
            onClick={handleRestart}
            disabled={isRestarting}
          >
            {isRestarting ? '...' : '↺ RETRY QUEST'}
          </PixelButton>
          <Link to={PATHS.quests}>
            <PixelButton variant="gold" className="w-full justify-center">
              ▶ NEW QUEST
            </PixelButton>
          </Link>
        </aside>
      </div>
    </div>
  )
}
