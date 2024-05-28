package middleware

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.uber.org/zap"
)

func SetZapLogger(l *zap.Logger) {
	zapLogger := fiberzap.NewLogger(fiberzap.LoggerConfig{
		SetLogger: l,
		ExtraKeys: []string{"requestId"},
	})
	log.SetLogger(zapLogger)
}

func Logger(l *zap.Logger) fiber.Handler {
	return fiberzap.New(fiberzap.Config{
		Logger: l,
		FieldsFunc: func(c *fiber.Ctx) []zap.Field {
			return []zap.Field{
				zap.String("requestId", c.Context().UserValue("requestId").(string)),
			}
		},
	})
}
