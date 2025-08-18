package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient     *redis.Client
	redisClientOnce sync.Once
)

func init() {
	redisClientOnce.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:            fmt.Sprintf("%s:%s", Env.REDIS_HOST, Env.REDIS_PORT_NUMBER),
			Password:        Env.REDIS_PASSWORD,
			DB:              0,
			Network:         "tcp",
			PoolSize:        20,
			MinIdleConns:    4,
			MaxIdleConns:    10,
			PoolTimeout:     30 * time.Second,
			ConnMaxIdleTime: 10 * time.Minute,
			ConnMaxLifetime: 30 * time.Minute,
			PoolFIFO:        true,
		})
		RedisClient = client
	})
}

func CloseRedisClient() {
	if err := RedisClient.Close(); err != nil {
		log.Println(err)
		return
	}
}
