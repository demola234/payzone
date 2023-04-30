package api

import "testing"

// Server serves HTTP requests for our banking service.

func TestGetAccount(t *testing.T) {

}

func generateRandomAccount() Account {
	return Account{
		ID:       1,
		Owner:    "owner",
		Balance:  0,
		Currency: "USD",
	}
}