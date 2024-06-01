package server

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/muhrifqii/curium_go_fiber/config"
	"github.com/muhrifqii/curium_go_fiber/internal/repository"
	"github.com/muhrifqii/curium_go_fiber/internal/repository/postgresql"
	"github.com/muhrifqii/curium_go_fiber/internal/rest"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/api_error"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/middleware"
	"github.com/muhrifqii/curium_go_fiber/usecase/user"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type (
	Server struct {
		app  *fiber.App
		args ServerArgs
	}

	ServerArgs struct {
		Config      config.ApiConfig
		Logger      *zap.Logger
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
	redisStorage := repository.NewStorageRedis(args.RedisClient)

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
	apiV1 := app.Group(apiPath)
	publicApiV1 := app.Group(apiPath)

	// prepare repository layer
	userRepository := postgresql.NewUserRepository(nil)

	// build service layer
	svc := user.NewService(userRepository)
	rest.NewUserHandler(apiV1, svc)
	rest.NewAuthnHandler(publicApiV1, nil)

	return &Server{
		app:  app,
		args: args,
	}
}

func errorHandler(c *fiber.Ctx, err error) error {
	return api_error.ApiErrorResponseHandler(c, err)
}

func (s *Server) Run() error {
	return s.app.Listen(s.args.Config.Port)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
