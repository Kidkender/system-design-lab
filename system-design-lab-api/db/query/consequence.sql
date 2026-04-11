-- name: GetConsequencesByChoiceIDs :many

SELECT
    choice_id,
    keys,
    value
FROM consequences
WHERE choice_id = ANY($1::uuid[]);
-- name: GetConsequence :one

SELECT
    id,
    choice_id,
    type,
    keys,
    value
FROM consequences
WHERE choice_id = $1::uuid;
-- name: CreateConsequence :one

INSERT INTO consequences (
    id,
    choice_id,
    type,
    keys,
    value
)
VALUES (
    $1::uuid,
    $2::uuid,
    $3,
    $4,
    $5
)
RETURNING id, choice_id, type, keys, value;


-- name: UpdateConsequence :one
UPDATE consequences
SET
    type = $2,
    keys = $3,
    value = $4
WHERE id = $1::uuid
RETURNING *;

-- name: DeleteConsequence :exec
DELETE FROM consequences
WHERE id = $1::uuid;
