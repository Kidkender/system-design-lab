import { cn } from '@/lib/utils'

interface PixelHeadingProps {
  children: React.ReactNode
  level?: 1 | 2 | 3
  className?: string
  color?: 'gold' | 'blood' | 'parchment' | 'mana'
}

const colorMap = {
  gold: 'text-[var(--gold)]',
  blood: 'text-[var(--blood)]',
  parchment: 'text-[var(--parchment)]',
  mana: 'text-[var(--mana)]',
}

const sizeMap = {
  1: 'text-2xl md:text-3xl',
  2: 'text-lg md:text-xl',
  3: 'text-sm md:text-base',
}

export function PixelHeading({ children, level = 1, className, color = 'gold' }: PixelHeadingProps) {
  const Tag = `h${level}` as 'h1' | 'h2' | 'h3'
  return (
    <Tag
      className={cn(
        "font-['Press_Start_2P'] leading-relaxed",
        sizeMap[level],
        colorMap[color],
        className,
      )}
    >
      {children}
    </Tag>
  )
}
