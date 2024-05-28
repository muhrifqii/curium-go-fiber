package rest

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/domain"
)

type (
	// Representation of User's Usecases
	//
	//go:generate mockery --name ArticleService
	UserService interface {
		GetUserByIdentifier(ctx context.Context, identifier string) (domain.User, error)
		CreateUser(ctx context.Context, user domain.User) error
	}

	UserHandler struct {
		userService UserService
	}
)

func NewUserHandler(router fiber.Router, svc UserService) {
	handler := &UserHandler{
		userService: svc,
	}

	userRoute := router.Group("/user")
	userRoute.Get("/:identifier", handler.GetUserByIdentifier)
}

func (h *UserHandler) GetUserByIdentifier(c *fiber.Ctx) error {
	identifier := c.Params("identifier")
	user, err := h.userService.GetUserByIdentifier(c.Context(), identifier)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := domain.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	err := h.userService.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
