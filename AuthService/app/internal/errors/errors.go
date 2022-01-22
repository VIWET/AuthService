package errors

import "errors"

var (
	ErrRecordNotFound       = errors.New("record not found")
	ErrEmailIsEmpty         = errors.New("email mustn't be empty")
	ErrEmailIsNotValid      = errors.New("email must be valid")
	ErrPasswordIsEmpty      = errors.New("password mustn't be empty")
	ErrPasswordConfirmation = errors.New("password confirmation must be equal password")
	ErrPasswordLength       = errors.New("password length must be more than six symbols")
	ErrOldPasswordEqualNew  = errors.New("new password the same as old one")
	ErrPasswordIsWrong      = errors.New("password is wrong")
	ErrOldPasswordIsWrong   = errors.New("old password is wrong")
	ErrUnauthorized         = errors.New("unauthorized")
)
