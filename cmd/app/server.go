package main

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
)

type Server struct {
	App *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		CaseSensitive:            true,
		DisableHeaderNormalizing: true,
		JSONEncoder:              sonic.Marshal,
		JSONDecoder:              sonic.Unmarshal,
	})
	// api := app.Group("/api")

	// apiV1 := api.Group("/v1")

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(404)
	})

	return &Server{
		App: app,
	}
}

func notFound(c fiber.Ctx) error {
	return c.SendStatus(404)
}
