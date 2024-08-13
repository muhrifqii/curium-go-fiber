package rest_utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/domain"
)

func ReturnOkResponse[T interface{}](c *fiber.Ctx, data T) error {
	r := domain.ApiResponse[T]{
		Status:  "success",
		Message: "",
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(r)
}

func ReturnCreatedResponse[T interface{}](c *fiber.Ctx, data T) error {
	r := domain.ApiResponse[T]{
		Status:  "success",
		Message: "",
		Data:    data,
	}
	return c.Status(fiber.StatusCreated).JSON(r)
}
