package dtos

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type CreateUserRequest struct {
	FirstName   string    `json:"first_name" binding:"required"`
	MiddleName  *string   `json:"middle_name"`
	LastName    string    `json:"last_name" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	UserName    string    `json:"user_name" binding:"required"`
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"first_name"`
	MiddleName  *string   `json:"middle_name,omitempty"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	UserName    string    `json:"user_name"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
