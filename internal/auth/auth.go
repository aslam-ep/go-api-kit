package auth

import "time"

// RegisterUserReq represents the request payload for creating a new user.
type RegisterUserReq struct {
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,e164"`
	Role     string `json:"role" validate:"required,oneof=user vendor"`
	Password string `json:"password" validate:"required,min=6"`
}

// RefreshToken represents a refresh token issued to a user for renewing access tokens.
type RefreshToken struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

// LoginReq represents the request payload for user login.
type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// LoginRes represents the response returned upon successful user login.
type LoginRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshTokenReq represents the request payload for refreshing an access token.
type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// RefreshTokenRes represents the response returned upon successful token refresh.
type RefreshTokenRes struct {
	AccessToken string `json:"access_token"`
}
