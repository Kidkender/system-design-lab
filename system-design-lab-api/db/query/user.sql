-- name: CreateUser :one
INSERT INTO users (id, username, email)
VALUES ($1::uuid, $2, $3)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1::uuid;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;
