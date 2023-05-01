package utils

const (
	USD = "USD"
	EUR = "EUR"
	NGN = "NGN"
	GBP = "GBP"
)

func IsSuppoertedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, NGN, GBP:
		return true
	}
	return false
}