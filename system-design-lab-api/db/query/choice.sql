-- name: GetChoicesByStepIDs :many

SELECT
    id,
    step_id,
    label,
    next_step_id,
    is_correct
FROM choices
WHERE step_id = ANY($1::uuid[]);
