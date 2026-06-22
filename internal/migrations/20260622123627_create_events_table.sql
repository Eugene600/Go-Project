-- +goose Up
SELECT 'up SQL query';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE events (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    event_name VARCHAR(100) NOT NULL,
    event_location VARCHAR(100) NOT NULL,

    event_start_date TIMESTAMPTZ(100) NOT NULL,
    event_end_date TIMESTAMPTZ(100) NOT NULL,

    organizer_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_events_users
        FOREIGN KEY (organizer_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

-- +goose Down
SELECT 'down SQL query';

DROP TABLE events;
