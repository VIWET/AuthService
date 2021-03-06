package service

import (
	"context"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/errors"
	"github.com/VIWET/Beeracle/AuthService/internal/jwt"
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
)

type userService struct {
	r repository.UserRepository
	m jwt.TokenManager
	c repository.CacheRepository
}

func NewUserService(r repository.UserRepository, m jwt.TokenManager, c repository.CacheRepository) User {
	return &userService{
		r: r,
		m: m,
		c: c,
	}
}

func (s *userService) SignUp(ctx context.Context, dto *domain.UserCreateDTO, role string, ua string) (jwt.Tokens, error) {
	tokens := jwt.Tokens{}
	if err := dto.Validate(); err != nil {
		return tokens, err
	}

	hash, err := GeneratePasswordHash(dto.Password)
	if err != nil {
		return tokens, err
	}

	u := domain.NewUser(dto.Email, role, hash)

	err = s.r.Create(u)
	if err != nil {
		return tokens, err
	}

	accessToken, err := s.m.GenerateToken(u.ID, u.Role, u.ProfileID)
	if err != nil {
		return tokens, err
	}

	refreshToken, err := s.m.GenerateRefreshToken()
	if err != nil {
		return tokens, err
	}

	rs := &domain.RefreshSession{
		ProfileID:   u.ProfileID,
		Role:        u.Role,
		UserAgent:   ua,
		Fingerprint: dto.Fingerprint,
	}

	s.c.Set(refreshToken, rs)

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken
	tokens.Exp = s.m.GetExpTime()

	return tokens, nil
}

func (s *userService) SignIn(ctx context.Context, dto *domain.UserSignIn, ua string) (jwt.Tokens, error) {
	tokens := jwt.Tokens{}

	if err := dto.Validate(); err != nil {
		return tokens, err
	}

	u, err := s.r.GetByEmail(dto.Email)
	if err != nil {
		return tokens, err
	}

	if err := CheckPassword(u, dto.Password); err != nil {
		return tokens, errors.ErrUnauthorized
	}

	accessToken, err := s.m.GenerateToken(u.ID, u.Role, u.ProfileID)
	if err != nil {
		return tokens, err
	}

	refreshToken, err := s.m.GenerateRefreshToken()
	if err != nil {
		return tokens, err
	}

	rs := &domain.RefreshSession{
		ProfileID:   u.ProfileID,
		Role:        u.Role,
		UserAgent:   ua,
		Fingerprint: dto.Fingerprint,
	}

	s.c.Set(refreshToken, rs)

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken
	tokens.Exp = s.m.GetExpTime()

	return tokens, nil
}

func (s *userService) Refresh(ctx context.Context, rt string, ua string, fp string) (jwt.Tokens, error) {
	tokens := jwt.Tokens{}

	rs, err := s.c.Get(rt)
	if err != nil {
		return tokens, errors.ErrUnauthorized
	}

	if rs.Fingerprint != fp || rs.UserAgent != ua {
		return tokens, errors.ErrUnauthorized
	}

	err = s.c.Delete(rt)
	if err != nil {
		return tokens, err
	}

	accessToken, err := s.m.GenerateToken(rs.ID, rs.Role, rs.ProfileID)
	if err != nil {
		return tokens, err
	}

	refreshToken, err := s.m.GenerateRefreshToken()
	if err != nil {
		return tokens, err
	}

	err = s.c.Set(refreshToken, rs)
	if err != nil {
		return tokens, err
	}

	tokens.AccessToken = accessToken
	tokens.RefreshToken = refreshToken
	tokens.Exp = s.m.GetExpTime()

	return tokens, nil
}

func (s *userService) Update(ctx context.Context, dto *domain.UserUpdateDTO) error {
	if err := dto.Validate(); err != nil {
		return err
	}

	id := ctx.Value(domain.UID("id")).(int)

	u, err := s.r.GetById(id)
	if err != nil {
		return err
	}

	if dto.Email != nil {
		u.Email = *dto.Email
	}

	if dto.NewPassword != nil {
		if err := CheckPassword(u, *dto.OldPassword); err != nil {
			if err == errors.ErrPasswordIsWrong {
				return errors.ErrOldPasswordIsWrong
			}
			return err
		}

		hash, err := GeneratePasswordHash(*dto.NewPassword)
		if err != nil {
			return err
		}

		u.PasswordHash = hash
	}

	if err := s.r.Update(u); err != nil {
		return err
	}

	return nil
}

func (s *userService) Delete(ctx context.Context, password string) error {
	id := ctx.Value(domain.UID("id")).(int)

	u, err := s.r.GetById(id)
	if err != nil {
		return err
	}

	if err := CheckPassword(u, password); err != nil {
		return err
	}

	if err := s.r.Delete(id); err != nil {
		return err
	}

	return nil
}
