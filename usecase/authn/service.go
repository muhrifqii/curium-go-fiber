package authn

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhrifqii/curium_go_fiber/domain"
	"github.com/muhrifqii/curium_go_fiber/internal/config"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/rest_utils"
	"github.com/muhrifqii/curium_go_fiber/internal/utils"
	"go.uber.org/zap"
)

type (
	UserRepository interface {
		GetByEmail(ctx context.Context, email string) (domain.User, error)
		GetByUsername(ctx context.Context, username string) (domain.User, error)
		GetByOrgEmail(c context.Context, orgID, email string) (domain.User, error)
		GetByOrgUsername(c context.Context, orgID, username string) (domain.User, error)

		CreateUser(ctx context.Context, user *domain.User) error

		IsOrgUserExistByIdentifier(c context.Context, orgID, email, username string) (bool, error)

		OnUserLoggedIn(ctx context.Context, id int64, time time.Time, ip, ua string) error
	}
)

type Service struct {
	userRepository UserRepository
	log            *zap.Logger
	jwtConf        config.JwtConfig
}

func NewService(zap *zap.Logger, jwtConf config.JwtConfig, userRepository UserRepository) *Service {
	return &Service{
		log:            zap,
		jwtConf:        jwtConf,
		userRepository: userRepository,
	}
}

func getOrgID(ctx context.Context) (string, error) {
	orgID, ok := ctx.Value("orgID").(string)
	if !ok || orgID == "" {
		return "", domain.ErrNotFound
	}
	return strings.ToLower(orgID), nil
}

func generateAccessToken(tokenExpiration int, userID int64, orgID, secret string) (string, time.Time, error) {
	accessTokenExpiry := time.Now().Add(time.Minute * time.Duration(tokenExpiration))
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"org": orgID,
		"iat": time.Now().Unix(),
		"exp": accessTokenExpiry.Unix(),
	})
	accessTokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return "", time.Time{}, err
	}
	return accessTokenString, accessTokenExpiry, nil
}

func generateRefreshToken(refreshTokenExpiration int, secret string) (string, time.Time, error) {
	refreshTokenExpiry := time.Now().Add(time.Hour * 24 * time.Duration(refreshTokenExpiration))
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": refreshTokenExpiry.Unix(),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", time.Time{}, err
	}
	return refreshTokenString, refreshTokenExpiry, nil
}

func (s *Service) Login(ctx context.Context, req domain.AuthnRequest) (domain.AuthnResponse, error) {
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
		return domain.AuthnResponse{}, domain.ErrInvalidCredentials
	}
	if err = utils.CheckPassword(user.Password, req.Password); err != nil {
		return domain.AuthnResponse{}, domain.ErrInvalidCredentials
	}

	accessTokenString, accessTokenExpiry, err := generateAccessToken(s.jwtConf.Expiration, user.ID, user.OrgID, s.jwtConf.Secret)
	if err != nil {
		return domain.AuthnResponse{}, err
	}

	refreshTokenString, refreshTokenExpiry, err := generateRefreshToken(s.jwtConf.RefreshExpirationInDays, s.jwtConf.RefreshSecret)
	if err != nil {
		return domain.AuthnResponse{}, err
	}

	return domain.AuthnResponse{
		AccessToken:           accessTokenString,
		AccessTokenExpiresAt:  accessTokenExpiry,
		RefreshToken:          refreshTokenString,
		RefreshTokenExpiresAt: refreshTokenExpiry,
	}, nil
}

func (s *Service) RegisterByEmail(ctx context.Context, req domain.RegisterByEmailRequest) error {
	orgID, err := getOrgID(ctx)
	if err != nil {
		return err
	}
	exist, err := s.userRepository.IsOrgUserExistByIdentifier(ctx, orgID, req.Email, req.Username)
	if err != nil {
		return err
	}
	if exist {
		return rest_utils.NewApiErrorResponse(http.StatusBadRequest, "User already exist")
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}
	user := &domain.User{
		BaseOrganizationModel: domain.BaseOrganizationModel{
			OrgID: orgID,
		},
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
