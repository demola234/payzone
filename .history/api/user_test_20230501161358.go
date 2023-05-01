package api

import "testing"

func TestCreateUser(t *testing.T) {

}

func generateRandomAccount() db.Accounts {
	return db.Accounts{
		ID:       int64(utils.RandomInt(1, 1000)),
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
}