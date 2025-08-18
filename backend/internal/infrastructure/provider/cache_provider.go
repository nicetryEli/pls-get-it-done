package provider_impl

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheProviderImpl struct {
	client *redis.Client
}

func NewCacheProviderImpl(client *redis.Client) *CacheProviderImpl {
	return &CacheProviderImpl{client: client}
}

func (provider *CacheProviderImpl) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return provider.client.Set(ctx, key, value, expiration).Err()
}

func (provider *CacheProviderImpl) Get(ctx context.Context, key string) (string, error) {
	return provider.client.Get(ctx, key).Result()
}

func (provider *CacheProviderImpl) Delete(ctx context.Context, key string) error {
	return provider.client.Del(ctx, key).Err()
}

func (provider *CacheProviderImpl) Exists(ctx context.Context, key string) (bool, error) {
	val, err := provider.client.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return val > 0, nil
}

func (provider *CacheProviderImpl) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return provider.client.Expire(ctx, key, expiration).Err()
}

func (provider *CacheProviderImpl) Ping(ctx context.Context) error {
	return provider.client.Ping(ctx).Err()
}
