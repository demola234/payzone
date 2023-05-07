package utils

const (
	USD = "USD"
	EUR = "EUR"
	NGN = "NGN"
	GBP = "GBP"
	CAD = "CAD"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, NGN, GBP, CAD:
		return true
	}
	return false
}