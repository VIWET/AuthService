package service

import (
	"context"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
	"github.com/VIWET/Beeracle/AuthService/pkg/jwt"
)

type userService struct {
	r repository.UserRepository
	m jwt.TokenManager
}

func NewUserService(r repository.UserRepository, m jwt.TokenManager) User {
	return &userService{
		r: r,
		m: m,
	}
}

func (s *userService) SignUp(ctx context.Context, dto *domain.UserCreateDTO, role string) (*domain.User, jwt.Tokens, error) {
	tokens := jwt.Tokens{}
	if err := dto.Validate(); err != nil {
		return nil, tokens, err
	}

	hash, err := GeneratePasswordHash(dto.Password)
	if err != nil {
		return nil, tokens, err
	}

	u := domain.NewUser(dto.Email, role, hash)

	err = s.r.Create(u)
	if err != nil {
		return nil, tokens, err
	}

	accessToken, err := s.m.GenerateToken(u.Profile_ID, u.Email, u.Role)
	if err != nil {
		return nil, tokens, err
	}

	refreshToken, err := s.m.GenerateRefreshToken()
	if err != nil {
		return nil, tokens, err
	}

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken

	return u, tokens, nil
}
