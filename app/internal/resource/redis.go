package resource

import (
	"context"
	"github.com/TJxiaobao/go-ddd-template/pkg/assert"
	"github.com/TJxiaobao/go-ddd-template/pkg/config"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"

	"github.com/TJxiaobao/go-ddd-template/pkg/repository"
	"github.com/go-redis/redis/v8"
)

var (
	redisCacheOnce              sync.Once
	singletonRedisCacheResource *RedisCache
)

type RedisCache struct {
	redisRepo *repository.Redis
}

func DefaultRedisCacheResource() *RedisCache {
	redisCacheOnce.Do(func() {
		singletonRedisCacheResource = &RedisCache{}
	})
	assert.NotNil(singletonRedisCacheResource)
	return singletonRedisCacheResource
}

func (c *RedisCache) MustOpen() {
	if c.redisRepo == nil {
		c.redisRepo = newRedisRepo()
	}
	assert.NotNil(c.redisRepo)
}

func (c *RedisCache) Close() {
	if c.redisRepo != nil && c.redisRepo.Client != nil {
		_ = c.redisRepo.Client.Close()
	}
}

func newRedisRepo() *repository.Redis {
	var (
		opt  repository.RedisOption
		repo *repository.Redis
	)

	err := config.GetConfig("redis", &opt)
	if err != nil {
		log.Errorf("init redis failed: %v", err)
		return nil
	}

	repo, err = repository.NewRedisRepository(&opt)
	if err != nil {
		log.Errorf("init redis failed: %v", err)
		return nil
	}

	return repo
}

func (c *RedisCache) Set(key string, value interface{}, duration time.Duration) error {
	return c.redisRepo.Client.Set(context.Background(), key, value, duration).Err()
}

func (c *RedisCache) SetIfNotExist(key string, value interface{}, duration time.Duration) (bool, error) {
	return c.redisRepo.Client.SetNX(context.Background(), key, value, duration).Result()
}

func (c *RedisCache) ZRange(key string, start, stop int64) ([]string, error) {
	return c.redisRepo.Client.ZRange(context.Background(), key, start, stop).Result()
}

func (c *RedisCache) ZAddNX(key string, members ...*redis.Z) error {
	return c.redisRepo.Client.ZAddNX(context.Background(), key, members...).Err()
}

func (c *RedisCache) IncrBy(key string, value int64) (int64, error) {
	return c.redisRepo.Client.IncrBy(context.Background(), key, value).Result()
}

func (c *RedisCache) Get(key string) (interface{}, error) {
	return c.redisRepo.Client.Get(context.Background(), key).Result()
}

func (c *RedisCache) Expire(key string, duration time.Duration) (bool, error) {
	return c.redisRepo.Client.Expire(context.Background(), key, duration).Result()
}

func (c *RedisCache) Contains(key string) (bool, error) {
	v, err := c.redisRepo.Client.Exists(context.Background(), key).Result()
	return v > 0, err
}

func (c *RedisCache) Remove(keys ...string) error {
	return c.redisRepo.Client.Del(context.Background(), keys...).Err()
}
