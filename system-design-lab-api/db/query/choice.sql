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
VALUES ($1::uuid, $2::uuid, $3, $4, $5)
RETURNING *;

-- name: UpdateChoice :one
UPDATE choices
SET label = $2, next_step_id = $3, is_correct = $4
WHERE id = $1::uuid
RETURNING *;
