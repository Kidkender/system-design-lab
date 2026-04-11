CREATE TYPE difficulty_level AS ENUM ('easy', 'medium', 'hard');

CREATE TYPE impact_metric AS ENUM ('latency', 'cost', 'scalability');

CREATE TYPE impact_type AS ENUM ('add', 'multiply', 'set');

CREATE TYPE explanation_level AS ENUM ('beginner', 'intermediate', 'advanced');

CREATE TYPE condition_type AS ENUM ('metric', 'flag', 'choice');

CREATE TYPE session_status AS ENUM ('in_progress', 'completed', 'failed');

CREATE TABLE scenarios (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    start_step_id UUID,
    difficulty difficulty_level NOT NULL DEFAULT 'medium',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE steps (
    id UUID PRIMARY KEY,
    scenario_id UUID NOT NULL REFERENCES scenarios(id) ON DELETE CASCADE,
    question TEXT NOT NULL,
    context TEXT,
    order_index INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_steps_scenario ON steps(scenario_id);

ALTER TABLE scenarios
ADD CONSTRAINT fk_start_step
FOREIGN KEY (start_step_id) REFERENCES steps(id);

CREATE TABLE choices (
    id UUID PRIMARY KEY,
    step_id UUID NOT NULL REFERENCES steps(id) ON DELETE CASCADE,
    label TEXT NOT NULL,
    next_step_id UUID REFERENCES steps(id),
    is_correct BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_choices_step ON choices(step_id);

CREATE TABLE impacts (
    id UUID PRIMARY KEY,
    choice_id UUID NOT NULL REFERENCES choices(id) ON DELETE CASCADE,
    metric impact_metric NOT NULL,
    type impact_type NOT NULL,
    value FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_impacts_choice ON impacts(choice_id);

CREATE TABLE consequences (
    id UUID PRIMARY KEY,
    choice_id UUID NOT NULL REFERENCES choices(id) ON DELETE CASCADE,
    type TEXT NOT NULL DEFAULT 'flag',
    keys TEXT NOT NULL,
    value BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_consequences_choice ON consequences(choice_id);

CREATE TABLE explanations (
    id UUID PRIMARY KEY,
    choice_id UUID NOT NULL REFERENCES choices(id) ON DELETE CASCADE,
    level explanation_level NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_explanations_choice ON explanations(choice_id);

CREATE TABLE conditions (
    id UUID PRIMARY KEY,
    step_id UUID NOT NULL REFERENCES steps(id) ON DELETE CASCADE,
    type condition_type NOT NULL,

    metric impact_metric,
    operator TEXT,
    value FLOAT,

    float_key TEXT,

    choice_id UUID REFERENCES choices(id),

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE users (
    id UUID PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE user_sessions (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    scenario_id UUID NOT NULL REFERENCES scenarios(id),
    current_step_id UUID REFERENCES steps(id),

    metrics JSONB NOT NULL DEFAULT '{}',
    flags JSONB NOT NULL DEFAULT '{}',

    status session_status NOT NULL DEFAULT 'in_progress',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE user_choices (
    id UUID PRIMARY KEY,
    session_id UUID NOT NULL REFERENCES user_sessions(id) ON DELETE CASCADE,
    step_id UUID NOT NULL REFERENCES steps(id) ON DELETE CASCADE,
    choice_id UUID NOT NULL REFERENCES choices(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
