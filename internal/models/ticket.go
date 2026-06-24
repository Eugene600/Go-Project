package models

import (
	"database/sql"
	"time"

	"github.com/gofrs/uuid/v5"
)

type Ticket struct {
	Id              uuid.UUID
	EventId         uuid.UUID
	TicketName      string
	Price           float64
	QuantityCreated int32
	QuantitySold    int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       sql.NullTime
}
