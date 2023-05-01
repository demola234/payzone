package api

import (
	"testing"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/utils"
)

func TestCreateUser(t *testing.T) {

}

func generateRandomUser() db.Users {
	return db.Accounts{
		ID:       int64(utils.RandomInt(1, 1000)),
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
}
