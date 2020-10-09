package hredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Kv struct {
	KeyPre string        // Key前缀
	Client *redis.Client // Redis Client
}

func (s *Kv) Set(ctx context.Context, key, value string, second int64) error {
	if ctx == nil {
		ctx = context.Background()
	}
	return s.Client.Set(ctx, s.KeyPre+key, value, time.Duration(second)*time.Second).Err()
}

func (s *Kv) Get(ctx context.Context, key string) string {
	if ctx == nil {
		ctx = context.Background()
	}
	return s.Client.Get(ctx, s.KeyPre+key).Val()
}

func (k *Kv) Equal(ctx context.Context, key string, val string) bool {
	if k.Get(ctx, key) == val {
		return true
	}
	return false
}

func (s *Kv) Inc(ctx context.Context, key string) int64 {
	if ctx == nil {
		ctx = context.Background()
	}
	return s.Client.Incr(ctx, s.KeyPre+key).Val()
}
