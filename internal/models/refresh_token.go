package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofrs/uuid/v5"
)

type RefreshToken struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	TokenHash string
	ExpiresAt time.Time
	IssuedAt  time.Time
}

func (r *RefreshToken) CreateRefreshToken(tx *sql.Tx, ctx context.Context) error {
	query :=
		`
	INSERT INTO refresh_tokens (
		user_id,
		token_hash, 
		expires_at
	) 
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id)
	DO UPDATE SET 
		token_hash = EXCLUDED.token_hash, 
		expires_at = EXCLUDED.expires_at, 
		issued_at = NOW()
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		r.UserId,
		r.TokenHash,
		r.ExpiresAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *RefreshToken) DeleteRefreshToken(tx *sql.Tx, ctx context.Context) error {
	query :=
		`
	DELETE FROM refresh_tokens WHERE token_hash = $1 OR user_id = $2
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		r.TokenHash,
		r.UserId,
	)

	if err != nil {
		return err
	}

	return nil
}
