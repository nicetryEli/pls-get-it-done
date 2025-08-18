package healthcheck_usecase

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type HealthcheckConsumer struct {
	kafkaConsumer *kafka.Reader
	logger        *zap.Logger
}

func NewHealthcheckConsumer(kafkaConsumer *kafka.Reader, logger *zap.Logger) *HealthcheckConsumer {
	return &HealthcheckConsumer{
		kafkaConsumer: kafkaConsumer,
		logger:        logger,
	}
}

func (consumer *HealthcheckConsumer) Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			time.Sleep(time.Second)
			if err := consumer.kafkaConsumer.Close(); err != nil {
				log.Println(err)
			}
			return
		default:
			msg, err := consumer.kafkaConsumer.FetchMessage(ctx)
			if err != nil {
				log.Println(err)
				time.Sleep(time.Second)
				continue
			}
			clock := time.Now()
			if err := consumer.kafkaConsumer.CommitMessages(ctx, msg); err != nil {
				log.Println(err)
				time.Sleep(time.Second)
				continue
			}
			consumer.logger.Info(
				"consumer",
				zap.String("key", string(msg.Key)),
				zap.String("topic", msg.Topic),
				zap.Int64("offset", msg.Offset),
				zap.Int("partition", msg.Partition),
				zap.Duration("took", time.Since(clock)),
			)
		}
	}
}
