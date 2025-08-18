package provider

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaProvider interface {
	ProduceMessage(ctx context.Context, message *kafka.Message) error
	ProduceMessages(ctx context.Context, messages []*kafka.Message) error
}
