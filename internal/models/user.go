package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	Id          uuid.UUID
	FirstName   string
	MiddleName  string
	LastName    string
	DateOfBirth time.Time
	UserName    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
