package authn

import (
	"context"
	"net/http"
	"strings"

	"github.com/muhrifqii/curium_go_fiber/domain"
	"github.com/muhrifqii/curium_go_fiber/internal/repository"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/api_error"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/dto"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type (
	Service struct {
		userRepository repository.UserRepository
		log            *zap.Logger
	}
)

func NewService(zap *zap.Logger, userRepository repository.UserRepository) *Service {
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

func (s *Service) Login(ctx context.Context, req dto.AuthnRequest) error {
	var (
		user domain.User
		err  error
	)

	if strings.Contains(req.Identifier, "@") {
		user, err = s.userRepository.GetByEmail(ctx, req.Identifier)
	} else {
		user, err = s.userRepository.GetByUsername(ctx, req.Identifier)
	}

	if err != nil {
		return err
	}
	if err = CheckPassword(user.Password, req.Password); err != nil {
		return api_error.NewApiErrorResponse(http.StatusNotFound, "Invalid user/password")
	}
	return nil
}

func (s *Service) RegisterByEmail(ctx context.Context, req dto.RegisterByEmailRequest) error {
	exist, err := s.userRepository.IsUserExistByIdentifier(ctx, req.Email, req.Username)
	if err != nil {
		return err
	}
	if exist {
		return api_error.NewApiErrorResponse(http.StatusBadRequest, "User already exist")
	}
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return err
	}
	user := &domain.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Status:   domain.UserStatusPending,
	}
	return s.userRepository.CreateUser(ctx, user)
}

func (s *Service) Logout(ctx context.Context) error {
	panic("not implemented") // TODO: Implement
}
