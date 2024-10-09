package errs

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrUserNotFound    = NewCustomError(NotFound, "user not found", nil)
	ErrFindingUsers    = NewCustomError(InternalError, "error finding users", nil)
	ErrFindingUserByID = NewCustomError(InternalError, "error finding user by id", nil)
	ErrCreatingUser    = NewCustomError(InternalError, "error creating user", nil)
	ErrUpdatingUser    = NewCustomError(InternalError, "error updating user", nil)
	ErrDeletingUser    = NewCustomError(InternalError, "error deleting user", nil)
	ErrValidatingUser  = NewCustomError(InvalidFormat, "invalid user", nil)
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

func (e *CustomError) SetDetailFromString(err string) *CustomError {
	e.Err = errors.New(err)
	return e
}

// Error implements the error interface for CustomError
func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("error: %s details: %v", e.Message, e.Err)
	}
	return fmt.Sprintf("error: %s", e.Message)
}

// Unwrap allows errors.Is and errors.As to work with CustomError
func (e *CustomError) Unwrap() error {
	return e.Err
}

func (e *CustomError) MarshalJSON() ([]byte, error) {
	type Alias CustomError // Create an alias to avoid recursion

	detail := make([]string, 0)
	if e.Err != nil {
		detail = strings.Split(e.Err.Error(), "\n")
	}

	// Marshal the error field as a string if it's not nil
	return json.Marshal(&struct {
		*Alias
		Err []string `json:"detail,omitempty"`
	}{
		Alias: (*Alias)(e),
		Err:   detail,
	})
}

// NewCustomError creates a new CustomError
func NewCustomError(code ErrorCode, message string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
