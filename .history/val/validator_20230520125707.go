package val

import "fmt"

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
		
	return 
}

func ValidatePassword(password string) error {
	return ValidateStringField(password, 6, 50)
}

func ValidateEmail(email string) error {
	return ValidateStringField(email, 5, 50)
}
