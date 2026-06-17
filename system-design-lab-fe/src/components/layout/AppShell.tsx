import { Link } from 'react-router-dom'
import { PATHS } from '@/lib/router/paths'

interface AppShellProps {
  children: React.ReactNode
}

export function AppShell({ children }: AppShellProps) {
  return (
    <div className="min-h-screen bg-[var(--dungeon-deep)] flex flex-col">
      <header className="pixel-border-gold bg-[var(--dungeon)] px-4 py-3 flex items-center justify-between">
        <Link to={PATHS.home} className="font-['Press_Start_2P'] text-[10px] text-[var(--gold)] no-underline hover:text-[var(--parchment)] transition-none">
          ⚔ SYS.DESIGN.LAB
        </Link>
        <nav className="flex gap-4">
          <Link to={PATHS.quests} className="font-['Press_Start_2P'] text-[8px] text-[var(--parchment-dim)] no-underline hover:text-[var(--gold)] transition-none">
            QUESTS
          </Link>
          <Link to={PATHS.progress} className="font-['Press_Start_2P'] text-[8px] text-[var(--parchment-dim)] no-underline hover:text-[var(--gold)] transition-none">
            PROGRESS
          </Link>
        </nav>
      </header>

      <main className="flex-1 p-4 md:p-8">
        {children}
      </main>

      <footer className="text-center py-3 border-t-4 border-[var(--dungeon-border)]">
        <p className="font-['Press_Start_2P'] text-[6px] text-[var(--parchment-dim)]">
          LEARN SYSTEM DESIGN THROUGH DECISIONS
        </p>
      </footer>
    </div>
  )
}
