package sqlstore

import (
	"database/sql"
	"fmt"

	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func TestSQLStore(t *testing.T, c *Config) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", c.GetConnectionString())
	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
		}

		db.Close()
	}
}
