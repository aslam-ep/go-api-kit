package utils

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateToken generates a JWT token for a user with a specified expiration time.
func GenerateToken(userID int64, secret string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": strconv.Itoa(int(userID)),
		"exp":     time.Now().Add(expiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateToken validates a JWT token and returns the claims if valid.
func ValidateToken(tokenStr, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
