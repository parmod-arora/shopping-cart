CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username text NOT NULL,
    password text NOT NULL,
    firstname text,
    lastname text,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone NOT NULL DEFAULT now()
);

ALTER TABLE users ADD UNIQUE (username);
CREATE INDEX users_username_idx ON users(username);
