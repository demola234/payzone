package main

import (
	"database/sql"
	"log"
	"payzone/api"
	db "payzone/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/payzone?sslmode=disable"
	serverAddress = ":8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

}
