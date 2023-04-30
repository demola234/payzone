package db

import (
	"database/sql"
	"log"
	"os"
	"payzone/utils"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	configs, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	var dbErr error
	testDB, err = sql.Open(, dbSource)
	if dbErr != nil {
		log.Fatal("cannot connect to db: ", dbErr)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
