package dtos

import (
	"time"

	"github.com/Eugene600/Go-Project/internal/models"
	"github.com/gofrs/uuid/v5"
)

type UserRequest struct {
	FirstName   string    `json:"first_name" binding:"required"`
	MiddleName  *string   `json:"middle_name"`
	LastName    string    `json:"last_name" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth" binding:"required"`
	UserName    string    `json:"user_name" binding:"required"`
}

type UserResponse struct {
	ID          uuid.UUID  `json:"id"`
	FirstName   string     `json:"first_name"`
	MiddleName  *string    `json:"middle_name,omitempty"`
	LastName    string     `json:"last_name"`
	DateOfBirth time.Time  `json:"date_of_birth"`
	UserName    string     `json:"user_name"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func MappedUserResponse(user models.User) UserResponse {
	var middleName *string
	var deletedAt *time.Time

	if user.MiddleName.Valid {
		middleName = &user.MiddleName.String
	}

	if user.DeletedAt.Valid {
		deletedAt = &user.DeletedAt.Time
	}

	return UserResponse{
		ID:          user.Id,
		FirstName:   user.FirstName,
		MiddleName:  middleName,
		LastName:    user.LastName,
		DateOfBirth: user.DateOfBirth,
		UserName:    user.UserName,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}

func MappedUserResponseList(users []models.User) []UserResponse {
	responses := []UserResponse{}

	for _, user := range users {
		responses = append(responses, MappedUserResponse(user))
	}

	return responses
}
