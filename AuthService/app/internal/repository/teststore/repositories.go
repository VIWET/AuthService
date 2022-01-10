package teststore

import (
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
)

func NewTestRepositories() *repository.Repositories {
	return &repository.Repositories{
		UserRepository: NewTestUserRepository(),
	}
}
