-- name: GetConsequencesByChoiceIDs :many

SELECT
    choice_id,
    keys,
    value
FROM consequences
WHERE choice_id = ANY($1::uuid[]);
