package main

import (
	"database/sql"
	"log"

	_ "github.com/golang/mock/mockgen/model"
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

	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(ServerAddress)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
