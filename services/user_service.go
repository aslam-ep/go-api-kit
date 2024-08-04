package services

import (
	"context"
	"errors"
	"time"

	"github.com/aslam-ep/go-e-commerce/config"
	"github.com/aslam-ep/go-e-commerce/models"
	"github.com/aslam-ep/go-e-commerce/repositories"
	"github.com/aslam-ep/go-e-commerce/utils"
)

type UserService interface {
	CreateUser(c context.Context, req *models.CreateUserReq) (*models.UserRes, error)
	UpdateUser(c context.Context, req *models.UpdateUserReq) (*models.UserRes, error)
	GetUserById(c context.Context, id int) (*models.UserRes, error)
	ResetUserPassword(c context.Context, req *models.ResetPasswordReq) (*models.MessageRes, error)
	DeleteUser(c context.Context, id int) (*models.MessageRes, error)
}

type userService struct {
	userRepo repositories.UserRepository
	timeout  time.Duration
}

func NewUserService(ur repositories.UserRepository) UserService {
	return &userService{
		userRepo: ur,
		timeout:  time.Duration(config.AppConfig.DBTimeout) * time.Second,
	}
}

func (s *userService) CreateUser(c context.Context, req *models.CreateUserReq) (*models.UserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &models.User{
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

	res := &models.UserRes{
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

func (s *userService) GetUserById(c context.Context, id int) (*models.UserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &models.UserRes{
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return res, nil
}

func (s *userService) UpdateUser(c context.Context, req *models.UpdateUserReq) (*models.UserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// Check user exist before updating
	_, err := s.userRepo.GetByID(ctx, int(req.ID))
	if err != nil {
		return nil, err
	}

	u := &models.User{
		ID:    req.ID,
		Name:  req.Name,
		Phone: req.Phone,
		Role:  req.Role,
	}

	updatedUser, err := s.userRepo.Update(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &models.UserRes{
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		Phone:     updatedUser.Phone,
		Role:      updatedUser.Role,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return res, nil
}

func (s *userService) ResetUserPassword(c context.Context, req *models.ResetPasswordReq) (*models.MessageRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// Check user exist before updating password
	user, err := s.userRepo.GetByID(ctx, int(req.ID))
	if err != nil {
		return nil, err
	}

	// Check current user db password and given password match
	if !utils.CheckPasswordHash(req.CurrentPassword, user.Password) {
		return nil, errors.New("current password doesn't match")
	}

	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return nil, err
	}

	err = s.userRepo.ResetPassword(ctx, user, hashedPassword)
	if err != nil {
		return nil, err
	}

	res := &models.MessageRes{
		Success: true,
		Message: "Password updated,",
	}

	return res, nil
}

func (s *userService) DeleteUser(c context.Context, id int) (*models.MessageRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// Check user exist before updating password
	user, err := s.userRepo.GetByID(ctx, int(id))
	if err != nil {
		return nil, err
	}

	err = s.userRepo.Delete(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &models.MessageRes{
		Success: true,
		Message: "User Deleted.",
	}

	return res, nil
}
