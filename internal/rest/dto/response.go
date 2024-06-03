package dto

import "github.com/gofiber/fiber/v2"

type ApiResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func ReturnOkResponse[T interface{}](c *fiber.Ctx, data T) error {
	r := ApiResponse[T]{
		Status:  "success",
		Message: "",
		Data:    data,
	}
	return c.Status(fiber.StatusOK).JSON(r)
}

func ReturnCreatedResponse[T interface{}](c *fiber.Ctx, data T) error {
	r := ApiResponse[T]{
		Status:  "success",
		Message: "",
		Data:    data,
	}
	return c.Status(fiber.StatusCreated).JSON(r)
}
