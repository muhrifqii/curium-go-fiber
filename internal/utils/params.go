package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/muhrifqii/curium_go_fiber/internal/repository"
	"go.uber.org/zap"
)

type (
	HandlerParams struct {
		Validator *validator.Validate
		Redis     *repository.StorageRedis
		Logger    *zap.Logger
	}
)
