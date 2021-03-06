package service

import (
	"context"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/errors"
	"github.com/VIWET/Beeracle/AuthService/internal/jwt"
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	SignUp(ctx context.Context, dto *domain.UserCreateDTO, role string, ua string) (jwt.Tokens, error)
	SignIn(ctx context.Context, dto *domain.UserSignIn, ua string) (jwt.Tokens, error)
	Refresh(ctx context.Context, rt string, ua string, fp string) (jwt.Tokens, error)
	Update(ctx context.Context, dto *domain.UserUpdateDTO) error
	Delete(ctx context.Context, password string) error
}

type Services struct {
	User User
}

func NewServices(repos *repository.Repositories, tm jwt.TokenManager) *Services {
	return &Services{
		User: NewUserService(repos.UserRepository, tm, repos.Cache),
	}
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
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errors.ErrPasswordIsWrong
		}
		return err
	}
	return nil
}
