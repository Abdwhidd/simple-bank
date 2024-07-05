package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	dbDiver  = "postgres"
	dbSource = "postgresql://root:mysecret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDiver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
