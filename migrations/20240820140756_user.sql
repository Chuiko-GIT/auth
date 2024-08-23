-- +goose Up
CREATE TYPE users_role AS ENUM ('ADMIN', 'USER');

CREATE TABLE users
(
    id               SERIAL PRIMARY KEY,
    name             TEXT      NOT NULL,
    email            TEXT      NOT NULL UNIQUE,
    password         TEXT      NOT NULL,
    password_confirm TEXT      NOT NULL,
    role             users_role NOT NULL DEFAULT 'USER'::users_role,
    created_at       TIMESTAMP  NOT NULL DEFAULT now(),
    updated_at       TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;

DROP TYPE IF EXISTS users_role;
