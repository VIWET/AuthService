package teststore

import (
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
	"github.com/VIWET/Beeracle/AuthService/internal/repository/testcache"
)

func NewTestRepositories() *repository.Repositories {
	return &repository.Repositories{
		UserRepository: NewTestUserRepository(),
		Cache:          testcache.NewTestCacheRepository(),
	}
}
