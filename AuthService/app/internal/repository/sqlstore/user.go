package sqlstore

import (
	"database/sql"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(u *domain.User) error {
	return r.db.QueryRow("INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING ID", u.Email, u.PasswordHash).Scan(&u.ID)
}

func (r *UserRepository) GetById(id int) (*domain.User, error) {
	u := &domain.User{}
	if err := r.db.QueryRow("SELECT id, email, password_hash FROM users WHERE id = $1", id).Scan(&u.ID, &u.Email, &u.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	u := &domain.User{}
	if err := r.db.QueryRow("SELECT id, email, password_hash FROM users WHERE email = $1", email).Scan(&u.ID, &u.Email, &u.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) Update(u *domain.User) error {
	err := r.db.QueryRow("UPDATE users SET email = $1, password_hash = $2 WHERE id = $3 RETURNING id", u.Email, u.PasswordHash, u.ID).Scan(&u.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return repository.ErrRecordNotFound
		}
		return err
	}
	return nil
}

func (r *UserRepository) Delete(id int) error {
	err := r.db.QueryRow("DELETE FROM users WHERE id = $1 RETURNING id", id).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return repository.ErrRecordNotFound
		}
		return err
	}
	return nil
}
