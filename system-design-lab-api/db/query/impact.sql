-- name: GetImpactsByChoiceIDs :many

SELECT
    choice_id,
    metric,
    type,
    value
FROM impacts
WHERE choice_id = ANY($1::uuid[]);
