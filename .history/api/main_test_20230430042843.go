package api

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/demola234/payzone/utils"
)

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		


	os.Exit(m.Run())
}
