package domain

import "errors"

type ApiErrorResponse struct {
	ApiResponse[map[string]interface{}]
	StatusCode int `json:"-"`
}

func (e ApiErrorResponse) Error() string {
	return e.Message
}

var (
	ErrNotFound           = errors.New("error_not_found")
	ErrInvalidCredentials = errors.New("error_invalid_credentials")
	ErrInvalidToken       = errors.New("error_invalid_token")
	ErrInvalidClient      = errors.New("error_invalid_client")
)

var ErrorMessagesByStatus = map[string]string{
	"error":                     "Internal Server Error",
	"error_not_found":           "Not found",
	"error_invalid_credentials": "Invalid user/password",
	"error_invalid_token":       "Invalid token",
	"error_invalid_client":      "Invalid client",
}

var Errors = []error{
	ErrNotFound,
	ErrInvalidCredentials,
}
