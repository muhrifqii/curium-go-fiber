package config

import (
	"fmt"
	"os"
	"strconv"
)

type (
	AppConfig struct {
		DevMode       bool
		LogConfig     LogConfig
		RedisAddresss string
	}

	LogConfig struct {
		LogFileName    string
		LogFileMaxSize int
		LogFileMaxDays int
	}

	DbConfig struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}

	ApiConfig struct {
		Port            string
		ApiPrefix       string
		AllowedOrigins  string
		HeaderRequestID string
		JwtConfig       JwtConfig
	}

	JwtConfig struct {
		Secret                  string
		RefreshSecret           string
		Expiration              int
		RefreshExpirationInDays int
		CookieName              string
	}
)

const (
	ContextKeyAuthnToken = "authn_token"
	ContextKeyRequestID  = "request_id"
)

func InitAppConfig() AppConfig {
	isDev, err := strconv.ParseBool(os.Getenv("DEV"))
	if err != nil {
		isDev = false
	}

	logMaxSize, err := strconv.ParseInt(os.Getenv("LOG_FILE_MAX_SIZE"), 10, 32)
	if err != nil {
		logMaxSize = 10
	}
	logMaxDays, err := strconv.ParseInt(os.Getenv("LOG_FILE_MAX_DAYS"), 10, 32)
	if err != nil {
		logMaxDays = 60
	}

	return AppConfig{
		DevMode: isDev,
		LogConfig: LogConfig{
			LogFileName:    os.Getenv("LOG_FILE"),
			LogFileMaxSize: int(logMaxSize),
			LogFileMaxDays: int(logMaxDays),
		},
		RedisAddresss: "localhost:" + os.Getenv("REDIS_PORT"),
	}
}

func InitDbConfig() DbConfig {
	return DbConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASS"),
		Name:     os.Getenv("DATABASE_NAME"),
	}
}

func InitApiConfig() ApiConfig {
	return ApiConfig{
		Port:            ":" + os.Getenv("SERVER_PORT"),
		ApiPrefix:       os.Getenv("SERVER_API_PREFIX"),
		AllowedOrigins:  os.Getenv("SERVER_ALLOW_ORIGINS"),
		HeaderRequestID: os.Getenv("HEADER_REQ_ID"),
		JwtConfig:       InitJwtConfig(),
	}
}

func InitJwtConfig() JwtConfig {
	expiration, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		expiration = 5
	}
	refreshExpiration, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRATION_DAYS"))
	if err != nil {
		refreshExpiration = 14
	}
	return JwtConfig{
		Secret:                  os.Getenv("JWT_SECRET"),
		RefreshSecret:           os.Getenv("JWT_REFRESH_SECRET"),
		Expiration:              expiration,
		RefreshExpirationInDays: refreshExpiration,
		CookieName:              os.Getenv("JWT_COOKIE_NAME"),
	}
}

func (d DbConfig) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", d.User, d.Password, d.Host, d.Port, d.Name)
}
