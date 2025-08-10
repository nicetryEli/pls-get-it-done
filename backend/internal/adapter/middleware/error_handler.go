package middleware

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	error_usecase "github.com/little-tonii/gofiber-base/internal/usecase/error"
)

func ErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		messages := strings.Split(err.Error(), ",")
		var e *fiber.Error
		if errors.As(err, &e) {
			code = e.Code
		}
		if code == fiber.StatusInternalServerError {
			messages[0] = error_usecase.InternalServerError
		}
		if len(messages) > 1 {
			return c.Status(code).JSON(error_usecase.ErrorsResp{
				Messages: messages,
			})
		}
		return c.Status(code).JSON(error_usecase.ErrorResp{
			Message: messages[0],
		})
	}
}
