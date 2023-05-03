package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChangePhoneNumberToInternationalFormat(t *testing.T) {
	phoneNumber := RandomPhoneNumber()
	phoneNumber = ChangePhoneNumberToInternationalFormat(phoneNumber)

	require.NotNil(t, phoneNumber)
	require.Equal(t, "234", phoneNumber[:3])
	require.Equal(t, 13, len(phoneNumber))
	require.Equal(t, "234", phoneNumber[:3])
	
}
