import { BrowserRouter, Routes, Route } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { AppShell } from './components/layout/AppShell'
import { LandingRoute } from './routes/LandingRoute'
import { ScenariosRoute } from './routes/ScenariosRoute'
import { SessionStartRoute } from './routes/SessionStartRoute'
import { PlayRoute } from './routes/PlayRoute'
import { SummaryRoute } from './routes/SummaryRoute'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: { retry: 1 },
    mutations: { retry: 0 },
  },
})

function NotFound() {
  return (
    <div className="flex items-center justify-center min-h-64 text-center">
      <div>
        <p className="font-['Press_Start_2P'] text-[var(--blood)] text-2xl mb-4">404</p>
        <p className="font-['VT323'] text-[var(--parchment-dim)] text-xl">You wandered off the map.</p>
      </div>
    </div>
  )
}

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<LandingRoute />} />
          <Route
            path="/*"
            element={
              <AppShell>
                <Routes>
                  <Route path="quests" element={<ScenariosRoute />} />
                  <Route path="quests/:scenarioId/begin" element={<SessionStartRoute />} />
                  <Route path="play/:sessionId" element={<PlayRoute />} />
                  <Route path="play/:sessionId/summary" element={<SummaryRoute />} />
                  <Route path="*" element={<NotFound />} />
                </Routes>
              </AppShell>
            }
          />
        </Routes>
      </BrowserRouter>
    </QueryClientProvider>
  )
}
