package hredis

import "github.com/go-redis/redis/v8"

// 示例redisURL: "redis://101.133.221.239:7777/0"
func NewClient(redisURL string) *redis.Client {
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opt)
}
