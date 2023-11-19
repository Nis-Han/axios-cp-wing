-- name: CreateUserVerification :one
WITH deleted_rows AS (
    DELETE FROM user_verification
    WHERE user_id = $1
    RETURNING *
)
INSERT INTO user_verification (
    id,
    user_id,
    created_at,
    valid_till,
    verification_key
) VALUES (
    uuid_generate_v4(),
    $1,
    current_timestamp,
    current_timestamp + interval '1 day',
    encode(gen_random_bytes(32), 'hex')
) RETURNING *;
