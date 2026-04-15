-- name: GetChoicesByStepIDs :many

SELECT
    id,
    step_id,
    label,
    next_step_id,
    is_correct
FROM choices
WHERE step_id = ANY($1::uuid[])
ORDER BY step_id, created_at;

-- name: GetChoice :one
SELECT
    id,
    step_id,
    label,
    next_step_id,
    is_correct
FROM choices
WHERE id = $1::uuid;

-- name: DeleteChoice :exec
DELETE FROM choices
WHERE id = $1::uuid;

-- name: CreateChoice :one
INSERT INTO choices (id, step_id, label, next_step_id, is_correct)
VALUES (@id::uuid, @step_id::uuid, @label, @next_step_id, @is_correct)
RETURNING *;

-- name: UpdateChoice :one
UPDATE choices
SET label = $2, next_step_id = $3, is_correct = $4
WHERE id = $1::uuid
RETURNING *;
