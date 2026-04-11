-- name: CreateCondition :one

INSERT INTO conditions (
    id,
    step_id,
    type,
    metric,
    operator,
    value,
    float_key,
    choice_id
)
VALUES (
    $1::uuid,
    $2::uuid,
    $3::condition_type,
    $4::impact_metric,
    $5::text,
    $6::float,
    $7::text,
    $8::uuid
)
RETURNING id, step_id, type, metric, operator, value, float_key, choice_id;

-- name: GetCondition :one
SELECT
    id,
    step_id,
    type,
    metric,
    operator,
    value,
    float_key,
    choice_id
FROM conditions
WHERE id = $1::uuid;
-- name: DeleteCondition :exec
DELETE FROM conditions
WHERE id = $1::uuid;

-- name: GetConditionsByStep :many
SELECT
    id,
    step_id,
    type,
    metric,
    operator,
    value,
    float_key,
    choice_id
FROM conditions
WHERE step_id = $1::uuid
ORDER BY created_at;

