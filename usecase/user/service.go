package user

import (
	"context"
	"strings"

	"github.com/muhrifqii/curium_go_fiber/domain"
	"github.com/muhrifqii/curium_go_fiber/internal/repository"
)

type (
	Service struct {
		userRepository repository.UserRepository
	}
)

func NewService(userRepository repository.UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) GetUserByIdentifier(ctx context.Context, identifier string) (domain.User, error) {
	if strings.Contains(identifier, "@") {
		return s.userRepository.GetByEmail(ctx, identifier)
	}
	return s.userRepository.GetByUsername(ctx, identifier)
}

func (s *Service) CreateUser(ctx context.Context, user domain.User) error {
	return s.userRepository.CreateUser(ctx, &user)
}
