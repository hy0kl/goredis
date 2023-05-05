package goredis

import "github.com/redis/go-redis/v9"

func Use(instance RedisInstance) *redis.Client {
	return getInstance(instance)
}
