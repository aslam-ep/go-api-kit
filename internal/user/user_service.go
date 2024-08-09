package user

import (
	"context"
	"errors"
	"time"

	"github.com/aslam-ep/go-e-commerce/config"
	"github.com/aslam-ep/go-e-commerce/utils"
)

// Service interface for the user service
type Service interface {
	// UpdateUser Updates an existing user's information based on the provided request and returns the updated user's details.
	UpdateUser(c context.Context, req *UpdateUserReq) (*User, error)

	// GetUserById Retrieves a user's details by their ID.
	GetUserByID(c context.Context, id int) (*User, error)

	// ChangeUserPassword Resets the user's password based on the provided request and returns a message indicating success or failure.
	ChangeUserPassword(c context.Context, req *ResetPasswordReq) (*utils.MessageRes, error)

	// DeleteUser Deletes a user by their ID and returns a message indicating success or failure.
	DeleteUser(c context.Context, id int) (*utils.MessageRes, error)
}

type service struct {
	userRepo Repository
	timeout  time.Duration
}

// NewService initialize and return the Service
func NewService(ur Repository) Service {
	return &service{
		userRepo: ur,
		timeout:  time.Duration(config.AppConfig.DBTimeout) * time.Second,
	}
}

func (s *service) GetUserByID(c context.Context, id int) (*User, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return res, nil
}

func (s *service) UpdateUser(c context.Context, req *UpdateUserReq) (*User, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// Check user exist before updating
	_, err := s.userRepo.GetByID(ctx, int(req.ID))
	if err != nil {
		return nil, err
	}

	u := &User{
		ID:    req.ID,
		Name:  req.Name,
		Phone: req.Phone,
		Role:  req.Role,
	}

	updatedUser, err := s.userRepo.Update(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &User{
		ID:        updatedUser.ID,
		Name:      updatedUser.Name,
		Email:     updatedUser.Email,
		Phone:     updatedUser.Phone,
		Role:      updatedUser.Role,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	return res, nil
}

func (s *service) ChangeUserPassword(c context.Context, req *ResetPasswordReq) (*utils.MessageRes, error) {
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

	err = s.userRepo.ChangePassword(ctx, int(user.ID), hashedPassword)
	if err != nil {
		return nil, err
	}

	res := &utils.MessageRes{
		Success: true,
		Message: "Password updated,",
	}

	return res, nil
}

func (s *service) DeleteUser(c context.Context, id int) (*utils.MessageRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// Check user exist before updating password
	user, err := s.userRepo.GetByID(ctx, int(id))
	if err != nil {
		return nil, err
	}

	err = s.userRepo.Delete(ctx, int(user.ID))
	if err != nil {
		return nil, err
	}

	res := &utils.MessageRes{
		Success: true,
		Message: "User Deleted.",
	}

	return res, nil
}
