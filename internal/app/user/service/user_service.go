package service

import (
	"context"
	"log"

	"github.com/Beretta350/golang-rest-template/internal/app/user/model"
	"github.com/Beretta350/golang-rest-template/internal/app/user/repository"
	"github.com/google/uuid"
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
	repo repository.UserMongoRepository
}

func NewUserService(repo repository.UserMongoRepository) UserService {
	return &userService{repo: repo}
}

// Create hashes the user's password and stores the user in the repository
func (s *userService) Create(ctx context.Context, user *model.User) error {
	log.Printf("package=service method=Create creating user with username: %v\n", user.Username)
	user.Id = uuid.NewString()
	err := user.Validate()
	if err != nil {
		logServiceError("Create", err)
		return err
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		logServiceError("Create", err)
		return err
	}
	user.Password = string(hashedPassword) // Store the hashed password

	err = s.repo.Create(ctx, user)
	if err != nil {
		logServiceError("Create", err)
		return err
	}

	log.Printf("package=service method=Create user %v created\n", user.Username)

	// Call the repository to create the user
	return err
}

// GetByID retrieves a user by ID and verifies the password
func (s *userService) GetByID(ctx context.Context, id string) (*model.User, error) {
	log.Printf("package=service method=GetByID getting user with ID: %v\n", id)
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		logServiceError("GetByID", err)
		return nil, err
	}

	return user, nil
}

// Update hashes the new password if it's provided, then updates the user
func (s *userService) Update(ctx context.Context, newUser *model.User) error {
	log.Printf("package=service method=Update updating user with ID: %v\n", newUser.Id)
	existentUser, err := s.GetByID(ctx, newUser.Id)
	if err != nil {
		logServiceError("Update", err)
		return err
	}

	if newUser.Password != "" { // Only hash if a new password is provided
		//Check if the password and username has all need
		err = newUser.Validate()
		if err != nil {
			logServiceError("Update", err)
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			logServiceError("Update", err)
			return err
		}
		newUser.Password = string(hashedPassword)
	} else {
		newUser.Password = existentUser.Password
	}

	err = s.repo.Update(ctx, newUser)
	if err != nil {
		logServiceError("Update", err)
		return err
	}

	log.Printf("package=service method=Update user %v updated\n", newUser.Id)
	return err
}

// Delete removes a user by ID
func (s *userService) Delete(ctx context.Context, id string) error {
	log.Println("package=service method=Delete deleting user with ID:", id)

	err := s.repo.Delete(ctx, id)
	if err != nil {
		logServiceError("Delete", err)
		return err
	}

	log.Printf("package=service method=Delete user %v deleted\n", id)
	return nil
}

// GetAll retrieves all users
func (s *userService) GetAll(ctx context.Context) ([]model.User, error) {
	log.Println("package=service method=GetAll getting all users")
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		logServiceError("GetAll", err)
		return nil, err
	}

	return users, nil
}

func logServiceError(method string, err error) {
	log.Printf("package=service method=%v %v\n", method, err)
}
