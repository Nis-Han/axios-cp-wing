-- +goose Up
CREATE TABLE IF NOT EXISTS user_verification (
    id UUID NOT NULL UNIQUE,
    user_id uuid NOT NULL,
    created_at timestamp NOT NULL,
    valid_till timestamp NOT NULL,
    verification_key varchar(32),
    PRIMARY KEY(id)
);

ALTER TABLE "tasks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

-- CREATE INDEX IF NOT EXISTS idx_tasks_created_by on tasks(created_by);
-- CREATE INDEX IF NOT EXISTS idx_tasks_platform on tasks(platform);

CREATE INDEX IF NOT EXISTS idx_user_verification_user_id on user_verification(user_id);
CREATE INDEX IF NOT EXISTS idx_user_verification_key on user_verification(verification_key);



-- +goose Down
DROP TABLE IF EXISTS tasks;
DROP INDEX IF EXISTS idx_user_verification_user_id;
DROP INDEX IF EXISTS idx_user_ verification_key;
