package service

import (
	"context"

	"github.com/Beretta350/golang-rest-template/internal/app/user/model"
	"github.com/Beretta350/golang-rest-template/internal/app/user/repository"
	"github.com/Beretta350/golang-rest-template/pkg/logging"
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

var log logging.Logger = logging.GetLogger()

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

// GetAll retrieves all users
func (s *userService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	log.LogInternal(ctx, "service", "GetAllUsers", "attempting to retrieve all users")
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		log.LogError(ctx, "service", "GetAllUsers", err)
		return nil, err
	}
	log.LogInternal(ctx, "service", "GetAllUsers", "successfully retrieved %d users", len(users))
	return users, nil
}

// GetUserByID retrieves a user by ID and verifies the password
func (s *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	log.LogInternal(ctx, "service", "GetUserByID", "attempting to retrieve user with ID: %v", id)
	if err := uuid.Validate(id); err != nil {
		log.LogError(ctx, "service", "GetUserByID", err)
		return nil, err
	}

	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		log.LogError(ctx, "service", "GetUserByID", err)
		return nil, err
	}

	log.LogInternal(ctx, "service", "GetUserByID", "successfully retrieved user with ID: %v", id)
	return user, nil
}

// CreateUser hashes the user's password and stores the user in the repository
func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	log.LogInternal(ctx, "service", "CreateUser", "attempting to create a new user with username: %v", user.Username)
	user.Id = uuid.NewString()
	if err := user.Validate(); err != nil {
		log.LogError(ctx, "service", "CreateUser", err)
		return err
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.LogError(ctx, "service", "CreateUser", err)
		return err
	}
	user.Password = string(hashedPassword) // Store the hashed password

	if err = s.repo.CreateUser(ctx, user); err != nil {
		log.LogError(ctx, "service", "CreateUser", err)
		return err
	}

	user.Password = ""
	log.LogInternal(ctx, "service", "CreateUser", "user created successfully: %v", user.Username)
	return nil
}

// UpdateUser hashes the new password if it's provided, then updates the user
func (s *userService) UpdateUser(ctx context.Context, newUser *model.User) error {
	log.LogInternal(ctx, "service", "UpdateUser", "attempting to update user with ID: %v", newUser.Id)
	existentUser, err := s.GetUserByID(ctx, newUser.Id)
	if err != nil {
		log.LogError(ctx, "service", "UpdateUser", err)
		return err
	}

	// Only hash if a new password is provided
	if newUser.Password != "" {
		if err = newUser.Validate(); err != nil {
			log.LogError(ctx, "service", "UpdateUser", err)
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			log.LogError(ctx, "service", "UpdateUser", err)
			return err
		}
		newUser.Password = string(hashedPassword)
		log.LogInternal(ctx, "service", "UpdateUser", "password updated for user with ID: %v", newUser.Id)
	} else {
		if err = newUser.ValidateUsername(); err != nil {
			log.LogError(ctx, "service", "UpdateUser", err)
			return err
		}
		newUser.Password = existentUser.Password
	}

	if err = s.repo.UpdateUser(ctx, newUser); err != nil {
		log.LogError(ctx, "service", "UpdateUser", err)
		return err
	}

	log.LogInternal(ctx, "service", "UpdateUser", "user updated successfully with ID: %v", newUser.Id)
	return nil
}

// DeleteUser removes a user by ID
func (s *userService) DeleteUser(ctx context.Context, id string) error {
	log.LogInternal(ctx, "service", "DeleteUser", "attempting to delete user with ID: %v", id)
	if err := uuid.Validate(id); err != nil {
		log.LogError(ctx, "service", "DeleteUser", err)
		return err
	}

	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		log.LogError(ctx, "service", "DeleteUser", err)
		return err
	}

	log.LogInternal(ctx, "service", "DeleteUser", "user deleted successfully with ID: %v", id)
	return nil
}
