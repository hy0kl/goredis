package redisdao

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/hy0kl/gconfig"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
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

var initOnce sync.Once

var instance = struct {
	sync.RWMutex
	redisInstances map[string]*redis.Client
}{redisInstances: make(map[string]*redis.Client, 0)}

func Init() {
	initOnce.Do(func() {
		initRedis()
	})
}

func initRedis() {
	instance.redisInstances = make(map[string]*redis.Client, 0)

	var redisGroup = []string{`RedisCache`, `RedisStorage`}
	for _, section := range redisGroup {
		var redisConf = gconfig.GetConfStringMap(section)
		options := redis.Options{
			PoolSize:    100,
			ReadTimeout: 5 * time.Second,
		}

		host := redisConf["host"]
		if host == "" {
			panic(fmt.Sprintf(`lost required host, section: %s`, section))
		}

		options.Addr = host
		options.Username = redisConf["username"]
		options.Password = redisConf["password"]

		poolSize := cast.ToInt(redisConf["poolSize"])
		if poolSize > 0 {
			options.PoolSize = poolSize
		}

		db := cast.ToInt(redisConf["db"])
		if db > 0 {
			options.DB = db
		}

		instance.redisInstances[section] = newRedisClient(options)
	}
}

func getInstance(server string) *redis.Client {
	instance.RLock()
	ins, ok := instance.redisInstances[server]
	instance.RUnlock()
	if ok && ins != nil {
		return ins
	}

	return nil
}
