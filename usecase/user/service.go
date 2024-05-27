package user

import (
	"context"

	"github.com/muhrifqii/curium_go_fiber/domain"
)

type UserRepository interface {
	GetByIdentifier(ctx context.Context, identifier string) (domain.User, error)
}

type Service struct {
	userRepository UserRepository
}

func NewService(userRepository UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}
