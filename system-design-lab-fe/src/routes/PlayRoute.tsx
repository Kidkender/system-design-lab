import { useState, useCallback } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { useSession, useSubmitChoice } from '@/features/session/hooks/useSession'
import { abandonSession } from '@/features/session/api'
import { MetricsPanel } from '@/features/metrics/components/MetricsPanel'
import { PixelDialogue } from '@/components/pixel/PixelDialogue'
import { PixelPanel } from '@/components/pixel/PixelPanel'
import { PixelButton } from '@/components/pixel/PixelButton'
import { PixelHeading } from '@/components/pixel/PixelHeading'
import { PATHS } from '@/lib/router/paths'
import type { SubmitChoiceResponse, Step, Choice } from '@/types/api'
import { cn } from '@/lib/utils'

type ExplanationLevel = 'beginner' | 'intermediate' | 'advanced'

const levelLabel: Record<ExplanationLevel, string> = {
  beginner: '🟢 BEGINNER',
  intermediate: '🟡 INTER.',
  advanced: '🔴 ADVANCED',
}

function ExplanationTabs({ explanations }: { explanations: Record<string, string> }) {
  const [level, setLevel] = useState<ExplanationLevel>('beginner')
  const content = explanations[level] ?? ''

  if (!content) return null

  return (
    <div className="flex flex-col gap-3">
      <div className="flex gap-2 flex-wrap">
        {(['beginner', 'intermediate', 'advanced'] as ExplanationLevel[]).map((l) => (
          <button
            key={l}
            onClick={() => setLevel(l)}
            className={cn(
              "font-['Press_Start_2P'] text-[6px] px-2 py-1 border-2 cursor-pointer transition-none",
              level === l
                ? 'border-[var(--gold)] bg-[var(--dungeon-light)] text-[var(--gold)]'
                : 'border-[var(--dungeon-border)] bg-[var(--dungeon)] text-[var(--parchment-dim)]',
            )}
          >
            {levelLabel[l]}
          </button>
        ))}
      </div>
      <PixelPanel variant="dark" className="text-sm text-[var(--parchment)] leading-relaxed">
        {content}
      </PixelPanel>
    </div>
  )
}

function FeedbackPanel({ result }: { result: SubmitChoiceResponse }) {
  return (
    <PixelPanel variant={result.isCorrect ? 'default' : 'blood'} className="flex flex-col gap-4">
      <p
        className="font-['Press_Start_2P'] text-[8px]"
        style={{ color: result.isCorrect ? 'var(--poison)' : 'var(--blood)' }}
      >
        {result.isCorrect ? '✓ CORRECT' : '✗ WRONG'}
      </p>
      <p className="text-[var(--parchment)] text-lg">{result.feedback}</p>
      {Object.keys(result.explanations).length > 0 && (
        <ExplanationTabs explanations={result.explanations} />
      )}
    </PixelPanel>
  )
}

function ChoiceMenu({
  step,
  selected,
  disabled,
  onSelect,
}: {
  step: Step
  selected: string | null
  disabled: boolean
  onSelect: (id: string) => void
}) {
  const handleKey = useCallback(
    (e: React.KeyboardEvent<HTMLButtonElement>, choice: Choice) => {
      if (e.key === 'Enter' || e.key === ' ') onSelect(choice.id)
    },
    [onSelect],
  )

  return (
    <div className="flex flex-col gap-3" role="radiogroup" aria-label="Choose your action">
      {step.choices.map((choice) => (
        <button
          key={choice.id}
          role="radio"
          aria-checked={selected === choice.id}
          disabled={disabled}
          onClick={() => onSelect(choice.id)}
          onKeyDown={(e) => handleKey(e, choice)}
          className={cn(
            "text-left w-full px-4 py-3 border-2 font-['VT323'] text-lg transition-none cursor-pointer",
            "flex items-start gap-3",
            selected === choice.id
              ? 'border-[var(--gold)] bg-[var(--dungeon-light)] text-[var(--gold)]'
              : 'border-[var(--dungeon-border)] bg-[var(--dungeon)] text-[var(--parchment)] hover:border-[var(--dungeon-light)] hover:bg-[var(--dungeon-light)]',
            disabled && 'opacity-50 cursor-not-allowed',
          )}
        >
          <span className="font-['Press_Start_2P'] text-[8px] mt-1 shrink-0" aria-hidden>
            {selected === choice.id ? '▶' : '○'}
          </span>
          <span>{choice.label}</span>
        </button>
      ))}
    </div>
  )
}

export function PlayRoute() {
  const { sessionId } = useParams<{ sessionId: string }>()
  const navigate = useNavigate()

  const [selectedChoiceId, setSelectedChoiceId] = useState<string | null>(null)
  const [lastResult, setLastResult] = useState<SubmitChoiceResponse | null>(null)
  const [currentStep, setCurrentStep] = useState<Step | null>(null)
  const [showHint, setShowHint] = useState(false)
  const [isAbandoning, setIsAbandoning] = useState(false)

  const { data: session, isLoading, isError } = useSession(sessionId ?? '')
  const submitMutation = useSubmitChoice(sessionId ?? '')

  if (!sessionId) return null

  if (isLoading) {
    return (
      <div className="flex items-center justify-center min-h-64">
        <p className="font-['Press_Start_2P'] text-[10px] text-[var(--parchment-dim)] blink">LOADING...</p>
      </div>
    )
  }

  if (isError || !session) {
    return (
      <PixelPanel variant="blood" className="max-w-md mx-auto mt-16 text-center">
        <p className="text-[var(--blood)] font-['Press_Start_2P'] text-[8px]">⚠ SESSION NOT FOUND</p>
      </PixelPanel>
    )
  }

  if (session.status === 'completed') {
    navigate(PATHS.summary(sessionId))
    return null
  }

  const displayStep = currentStep ?? session.currentStep

  async function handleAbandon() {
    if (!sessionId) return
    setIsAbandoning(true)
    try {
      await abandonSession(sessionId)
      navigate(PATHS.quests)
    } finally {
      setIsAbandoning(false)
    }
  }

  async function handleSubmit() {
    if (!selectedChoiceId || !sessionId) return

    const result = await submitMutation.mutateAsync({ choiceId: selectedChoiceId })
    setLastResult(result)
    setSelectedChoiceId(null)

    if (result.isCompleted) {
      setTimeout(() => navigate(PATHS.summary(sessionId!)), 2500)
    } else if (result.nextStep) {
      setCurrentStep(result.nextStep)
    }
  }

  return (
    <div className="flex flex-col lg:flex-row gap-6 max-w-6xl mx-auto">

      {/* Main gameplay */}
      <div className="flex-1 flex flex-col gap-6">
        <PixelHeading level={2} className="text-sm">⚔ MAKE YOUR DECISION</PixelHeading>

        {/* Question */}
        <PixelDialogue speakerName="SYSTEM MASTER">
          <span className="text-xl">{displayStep.question}</span>
          {displayStep.context && (
            <p className="text-[var(--parchment-dim)] text-sm mt-2">{displayStep.context}</p>
          )}
          {displayStep.hint && (
            <div className="mt-3">
              <button
                onClick={() => setShowHint((v) => !v)}
                className="font-['Press_Start_2P'] text-[6px] text-[var(--gold)] underline cursor-pointer"
              >
                {showHint ? '▲ HIDE HINT' : '▼ SHOW HINT'}
              </button>
              {showHint && (
                <p className="mt-2 text-[var(--gold)] text-sm italic">{displayStep.hint}</p>
              )}
            </div>
          )}
        </PixelDialogue>

        {/* Feedback from last submit */}
        {lastResult && !submitMutation.isPending && (
          <FeedbackPanel result={lastResult} />
        )}

        {/* Choices — only show if no pending feedback, or after feedback shown */}
        {(!lastResult || lastResult.nextStep) && (
          <>
            <ChoiceMenu
              step={displayStep}
              selected={selectedChoiceId}
              disabled={submitMutation.isPending}
              onSelect={setSelectedChoiceId}
            />

            <PixelButton
              variant="gold"
              size="lg"
              onClick={handleSubmit}
              disabled={!selectedChoiceId || submitMutation.isPending}
              className="self-start"
            >
              {submitMutation.isPending ? '▶ THINKING...' : '▶ CONFIRM DECISION'}
            </PixelButton>
          </>
        )}

        {lastResult?.isCompleted && (
          <PixelPanel variant="gold" className="text-center">
            <p className="font-['Press_Start_2P'] text-[8px] text-[var(--gold)] blink">
              ★ QUEST COMPLETE — TALLYING RESULTS... ★
            </p>
          </PixelPanel>
        )}
      </div>

      {/* Metrics HUD */}
      <aside className="lg:w-56 flex flex-col gap-4">
        <MetricsPanel metrics={session.metrics} />

        {session.mode === 'interview' && session.timeLimitSeconds != null && (
          <PixelPanel variant="dark" className="text-center">
            <p className="font-['Press_Start_2P'] text-[6px] text-[var(--blood)] mb-1">⏱ INTERVIEW</p>
            <p className="font-['VT323'] text-xl text-[var(--gold)]">
              {Math.max(0, session.timeLimitSeconds - session.timeElapsedSeconds)}s left
            </p>
          </PixelPanel>
        )}

        <PixelPanel variant="dark" className="text-center">
          <p className="font-['Press_Start_2P'] text-[6px] text-[var(--parchment-dim)] mb-1">SESSION</p>
          <p className="font-['VT323'] text-sm text-[var(--parchment-dim)]">{sessionId.slice(0, 12)}...</p>
          <p className="font-['Press_Start_2P'] text-[6px] text-[var(--dungeon-border)] mt-2 capitalize">
            {session.status}
          </p>
        </PixelPanel>

        <PixelButton
          variant="blood"
          size="sm"
          className="w-full justify-center"
          onClick={handleAbandon}
          disabled={isAbandoning}
        >
          {isAbandoning ? '...' : '✗ ABANDON QUEST'}
        </PixelButton>
      </aside>
    </div>
  )
}
