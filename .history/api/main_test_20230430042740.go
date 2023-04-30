package api

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config, err := LoadConfig("../../")

	os.Exit(m.Run())
}
