package service

import (
	"context"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
)

type UserService struct {
	r repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{
		r: r,
	}
}

func (s *UserService) SignUp(ctx context.Context, dto *domain.UserCreateDTO) (*domain.User, error) {
	if err := dto.Validate(); err != nil {
		return nil, err
	}

	hash, err := GeneratePasswordHash(dto.Password)
	if err != nil {
		return nil, err
	}

	u := domain.NewUser(dto.Email, hash)

	return u, s.r.Create(u)
}

func (s *UserService) GetById(id int) (*domain.User, error) {
	return nil, nil
}

func (s *UserService) GetByEmail(email string) (*domain.User, error) {
	return nil, nil
}

func (s *UserService) Update(dto *domain.UserUpdateDTO) error {
	return nil
}

func (s *UserService) Delete(id int) error {
	return nil
}
