import { cn } from '@/lib/utils'

interface PixelButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'default' | 'gold' | 'blood'
  size?: 'sm' | 'md' | 'lg'
}

export function PixelButton({
  children,
  className,
  variant = 'default',
  size = 'md',
  ...props
}: PixelButtonProps) {
  return (
    <button
      className={cn(
        'pixel-btn',
        variant === 'gold' && 'pixel-btn-gold',
        variant === 'blood' && 'bg-[var(--blood-dark)] text-[var(--parchment)] shadow-[[-4px_0_0_var(--blood),0_-4px_0_var(--blood),4px_0_0_var(--shadow-color),0_4px_0_var(--shadow-color)]]',
        size === 'sm' && 'text-[8px] px-3 py-2',
        size === 'lg' && 'text-[12px] px-6 py-4',
        'disabled:opacity-50 disabled:cursor-not-allowed disabled:transform-none',
        className,
      )}
      {...props}
    >
      {children}
    </button>
  )
}
