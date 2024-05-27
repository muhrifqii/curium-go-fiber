package config

import (
	"os"
	"strconv"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type ApiConfig struct {
	Port           string
	ApiPrefix      string
	AllowedOrigins string
}

type JwtConfig struct {
	Secret     string
	Expiration int
	CookieName string
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
		Port:           ":" + os.Getenv("SERVER_PORT"),
		ApiPrefix:      os.Getenv("SERVER_API_PREFIX"),
		AllowedOrigins: os.Getenv("SERVER_ALLOW_ORIGINS"),
	}
}

func InitJwtConfig() JwtConfig {
	expiration, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION"))
	if err != nil {
		expiration = 5
	}
	return JwtConfig{
		Secret:     os.Getenv("JWT_SECRET"),
		Expiration: expiration,
		CookieName: os.Getenv("JWT_COOKIE_NAME"),
	}
}
