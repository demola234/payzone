package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashPassword, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hashPassword)


	err = CheckPassword(password, hashPassword)
	require.NoError(t, err)

	password := RandomString(6)
}
