package service_test

import (
	"context"
	"testing"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/errors"
	"github.com/VIWET/Beeracle/AuthService/internal/repository/teststore"
	"github.com/VIWET/Beeracle/AuthService/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestUserService_SignUp(t *testing.T) {
	r := teststore.NewTestUserRepository()

	s := service.NewUserService(r)

	tests := []struct {
		Name            string
		Email           string
		Password        string
		PasswordConfirm string
		Valid           bool
		Err             error
	}{
		{
			Name:            "Valid",
			Email:           "example@exml.com",
			Password:        "example",
			PasswordConfirm: "example",
			Valid:           true,
			Err:             nil,
		},
		{
			Name:            "Empty email",
			Email:           "",
			Password:        "example",
			PasswordConfirm: "example",
			Valid:           false,
			Err:             errors.ErrEmailIsEmpty,
		},
		{
			Name:            "Invalid email",
			Email:           "examexml.com",
			Password:        "example",
			PasswordConfirm: "example",
			Valid:           false,
			Err:             errors.ErrEmailIsNotValid,
		},
		{
			Name:            "Empty password",
			Email:           "example@exml.com",
			Password:        "",
			PasswordConfirm: "",
			Valid:           false,
			Err:             errors.ErrPasswordIsEmpty,
		},
		{
			Name:            "Short password ",
			Email:           "example@exml.com",
			Password:        "exa",
			PasswordConfirm: "example",
			Valid:           false,
			Err:             errors.ErrPasswordLength,
		},
		{
			Name:            "Password isn't equal password confirmation",
			Email:           "example@exml.com",
			Password:        "example15",
			PasswordConfirm: "example",
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
			u, err := s.SignUp(ctx, dto)
			assert.NoError(t, err)
			assert.NotNil(t, u)
		} else {
			u, err := s.SignUp(ctx, dto)
			assert.EqualError(t, err, c.Err.Error())
			assert.Nil(t, u)
		}
	}
}
