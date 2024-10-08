package errs

import "net/http"

type ErrorCode string

const (
	NotFound      ErrorCode = "ERR404"
	InternalError ErrorCode = "ERR500"
)

func (ec ErrorCode) StatusCode() int {
	switch ec {
	case NotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError // Default to 500
	}
}
