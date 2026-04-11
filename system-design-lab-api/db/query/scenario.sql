-- name: GetScenario :one

SELECT
    id,
    title,
    description,
    start_step_id
FROM scenarios
WHERE id =$1::uuid;
