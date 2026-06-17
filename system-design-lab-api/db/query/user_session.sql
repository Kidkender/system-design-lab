-- name: CreateUserSession :one
INSERT INTO user_sessions (id, user_id, scenario_id, current_step_id, metrics, flags, status, mode)
VALUES (@id::uuid, @user_id::uuid, @scenario_id::uuid, @current_step_id::uuid, @metrics, @flags, 'in_progress', @mode)
RETURNING *;

-- name: GetUserSession :one
SELECT * FROM user_sessions
WHERE id = $1::uuid;

-- name: UpdateUserSession :one
UPDATE user_sessions
SET current_step_id = @current_step_id,
    metrics         = @metrics,
    flags           = @flags,
    status          = @status,
    completed_at    = CASE WHEN @status::session_status IN ('completed', 'failed', 'abandoned') THEN NOW() ELSE completed_at END
WHERE id = @id::uuid
RETURNING *;

-- name: AbandonSession :one
UPDATE user_sessions
SET status = 'abandoned', completed_at = NOW()
WHERE id = $1::uuid AND status = 'in_progress'
RETURNING *;

-- name: CreateUserChoice :one
INSERT INTO user_choices (id, session_id, step_id, choice_id)
VALUES (@id::uuid, @session_id::uuid, @step_id::uuid, @choice_id::uuid)
RETURNING *;

-- name: GetUserChoicesBySession :many
SELECT id, session_id, step_id, choice_id, created_at
FROM user_choices
WHERE session_id = $1::uuid
ORDER BY created_at;

-- name: ListSessionsByUserID :many
SELECT *
FROM user_sessions
WHERE user_id = $1::uuid
ORDER BY created_at;

-- name: GetUserSessionForUpdate :one
SELECT *
FROM user_sessions
WHERE id = $1::uuid
FOR UPDATE;