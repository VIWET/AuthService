package sqlstore_test

import (
	"os"
	"testing"

	"github.com/VIWET/Beeracle/AuthService/internal/repository/sqlstore"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	config *sqlstore.Config
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	config = sqlstore.NewConfig()

	config.Host = os.Getenv("TEST_DBHOST")
	config.Port = os.Getenv("TEST_DBPORT")
	config.Name = os.Getenv("TEST_DBNAME")
	config.User = os.Getenv("TEST_DBUSER")
	config.Pwd = os.Getenv("TEST_DBPASSWORD")

	os.Exit(m.Run())
}
