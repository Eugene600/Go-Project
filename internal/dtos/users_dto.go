package dtos

import "time"

type CreateUserRequest struct {
	FirstName   string    `json:"first_name" binding:"required"`
	MiddleName  *string   `json:"middle_name"`
	LastName    string    `json:"last_name" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	UserName    string    `json:"user_name" binding:"required"`
}
