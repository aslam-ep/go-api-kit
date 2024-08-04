package user

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"udapted_at,omitempty"`
}

type CreateUserReq struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164"`
	Role     string `json:"role" validate:"required,oneof=user vendor"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserRes struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"udapted_at"`
}

type UpdateUserReq struct {
	ID    int64  `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required,min=3,max=100"`
	Phone string `json:"phone" validate:"required,e164"`
	Role  string `json:"role" validate:"required,oneof=user vendor"`
}

type ResetPasswordReq struct {
	ID              int64  `json:"id"`
	CurrentPassword string `json:"current_password" validate:"required,min=6"`
	NewPassword     string `json:"new_password" validate:"required,min=6"`
}
