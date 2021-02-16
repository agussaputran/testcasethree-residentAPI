package helper

import "github.com/gomodule/redigo/redis"

// DelRedisCache func
func DelRedisCache() {
	pool := redis.NewPool(func() (redis.Conn, error) {
		return redis.Dial("tcp", "localhost:6379")
	}, 10)
	pool.MaxActive = 10

	conn := pool.Get()
	defer conn.Close()

	_, _ = conn.Do("DEL", "report_response")
}
