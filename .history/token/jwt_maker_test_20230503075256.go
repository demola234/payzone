package token

import (
	"testing"

	"github.com/demola234/payzone/utils"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
}
