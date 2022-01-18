package jwt_test

import (
	"fmt"
	"testing"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/jwt"

	"github.com/stretchr/testify/assert"
)

func TestManager_GenerateUserToken(t *testing.T) {
	m, err := jwt.NewTokenManager("Secret key")
	if err != nil {
		t.Fatal(err)
	}

	u := domain.TestUser()

	ss, err := m.GenerateToken(u.ID, u.Email, "user")
	assert.NoError(t, err)
	assert.NotNil(t, ss)
}

func TestManager_ParseToken(t *testing.T) {
	m, err := jwt.NewTokenManager("Secret key")
	if err != nil {
		t.Fatal(err)
	}

	u := domain.TestUser()

	ss, err := m.GenerateToken(u.ProfileID, u.Email, "user")
	assert.NoError(t, err)
	assert.NotNil(t, ss)

	id, sub, err := m.ParseToken(ss)
	assert.NoError(t, err)
	fmt.Println(id, sub)
}
