DROP TABLE IF EXISTS user_choices;
DROP TABLE IF EXISTS user_sessions;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS conditions;
DROP TABLE IF EXISTS explanations;
DROP TABLE IF EXISTS consequences;
DROP TABLE IF EXISTS impacts;
DROP TABLE IF EXISTS choices;

ALTER TABLE scenarios DROP CONSTRAINT IF EXISTS fk_start_step;

DROP TABLE IF EXISTS steps;
DROP TABLE IF EXISTS scenarios;

DROP TYPE IF EXISTS session_status;
DROP TYPE IF EXISTS condition_type;
DROP TYPE IF EXISTS explanation_level;
DROP TYPE IF EXISTS impact_type;
DROP TYPE IF EXISTS impact_metric;
DROP TYPE IF EXISTS difficulty_level;
