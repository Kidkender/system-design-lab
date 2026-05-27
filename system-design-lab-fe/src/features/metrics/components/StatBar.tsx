import { cn } from '@/lib/utils'

interface StatBarProps {
  label: string
  value: number
  max: number
  color: 'red' | 'gold' | 'green' | 'blue'
  icon?: string
}

const colorMap = {
  red: 'bg-[var(--blood)]',
  gold: 'bg-[var(--gold)]',
  green: 'bg-[var(--poison)]',
  blue: 'bg-[var(--mana)]',
}

const SEGMENTS = 10

export function StatBar({ label, value, max, color, icon }: StatBarProps) {
  const ratio = Math.min(Math.max(value / max, 0), 1)
  const filledSegments = Math.round(ratio * SEGMENTS)

  return (
    <div className="flex flex-col gap-1">
      <div className="flex items-center justify-between">
        <span className="font-['Press_Start_2P'] text-[8px] text-[var(--parchment-dim)]">
          {icon && <span className="mr-1">{icon}</span>}
          {label}
        </span>
        <span className="font-['Press_Start_2P'] text-[8px] text-[var(--parchment)]">
          {value}/{max}
        </span>
      </div>

      <div className="flex gap-[2px] h-3 pixel-border p-[2px]">
        {Array.from({ length: SEGMENTS }).map((_, i) => (
          <div
            key={i}
            className={cn(
              'flex-1 transition-[background-color] duration-[400ms]',
              i < filledSegments ? colorMap[color] : 'bg-[var(--dungeon-deep)]',
            )}
            style={{ transitionTimingFunction: 'steps(4, end)' }}
          />
        ))}
      </div>
    </div>
  )
}
