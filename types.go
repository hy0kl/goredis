package goredis

type RedisInstance string

const (
	RedisInsCommon  RedisInstance = "common"
	RedisInsCache   RedisInstance = "cache"
	RedisInsStorage RedisInstance = "storage"
)

type RedisConfig struct {
	RedisIns RedisInstance

	// Options

	// host:port address.
	Addr string

	Username string
	Password string
	DB       int

	// Timeout for socket reads. If reached, commands will fail
	// with a timeout instead of blocking. Supported values:
	//   - `0` - default timeout (3 seconds).
	//   - `-1` - no timeout (block indefinitely).
	//   - `-2` - disables SetReadDeadline calls completely.
	ReadTimeout int
	// Timeout for socket writes. If reached, commands will fail
	// with a timeout instead of blocking.  Supported values:
	//   - `0` - default timeout (3 seconds).
	//   - `-1` - no timeout (block indefinitely).
	//   - `-2` - disables SetWriteDeadline calls completely.
	WriteTimeout int

	// Maximum number of socket connections.
	// Default is 10 connections per every available CPU as reported by runtime.GOMAXPROCS.
	PoolSize int

	// Maximum number of idle connections.
	// Default is 0. the idle connections are not closed by default.
	MaxIdleConns int
}
