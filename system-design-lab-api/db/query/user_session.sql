-- name: CreateUserSession :one
INSERT INTO user_sessions (id, user_id, scenario_id, current_step_id, metrics, flags, status)
VALUES (@id::uuid, @user_id::uuid, @scenario_id::uuid, @current_step_id::uuid, @metrics, @flags, 'in_progress')
RETURNING *;

-- name: GetUserSession :one
SELECT * FROM user_sessions
WHERE id = $1::uuid;

-- name: UpdateUserSession :one
UPDATE user_sessions
SET current_step_id = @current_step_id,
    metrics         = @metrics,
    flags           = @flags,
    status          = @status
WHERE id = @id::uuid
RETURNING *;

-- name: CreateUserChoice :one
INSERT INTO user_choices (id, session_id, step_id, choice_id)
VALUES (@id::uuid, @session_id::uuid, @step_id::uuid, @choice_id::uuid)
RETURNING *;
