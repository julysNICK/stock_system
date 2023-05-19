package main

import (
	"database/sql"
	"log"

	"github.com/julysNICK/stock_system/api"
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

	server := api.NewServer(store)

	err = server.Start(ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
