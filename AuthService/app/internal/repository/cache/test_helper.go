package cache

import (
	"testing"

	"github.com/go-redis/redis/v7"
)

func TestRedisCache(t *testing.T, c *Config) *redis.Client {
	t.Helper()

	redis := redis.NewClient(&redis.Options{
		Addr:     c.GetAddr(),
		Password: c.Pwd,
		DB:       c.DB,
	})

	status := redis.Ping()
	if err := status.Err(); err != nil {
		t.Fatal(err)
	}

	return redis
}
