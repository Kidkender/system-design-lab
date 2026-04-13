-- name: GetStepsByScenario :many

SELECT
    id,
    question,
    context
FROM steps
WHERE scenario_id = $1::uuid;

-- name: GetStepsByScenarioPaginated :many
SELECT
    id,
    question,
    context
FROM steps
WHERE scenario_id = $1::uuid
ORDER BY order_index
LIMIT $2 OFFSET $3;

-- name: GetStepsPaginated :many
SELECT
    id,
    question,
    context
FROM steps
ORDER BY order_index
LIMIT $1 OFFSET $2;

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

-- name: ExistsStepOrderIndex :one
SELECT EXISTS(
    SELECT 1
    FROM steps
    WHERE scenario_id = $1::uuid
    AND order_index = $2
);

-- name: SetStartStepIfNull :exec
UPDATE scenarios 
SET start_step_id = $2
WHERE id = $1::uuid
    AND start_step_id IS NULL;