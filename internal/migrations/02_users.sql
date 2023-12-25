-- +goose Up
CREATE TABLE if not exists myProject.users (
    user_id text NOT NULL,
    last_name text NOT NULL,
    first_name text NOT NULL,
    email text NOT NULL,
    phone text NOT NULL,
    PRIMARY KEY (user_id)
);