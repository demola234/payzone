package token

import (
	"testing"

	"github.com/demola234/payzone/utils"
)

func TestJWTMaker(t *testing.T) {
	NewJWTMaker(utils.RandomString(32))
}
