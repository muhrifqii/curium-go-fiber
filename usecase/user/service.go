package user

import (
	"context"

	"github.com/muhrifqii/curium_go_fiber/domain"
)

type (
	UserRepository interface {
		GetByIdentifier(ctx context.Context, identifier string) (domain.User, error)
		CreateUser(ctx context.Context, user domain.User) error
	}

	Service struct {
		userRepository UserRepository
	}
)

func NewService(userRepository UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) GetUserByIdentifier(ctx context.Context, identifier string) (domain.User, error) {
	return s.userRepository.GetByIdentifier(ctx, identifier)
}

func (s *Service) CreateUser(ctx context.Context, user domain.User) error {
	return s.userRepository.CreateUser(ctx, user)
}
