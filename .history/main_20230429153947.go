package main

import (
	"database/sql"
	"log"
	"payzone/api"
	db "payzone/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	configs, err := utils.loadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(configs.DBDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
