package sqlstore

import (
	"database/sql"
	"time"

	"github.com/VIWET/Beeracle/AuthService/internal/repository"
	"github.com/VIWET/Beeracle/AuthService/internal/repository/cache"
	"github.com/go-redis/redis/v7"
)

func NewRepositories(db *sql.DB, c *redis.Client, exp time.Duration) *repository.Repositories {
	return &repository.Repositories{
		UserRepository: NewUserRepository(db),
		Cache:          cache.NewRedisCacheRepository(c, exp),
	}
}
