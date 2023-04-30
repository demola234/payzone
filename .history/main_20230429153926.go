package main

import (
	"database/sql"
	"log"
	"payzone/api"
	"payzone/utils"
	db "payzone/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.loadConfig(".")
	if err != nil {
		
	conn, err := sql.Open(dbDriver, dbSource)
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
