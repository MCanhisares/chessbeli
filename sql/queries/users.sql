-- name: CreateUser :one
INSERT INTO users(id, email, password, created_at, updated_at, username)
VALUES ($1, $2, crypt($3, gen_salt('bf')), $4, $5, $6)
RETURNING *;

-- name: Authenticate :one
SELECT id, username
FROM users
WHERE email = $1
AND password = crypt($2, password);

-- name: SelectUser :one
SELECT *
FROM users
WHERE email = $1;