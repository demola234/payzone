package utils

import "testing"

func TestChangePhoneNumberToInternationalFormat(t *testing.T) {
	phoneNumber := "08012345678"
	phoneNumber = ChangePhoneNumberToInternationalFormat(phoneNumber)

}

