package sqlstore_test

import (
	"testing"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/errors"
	"github.com/VIWET/Beeracle/AuthService/internal/repository/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestSQLStore(t, config)
	defer teardown("users", "users_breweries", "users_profiles")

	r := sqlstore.NewUserRepository(db)

	u := domain.TestUser()

	assert.NoError(t, r.Create(u))
	assert.NotEqual(t, 0, u.ID)

	b := domain.TestBrewery()

	assert.NoError(t, r.Create(b))
	assert.NotEqual(t, 0, b.ID)
}

func TestUserRepository_GetByID(t *testing.T) {
	db, teardown := sqlstore.TestSQLStore(t, config)
	defer teardown("users", "users_breweries", "users_profiles")

	r := sqlstore.NewUserRepository(db)

	_, err := r.GetById(0)
	assert.EqualError(t, err, errors.ErrRecordNotFound.Error())

	u := domain.TestUser()

	if err := r.Create(u); err != nil {
		t.Fatal("error on creating:", err)
	}

	ut, err := r.GetById(u.ID)
	assert.NoError(t, err)
	assert.Equal(t, u, ut)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	db, teardown := sqlstore.TestSQLStore(t, config)
	defer teardown("users", "users_breweries", "users_profiles")

	r := sqlstore.NewUserRepository(db)

	_, err := r.GetByEmail("example@exml.com")
	assert.EqualError(t, err, errors.ErrRecordNotFound.Error())

	u := domain.TestUser()

	if err := r.Create(u); err != nil {
		t.Fatal("error on creating:", err)
	}

	ut, err := r.GetByEmail(u.Email)
	assert.NoError(t, err)
	assert.Equal(t, u, ut)
}

func TestUserRepository_Update(t *testing.T) {
	db, teardown := sqlstore.TestSQLStore(t, config)
	defer teardown("users", "users_breweries", "users_profiles")

	r := sqlstore.NewUserRepository(db)
	u := domain.TestUser()

	assert.EqualError(t, r.Update(u), errors.ErrRecordNotFound.Error())

	oldPwd := "example1"
	u.PasswordHash = oldPwd
	if err := r.Create(u); err != nil {
		t.Fatal("error on creating:", err)
	}
	u.PasswordHash = "example2"

	assert.NoError(t, r.Update(u))
	ut, err := r.GetById(u.ID)
	if err != nil {
		t.Fatal("error on getting by id: ", err)
	}
	assert.NotEqual(t, oldPwd, ut.PasswordHash)
}

func TestUserRepository_Delete(t *testing.T) {
	db, teardown := sqlstore.TestSQLStore(t, config)
	defer teardown("users", "users_breweries", "users_profiles")

	r := sqlstore.NewUserRepository(db)
	u := domain.TestUser()

	assert.EqualError(t, r.Delete(u.ID), errors.ErrRecordNotFound.Error())

	if err := r.Create(u); err != nil {
		t.Fatal("error on creating:", err)
	}

	assert.NoError(t, r.Delete(u.ID))
	_, err := r.GetById(u.ID)
	assert.EqualError(t, err, errors.ErrRecordNotFound.Error())
}
