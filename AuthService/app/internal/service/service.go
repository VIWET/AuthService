package service

import (
	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	Create(*domain.UserCreateDTO) (*domain.User, error)
	GetById(int) (*domain.User, error)
	GetByEmail(string) (*domain.User, error)
	Update(*domain.UserUpdateDTO) error
	Delete(int) error
}

type Services struct {
	User User
}

func GeneratePasswordHash(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPassword(u *domain.User, p string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(p))
	if err != nil {
		return err
	}
	return nil
}
