package utils

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if (err != nil) {
		return "", fmt.P
	}
}