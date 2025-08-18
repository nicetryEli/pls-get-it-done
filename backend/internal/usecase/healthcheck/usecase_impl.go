package healthcheck_usecase

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/little-tonii/gofiber-base/internal/domain/provider"
	error_usecase "github.com/little-tonii/gofiber-base/internal/usecase/error"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type HealthcheckUsecaseImpl struct {
	postgresTxProvider provider.TransactionProvider
	filestoreProvider  provider.FilestoreProvider
	cacheProvider      provider.CacheProvider
	kafkaProvider      provider.KafkaProvider
}

func NewHealthcheckUsecaseImpl(
	postgresTxProvider provider.TransactionProvider,
	filestoreProvider provider.FilestoreProvider,
	cacheProvider provider.CacheProvider,
	kafkaProvider provider.KafkaProvider,
) *HealthcheckUsecaseImpl {
	return &HealthcheckUsecaseImpl{
		postgresTxProvider: postgresTxProvider,
		filestoreProvider:  filestoreProvider,
		cacheProvider:      cacheProvider,
		kafkaProvider:      kafkaProvider,
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
	err = usecase.cacheProvider.Ping(ctx)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, error_usecase.ServiceUnavailable)
	}
	_, err = usecase.filestoreProvider.GetBucketNames(ctx)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, error_usecase.ServiceUnavailable)
	}
	err = usecase.kafkaProvider.ProduceMessage(ctx, &kafka.Message{
		Topic: "healthcheck",
		Key:   []byte("check"),
		Value: []byte("ping"),
	})
	if err != nil {
		return nil, fiber.NewError(fiber.StatusServiceUnavailable, error_usecase.ServiceUnavailable)
	}
	return &GetHealthStatusResp{
		Message: "pong",
	}, nil
}
