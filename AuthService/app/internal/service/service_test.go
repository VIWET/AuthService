package service_test

import (
	"testing"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/service"
	"github.com/stretchr/testify/assert"
)

func TestService_GeneratePasswordHash(t *testing.T) {
	hash, err := service.GeneratePasswordHash("example")
	assert.NoError(t, err)
	assert.NotNil(t, hash)
}

func TestService_CheckPassword(t *testing.T) {
	u := domain.TestUser()
	hash, _ := service.GeneratePasswordHash("example")
	u.PasswordHash = hash

	err := service.CheckPassword(u, "shdbfhbsf")
	assert.Error(t, err)

	err = service.CheckPassword(u, "example")
	assert.NoError(t, err)

}
