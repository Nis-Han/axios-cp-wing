-- +goose Up
CREATE TABLE IF NOT EXISTS tasks (
    id UUID NOT NULL UNIQUE,
    created_by varchar(50) NOT NULL,
    created_at timestamp NOT NULL,
    last_edited_by varchar(50) NOT NULL,
    last_edited_at timestamp NOT NULL,
    title text NOT NULL UNIQUE,
    link text NOT NULL,
    tags text[],
    platform varchar(50) NOT NULL
);

CREATE INDEX IF NOT EXISTS idx_tasks_created_by on tasks(created_by);
CREATE INDEX IF NOT EXISTS idx_tasks_platform on tasks(platform);


-- +goose Down
DROP TABLE IF EXISTS tasks;
DROP INDEX IF EXISTS idx_tasks_created_by;
DROP INDEX IF EXISTS idx_tasks_platform;
