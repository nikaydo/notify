-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    login TEXT UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE events (
    uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    Description TEXT NOT NULL,
    MaxMembers INT NOT NULL,
    CurrectMembers INT NOT NULL,
    Members INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE events;
-- +goose StatementEnd
