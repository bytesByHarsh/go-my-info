-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    name VARCHAR(100) NOT NULL,
    phone_num VARCHAR(15) NOT NULL,
    email VARCHAR(200) NOT NULL UNIQUE,
    username VARCHAR(200) NOT NULL UNIQUE,
    profile_img TEXT NOT NULL,
    role INTEGER NOT NULL,
    hashed_password VARCHAR(100) NOT NULL
);

-- +goose Down
DROP TABLE users;