package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)


type Event struct {
	Id uuid.UUID
	Name string
	Location string
	StartTime time.Time
	EndTime time.Time
	OrganizerId uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}