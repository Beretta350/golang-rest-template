package errs

import (
	"fmt"
)

var (
	ErrUserNotFound    = NewCustomError(NotFound, "user not found", nil)
	ErrFindingUsers    = NewCustomError(InternalError, "error finding users", nil)
	ErrFindingUserByID = NewCustomError(InternalError, "error finding user by id", nil)
	ErrCreatingUser    = NewCustomError(InternalError, "error creating user", nil)
	ErrUpdatingUser    = NewCustomError(InternalError, "error updating user", nil)
	ErrDeletingUser    = NewCustomError(InternalError, "error deleting user", nil)
)

// CustomError represents a more detailed error with a message and code
type CustomError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Err     error     `json:"detail,omitempty"`
}

func (e *CustomError) SetDetail(err error) *CustomError {
	e.Err = err
	return e
}

// Error implements the error interface for CustomError
func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("error: %s, details: %v", e.Message, e.Err)
	}
	return fmt.Sprintf("error: %s", e.Message)
}

// Unwrap allows errors.Is and errors.As to work with CustomError
func (e *CustomError) Unwrap() error {
	return e.Err
}

// NewCustomError creates a new CustomError
func NewCustomError(code ErrorCode, message string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
