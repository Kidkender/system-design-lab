-- name: CreateUser :one
INSERT INTO users (id, username, email, password)
VALUES ($1::uuid, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1::uuid;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;
