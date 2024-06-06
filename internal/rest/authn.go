package rest

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/api_error"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/dto"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/middleware"
	"github.com/muhrifqii/curium_go_fiber/internal/utils"
	"go.uber.org/zap"
)

type (
	AuthnService interface {
		Login(ctx context.Context, req dto.AuthnRequest) error
		RegisterByEmail(ctx context.Context, req dto.RegisterByEmailRequest) error
		Logout(ctx context.Context) error
	}

	AuthnHandler struct {
		authnService AuthnService
		validator    *validator.Validate
		log          *zap.Logger
	}
)

func NewAuthnHandler(router fiber.Router, svc AuthnService, params utils.HandlerParams) {
	handler := &AuthnHandler{
		authnService: svc,
		validator:    params.Validator,
		log:          params.Logger,
	}

	authnRoute := router.Group("", middleware.RateLimiter(10, params.Redis))
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
	// by default using register by email
	var req dto.RegisterByEmailRequest
	if err := c.BodyParser(&req); err != nil {
		return api_error.NewApiErrorResponse(fiber.StatusBadRequest, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return err
	}
	if err := h.authnService.RegisterByEmail(c.Context(), req); err != nil {
		return err
	}
	return dto.ReturnCreatedResponse[interface{}](c, nil)
}

func (h *AuthnHandler) Logout(c *fiber.Ctx) error {
	return nil
}

func (h *AuthnHandler) Refresh(c *fiber.Ctx) error {
	return nil
}
