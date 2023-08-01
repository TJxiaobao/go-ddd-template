package cache

import (
	"time"
)

type Cache interface {
	Set(key string, value interface{}, duration time.Duration) error
	SetIfNotExist(key string, value interface{}, duration time.Duration) (bool, error)
	IncrBy(key string, value int64) (int64, error)
	Get(key string) (interface{}, error)
	Contains(key string) (bool, error)
	Remove(keys ...string) error
}
