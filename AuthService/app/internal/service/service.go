package service

import (
	"context"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/jwt"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	SignUp(context.Context, *domain.UserCreateDTO, string) (*domain.User, jwt.Tokens, error)
	// SignIn(context.Context, *domain.UserSignIn) (string, string, error)
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
