-- +goose Up
SELECT 'up SQL query';
ALTER TABLE users
ADD COLUMN password_hash VARCHAR(255) NOT NULL;

-- +goose Down
SELECT 'down SQL query';

ALTER TABLE users
DROP COLUMN password_hash;
