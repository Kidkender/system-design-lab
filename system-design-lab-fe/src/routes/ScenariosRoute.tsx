import { useState } from 'react'
import { Link } from 'react-router-dom'
import { useScenarios } from '@/features/scenarios/hooks/useScenarios'
import { PATHS } from '@/lib/router/paths'
import { PixelHeading } from '@/components/pixel/PixelHeading'
import { PixelPanel } from '@/components/pixel/PixelPanel'
import { PixelButton } from '@/components/pixel/PixelButton'
import type { ScenarioListItem } from '@/types/api'

function ScenarioCard({ scenario }: { scenario: ScenarioListItem }) {
  return (
    <PixelPanel className="flex flex-col gap-4">
      <PixelHeading level={3} className="text-xs leading-loose">
        {scenario.title}
      </PixelHeading>
      <p className="text-[var(--parchment-dim)] text-sm flex-1">{scenario.description}</p>
      <Link to={PATHS.questBegin(scenario.id)}>
        <PixelButton variant="gold" className="w-full justify-center">
          ▶ ACCEPT QUEST
        </PixelButton>
      </Link>
    </PixelPanel>
  )
}

export function ScenariosRoute() {
  const [page, setPage] = useState(1)
  const { data, isLoading, isError } = useScenarios(page)

  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-64">
        <p className="font-['Press_Start_2P'] text-[10px] text-[var(--parchment-dim)] blink">
          LOADING QUESTS...
        </p>
      </div>
    )
  }

  if (isError || !data) {
    return (
      <PixelPanel variant="blood" className="max-w-md mx-auto mt-16 text-center">
        <p className="font-['Press_Start_2P'] text-[8px] text-[var(--blood)] mb-4">⚠ CONNECTION FAILED</p>
        <p className="text-[var(--parchment-dim)]">Could not reach the dungeon server. Is the backend running?</p>
      </PixelPanel>
    )
  }

  return (
    <div className="flex flex-col gap-8 max-w-5xl mx-auto">
      <div className="text-center">
        <PixelHeading level={1} className="text-xl mb-4">⚔ QUEST BOARD ⚔</PixelHeading>
        <p className="text-[var(--parchment-dim)]">Choose a system to design. Your decisions shape the architecture.</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {data.data.map((scenario) => (
          <ScenarioCard key={scenario.id} scenario={scenario} />
        ))}
      </div>

      {data.totalPages > 1 && (
        <div className="flex justify-center items-center gap-4">
          <PixelButton size="sm" onClick={() => setPage((p) => Math.max(1, p - 1))} disabled={page === 1}>
            ◀ PREV
          </PixelButton>
          <span className="font-['Press_Start_2P'] text-[8px] text-[var(--parchment-dim)]">
            {page} / {data.totalPages}
          </span>
          <PixelButton size="sm" onClick={() => setPage((p) => Math.min(data.totalPages, p + 1))} disabled={page === data.totalPages}>
            NEXT ▶
          </PixelButton>
        </div>
      )}
    </div>
  )
}
