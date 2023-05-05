package goredis

import (
	"context"
	"testing"
	"time"
)

var (
	testRdsKey   = "redis_test_name"
	testRdsValue = `work`
)

func TestNewSimpleRedis(t *testing.T) {
	ctx := context.Background()

	setup := []RedisConfig{
		{
			RedisIns: RedisInsCache,
			Addr:     `127.0.0.1:6379`,
		},
		{
			RedisIns: RedisInsStorage,
			Addr:     `127.0.0.1:6379`,
			DB:       1,
		},
	}

	Initialize(setup...)

	client := Use(RedisInsCache)
	err := client.Set(ctx, testRdsKey, testRdsValue, 30*time.Second).Err()
	if err != nil {
		t.Error("fail")
	} else {
		t.Log("pass")
	}

	v := client.Get(ctx, testRdsKey).Val()
	if v == testRdsValue {
		t.Log("pass")
	} else {
		t.Error("fail")
	}

	delayClient := Use(RedisInsStorage)
	err = delayClient.Set(ctx, testRdsKey, testRdsValue, 30*time.Second).Err()
	if err != nil {
		t.Error("fail")
	} else {
		t.Log("pass")
	}

	v = client.Get(ctx, testRdsKey).Val()
	if v == testRdsValue {
		t.Log("pass")
	} else {
		t.Error("fail")
	}
}
