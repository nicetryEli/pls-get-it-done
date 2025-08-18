package config

import (
	"fmt"
	"sync"

	"github.com/segmentio/kafka-go"
)

var (
	KafkaProducer *kafka.Writer
	kafkaOnce     sync.Once
)

func init() {
	kafkaOnce.Do(func() {
		kafkaProducer := &kafka.Writer{
			Addr: kafka.TCP(
				fmt.Sprintf("%s:%s", Env.KAFKA_NODE_0_HOST, Env.KAFKA_NODE_0_PORT),
			),
			RequiredAcks:           kafka.RequireAll,
			AllowAutoTopicCreation: true,
			Compression:            kafka.Snappy,
			Balancer:               &kafka.RoundRobin{},
		}
		KafkaProducer = kafkaProducer
	})
}
