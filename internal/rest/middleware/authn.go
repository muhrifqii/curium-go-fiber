package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RequireAuthn() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return jwtAware.New()
	}
}

func RequireApiKey() fiber.Handler {

}
