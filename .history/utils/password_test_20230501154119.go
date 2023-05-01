package utils

import "testing"

func TestPassword(t *testing.T) {
	password := RandomString(6)

	hashPassword, err HashPassword(password)
}
