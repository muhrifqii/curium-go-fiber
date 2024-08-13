package server

import (
	"context"
	"os"

	"github.com/bytedance/sonic"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
	"github.com/muhrifqii/curium_go_fiber/internal/config"
	"github.com/muhrifqii/curium_go_fiber/internal/repository/postgresql"
	redisInternal "github.com/muhrifqii/curium_go_fiber/internal/repository/redis"
	"github.com/muhrifqii/curium_go_fiber/internal/rest"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/middleware"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/rest_utils"
	"github.com/muhrifqii/curium_go_fiber/usecase/authn"
	"github.com/muhrifqii/curium_go_fiber/usecase/provision"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type (
	Server struct {
		app         *fiber.App
		args        ServerArgs
		provisioner *provision.Service
	}

	ServerArgs struct {
		Config      config.ApiConfig
		Logger      *zap.Logger
		Validator   *validator.Validate
		RedisClient *redis.Client
		DB          *sqlx.DB
	}
)

func NewServer(args ServerArgs) *Server {

	app := fiber.New(fiber.Config{
		CaseSensitive:            true,
		DisableHeaderNormalizing: true,
		JSONEncoder:              sonic.Marshal,
		JSONDecoder:              sonic.Unmarshal,
		ErrorHandler:             errorHandler,
	})

	// build redis client on fiber.Storage
	redisStorage := redisInternal.NewStorageRedis(args.RedisClient)

	// instantiate schema decoder
	encoderDecoder := rest_utils.SchemaEncoderDecoder{
		Encoder: schema.NewEncoder(),
		Decoder: schema.NewDecoder(),
	}

	// prepare middleware
	app.Use(middleware.Recover())
	app.Use(middleware.Cors(args.Config))
	app.Use(middleware.RequestID(args.Config))
	app.Use(middleware.Logger(args.Logger))
	app.Use(middleware.RateLimiter(50, redisStorage))
	app.Use(middleware.ActuatorHealthCheck())

	middleware.SetZapLogger(args.Logger)

	// prepare route group
	apiPath := args.Config.ApiPrefix + "/v1"
	// apiV1 := app.Group(apiPath)
	publicApiV1 := app.Group(apiPath)

	// prepare repository layer
	userRepository := postgresql.NewUserRepository(args.DB, args.Logger)
	orgRepository := postgresql.NewOrganizationRepository(args.DB, args.Logger)
	clientRepository := postgresql.NewOauthClientRepository(args.DB, args.Logger)

	// build service layer
	authnSvc := authn.NewService(args.Logger, args.Config.JwtConfig, userRepository)
	provisioner := provision.NewService(args.Logger, orgRepository, userRepository, clientRepository)
	handlerParams := rest_utils.HandlerParams{
		Validator:            args.Validator,
		Logger:               args.Logger,
		Redis:                redisStorage,
		SchemaEncoderDecoder: encoderDecoder,
	}
	rest.NewAuthnHandler(publicApiV1, authnSvc, handlerParams, args.Config.JwtConfig)

	return &Server{
		app:         app,
		args:        args,
		provisioner: provisioner,
	}
}

func errorHandler(c *fiber.Ctx, err error) error {
	return rest_utils.ApiErrorResponseHandler(c, err)
}

func (s *Server) ProvisionSystemUser() {
	migrationUsername := os.Getenv("SYSTEM_USER_MIGRATION_USERNAME")
	migrationPassword := os.Getenv("SYSTEM_USER_MIGRATION_PASSWORD")

	err := s.provisioner.CreateSystemUser(context.Background(), migrationUsername, migrationPassword)
	if err != nil {
		s.args.Logger.Error("Failed to create system user", zap.Error(err))
	}
}

func (s *Server) ProvisionDefaultAuthClient() {
	clientID := os.Getenv("DEFAULT_CLIENT_ID")
	clientSecret := os.Getenv("DEFAULT_CLIENT_SECRET")
	redirectUrisStr := os.Getenv("DEFAULT_CLIENT_REDIRECT_URIS")
	grantTypeStr := os.Getenv("DEFAULT_CLIENT_GRANT_TYPE")

	err := s.provisioner.CreateDefaultAuthClient(context.Background(), clientID, clientSecret, redirectUrisStr, grantTypeStr)
	if err != nil {
		s.args.Logger.Error("Failed to create default auth client", zap.Error(err))
	}
}

func (s *Server) Run() error {
	return s.app.Listen(s.args.Config.Port)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
