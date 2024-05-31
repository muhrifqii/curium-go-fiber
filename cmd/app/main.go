package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhrifqii/curium_go_fiber/config"
	"github.com/redis/go-redis/v9"
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
	rdb    *redis.Client
}

func InitializeApp() *AppProvider {
	appConf := config.InitAppConfig()

	logger := InitializeLog(appConf)

	server := InitializeServer(logger)

	rdb, err := InitializeRedis(appConf)
	if err != nil {
		logger.Fatal("Could not connect to Redis", zap.Error(err))
	}

	return &AppProvider{
		log:    logger,
		server: server,
		rdb:    rdb,
	}
}

func InitializeLog(appConf config.AppConfig) *zap.Logger {
	var (
		logLevel   zapcore.Level
		stackLevel zapcore.Level = zapcore.ErrorLevel
		coreLogger zapcore.Core
	)

	logConf := appConf.LogConfig
	if appConf.DevMode {
		logLevel = zapcore.DebugLevel
	} else {
		logLevel = zapcore.InfoLevel
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

	jsonEncoder := zapcore.NewJSONEncoder(encoderCfg)

	loggerWritter := zapcore.NewCore(
		jsonEncoder,
		writer,
		zap.NewAtomicLevelAt(logLevel),
	)

	if appConf.DevMode {
		loggerConsole := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(logLevel),
		)
		coreLogger = zapcore.NewTee(
			loggerWritter,
			loggerConsole,
		)
	} else {
		coreLogger = loggerWritter
	}

	return zap.New(coreLogger, zap.AddCaller(), zap.AddStacktrace(stackLevel))
}

func InitializeRedis(appConf config.AppConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     appConf.RedisAddresss,
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}

func InitializeServer(logger *zap.Logger) *Server {
	apiConf := config.InitApiConfig()

	return NewServer(apiConf, logger)
}

func main() {
	app := InitializeApp()
	log := app.log

	defer log.Sync()
	defer app.rdb.Close()

	if err := app.server.Run(); err != nil {
		log.Fatal("Failed to start server:", zap.Error(err))
	}
}
