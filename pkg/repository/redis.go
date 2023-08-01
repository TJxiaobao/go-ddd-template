package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type Redis struct {
	Client *redis.Client
}

type RedisOption struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Auth     string `json:"auth" yaml:"auth"`
	PoolSize int    `json:"pool_size" yaml:"pool_size"`
}

// NewRedisRepository 初始化Redis连接
func NewRedisRepository(opt *RedisOption) (*Redis, error) {
	client, err := initRedisClient(opt)
	if err != nil {
		return nil, err
	}
	return &Redis{
		Client: client,
	}, nil
}

func initRedisClient(opt *RedisOption) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", opt.Host, opt.Port),
		Password: opt.Auth,
		OnConnect: func(c context.Context, cn *redis.Conn) error {
			log.Info("Redis connection established.")
			return nil
		},
		PoolSize:     opt.PoolSize,
		MinIdleConns: opt.PoolSize,
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}
	return client, nil
}
