package authn

import (
	"context"

	"github.com/muhrifqii/curium_go_fiber/usecase/user"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type (
	Service struct {
		userRepository user.UserRepository
		log            *zap.Logger
	}
)

func NewService(zap *zap.Logger, userRepository user.UserRepository) *Service {
	return &Service{
		log:            zap,
		userRepository: userRepository,
	}
}

// HashPassword hashes the password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a hashed password with a plaintext password.
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *Service) Login(ctx context.Context) error {
	return nil
}

func (s *Service) Register(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (s *Service) Logout(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}

func (s *Service) RefreshToken(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}
