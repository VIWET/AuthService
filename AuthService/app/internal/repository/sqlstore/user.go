package sqlstore

import (
	"database/sql"

	"github.com/VIWET/Beeracle/AuthService/internal/domain"
	"github.com/VIWET/Beeracle/AuthService/internal/errors"
	"github.com/VIWET/Beeracle/AuthService/internal/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(u *domain.User) error {
	var role_id int
	switch u.Role {
	case "admin":
		role_id = 1
	case "user":
		role_id = 2
	case "brewery":
		role_id = 3
	}
	err := r.db.QueryRow("INSERT INTO users (email, password_hash, role_id) VALUES ($1, $2, $3) RETURNING ID", u.Email, u.PasswordHash, role_id).Scan(&u.ID)
	if err != nil {
		return err
	}

	if role_id == 3 {
		err = r.db.QueryRow("INSERT INTO users_breweries (user_id) VALUES ($1) RETURNING ID", u.ID).Scan(&u.Profile_ID)
		if err != nil {
			return err
		}
	} else {
		err = r.db.QueryRow("INSERT INTO users_profiles (user_id) VALUES ($1) RETURNING ID", u.ID).Scan(&u.Profile_ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *userRepository) GetById(id int) (*domain.User, error) {
	u := &domain.User{}
	if err := r.db.QueryRow(
		"SELECT u.id, u.email, u.password_hash, r.role_name, CASE WHEN r.role_name = 'brewery' THEN ub.id ELSE up.id END "+
			"FROM users AS u "+
			"LEFT JOIN roles AS r ON u.role_id = r.id "+
			"LEFT JOIN users_breweries AS ub ON ub.user_id = u.id "+
			"LEFT JOIN users_profiles AS up ON up.user_id = u.id "+
			"WHERE u.id = $1",
		id,
	).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Role, &u.Profile_ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	u := &domain.User{}
	if err := r.db.QueryRow(
		"SELECT u.id, u.email, u.password_hash, r.role_name, CASE WHEN r.role_name = 'brewery' THEN ub.id ELSE up.id END "+
			"FROM users AS u "+
			"LEFT JOIN roles AS r ON u.role_id = r.id "+
			"LEFT JOIN users_breweries AS ub ON ub.user_id = u.id "+
			"LEFT JOIN users_profiles AS up ON up.user_id = u.id "+
			"WHERE u.email = $1 ",
		email,
	).Scan(&u.ID, &u.Email, &u.PasswordHash, &u.Role, &u.Profile_ID); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.ErrRecordNotFound
		}
		return nil, err
	}

	return u, nil
}

func (r *userRepository) Update(u *domain.User) error {
	err := r.db.QueryRow("UPDATE users SET email = $1, password_hash = $2 WHERE id = $3 RETURNING id", u.Email, u.PasswordHash, u.ID).Scan(&u.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.ErrRecordNotFound
		}
		return err
	}
	return nil
}

func (r *userRepository) Delete(id int) error {
	err := r.db.QueryRow("DELETE FROM users WHERE id = $1 RETURNING id", id).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.ErrRecordNotFound
		}
		return err
	}
	return nil
}
