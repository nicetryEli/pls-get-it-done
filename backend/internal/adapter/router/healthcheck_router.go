package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/little-tonii/gofiber-base/internal/adapter/handler"
)

func RegisterHealthCheckRouter(baseGroup fiber.Router, healthcheckHandler *handler.HealthcheckHandler) {
	group := baseGroup.Group("/healthcheck")
	{
		group.Get("", healthcheckHandler.GetHealthStatus())
	}
}
