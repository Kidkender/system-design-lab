-- name: GetStepsByScenario :many

SELECT
    id,
    question,
    context
FROM steps
WHERE scenario_id = $1::uuid;
