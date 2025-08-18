package handler

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/little-tonii/gofiber-base/internal/adapter/middleware"
	healthcheck_usecase "github.com/little-tonii/gofiber-base/internal/usecase/healthcheck"
)

type HealthcheckHandler struct {
	healthcheckUsecase healthcheck_usecase.HealthcheckUsecase
}

func NewHealthcheckHandler(healthcheckUsecase healthcheck_usecase.HealthcheckUsecase) *HealthcheckHandler {
	return &HealthcheckHandler{
		healthcheckUsecase: healthcheckUsecase,
	}
}

// GetHealthStatus 	godoc
//
// @Summary		get health status
// @Tags		healthcheck
// @Accept		json
// @Produce		json
// @Success		200	{object}	healthcheck_usecase.GetHealthStatusResp
// @Failure		408	{object}	error_usecase.ErrorResp
// @Failure		500	{object}	error_usecase.ErrorResp
// @Failure		503	{object}	error_usecase.ErrorResp
// @Router		/healthcheck [get]
func (handler *HealthcheckHandler) GetHealthStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return middleware.TimeoutWrapper(&middleware.TimeoutWrapperConfig[healthcheck_usecase.GetHealthStatusResp]{
			FiberCtx:       c,
			ExpectedStatus: fiber.StatusOK,
			Fn: func(ctx context.Context) (*healthcheck_usecase.GetHealthStatusResp, error) {
				resp, err := handler.healthcheckUsecase.CheckHeathStatus(ctx)
				return resp, err
			},
		})
	}
}
