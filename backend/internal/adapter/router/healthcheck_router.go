package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/little-tonii/gofiber-base/internal/adapter/handler"
)

type HealthCheckRouterConfig struct {
	BaseGroup          fiber.Router
	HealthCheckHandler *handler.HealthcheckHandler
}

func RegisterHealthCheckRouter(config *HealthCheckRouterConfig) {
	group := config.BaseGroup.Group("/healthcheck")
	{
		group.Get("", config.HealthCheckHandler.GetHealthStatus())
	}
}
