package repository

import "github.com/VIWET/Beeracle/AuthService/internal/domain"

type UserRepository interface {
	Create(*domain.User) error
	GetById(int) (domain.User, error)
	GetByEmail(string) (domain.User, error)
	Update(*domain.User) error
	Delete(int) error
}

type BreweryRepository interface {
	Create(*domain.Brewery) error
	GetById(int) (domain.Brewery, error)
	GetByEmail(string) (domain.Brewery, error)
	Update(*domain.Brewery) error
	Delete(int) error
}

type Repositories struct {
	UserRepository    UserRepository
	BreweryRepository BreweryRepository
}
