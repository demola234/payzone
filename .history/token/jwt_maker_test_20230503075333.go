package token

import (
	"testing"

	"github.com/demola234/payzone/utils"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)


	username := utils.RandomOwner()
	duration := time.Duration(utils.RandomInt(0, 256)) * time.Minute

}
