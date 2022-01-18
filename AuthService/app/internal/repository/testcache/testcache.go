package testcache

import (
	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/errors"
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
)

type TestCacheRepository struct {
	db map[string]*domain.RefreshSession
}

func NewTestCacheRepository() repository.CacheRepository {
	return &TestCacheRepository{
		db: make(map[string]*domain.RefreshSession),
	}
}

func (c *TestCacheRepository) Set(rt string, rs *domain.RefreshSession) error {
	c.db[rt] = rs
	return nil
}

func (c *TestCacheRepository) Get(rt string) (*domain.RefreshSession, error) {
	rs, ok := c.db[rt]
	if !ok {
		return nil, errors.ErrRecordNotFound
	}

	return rs, nil
}
