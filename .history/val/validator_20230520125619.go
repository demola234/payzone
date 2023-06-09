package val

func ValidateStringField(value string, minLength int, maxLength int) error {
	n := len(value)
		if n < minLength || n > maxLength {
			return 
		}

		return nil
}
