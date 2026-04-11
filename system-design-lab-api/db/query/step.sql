-- name: GetStepsByScenario :many

SELECT
    id,
    question,
    context
FROM steps
WHERE scenario_id = $1::uuid;

-- name: DeleteStep :exec
DELETE FROM steps
WHERE id = $1::uuid;

-- name: CreateStep :one
INSERT INTO steps (id, scenario_id, question, context, order_index)
VALUES ($1::uuid, $2::uuid, $3, $4, $5)
RETURNING *;

-- name: UpdateStep :one
UPDATE steps
SET question = $2, context = $3, order_index = $4
WHERE id = $1::uuid
RETURNING *;
