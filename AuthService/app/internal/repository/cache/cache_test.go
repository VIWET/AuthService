package cache_test

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/VIWET/Beeracle/AuthService/internal/repository/cache"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var (
	config *cache.Config
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error loading .env file")
	}
	config = cache.NewConfig()

	config.Host = os.Getenv("TEST_DBHOST")
	config.Port = os.Getenv("TEST_DBPORT")
	db, err := strconv.Atoi(os.Getenv("TEST_DB"))
	if err != nil {
		logrus.Fatal(err)
		return
	}

	config.DB = db
	config.Pwd = os.Getenv("TEST_DBPASSWORD")

	exp, err := strconv.Atoi(os.Getenv("TEST_EXP"))
	if err != nil {
		logrus.Fatal(err)
		return
	}

	config.Expires = time.Duration(exp) * time.Second
	os.Exit(m.Run())
}
