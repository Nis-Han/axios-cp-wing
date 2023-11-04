-- name: GetUser :one
SELECT * FROM users
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (
    id,
    email,
    hashed_password,
    salt,
    first_name,
    last_name,
    auth_token,
    is_admin_user
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING *;

-- name: GetUserAuthTokenwithEmail :one
SELECT * FROM users
WHERE users.auth_token = $1
AND users.email = $2;