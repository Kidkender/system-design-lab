import { StatBar } from './StatBar'
import type { Metrics } from '@/types/api'
import { PixelPanel } from '@/components/pixel/PixelPanel'

interface MetricConfig {
  icon: string
  color: 'red' | 'gold' | 'green' | 'blue'
  max: number
  label: string
}

const METRIC_CONFIG: Record<string, MetricConfig> = {
  latency:     { icon: '⚡', color: 'red',   max: 500, label: 'LATENCY' },
  cost:        { icon: '💰', color: 'gold',  max: 5,   label: 'COST' },
  scalability: { icon: '🔱', color: 'green', max: 10,  label: 'SCALABILITY' },
}

function getConfig(key: string): MetricConfig {
  return METRIC_CONFIG[key] ?? { icon: '📊', color: 'blue', max: 100, label: key.toUpperCase() }
}

interface MetricsPanelProps {
  metrics: Metrics
}

export function MetricsPanel({ metrics }: MetricsPanelProps) {
  const entries = Object.entries(metrics)

  return (
    <PixelPanel className="flex flex-col gap-4 min-w-[200px]">
      <p className="font-['Press_Start_2P'] text-[8px] text-[var(--parchment-dim)] text-center mb-2">
        SYSTEM STATS
      </p>
      {entries.map(([key, value]) => {
        const cfg = getConfig(key)
        return (
          <StatBar
            key={key}
            label={cfg.label}
            icon={cfg.icon}
            value={value}
            max={cfg.max}
            color={cfg.color}
          />
        )
      })}
      {entries.length === 0 && (
        <p className="text-[var(--parchment-dim)] text-sm text-center">No metrics yet</p>
      )}
    </PixelPanel>
  )
}
