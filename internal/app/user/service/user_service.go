package service

import (
	"context"

	"github.com/Beretta350/golang-rest-template/internal/app/user/model"
	"github.com/Beretta350/golang-rest-template/internal/app/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAll(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// Create hashes the user's password and stores the user in the repository
func (s *userService) Create(ctx context.Context, user *model.User) error {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword) // Store the hashed password

	// Call the repository to create the user
	return s.repo.Create(ctx, user)
}

// GetByID retrieves a user by ID and verifies the password
func (s *userService) GetByID(ctx context.Context, id string) (*model.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update hashes the new password if it's provided, then updates the user
func (s *userService) Update(ctx context.Context, user *model.User) error {
	if user.Password != "" { // Only hash if a new password is provided
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return s.repo.Update(ctx, user)
}

// Delete removes a user by ID
func (s *userService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// GetAll retrieves all users
func (s *userService) GetAll(ctx context.Context) ([]model.User, error) {
	return s.repo.GetAll(ctx)
}
