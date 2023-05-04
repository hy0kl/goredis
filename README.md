# goredis
goredis 基于 `github.com/go-redis/redis` 封装，将 `redis` 实例工厂化，便于接入项目使用。


## 安装
```shell script
go get github.com/hy0kl/goredis
```

## 配置
```ini
[RedisCache]
host = 127.0.0.1:6379
username =
password =
idleTimeout = 240
poolSize = 100
db = 0
```

## 使用示例

```go
import (
    "context"
    
    "github.com/go-redis/redis/v9"
    "github.com/hy0kl/goredis"
)

var ctx = context.Background()

func ExampleClient() {
    rdb := goredis.NewSimpleRedis("RedisCache")

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
}
```
