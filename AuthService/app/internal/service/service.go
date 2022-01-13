package service

import (
	"context"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	SignUp(context.Context, *domain.UserCreateDTO) (*domain.User, error)
	SignIn(context.Context, *domain.UserSignIn) error
	GetById(context.Context, int) (*domain.User, error)
	GetByEmail(context.Context, string) (*domain.User, error)
	Update(context.Context, *domain.UserUpdateDTO) error
	Delete(context.Context, int) error
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
