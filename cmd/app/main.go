package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhrifqii/curium_go_fiber/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
		logLevel zapcore.Level
	)
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.TimeKey = "timestamp"

	if appConf.DevMode {
		logLevel = zapcore.DebugLevel
	} else {
		logLevel = zapcore.InfoLevel
	}

	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevel),
		Development:       appConf.DevMode,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths:       []string{"stdout"},
		ErrorOutputPaths:  []string{"stderr"},
		InitialFields:     map[string]interface{}{"pid": os.Getpid()},
	}
	return cfg.Build()
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
		log.Fatal("Failed to initialize server:", zap.Error(err))
	}
	if err := app.server.Run(); err != nil {
		log.Fatal("Failed to start server:", zap.Error(err))
	}
}
