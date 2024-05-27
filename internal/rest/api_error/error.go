package api_error

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/dto"
)

type ApiErrorResponse struct {
	dto.ApiResponse[interface{}]
	StatusCode int `json:"-"`
}

func (e *ApiErrorResponse) Error() string {
	return e.Message
}

func ApiErrorResponseHandler(c *fiber.Ctx, err error) error {

	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	var r *ApiErrorResponse
	if errors.As(err, &r) {
		code = r.StatusCode
	} else {
		r = &ApiErrorResponse{}
		r.Status = "error"
		r.Message = err.Error()
		r.Data = nil

	}

	return c.Status(code).JSON(r)
}
