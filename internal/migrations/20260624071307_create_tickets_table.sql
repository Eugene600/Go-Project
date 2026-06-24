-- +goose Up
SELECT 'up SQL query';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE tickets (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    event_id UUID NOT NULL,

    ticket_name VARCHAR(100) NOT NULL,

    ticket_price NUMERIC(10,2) NOT NULL,

    quantity_created INTEGER NOT NULL CHECK (quantity_created >= 0),

    quantity_sold INTEGER NOT NULL DEFAULT 0
        CHECK (
            quantity_sold >= 0
            AND quantity_sold <= quantity_created
        ),

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_tickets_events
        FOREIGN KEY (event_id)
        REFERENCES events(id)
        ON DELETE CASCADE
);

-- +goose Down
SELECT 'down SQL query';

DROP TABLE tickets;
