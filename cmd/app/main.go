package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/muhrifqii/curium_go_fiber/config"
	"go.uber.org/zap"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

type AppProvider struct {
	log    *zap.Logger
	server *Server
}

func InitializeApp() (*AppProvider, error) {
	appConf := config.InitAppConfig()

	logger, err := InitializeLog(appConf)
	if err != nil {
		return nil, err
	}

	server := InitializeServer(logger)

	return &AppProvider{
		log:    logger,
		server: server,
	}, nil
}

func InitializeLog(appConf config.AppConfig) (*zap.Logger, error) {
	var (
		logger *zap.Logger
		err    error
	)
	if appConf.DevMode {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	return logger, err
}

func InitializeServer(logger *zap.Logger) *Server {
	apiConf := config.InitApiConfig()

	return NewServer(apiConf, logger)
}

func main() {
	app, err := InitializeApp()
	log := app.log

	defer log.Sync()

	if err != nil {
		log.Fatal("Failed to initialize server: %v", zap.Error(err))
	}
	if err := app.server.Run(); err != nil {
		log.Fatal("Failed to start server: %v", zap.Error(err))
	}
}
