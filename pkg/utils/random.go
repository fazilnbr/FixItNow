package utils

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
)

const (
	alpabet = "abcdefghijklmnopqrstuvwxyz"
	dbURI   = "postgresql://developer:dev123@localhost:5432/workey?sslmode=disable"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate randome inteager value between min and max

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)

}

// RandomString generate random string of lenth num

func RandomString(num int) string {
	var sb strings.Builder
	k := len(alpabet)
	for i := 0; i < num; i++ {
		c := alpabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func MockGormDB() (*sql.DB, sqlmock.Sqlmock) {
	_, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		logrus.Fatal(err)
	}

	gormDB, err := sql.Open("postgres", dbURI)
	if err != nil {
		logrus.Fatal(err)
	}

	return gormDB, mock
}
