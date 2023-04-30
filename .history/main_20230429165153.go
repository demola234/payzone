package main

import (
	"database/sql"
	"log"

	db "payzone/db/sqlc"

	"github.com/demola234/payzone/api"
	"github.com/demola234/payzone/utils"
	_ "github.com/lib/pq"
)

func main() {
	configs, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(configs.DBDriver, configs.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(configs.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
