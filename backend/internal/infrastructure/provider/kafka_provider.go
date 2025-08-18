package provider_impl

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaProviderImpl struct {
	kafkaProducer *kafka.Writer
}

func NewKafkaProviderImpl(kafkaProducer *kafka.Writer) *KafkaProviderImpl {
	return &KafkaProviderImpl{
		kafkaProducer: kafkaProducer,
	}
}

func (provider *KafkaProviderImpl) ProduceMessage(ctx context.Context, message *kafka.Message) error {
	return provider.kafkaProducer.WriteMessages(ctx, *message)
}

func (provider *KafkaProviderImpl) ProduceMessages(ctx context.Context, messages []*kafka.Message) error {
	msgs := make([]kafka.Message, 0, len(messages))
	for _, msg := range messages {
		msgs = append(msgs, *msg)
	}
	return provider.kafkaProducer.WriteMessages(ctx, msgs...)
}
