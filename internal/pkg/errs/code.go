package errs

import "net/http"

type ErrorCode string

const (
	InvalidFormat ErrorCode = "ERR400"
	NotFound      ErrorCode = "ERR404"
	InternalError ErrorCode = "ERR500"
)

func (ec ErrorCode) StatusCode() int {
	switch ec {
	case NotFound:
		return http.StatusNotFound
	case InvalidFormat:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError // Default to 500
	}
}
