package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Beretta350/golang-rest-template/internal/app/common/constants"
	"github.com/Beretta350/golang-rest-template/internal/app/user/model"
	"github.com/Beretta350/golang-rest-template/pkg/errs"
	"github.com/Beretta350/golang-rest-template/pkg/logging"
)

// UserRepository defines the methods for user data access.
type UserRepository interface {
	// GetAllUsers retrieves all users from the database.
	GetAllUsers(ctx context.Context) ([]model.User, error)
	// GetUserByID retrieves a user by their ID.
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	// CreateUser adds a new user to the database.
	CreateUser(ctx context.Context, user *model.User) error
	// UpdateUser modifies an existing user's details.
	UpdateUser(ctx context.Context, user *model.User) error
	// DeleteUser removes a user from the database.
	DeleteUser(ctx context.Context, id string) error
}

// userRepository is the SQL-based implementation.
type userRepository struct {
	db  *sql.DB
	log logging.Logger
}

// NewSQLUserRepository creates a new instance of userRepository.
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db, log: logging.GetLogger()}
}

func (r *userRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	rows, err := r.db.QueryContext(ctx, "SELECT username, created_at, updated_at FROM users")
	if err != nil {
		r.log.LogError(ctx, "repository", "GetAllUsers", err)
		return nil, errs.ErrFindingUsers.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Username, &user.CreatedAt, &user.UpdatedAt); err != nil {
			r.log.LogError(ctx, "repository", "GetAllUsers", err)
			return nil, errs.ErrFindingUserByID.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		r.log.LogError(ctx, "repository", "GetAllUsers", err)
		return nil, errs.ErrFindingUserByID.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
	}

	return users, nil
}

func (r *userRepository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.
		QueryRowContext(ctx, "SELECT id, username, created_at, updated_at FROM users WHERE id= $1", id).
		Scan(&user.Id, &user.Username, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		r.log.LogError(ctx, "repository", "GetUserByID", errs.ErrUserNotFound)
		return nil, errs.ErrUserNotFound
	} else if err != nil {
		r.log.LogError(ctx, "repository", "GetUserByID", err)
		return nil, errs.ErrFindingUserByID.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
	}

	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.db.ExecContext(ctx,
		"INSERT INTO users (id, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		user.Id, user.Username, user.Password, user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		r.log.LogError(ctx, "repository", "CreateUser", err)
		return errs.ErrCreatingUser.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
	}

	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	user.UpdatedAt = time.Now()

	_, err := r.db.ExecContext(ctx,
		"UPDATE users SET username = $1, password = $2, updated_at = $3 WHERE id = $4",
		user.Username, user.Password, user.UpdatedAt, user.Id,
	)
	if err != nil {
		r.log.LogError(ctx, "repository", "UpdateUser", err)
		return errs.ErrUpdatingUser.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	result, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		r.log.LogError(ctx, "repository", "DeleteUser", err)
		return errs.ErrDeletingUser.SetDetailFromString(constants.UnexpectedDatabaseErrorMessage)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		r.log.LogError(ctx, "repository", "DeleteUser", errs.ErrUserNotFound)
		return errs.ErrDeletingUser.SetDetailFromString(constants.NoUsersToDelete)
	}

	return nil
}
