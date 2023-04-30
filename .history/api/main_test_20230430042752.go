package api

import (
	"os"
	"testing"

	"github.com/demola234/payzone/utils"
)

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../")
	

	os.Exit(m.Run())
}
