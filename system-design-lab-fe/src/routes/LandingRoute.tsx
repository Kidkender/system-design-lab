import { Link } from 'react-router-dom'
import { PATHS } from '@/lib/router/paths'
import { PixelHeading } from '@/components/pixel/PixelHeading'
import { PixelButton } from '@/components/pixel/PixelButton'
import { PixelPanel } from '@/components/pixel/PixelPanel'

export function LandingRoute() {
  const lastSessionId = localStorage.getItem('lastSessionId')

  return (
    <div className="min-h-screen bg-[var(--dungeon-deep)] flex flex-col items-center justify-center gap-12 p-6">

      {/* Title */}
      <div className="text-center flex flex-col gap-6">
        <p className="font-['Press_Start_2P'] text-[10px] text-[var(--parchment-dim)] tracking-widest">
          ✦ WELCOME TO ✦
        </p>
        <PixelHeading level={1} className="text-3xl md:text-5xl leading-loose text-shadow">
          SYSTEM DESIGN
        </PixelHeading>
        <PixelHeading level={1} className="text-3xl md:text-5xl leading-loose" color="blood">
          LAB
        </PixelHeading>
        <p className="font-['VT323'] text-xl text-[var(--parchment-dim)] max-w-md mx-auto">
          Learn system design by making real decisions. Every choice changes your system's fate.
        </p>
      </div>

      {/* CTAs */}
      <div className="flex flex-col items-center gap-4">
        <Link to={PATHS.quests}>
          <PixelButton variant="gold" size="lg">
            ▶ CHOOSE YOUR QUEST
          </PixelButton>
        </Link>

        {lastSessionId && (
          <Link to={PATHS.play(lastSessionId)}>
            <PixelButton variant="default" size="md">
              ↩ CONTINUE LAST QUEST
            </PixelButton>
          </Link>
        )}
      </div>

      {/* Info panels */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-3xl w-full">
        <PixelPanel className="text-center">
          <p className="font-['Press_Start_2P'] text-[8px] text-[var(--gold)] mb-3">⚔ DECIDE</p>
          <p className="text-[var(--parchment-dim)] text-sm">
            Every step forces a choice. No free drawing. Real architecture decisions.
          </p>
        </PixelPanel>
        <PixelPanel className="text-center" variant="gold">
          <p className="font-['Press_Start_2P'] text-[8px] text-[var(--gold)] mb-3">📊 CONSEQUENCE</p>
          <p className="text-[var(--parchment-dim)] text-sm">
            Watch your system's Latency, Cost, and Scalability shift with each decision.
          </p>
        </PixelPanel>
        <PixelPanel className="text-center">
          <p className="font-['Press_Start_2P'] text-[8px] text-[var(--gold)] mb-3">📖 LEARN</p>
          <p className="text-[var(--parchment-dim)] text-sm">
            Explanations at 3 levels: Beginner, Intermediate, and Advanced.
          </p>
        </PixelPanel>
      </div>
    </div>
  )
}
