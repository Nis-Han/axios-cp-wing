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
    gen_random_uuid (),
    $1,
    current_timestamp,
    current_timestamp + interval '1 day',
    LEFT(MD5(random()::text), 32)
) RETURNING *;

-- name: GetUserVerificationEntryFromUserID :one
SELECT * FROM user_verification
WHERE user_id = $1;