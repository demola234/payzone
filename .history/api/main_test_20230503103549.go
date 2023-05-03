package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func newTest

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
