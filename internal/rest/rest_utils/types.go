package rest_utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/schema"
	"go.uber.org/zap"
)

type SchemaEncoderDecoder struct {
	Encoder *schema.Encoder
	Decoder *schema.Decoder
}

type (
	HandlerParams struct {
		Logger               *zap.Logger
		Validator            *validator.Validate
		Redis                fiber.Storage
		SchemaEncoderDecoder SchemaEncoderDecoder
	}
)
