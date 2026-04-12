-- name: GetScenario :one

SELECT
    id,
    title,
    description,
    start_step_id
FROM scenarios
WHERE id =$1::uuid;
-- name: GetScenarios :many

SELECT
    id,
    title,
    description,
    start_step_id
FROM scenarios
ORDER BY created_at DESC;

-- name: ListScenariosPaginated :many
SELECT *
FROM scenarios
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: CreateScenario :one

INSERT INTO scenarios (id, title, description, difficulty)
VALUES (@id::uuid, @title, @description, @difficulty)
RETURNING id, title, description, start_step_id, difficulty;
-- name: UpdateScenario :one

UPDATE scenarios
SET title = $2, description = $3, start_step_id = $4, difficulty = $5
WHERE id = $1::uuid
RETURNING id, title, description, start_step_id, difficulty;

-- name: DeleteScenario :exec
DELETE FROM scenarios
WHERE id = $1::uuid;

-- name: UpdateStartStep :one
UPDATE scenarios
SET start_step_id = $2
WHERE id = $1::uuid
RETURNING id, title, description, start_step_id, difficulty;