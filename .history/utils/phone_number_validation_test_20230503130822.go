package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChangePhoneNumberToInternationalFormat(t *testing.T) {
	phoneNumber := "08012345678"
	phoneNumber = ChangePhoneNumberToInternationalFormat(phoneNumber)

	if phoneNumber != "2348012345678" {
		t.Errorf("Expected 2348012345678, got %s", phoneNumber)
	}

	require.NotNil()

}
