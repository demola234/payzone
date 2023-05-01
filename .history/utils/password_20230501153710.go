package utils

import "golang.org/x/crypto/bcrypt"

func hashPassword(password string) (string, error) {
	hashpassword := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}