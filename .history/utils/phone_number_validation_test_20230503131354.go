package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChangePhoneNumberToInternationalFormat(t *testing.T) {
	phoneNumber := 
	phoneNumber = ChangePhoneNumberToInternationalFormat(phoneNumber)

	if phoneNumber != "2348012345678" {
		t.Errorf("Expected 2348012345678, got %s", phoneNumber)
	}

	require.NotNil(t, phoneNumber)
	require.Equal(t, "2348012345678", phoneNumber)
}
