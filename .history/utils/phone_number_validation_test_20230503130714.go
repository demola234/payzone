package utils


func TestChangePhoneNumberToInternationalFormat(t *testing.T) {
	phoneNumber := "08012345678"
	phoneNumber = ChangePhoneNumberToInternationalFormat(phoneNumber)
	assert.Equal(t, "2348012345678", phoneNumber)
}