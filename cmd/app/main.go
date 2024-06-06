package main

import (
	"context"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/muhrifqii/curium_go_fiber/cmd/server"
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
	server *server.Server
	rdb    *redis.Client
}

func InitializeApp() *AppProvider {
	appConf := config.InitAppConfig()

	logger := InitializeLog(appConf)
	zap.ReplaceGlobals(logger)

	validator := InitializeValidator()

	rdb, err := InitializeRedis(appConf)
	if err != nil {
		logger.Fatal("Could not connect to Redis", zap.Error(err))
	}

	db, err := InitializeDB(config.InitDbConfig())
	if err != nil {
		logger.Fatal("Could not connect to DB", zap.Error(err))
	}

	server := InitializeServer(logger, validator, rdb, db)

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

func InitializeValidator() *validator.Validate {
	return validator.New()
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

func InitializeServer(logger *zap.Logger, validator *validator.Validate, rdb *redis.Client, db *sqlx.DB) *server.Server {
	args := server.ServerArgs{
		Config:      config.InitApiConfig(),
		Logger:      logger,
		DB:          db,
		RedisClient: rdb,
		Validator:   validator,
	}

	return server.NewServer(args)
}

func InitializeDB(conf config.DbConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", conf.String())

	if err != nil {
		return nil, err
	}

	return db, nil
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
