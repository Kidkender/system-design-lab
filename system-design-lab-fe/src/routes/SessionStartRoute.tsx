import { useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { createUser } from '@/features/user/api'
import { createSession } from '@/features/session/api'
import { PATHS } from '@/lib/router/paths'
import { PixelHeading } from '@/components/pixel/PixelHeading'
import { PixelPanel } from '@/components/pixel/PixelPanel'
import { PixelButton } from '@/components/pixel/PixelButton'
import { PixelDialogue } from '@/components/pixel/PixelDialogue'

export function SessionStartRoute() {
  const { scenarioId } = useParams<{ scenarioId: string }>()
  const navigate = useNavigate()

  const [username, setUsername] = useState(localStorage.getItem('username') ?? '')
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  async function handleBegin() {
    if (!username.trim() || !scenarioId) return
    setIsLoading(true)
    setError(null)

    try {
      const user = await createUser(username.trim())
      localStorage.setItem('userId', user.ID)
      localStorage.setItem('username', username.trim())

      const session = await createSession(user.ID, scenarioId)
      localStorage.setItem('lastSessionId', session.id)

      navigate(PATHS.play(session.id))
    } catch {
      setError('The dungeon master rejected your entry. Try again.')
    } finally {
      setIsLoading(false)
    }
  }

  function handleKeyDown(e: React.KeyboardEvent) {
    if (e.key === 'Enter') handleBegin()
  }

  return (
    <div className="flex flex-col items-center justify-center min-h-[80vh] gap-8 p-6">
      <PixelHeading level={1} className="text-lg text-center leading-loose">
        ⚔ ENTER YOUR NAME, ADVENTURER ⚔
      </PixelHeading>

      <PixelPanel variant="gold" className="w-full max-w-md flex flex-col gap-6">
        <PixelDialogue speakerName="DUNGEON MASTER">
          Before you begin your quest, tell me your name. It shall be recorded in the annals of system design history.
        </PixelDialogue>

        <div className="flex flex-col gap-3">
          <label className="font-['Press_Start_2P'] text-[8px] text-[var(--parchment-dim)]">
            YOUR NAME:
          </label>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            onKeyDown={handleKeyDown}
            maxLength={20}
            placeholder="e.g. ARCH_WIZARD"
            className="w-full bg-[var(--dungeon-deep)] text-[var(--parchment)] font-['VT323'] text-xl px-3 py-2 pixel-border outline-none focus:pixel-border-gold"
            autoFocus
          />
          <p className="font-['Press_Start_2P'] text-[6px] text-[var(--parchment-dim)]">
            MAX 20 CHARACTERS
          </p>
        </div>

        {error && (
          <p className="font-['Press_Start_2P'] text-[8px] text-[var(--blood)] text-center">
            ⚠ {error}
          </p>
        )}

        <PixelButton
          variant="gold"
          size="lg"
          className="w-full justify-center"
          onClick={handleBegin}
          disabled={!username.trim() || isLoading}
        >
          {isLoading ? '...' : '▶ BEGIN QUEST'}
        </PixelButton>
      </PixelPanel>
    </div>
  )
}
