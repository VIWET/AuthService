package jwt_test

import (
	"testing"
	"time"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/jwt"

	"github.com/stretchr/testify/assert"
)

func TestManager_GenerateUserToken(t *testing.T) {
	m, err := jwt.NewTokenManager("Secret key", time.Minute)
	if err != nil {
		t.Fatal(err)
	}

	u := domain.TestUser()

	ss, err := m.GenerateToken(u.ID, "user", u.ProfileID)
	assert.NoError(t, err)
	assert.NotNil(t, ss)
}

func TestManager_ParseToken(t *testing.T) {
	m, err := jwt.NewTokenManager("Secret key", time.Minute)
	if err != nil {
		t.Fatal(err)
	}

	u := domain.TestUser()

	ss, err := m.GenerateToken(u.ID, "user", u.ProfileID)
	assert.NoError(t, err)
	assert.NotNil(t, ss)

	c, err := m.ParseToken(ss)
	assert.NoError(t, err)
	assert.Equal(t, u.ProfileID, c.ProfileID)
}
