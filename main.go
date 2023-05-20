package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/julysNICK/stock_system/api"
	"github.com/julysNICK/stock_system/cron"
	db "github.com/julysNICK/stock_system/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	DBDrive       = "postgres"
	DBURL         = "postgresql://root:secret@localhost:5432/stock_system?sslmode=disable"
	ServerAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(DBDrive, DBURL)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStoreDB(conn)

	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	go func() {
		for {
			cron.CheckStockAlerts(store)

			time.Sleep(1 * time.Minute)
		}
	}()

	err = server.Start(ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
