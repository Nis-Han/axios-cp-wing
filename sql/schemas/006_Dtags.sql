-- +goose Up
CREATE TABLE IF NOT EXISTS Dtags (
    id UUID NOT NULL UNIQUE,
    created_by uuid NOT NULL,
    created_at timestamp NOT NULL,
    title text NOT NULL UNIQUE,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS Dtags;
