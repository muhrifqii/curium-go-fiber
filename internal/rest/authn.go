package rest

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/middleware"
)

type (
	AuthnService interface {
		Login(ctx context.Context) error
		Register(ctx context.Context) error
		Logout(ctx context.Context) error
		RefreshToken(ctx context.Context) error
	}

	AuthnHandler struct {
		authnService AuthnService
	}
)

func NewAuthnHandler(router fiber.Router, svc AuthnService) {
	handler := &AuthnHandler{
		authnService: svc,
	}

	authnRoute := router.Group("", middleware.RateLimiter(10))
	authnRoute.Post("/authenticate", handler.Login)
	authnRoute.Post("/register", handler.Register)

	protectedAuthnRoute := router.Group("")
	protectedAuthnRoute.Put("/authenticate", handler.Refresh)
	protectedAuthnRoute.Delete("/authenticate", handler.Logout)

}

func (h *AuthnHandler) Login(c *fiber.Ctx) error {
	return nil
}

func (h *AuthnHandler) Register(c *fiber.Ctx) error {
	return nil
}

func (h *AuthnHandler) Logout(c *fiber.Ctx) error {
	return nil
}

func (h *AuthnHandler) Refresh(c *fiber.Ctx) error {
	return nil
}
