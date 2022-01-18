package cache

import (
	"encoding/json"
	"time"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
	"github.com/go-redis/redis/v7"
)

type redisCacheRepository struct {
	client *redis.Client
	exp    time.Duration
}

func NewRedisCacheRepository(c *redis.Client, exp time.Duration) repository.CacheRepository {
	return &redisCacheRepository{
		client: c,
		exp:    exp,
	}
}

func (c *redisCacheRepository) Set(rt string, rs *domain.RefreshSession) error {
	json, err := json.Marshal(rs)
	if err != nil {
		return err
	}

	status := c.client.Set(rt, json, c.exp)
	if err := status.Err(); err != nil {
		return err
	}

	return nil
}

func (c *redisCacheRepository) Get(rt string) (*domain.RefreshSession, error) {
	val, err := c.client.Get(rt).Result()
	if err != nil {
		return nil, err
	}

	rs := &domain.RefreshSession{}

	err = json.Unmarshal([]byte(val), rs)
	if err != nil {
		return nil, err
	}

	return rs, nil
}
