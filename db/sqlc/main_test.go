package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/julysNICK/stock_system/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries

var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := utils.LoadConfig("../../")

	if err != nil {
		fmt.Println("cannot load config: ", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DB_URL)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
