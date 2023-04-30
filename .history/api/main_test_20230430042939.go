package api

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/demola234/payzone/utils"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
