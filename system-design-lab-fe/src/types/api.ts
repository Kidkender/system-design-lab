/* ---------- Pagination ---------- */
export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  limit: number
  totalPages: number
}

/* ---------- Scenario ---------- */
export interface ScenarioListItem {
  id: string
  title: string
  description: string
}

/* ---------- Step / Choice ---------- */
export interface Choice {
  id: string
  label: string
  nextStepId: string | null
  isCorrect: boolean
  impacts: Impact[]
  consequences: Consequence[]
  explanations: Record<string, string>
}

export interface Impact {
  metric: string
  type: 'add' | 'multiply' | 'set'
  value: number
}

export interface Consequence {
  key: string
  value: boolean
}

export interface Step {
  id: string
  question: string
  context: string | null
  choices: Choice[]
}

export interface ScenarioDetail {
  id: string
  title: string
  description: string
  steps: Step[]
}

/* ---------- Session ---------- */
export type SessionStatus = 'in_progress' | 'completed' | 'abandoned'

export type Metrics = Record<string, number>

export interface Session {
  id: string
  scenarioId: string
  currentStep: Step
  metrics: Metrics
  flags: Record<string, boolean>
  status: SessionStatus
}

/* ---------- Submit Choice ---------- */
export interface SubmitChoiceRequest {
  choiceId: string
}

export interface SubmitChoiceResponse {
  isCorrect: boolean
  feedback: string
  explanations: Record<string, string>
  metrics: Metrics
  nextStep: Step | null
  isCompleted: boolean
}

/* ---------- Summary ---------- */
export interface ChoiceLogEntry {
  stepId: string
  question: string
  choiceId: string
  label: string
  isCorrect: boolean
  explanations: Record<string, string>
}

export interface SessionSummary {
  id: string
  scenarioId: string
  status: SessionStatus
  metrics: Metrics
  flags: Record<string, boolean>
  choices: ChoiceLogEntry[]
  createdAt: string
}

/* ---------- User ---------- */
// Backend returns raw db.User (no JSON tags) → PascalCase fields
export interface User {
  ID: string
  Username: string
  Email: string
}

export interface CreateUserRequest {
  name: string
  email: string
  password: string
}

/* ---------- Session Start ---------- */
export interface StartSessionRequest {
  userId: string
  scenarioId: string
}
