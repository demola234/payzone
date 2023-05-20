package main

import (
	"database/sql"
	"log"

	"github.com/demola234/payzone/api"
	db "github.com/demola234/payzone/db/sqlc"
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
	runGinServer
	

}


func runGinServer(configs utils.Config, store db.Store) {
	server, err := api.NewServer(configs, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(configs.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}