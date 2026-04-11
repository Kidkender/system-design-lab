-- name: GetExplanationsByChoiceIDs :many

SELECT
    choice_id,
    level,
    content
FROM explanations
WHERE choice_id = ANY($1::uuid[]);
