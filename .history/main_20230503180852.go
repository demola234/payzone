package main

import (
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// func main() {
// 	configs, err := utils.LoadConfig(".")
// 	if err != nil {
// 		log.Fatal("cannot load config: ", err)
// 	}

// 	conn, err := sql.Open(configs.DBDriver, configs.DBSource)
// 	if err != nil {
// 		log.Fatal("cannot connect to db: ", err)
// 	}

// 	store := db.NewStore(conn)
// 	server, err := api.NewServer(configs, store)
// 	if err != nil {
// 		log.Fatal("cannot create server: ", err)
// 	}

// 	err = server.Start(configs.ServerAddress)
// 	if err != nil {
// 		log.Fatal("cannot start server: ", err)
// 	}

// }

func main() {
	resp, err := http.Get("https://exchange-rates.abstractapi.com/v1/live/?api_key=09ad830c8bf643ea93b2349e5b9cdd98&base=USD&target=EUR")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
