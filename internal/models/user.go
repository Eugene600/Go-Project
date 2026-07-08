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

func (u *User) GetUserByUsername(tx *sql.Tx, ctx context.Context, username string) error {
	query := `
	SELECT
		id,
		first_name,
		middle_name,
		last_name,
		date_of_birth,
		username,
		created_at,
		updated_at,
		deleted_at
	FROM users
	WHERE username = $1	
	`

	return tx.QueryRowContext(ctx, query, username).Scan(
		&u.Id,
		&u.FirstName,
		&u.MiddleName,
		&u.LastName,
		&u.DateOfBirth,
		&u.UserName,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	)
}

func GetAllUsers(tx *sql.Tx, ctx context.Context) ([]User, error) {
	query := `
	SELECT
		id,
		first_name,
		middle_name,
		last_name,
		date_of_birth,
		username,
		created_at,
		updated_at,
		deleted_at
	FROM users
	`

	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.MiddleName,
			&user.LastName,
			&user.DateOfBirth,
			&user.UserName,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
