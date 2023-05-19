package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/stock_system?sslmode=disable"
)

var testQueries *Queries

var testDb *sql.DB

func TestMain(m *testing.M) {

	var err error
	testDb, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
