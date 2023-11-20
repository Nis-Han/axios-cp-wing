-- +goose Up
CREATE TABLE IF NOT EXISTS user_verification (
    id UUID NOT NULL UNIQUE,
    user_id uuid NOT NULL,
    created_at timestamp NOT NULL,
    valid_till timestamp NOT NULL,
    verification_key varchar(32) NOT NULL,
    PRIMARY KEY(id)
);

ALTER TABLE "user_verification" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

CREATE INDEX IF NOT EXISTS idx_user_verification_user_id on user_verification(user_id);
CREATE INDEX IF NOT EXISTS idx_user_verification_key on user_verification(verification_key);



-- +goose Down
DROP TABLE IF EXISTS user_verification;
DROP INDEX IF EXISTS idx_user_verification_user_id;
DROP INDEX IF EXISTS idx_user_verification_key;
