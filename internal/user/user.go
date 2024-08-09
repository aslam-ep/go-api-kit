package user

import (
	"time"
)

// User represents the user entity in the system.
type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// UpdateUserReq represents the request payload for updating user details.
type UpdateUserReq struct {
	ID    int64  `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,min=3,max=100"`
	Phone string `json:"phone" validate:"required,e164"`
	Role  string `json:"role" validate:"required,oneof=user vendor"`
}

// ResetPasswordReq represents the request payload for resetting a user's password.
type ResetPasswordReq struct {
	ID              int64  `json:"id"`
	CurrentPassword string `json:"current_password" validate:"required,min=6"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}
