package sqlstore

import (
	"database/sql"

	"github.com/VIWET/Beeracle/AuthService/internal/repository"
)

func NewSQLRepositories(db *sql.DB) *repository.Repositories {
	return &repository.Repositories{
		UserRepository: NewUserRepository(db),
	}
}
