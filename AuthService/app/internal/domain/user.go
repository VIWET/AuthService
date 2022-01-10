package domain

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
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
