package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/golang/mock/mockgen/model"
	"github.com/julysNICK/stock_system/api"
	db "github.com/julysNICK/stock_system/db/sqlc"
	"github.com/julysNICK/stock_system/utils"
	_ "github.com/lib/pq"
)


func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		fmt.Println("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DB_URL)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStoreDB(conn)

	server, err := api.NewServer(config,store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	// go func() {
	// 	for {
	// 		cron.CheckStockAlerts(store)

	// 		time.Sleep(1 * time.Minute)
	// 	}
	// }()

	err = server.Start(config.HTTP_SERVER_ADDRESS)

	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
