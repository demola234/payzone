package val

import (
	"fmt"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidEmail    = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`).MatchString
	isValidPassword = regexp.MustCompile(`^[a-zA-Z0-9!@#$&*]+$`).MatchString
)

func ValidateStringField(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("length must be between %d and %d", minLength, maxLength)
	}

	return nil
}

func ValidateUsername(username string) error {
	err := ValidateStringField(username, 3, 50)
	if err != nil {
		return err
	}
	if !isValidUsername(username) {
		return fmt.Errorf("username must not contain special characters")
	}
	return nil
}

func ValidatePassword(password string) error {
	err := ValidateStringField(password, 6, 50)
	if err != nil {
		return err
	}
	if !isValidPassword(password) {
		return fmt.Errorf("password must not contain special characters")
	}
	return nil
}

func ValidateEmail(email string) error {
	if !isValidEmail(email) {
		return fmt.Errorf("invalid email address")
		
	return ValidateStringField(email, 5, 50)
}
