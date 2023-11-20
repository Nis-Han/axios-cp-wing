-- +goose Up
CREATE TABLE IF NOT EXISTS platform (
    id UUID NOT NULL UNIQUE,
    created_by uuid NOT NULL,
    created_at timestamp NOT NULL,
    title text NOT NULL UNIQUE,
    paltform_link text NOT NULL,
    PRIMARY KEY(id)
);

ALTER TABLE "platform" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");

CREATE INDEX IF NOT EXISTS idx_platform_id on platform(id);
CREATE INDEX IF NOT EXISTS idx_platform_title on platform(title);


-- +goose Down
DROP TABLE IF EXISTS platform;
DROP INDEX IF EXISTS idx_platform_id;
DROP INDEX IF EXISTS idx_platform_title;
