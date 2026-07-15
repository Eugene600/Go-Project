-- +goose Up
SELECT 'up SQL query';
CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    user_id UUID UNIQUE NOT NULL,

    token_hash TEXT NOT NULL,

    expires_at TIMESTAMPTZ NOT NULL,

    issued_at TIMESTAMPTZ DEFAULT NOW(),

    CONSTRAINT fk_refresh_tokens_users
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE refresh_tokens;
