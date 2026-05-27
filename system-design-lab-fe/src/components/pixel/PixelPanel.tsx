import { cn } from '@/lib/utils'

interface PixelPanelProps {
  children: React.ReactNode
  className?: string
  variant?: 'default' | 'gold' | 'blood' | 'dark'
}

export function PixelPanel({ children, className, variant = 'default' }: PixelPanelProps) {
  return (
    <div
      className={cn(
        'p-4',
        variant === 'default' && 'pixel-border bg-[var(--dungeon)]',
        variant === 'gold' && 'pixel-border-gold bg-[var(--dungeon)]',
        variant === 'blood' && 'pixel-border-blood bg-[var(--dungeon)]',
        variant === 'dark' && 'pixel-border bg-[var(--dungeon-deep)]',
        className,
      )}
    >
      {children}
    </div>
  )
}
