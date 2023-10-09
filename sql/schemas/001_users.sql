-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL,
    email varchar(100) NOT NULL,
    hashed_password varchar(256) NOT NULL,
    salt varchar(100) NOT NULL,
    first_name varchar(100) NOT NULL,
    last_name varchar(100) NOT NULL,
    auth_token varchar(256) NOT NULL,
    is_admin_user boolean NOT NULL,
    PRIMARY KEY(id)
);

CREATE INDEX IF NOT EXISTS idx_issers_email ON users (email);

-- +goose Down
DROP TABLE IF EXISTS users;
DROP INDEX IF EXISTS idx_issers_email;
