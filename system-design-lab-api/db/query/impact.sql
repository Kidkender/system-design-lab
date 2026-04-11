-- name: GetImpactsByChoiceIDs :many

SELECT
    choice_id,
    metric,
    type,
    value
FROM impacts
WHERE choice_id = ANY($1::uuid[]);

-- name: CreateImpact :one
INSERT INTO impacts (
    id,
    choice_id,
    metric,
    type,
    value
) VALUES (
    $1::uuid,
    $2::uuid,
    $3,
    $4,
    $5
) RETURNING *;

-- name: UpdateImpact :one
UPDATE impacts
SET
    metric = $2,
    type = $3,
    value = $4
WHERE id = $1::uuid
RETURNING *;

-- name: DeleteImpact :exec
DELETE FROM impacts
WHERE id = $1::uuid;
