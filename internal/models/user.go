package models

import (
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
