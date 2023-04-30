package api

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../")

	os.Exit(m.Run())
}
