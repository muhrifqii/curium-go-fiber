package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

func ActuatorHealthCheck() fiber.Handler {
	return healthcheck.New(healthcheck.Config{
		LivenessEndpoint:  "/healthy-and-wealthy",
		ReadinessEndpoint: "/up-up-and-ready",
	})
}
