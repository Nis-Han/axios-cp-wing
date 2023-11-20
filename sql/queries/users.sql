-- name: GetUserFromEmail :one
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
    is_admin_user,
    verified_user
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
) RETURNING *;

-- name: GetUserFromAuthToken :one
SELECT * FROM users
WHERE users.auth_token = $1;

-- name: GetAllAdminUsers :many
SELECT 
    email, first_name, last_name 
FROM users
WHERE
    is_admin_user = TRUE;

-- name: GetAllUsers :many
SELECT 
    email, first_name, last_name 
FROM users;

-- name: EditAdminAccess :one
UPDATE users
SET is_admin_user = $1
WHERE email = $2
RETURNING *;

-- name: SetUserVerificationTrue :one
UPDATE users
SET verified_user = TRUE
WHERE id = $1
RETURNING *;