package middleware

import (
	"github.com/gofiber/fiber/v2"
	error_usecase "github.com/little-tonii/gofiber-base/internal/usecase/error"
)

func NotFoundHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusNotFound, error_usecase.RouteNotFound)
	}
}
