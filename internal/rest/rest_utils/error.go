package rest_utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/domain"
)

func NewApiErrorResponse(statusCode int, message string, data ...map[string]interface{}) *domain.ApiErrorResponse {
	var additionalData map[string]interface{}
	if len(data) == 0 {
		additionalData = nil
	} else {
		additionalData = data[0]
	}
	return &domain.ApiErrorResponse{
		ApiResponse: domain.ApiResponse[map[string]interface{}]{
			Status:  "error",
			Message: message,
			Data:    additionalData,
		},
		StatusCode: statusCode,
	}
}

func handleFiberAndInternalError(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}
	r := NewApiErrorResponse(code, err.Error())
	return c.Status(code).JSON(r)
}

func handleDomainError(c *fiber.Ctx, err error) (bool, error) {
	for _, e := range domain.Errors {
		if errors.Is(err, e) {
			r := NewApiErrorResponse(fiber.StatusBadRequest, domain.ErrorMessagesByStatus[err.Error()])
			r.Status = err.Error()
			return true, c.Status(r.StatusCode).JSON(r)
		}
	}
	return false, nil
}

func ApiErrorResponseHandler(c *fiber.Ctx, err error) error {

	var apiErr *domain.ApiErrorResponse
	if errors.As(err, &apiErr) {
		return c.Status(apiErr.StatusCode).JSON(apiErr)
	}

	if handled, err := handleDomainError(c, err); handled {
		return err
	}

	var valErr validator.ValidationErrors
	if errors.As(err, &valErr) {
		apiErr = NewApiErrorResponse(
			fiber.StatusBadRequest,
			"Validation error for the given request",
			make(map[string]interface{}),
		)
		apiErr.Status = "error_validation"
		for _, e := range valErr {
			apiErr.Data[e.Field()] = e.Tag()
		}
		return c.Status(apiErr.StatusCode).JSON(apiErr)
	}

	return handleFiberAndInternalError(c, err)
}

func JwtErrorResponseHandler(c *fiber.Ctx, err error) error {
	var r = NewApiErrorResponse(fiber.StatusBadRequest, err.Error())
	if err.Error() == "Missing or malformed JWT" {
		return r
	}
	r.Message = "Unauthorized"
	r.Status = "error_unauthorized"
	r.StatusCode = fiber.StatusUnauthorized
	return r
}
