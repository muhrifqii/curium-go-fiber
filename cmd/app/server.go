package main

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/config"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/api_error"
)

type Server struct {
	App *fiber.App
}

func NewServer(conf config.ApiConfig) *Server {
	app := fiber.New(fiber.Config{
		CaseSensitive:            true,
		DisableHeaderNormalizing: true,
		JSONEncoder:              sonic.Marshal,
		JSONDecoder:              sonic.Unmarshal,
		ErrorHandler:             errorHandler,
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	return &Server{
		App: app,
	}
}

func errorHandler(c *fiber.Ctx, err error) error {
	return api_error.ApiErrorResponseHandler(c, err)
}
