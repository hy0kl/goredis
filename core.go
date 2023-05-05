package goredis

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

func newRedisClient(option redis.Options) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     option.Addr,
		Username: option.Username,
		Password: option.Password, // no password set
		DB:       option.DB,       // use default DB
		PoolSize: option.PoolSize,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("get redis connect error: %v, server: %v ", err, option.Addr)
	}

	return client
}

var instance = struct {
	sync.RWMutex
	redisInstances map[RedisInstance]*redis.Client
}{
	redisInstances: make(map[RedisInstance]*redis.Client, 0),
}

// Initialize 此方法必须在 main 函数中显示调用
func Initialize(setup ...RedisConfig) {
	instance.redisInstances = make(map[RedisInstance]*redis.Client, 0)

	for _, conf := range setup {
		options := redis.Options{
			PoolSize:    100,
			ReadTimeout: 5 * time.Second,
		}

		if conf.Addr == "" {
			panic(fmt.Sprintf(`lost required redis addr, conf: %#v`, conf))
		}

		options.Addr = conf.Addr
		options.Username = conf.Username
		options.Password = conf.Password

		if conf.PoolSize > 0 {
			options.PoolSize = conf.PoolSize
		}

		if conf.DB > 0 {
			options.DB = conf.DB
		}

		instance.redisInstances[conf.RedisIns] = newRedisClient(options)
	}
}

func getInstance(server RedisInstance) *redis.Client {
	instance.RLock()
	ins, ok := instance.redisInstances[server]
	instance.RUnlock()

	if ok && ins != nil {
		return ins
	} else {
		panic(fmt.Sprintf(`redis instance not found, server: %v`, server))
	}

	return nil
}
