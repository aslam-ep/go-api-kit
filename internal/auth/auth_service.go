package auth

import (
	"context"
	"errors"
	"time"

	"github.com/aslam-ep/go-e-commerce/config"
	"github.com/aslam-ep/go-e-commerce/internal/user"
	"github.com/aslam-ep/go-e-commerce/utils"
)

type AuthService interface {
	Authenticate(ctx context.Context, req *LoginReq) (*LoginRes, error)
	RefreshToken(ctx context.Context, req *RefreshTokenReq) (*RefreshTokenRes, error)
}

type authService struct {
	userRepo user.UserRepository
	authRepo AuthRepository
	timeout  time.Duration
	secret   string
}

func NewAuthService(ur user.UserRepository, ar AuthRepository) AuthService {
	return &authService{
		userRepo: ur,
		authRepo: ar,
		timeout:  time.Duration(config.AppConfig.DBTimeout) * time.Second,
		secret:   config.AppConfig.JWTSecret,
	}
}

func (s *authService) Authenticate(c context.Context, req *LoginReq) (*LoginRes, error) {
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

func (s *authService) RefreshToken(c context.Context, req *RefreshTokenReq) (*RefreshTokenRes, error) {
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
