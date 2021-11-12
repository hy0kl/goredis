package redisdao

import (
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(instance string) *redis.Client {
	Init()
	return getInstance(instance)
}
