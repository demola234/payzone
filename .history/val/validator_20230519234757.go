package val

func validateUsername(username string) error {
	if len(username) < 3 {
		return ErrUsernameTooShort
	}
	if len(username) > 25 {
		return ErrUsernameTooLong
	}
	return nil
}