package service_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/errors"
	"github.com/VIWET/Beeracle/AuthService/internal/jwt"
	"github.com/VIWET/Beeracle/AuthService/internal/repository/testcache"
	"github.com/VIWET/Beeracle/AuthService/internal/repository/teststore"
	"github.com/VIWET/Beeracle/AuthService/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestUserService_SignUp(t *testing.T) {
	r := teststore.NewTestUserRepository()

	cr := testcache.NewTestCacheRepository()

	m, err := jwt.NewTokenManager("123")
	if err != nil {
		t.Fatal(err)
	}

	s := service.NewUserService(r, m, cr)

	tests := []struct {
		Name            string
		Email           string
		Password        string
		PasswordConfirm string
		Role            string
		Valid           bool
		Err             error
	}{
		{
			Name:            "Valid",
			Email:           "example1@exml.com",
			Password:        "example",
			PasswordConfirm: "example",
			Role:            "user",
			Valid:           true,
			Err:             nil,
		},
		{
			Name:            "Valid",
			Email:           "brew@exml.com",
			Password:        "example",
			PasswordConfirm: "example",
			Role:            "brewery",
			Valid:           true,
			Err:             nil,
		},
		{
			Name:            "Valid",
			Email:           "example2@exml.com",
			Password:        "example",
			PasswordConfirm: "example",
			Role:            "user",
			Valid:           true,
			Err:             nil,
		},
		{
			Name:            "Empty email",
			Email:           "",
			Password:        "example",
			PasswordConfirm: "example",
			Role:            "user",
			Valid:           false,
			Err:             errors.ErrEmailIsEmpty,
		},
		{
			Name:            "Invalid email",
			Email:           "examexml.com",
			Password:        "example",
			PasswordConfirm: "example",
			Role:            "user",
			Valid:           false,
			Err:             errors.ErrEmailIsNotValid,
		},
		{
			Name:            "Empty password",
			Email:           "example@exml.com",
			Password:        "",
			PasswordConfirm: "",
			Role:            "user",
			Valid:           false,
			Err:             errors.ErrPasswordIsEmpty,
		},
		{
			Name:            "Short password ",
			Email:           "example@exml.com",
			Password:        "exa",
			PasswordConfirm: "example",
			Role:            "user",
			Valid:           false,
			Err:             errors.ErrPasswordLength,
		},
		{
			Name:            "Password isn't equal password confirmation",
			Email:           "example@exml.com",
			Password:        "example15",
			PasswordConfirm: "example",
			Role:            "user",
			Valid:           false,
			Err:             errors.ErrPasswordConfirmation,
		},
	}

	ctx := context.Background()

	for _, c := range tests {
		dto := &domain.UserCreateDTO{
			Email:           c.Email,
			Password:        c.Password,
			PasswordConfirm: c.PasswordConfirm,
		}
		if c.Valid {
			u, tokens, err := s.SignUp(ctx, dto, c.Role)
			assert.NoError(t, err)
			assert.NotNil(t, u)
			assert.NotEmpty(t, tokens.AccessToken)
			assert.NotEmpty(t, tokens.RefreshToken)
			profIdS, role, _ := m.ParseToken(tokens.AccessToken)
			profId, err := strconv.Atoi(profIdS)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, u.ProfileID, profId)
			assert.Equal(t, u.Role, role)
		} else {
			u, tokens, err := s.SignUp(ctx, dto, c.Role)
			assert.EqualError(t, err, c.Err.Error())
			assert.Nil(t, u)
			assert.Empty(t, tokens.AccessToken)
			assert.Empty(t, tokens.RefreshToken)
		}
	}
}
