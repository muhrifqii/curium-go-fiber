package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/muhrifqii/curium_go_fiber/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
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

	logger := InitializeLog(appConf)

	server := InitializeServer(logger)

	return &AppProvider{
		log:    logger,
		server: server,
	}, nil
}

func InitializeLog(appConf config.AppConfig) *zap.Logger {
	var (
		logLevel   zapcore.Level
		stackLevel zapcore.Level
	)

	logConf := appConf.LogConfig
	if appConf.DevMode {
		logLevel = zapcore.DebugLevel
		stackLevel = zapcore.WarnLevel
	} else {
		logLevel = zapcore.InfoLevel
		stackLevel = zapcore.ErrorLevel
	}

	writer := zapcore.AddSync(&lumberjack.Logger{
		Filename: logConf.LogFileName,
		MaxSize:  logConf.LogFileMaxSize,
		MaxAge:   logConf.LogFileMaxDays,
		Compress: !appConf.DevMode,
	})

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.TimeKey = "timestamp"

	coreLogger := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		writer,
		zap.NewAtomicLevelAt(logLevel),
	)
	return zap.New(coreLogger, zap.AddCaller(), zap.AddStacktrace(stackLevel))
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
