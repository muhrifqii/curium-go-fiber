package provision

import (
	"context"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/muhrifqii/curium_go_fiber/domain"
	"github.com/muhrifqii/curium_go_fiber/internal/utils"
	"go.uber.org/zap"
)

type (
	OrganizationRepository interface {
		GetByID(c context.Context, id string) (domain.Organization, error)
		CreateOrganization(c context.Context, org *domain.Organization) error
		UpdateOrganization(c context.Context, org *domain.Organization) error
	}

	UserRepository interface {
		GetByUsername(c context.Context, username string) (domain.User, error)
		CreateUser(ctx context.Context, user *domain.User) error
		IsOrgUserExistByIdentifier(c context.Context, orgID, email, username string) (bool, error)
	}

	OAuthClientRepository interface {
		Exists(c context.Context, clientID, clientSecret string) (bool, error)
		CreateClient(ctx context.Context, client *domain.OAuthClient) error
		UpdateClient(c context.Context, client *domain.OAuthClient) error
	}
)

type Service struct {
	organizationRepository OrganizationRepository
	userRepository         UserRepository
	clientRepository       OAuthClientRepository
	log                    *zap.Logger
}

func NewService(
	zap *zap.Logger,
	organizationRepository OrganizationRepository,
	userRepository UserRepository,
	clientRepository OAuthClientRepository,
) *Service {
	return &Service{
		organizationRepository: organizationRepository,
		userRepository:         userRepository,
		log:                    zap,
		clientRepository:       clientRepository,
	}
}

func (s *Service) CreateSystemUser(ctx context.Context, username, password string) error {
	orgID := "system"
	exist, err := s.userRepository.IsOrgUserExistByIdentifier(ctx, orgID, username+"@muhrifqii.com", username)
	if err != nil {
		return err
	} else if exist {
		return nil
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user := &domain.User{
		Username: username,
		Email:    username + "@muhrifqii.com",
		Password: hashedPassword,
		Status:   domain.UserStatusSystem,
	}
	user.OrgID = orgID
	return s.userRepository.CreateUser(ctx, user)
}

func (s *Service) CreateDefaultAuthClient(ctx context.Context, clientID, clientSecret, redirectUris, grantType string) error {
	if clientID == "" || clientSecret == "" || redirectUris == "" || grantType == "" {
		return nil
	}

	grantType = strings.Replace(grantType, ",", " ", -1)
	redirectUrisArray := strings.Split(redirectUris, ",")
	redirectUrisJson, err := sonic.MarshalString(redirectUrisArray)
	if err != nil {
		return err
	}
	req := &domain.OAuthClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		ClientName:   "Default Client - " + clientID,
		RedirectURIs: redirectUrisJson,
		GrantTypes:   grantType,
	}

	exist, err := s.clientRepository.Exists(ctx, clientID, clientSecret)
	if err != nil {
		return err
	}

	if !exist {
		return s.clientRepository.CreateClient(ctx, req)
	} else if (redirectUrisJson != req.RedirectURIs) || (grantType != req.GrantTypes) {
		s.log.Info("Updating client", zap.String("client_id", clientID))
		return s.clientRepository.UpdateClient(ctx, req)
	}

	return nil
}

func (s *Service) OnboardOrganization(ctx context.Context, req domain.CreateOrganizationRequest) error {
	ownerReq := req.Owner
	org := &domain.Organization{
		Identifier: req.Identifier,
		Name:       req.Name,
		Address:    req.Address,
		Email:      req.Email,
		Phone:      req.Phone,
	}
	owner := &domain.User{
		Username:  ownerReq.Username,
		Email:     ownerReq.Email,
		Password:  ownerReq.Password,
		FirstName: ownerReq.FirstName,
		LastName:  ownerReq.LastName,
		Phone:     ownerReq.Phone,
		Gender:    ownerReq.Gender,
	}

	if err := s.organizationRepository.CreateOrganization(ctx, org); err != nil {
		return err
	}
	if err := s.userRepository.CreateUser(ctx, owner); err != nil {
		return err
	}
	return nil
}
