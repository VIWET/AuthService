package repository

import (
	"github.com/VIWET/Beeracle/AuthService/internal/domain"
)

type UserRepository interface {
	Create(*domain.User) error
	GetById(int) (*domain.User, error)
	GetByEmail(string) (*domain.User, error)
	Update(*domain.User) error
	Delete(int) error
}

type CacheRepository interface {
	Set(string, *domain.RefreshSession) error
	Get(string) (*domain.RefreshSession, error)
	Delete(string) error
}

type Repositories struct {
	UserRepository UserRepository
	Cache          CacheRepository
}
