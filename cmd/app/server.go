package main

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/muhrifqii/curium_go_fiber/config"
	"github.com/muhrifqii/curium_go_fiber/internal/repository/postgresql"
	"github.com/muhrifqii/curium_go_fiber/internal/rest"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/api_error"
	"github.com/muhrifqii/curium_go_fiber/internal/rest/middleware"
	"github.com/muhrifqii/curium_go_fiber/usecase/user"
	"go.uber.org/zap"
)

type Server struct {
	app    *fiber.App
	config config.ApiConfig
}

func NewServer(conf config.ApiConfig, logger *zap.Logger) *Server {
	app := fiber.New(fiber.Config{
		CaseSensitive:            true,
		DisableHeaderNormalizing: true,
		JSONEncoder:              sonic.Marshal,
		JSONDecoder:              sonic.Unmarshal,
		ErrorHandler:             errorHandler,
	})

	app.Use(middleware.Recover())
	app.Use(middleware.Cors(conf))
	app.Use(middleware.RequestID(conf))
	app.Use(middleware.Logger(logger))
	app.Use(middleware.RateLimiter(50))
	app.Use(middleware.ActuatorHealthCheck())

	middleware.SetZapLogger(logger)

	apiPath := conf.ApiPrefix + "/v1"
	apiV1 := app.Group(apiPath)
	publicApiV1 := app.Group(apiPath)

	// prepare repository layer
	userRepository := postgresql.NewUserRepository(nil)

	// build service layer
	svc := user.NewService(userRepository)
	rest.NewUserHandler(apiV1, svc)
	rest.NewAuthnHandler(publicApiV1, nil)

	return &Server{
		app:    app,
		config: conf,
	}
}

func errorHandler(c *fiber.Ctx, err error) error {
	return api_error.ApiErrorResponseHandler(c, err)
}

func (s *Server) Run() error {
	return s.app.Listen(s.config.Port)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
