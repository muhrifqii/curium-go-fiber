package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/muhrifqii/curium_go_fiber/config"
)

func Cors(conf config.ApiConfig) fiber.Handler {
	return cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     conf.AllowedOrigins,
	})
}
