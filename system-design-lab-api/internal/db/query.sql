-- name: GetScenario :one

SELECT
    id,
    title,
    description,
    start_step_id
FROM scenarios
WHERE id =$1::uuid;
-- name: GetStepsByScenario :many

SELECT
    id,
    question,
    context
FROM steps
WHERE scenario_id = $1::uuid;
-- name: GetChoicesByStepIDs :many

SELECT
    id,
    step_id,
    label,
    next_step_id,
    is_correct
FROM choices
WHERE step_id = ANY($1::uuid[]);
-- name: GetImpactsByChoiceIDs :many

SELECT
    choice_id,
    metric,
    type,
    value
FROM impacts
WHERE choice_id = ANY($1::uuid[]);
-- name: GetExplanationsByChoiceIDs :many

SELECT
    choice_id,
    level,
    content
FROM explanations
WHERE choice_id = ANY($1::uuid[]);
-- name: GetConsequencesByChoiceIDs :many

SELECT
    choice_id,
    keys,
    value
FROM consequences
WHERE choice_id = ANY($1::uuid[]);
