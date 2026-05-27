import { cn } from '@/lib/utils'

interface PixelDialogueProps {
  children: React.ReactNode
  speakerName?: string
  showContinue?: boolean
  className?: string
}

export function PixelDialogue({ children, speakerName, showContinue, className }: PixelDialogueProps) {
  return (
    <div className={cn('relative pixel-border bg-[var(--dungeon-deep)] p-4', className)}>
      {speakerName && (
        <div className="absolute -top-5 left-4 bg-[var(--dungeon-deep)] px-2 border-2 border-[var(--dungeon-border)]">
          <span className="font-['Press_Start_2P'] text-[8px] text-[var(--gold)]">{speakerName}</span>
        </div>
      )}

      <div className="text-[var(--parchment)] leading-relaxed min-h-[60px]">
        {children}
      </div>

      {showContinue && (
        <span className="absolute bottom-3 right-4 text-[var(--gold)] blink font-['Press_Start_2P'] text-[8px]">
          ▼
        </span>
      )}
    </div>
  )
}
