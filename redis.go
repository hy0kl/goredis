package redisdao

import "github.com/redis/go-redis/v9"

func NewRedisClient(instance string) *redis.Client {
	Init()
	return getInstance(instance)
}
