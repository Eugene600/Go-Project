package models

import (
	"time"

	"github.com/gofrs/uuid/v5"
)


type Ticket struct {
	Id uuid.UUID
	EventId uuid.UUID
	Name string
	Price int32
	QuantityCreated int32
	QuantityRemaining int32
	QuantitySold int32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}