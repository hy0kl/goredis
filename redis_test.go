package redisdao

import (
	"context"
	"testing"
	"time"

	"github.com/hy0kl/gconfig"
)

var (
	testRdsKey   = "redis_test_name"
	testRdsValue = `work`
)

func TestNewSimpleRedis(t *testing.T) {
	// load local conf.ini
	gconfig.SetConfigFile("./conf/conf.ini")

	ctx := context.Background()
	client := NewRedisClient("RedisCache")

	err := client.Set(ctx, testRdsKey, testRdsValue, 30*time.Second).Err()
	if err != nil {
		t.Error("fail")
	} else {
		t.Log("pass")
	}

	v := client.Get(context.Background(), testRdsKey).Val()
	if v == testRdsValue {
		t.Log("pass")
	} else {
		t.Error("fail")
	}

	delayClient := NewRedisClient("RedisStorage")
	err = delayClient.Set(ctx, testRdsKey, testRdsValue, 30*time.Second).Err()
	if err != nil {
		t.Error("fail")
	} else {
		t.Log("pass")
	}

	v = client.Get(context.Background(), testRdsKey).Val()
	if v == testRdsValue {
		t.Log("pass")
	} else {
		t.Error("fail")
	}
}
