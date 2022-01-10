package teststore

import (
	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
)

type TestUserRepository struct {
	db map[int]*domain.User
}

func NewTestUserRepository() *TestUserRepository {
	return &TestUserRepository{
		db: make(map[int]*domain.User),
	}
}

func (r *TestUserRepository) Create(u *domain.User) error {
	u.ID = len(r.db) + 1
	r.db[u.ID] = u

	return nil
}

func (r *TestUserRepository) GetById(id int) (*domain.User, error) {
	u, ok := r.db[id]
	if !ok {
		return nil, repository.ErrRecordNotFound
	}

	return u, nil
}

func (r *TestUserRepository) GetByEmail(email string) (*domain.User, error) {
	for _, u := range r.db {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, repository.ErrRecordNotFound
}

func (r *TestUserRepository) Update(u *domain.User) error {
	u, ok := r.db[u.ID]
	if !ok {
		return repository.ErrRecordNotFound
	}

	r.db[u.ID] = u
	return nil
}

func (r *TestUserRepository) Delete(id int) error {
	u, ok := r.db[id]
	if !ok {
		return repository.ErrRecordNotFound
	}

	delete(r.db, u.ID)
	return nil
}
