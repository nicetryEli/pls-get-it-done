package middleware

import (
	"slices"
	"strings"

	"github.com/gofiber/fiber/v2"
	error_usecase "github.com/little-tonii/gofiber-base/internal/usecase/error"
)

type AllowRolesConfig struct {
	Roles []string
}

func AllowRoles(config *AllowRolesConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if strings.Compare(config.Roles[0], "*") == 0 {
			return c.Next()
		}
		claims := c.Locals(CLAIMS).(*Claims)
		if slices.Contains(config.Roles, claims.UserRole) {
			return c.Next()
		}
		return fiber.NewError(fiber.StatusForbidden, error_usecase.Forbidden)
	}
}
