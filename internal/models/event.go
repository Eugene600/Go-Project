package models

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid/v5"
)


type Event struct {
	Id uuid.UUID
	EventName string
	EventLocation string
	EventStartDate time.Time
	EventEndDate time.Time
	OrganizerId uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}