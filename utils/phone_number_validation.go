package utils

func ChangePhoneNumberToInternationalFormat(phoneNumber string) string {
	if phoneNumber[0] == '0' {
		phoneNumber = phoneNumber[1:]
		phoneNumber = "234" + phoneNumber
	}

	return phoneNumber
}
