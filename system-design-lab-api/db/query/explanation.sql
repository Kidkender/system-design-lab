-- name: GetExplanationsByChoiceIDs :many

SELECT
    choice_id,
    level,
    content
FROM explanations
WHERE choice_id = ANY($1::uuid[]);

-- name: CreateExplanation :one
INSERT INTO explanations (
    id,
    choice_id,
    level,
    content
) VALUES (
    $1::uuid,
    $2::uuid,
    $3,
    $4
) RETURNING *;

-- name: UpdateExplanation :one
UPDATE explanations
SET
    level = $2,
    content = $3
WHERE id = $1::uuid
RETURNING *;

-- name: DeleteExplanation :exec
DELETE FROM explanations
WHERE id = $1::uuid;