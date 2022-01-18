package teststore

import (
	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/errors"
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
)

type testUserRepository struct {
	profiles  int
	breweries int
	db        map[int]*domain.User
}

func NewTestUserRepository() repository.UserRepository {
	return &testUserRepository{
		profiles:  1,
		breweries: 1,
		db:        make(map[int]*domain.User),
	}
}

func (r *testUserRepository) Create(u *domain.User) error {
	u.ID = len(r.db) + 1
	if u.Role == "brewery" {
		u.ProfileID = r.breweries
		r.breweries++
	} else {
		u.ProfileID = r.profiles
		r.profiles++
	}

	r.db[u.ID] = u

	return nil
}

func (r *testUserRepository) GetById(id int) (*domain.User, error) {
	u, ok := r.db[id]
	if !ok {
		return nil, errors.ErrRecordNotFound
	}

	return u, nil
}

func (r *testUserRepository) GetByEmail(email string) (*domain.User, error) {
	for _, u := range r.db {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, errors.ErrRecordNotFound
}

func (r *testUserRepository) Update(u *domain.User) error {
	u, ok := r.db[u.ID]
	if !ok {
		return errors.ErrRecordNotFound
	}

	r.db[u.ID] = u
	return nil
}

func (r *testUserRepository) Delete(id int) error {
	u, ok := r.db[id]
	if !ok {
		return errors.ErrRecordNotFound
	}

	delete(r.db, u.ID)
	return nil
}
