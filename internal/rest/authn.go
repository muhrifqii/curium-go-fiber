package rest

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/domain"
	"github.com/muhrifqii/curium_go_fiber/internal/config"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/middleware"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/rest_utils"
	"go.uber.org/zap"
)

type (
	AuthnService interface {
		Login(ctx context.Context, req domain.AuthnRequest) (domain.AuthnResponse, error)
		RegisterByEmail(ctx context.Context, req domain.RegisterByEmailRequest) error
		Logout(ctx context.Context) error
	}

	AuthnHandler struct {
		authnService AuthnService
		validator    *validator.Validate
		log          *zap.Logger
		conf         config.JwtConfig
	}
)

func NewAuthnHandler(router fiber.Router, svc AuthnService, params rest_utils.HandlerParams, jwtConf config.JwtConfig) {
	handler := &AuthnHandler{
		authnService: svc,
		validator:    params.Validator,
		log:          params.Logger,
		conf:         jwtConf,
	}

	authnRoute := router.Group("")
	authnRoute.Post("/authenticate", handler.Login)
	authnRoute.Post("/register", handler.Register)

	protectedAuthnRoute := router.Group("", middleware.RequireAuthn(jwtConf))
	protectedAuthnRoute.Put("/authenticate", handler.Refresh)
	protectedAuthnRoute.Delete("/authenticate", handler.Logout)

}

func (h *AuthnHandler) Login(c *fiber.Ctx) error {
	var req domain.AuthnRequest
	if err := c.BodyParser(&req); err != nil {
		return rest_utils.NewApiErrorResponse(fiber.StatusBadRequest, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return err
	}
	response, err := h.authnService.Login(c.Context(), req)
	if err != nil {
		return err
	}
	c.Cookie(&fiber.Cookie{
		Name:     h.conf.CookieName,
		Value:    response.RefreshToken,
		Expires:  response.RefreshTokenExpiresAt,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     c.Path(),
	})
	return rest_utils.ReturnOkResponse(c, response)
}

func (h *AuthnHandler) Register(c *fiber.Ctx) error {
	// by default using register by email
	var req domain.RegisterByEmailRequest
	if err := c.BodyParser(&req); err != nil {
		return rest_utils.NewApiErrorResponse(fiber.StatusBadRequest, err.Error())
	}
	if err := h.validator.Struct(&req); err != nil {
		return err
	}
	if err := h.authnService.RegisterByEmail(c.Context(), req); err != nil {
		return err
	}
	return rest_utils.ReturnCreatedResponse[interface{}](c, nil)
}

func (h *AuthnHandler) Logout(c *fiber.Ctx) error {
	return nil
}

func (h *AuthnHandler) Refresh(c *fiber.Ctx) error {
	return nil
}
