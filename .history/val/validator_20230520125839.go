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
	return nil
}

func ValidatePassword(password string) error {
	err := ValidateStringField(password, 6, 50)
	if err != nil {
		return err
	}
	return 
}

func ValidateEmail(email string) error {
	return ValidateStringField(email, 5, 50)
}
