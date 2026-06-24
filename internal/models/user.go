package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	Id          uuid.UUID
	FirstName   string
	MiddleName  sql.NullString
	LastName    string
	DateOfBirth time.Time
	UserName    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

func (u *User) CreateUser(tx *sql.Tx, ctx context.Context) error {
	query := `
	INSERT INTO users (
		first_name,
		middle_name,
		last_name,
		date_of_birth,
		username
	)
	VALUES ($1, $2, $3, $4, $5)
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		u.FirstName,
		u.MiddleName,
		u.LastName,
		u.DateOfBirth,
		u.UserName,
	)

	if err != nil {
		return err
	}

	return nil
}
