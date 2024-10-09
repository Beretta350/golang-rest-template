package service

import (
	"context"

	"github.com/Beretta350/golang-rest-template/internal/app/common/logging"
	"github.com/Beretta350/golang-rest-template/internal/app/user/model"
	"github.com/Beretta350/golang-rest-template/internal/app/user/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers(ctx context.Context) ([]model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, id string) error
}

type userService struct {
	repo repository.UserMongoRepository
}

func NewUserService(repo repository.UserMongoRepository) UserService {
	return &userService{repo: repo}
}

// GetAll retrieves all users
func (s *userService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	logging.LogService(ctx, "GetAllUsers", "attempting to retrieve all users")
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		logging.LogServiceError(ctx, "GetAllUsers", err)
		return nil, err
	}
	logging.LogService(ctx, "GetAllUsers", "successfully retrieved %d users", len(users))
	return users, nil
}

// GetUserByID retrieves a user by ID and verifies the password
func (s *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	logging.LogService(ctx, "GetUserByID", "attempting to retrieve user with ID: %v", id)
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		logging.LogServiceError(ctx, "GetUserByID", err)
		return nil, err
	}

	logging.LogService(ctx, "GetUserByID", "successfully retrieved user with ID: %v", id)
	return user, nil
}

// CreateUser hashes the user's password and stores the user in the repository
func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	logging.LogService(ctx, "CreateUser", "attempting to create a new user with username: %v", user.Username)
	user.Id = uuid.NewString()
	if err := user.Validate(); err != nil {
		logging.LogServiceError(ctx, "CreateUser", err)
		return err
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logging.LogServiceError(ctx, "CreateUser", err)
		return err
	}
	user.Password = string(hashedPassword) // Store the hashed password

	if err = s.repo.CreateUser(ctx, user); err != nil {
		logging.LogServiceError(ctx, "CreateUser", err)
		return err
	}

	user.Password = ""
	logging.LogService(ctx, "CreateUser", "user created successfully: %v", user.Username)
	return nil
}

// UpdateUser hashes the new password if it's provided, then updates the user
func (s *userService) UpdateUser(ctx context.Context, newUser *model.User) error {
	logging.LogService(ctx, "UpdateUser", "attempting to update user with ID: %v", newUser.Id)
	existentUser, err := s.GetUserByID(ctx, newUser.Id)
	if err != nil {
		logging.LogServiceError(ctx, "UpdateUser", err)
		return err
	}

	// Only hash if a new password is provided
	if newUser.Password != "" {
		if err = newUser.Validate(); err != nil {
			logging.LogServiceError(ctx, "UpdateUser", err)
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			logging.LogServiceError(ctx, "UpdateUser", err)
			return err
		}
		newUser.Password = string(hashedPassword)
		logging.LogService(ctx, "UpdateUser", "password updated for user with ID: %v", newUser.Id)
	} else {
		newUser.Password = existentUser.Password
	}

	if err = s.repo.UpdateUser(ctx, newUser); err != nil {
		logging.LogServiceError(ctx, "UpdateUser", err)
		return err
	}

	logging.LogService(ctx, "UpdateUser", "user updated successfully with ID: %v", newUser.Id)
	return nil
}

// DeleteUser removes a user by ID
func (s *userService) DeleteUser(ctx context.Context, id string) error {
	logging.LogService(ctx, "DeleteUser", "attempting to delete user with ID: %v", id)

	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		logging.LogServiceError(ctx, "DeleteUser", err)
		return err
	}

	logging.LogService(ctx, "DeleteUser", "user deleted successfully with ID: %v", id)
	return nil
}
