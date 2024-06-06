package api_error

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/dto"
)

type ApiErrorResponse struct {
	dto.ApiResponse[map[string]string]
	StatusCode int `json:"-"`
}

func (e ApiErrorResponse) Error() string {
	return e.Message
}

func NewApiErrorResponse(statusCode int, message string, data ...map[string]string) *ApiErrorResponse {
	var additionalData map[string]string
	if len(data) == 0 {
		additionalData = nil
	} else {
		additionalData = data[0]
	}
	return &ApiErrorResponse{
		ApiResponse: dto.ApiResponse[map[string]string]{
			Status:  "error",
			Message: message,
			Data:    additionalData,
		},
		StatusCode: statusCode,
	}
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

func JwtErrorResponseHandler(c *fiber.Ctx, err error) error {
	var r = NewApiErrorResponse(fiber.StatusBadRequest, err.Error())
	if err.Error() == "Missing or malformed JWT" {
		return r
	}
	r.Message = "Unauthorized"
	r.StatusCode = fiber.StatusUnauthorized
	return r
}
