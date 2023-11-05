-- name: CreateTask :one
INSERT INTO tasks (
    id,
    created_by,
    created_at,
    last_edited_by,
    last_edited_at,
    title,
    link,
    tags,
    platform

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