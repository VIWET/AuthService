package domain

import (
	"regexp"

	"github.com/VIWET/Beeracle/AuthService/internal/errors"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateDTO struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type UserUpdateDTO struct {
	Email              string `json:"email"`
	OldPassword        string `json:"oldPassword"`
	NewPassword        string `json:"newPassword"`
	NewPasswordConfirm string `json:"newPasswordConfirm"`
}

func TestUser() *User {
	return &User{
		Email:        "example@exml.com",
		PasswordHash: "n28ygr923hr8r6g83rh923ygr283gr9u23hr",
	}
}

func NewUser(email string, passwordHash string) *User {
	return &User{
		Email:        email,
		PasswordHash: passwordHash,
	}
}

func (dto *UserCreateDTO) Validate() error {
	if dto.Email == "" {
		return errors.ErrEmailIsEmpty
	}

	if !emailRegex.MatchString(dto.Email) {
		return errors.ErrEmailIsNotValid
	}

	if dto.Password == "" {
		return errors.ErrPasswordIsEmpty
	}

	if len(dto.Password) < 6 {
		return errors.ErrPasswordLength
	}

	if dto.Password != dto.PasswordConfirm {
		return errors.ErrPasswordConfirmation
	}

	return nil
}

func (dto *UserUpdateDTO) Validate() error {
	if dto.Email == "" {
		return errors.ErrEmailIsEmpty
	}

	if !emailRegex.MatchString(dto.Email) {
		return errors.ErrEmailIsNotValid
	}

	if dto.OldPassword == dto.NewPassword {
		return errors.ErrOldPasswordEqualNew
	}

	if dto.NewPassword == "" {
		return errors.ErrPasswordIsEmpty
	}

	if len(dto.NewPassword) < 6 {
		return errors.ErrPasswordLength
	}

	if dto.NewPassword != dto.NewPasswordConfirm {
		return errors.ErrPasswordConfirmation
	}

	return nil
}
