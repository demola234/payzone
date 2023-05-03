package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChangePhoneNumberToInternationalFormat(t *testing.T) {
	phoneNumber := RandomPhoneNumber()
	phoneNumber = ChangePhoneNumberToInternationalFormat(phoneNumber)

	require.NotNil(t, phoneNumber)
	require.Equal(t, "2348012345678", phoneNumber)
}
