package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type LoggingConfig struct {
	Logger *zap.Logger
}

func LoggingMiddleware(config *LoggingConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		clock := time.Now()
		err := c.Next()
		latency := time.Since(clock)
		method := c.Method()
		statusCode := c.Response().StatusCode()
		urlPath := c.OriginalURL()

		if err != nil {
			if fiberErr, ok := err.(*fiber.Error); ok {
				statusCode = fiberErr.Code
			}
		}

		config.Logger.Info(
			"http",
			zap.String("method", method),
			zap.Int("status", statusCode),
			zap.Duration("latency", latency),
			zap.String("path", urlPath),
			zap.Error(err),
		)

		return err
	}
}
