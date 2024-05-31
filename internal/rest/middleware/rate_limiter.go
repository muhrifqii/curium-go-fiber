package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter(max int, storage fiber.Storage) fiber.Handler {
	return limiter.New(limiter.Config{
		Max:               max,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		Storage: storage,
	})
}

func RateLimiterWithKey(max int, storage fiber.Storage, keyFn func(*fiber.Ctx) string) fiber.Handler {
	if keyFn == nil {
		keyFn = func(c *fiber.Ctx) string {
			return c.IP()
		}
	}
	return limiter.New(limiter.Config{
		Max:               max,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		KeyGenerator:      keyFn,
		Storage:           storage,
	})
}
