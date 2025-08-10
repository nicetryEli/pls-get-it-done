package healthcheck_usecase

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/little-tonii/gofiber-base/internal/domain/persistence"
	error_usecase "github.com/little-tonii/gofiber-base/internal/usecase/error"
	"gorm.io/gorm"
)

type HealthcheckUsecaseImpl struct {
	postgresTxProvider persistence.TransactionProvider
	filestoreProvider  persistence.FilestoreProvider
}

func NewHealthcheckUsecaseImpl(postgresTxProvider persistence.TransactionProvider, filestoreProvider persistence.FilestoreProvider) *HealthcheckUsecaseImpl {
	return &HealthcheckUsecaseImpl{
		postgresTxProvider: postgresTxProvider,
		filestoreProvider:  filestoreProvider,
	}
}

func (usecase *HealthcheckUsecaseImpl) CheckHeathStatus(ctx context.Context) (*GetHealthStatusResp, error) {
	err := usecase.postgresTxProvider.Transaction(ctx, func(tx *gorm.DB) error {
		if err := tx.Exec("SELECT 1").Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, error_usecase.ServiceUnavailable)
	}
	_, err = usecase.filestoreProvider.GetBucketNames(ctx)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, error_usecase.ServiceUnavailable)
	}
	return &GetHealthStatusResp{
		Message: "pong",
	}, nil
}
