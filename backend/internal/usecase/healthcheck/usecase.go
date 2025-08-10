package healthcheck_usecase

import (
	"context"
)

type HealthcheckUsecase interface {
	CheckHeathStatus(ctx context.Context) (*GetHealthStatusResp, error)
}
