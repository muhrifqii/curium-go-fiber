package middleware

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/config"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/api_error"
)

func RequireAuthn(conf config.JwtConfig) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(conf.Secret),
		},
		ErrorHandler: api_error.JwtErrorResponseHandler,
	})
}

func RequireApiKey() fiber.Handler {
	return nil
}
