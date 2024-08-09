package auth

import (
	"context"
	"errors"
	"time"

	"github.com/aslam-ep/go-e-commerce/config"
	"github.com/aslam-ep/go-e-commerce/internal/user"
	"github.com/aslam-ep/go-e-commerce/utils"
)

// Service interface defines the methods required for authentication services.
type Service interface {
	// RegisterUser Creates a new user based on the provided request and returns the created user's details.
	RegisterUser(c context.Context, req *RegisterUserReq) (*user.User, error)

	// Authenticate checks the provided login credentials and returns a login response.
	Authenticate(ctx context.Context, req *LoginReq) (*LoginRes, error)

	// RefreshToken verifies the provided refresh token and issues a new access token.
	RefreshToken(ctx context.Context, req *RefreshTokenReq) (*RefreshTokenRes, error)
}

type service struct {
	userRepo user.Repository
	authRepo Repository
	timeout  time.Duration
	secret   string
}

// NewService creates a new instance of the authentication service.
func NewService(ur user.Repository, ar Repository) Service {
	return &service{
		userRepo: ur,
		authRepo: ar,
		timeout:  time.Duration(config.AppConfig.DBTimeout) * time.Second,
		secret:   config.AppConfig.JWTSecret,
	}
}

func (s *service) RegisterUser(c context.Context, req *RegisterUserReq) (*user.User, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &user.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Role:     req.Role,
		Password: hashedPassword,
	}

	createdUser, err := s.userRepo.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &user.User{
		ID:        createdUser.ID,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		Phone:     createdUser.Phone,
		Role:      createdUser.Role,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}

	return res, nil
}

func (s *service) Authenticate(c context.Context, req *LoginReq) (*LoginRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := utils.GenerateToken(user.ID, s.secret, time.Minute*15)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateToken(user.ID, s.secret, time.Hour*24*7)
	if err != nil {
		return nil, err
	}

	_, err = s.authRepo.Save(ctx, &RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
	})

	if err != nil {
		return nil, err
	}

	res := &LoginRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return res, nil
}

func (s *service) RefreshToken(c context.Context, req *RefreshTokenReq) (*RefreshTokenRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	refreshToken, err := s.authRepo.FindByToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetByID(ctx, int(refreshToken.UserID))
	if err != nil {
		return nil, err
	}

	newAccessToken, err := utils.GenerateToken(user.ID, s.secret, time.Minute*15)
	if err != nil {
		return nil, err
	}

	res := &RefreshTokenRes{
		AccessToken: newAccessToken,
	}

	return res, nil
}
