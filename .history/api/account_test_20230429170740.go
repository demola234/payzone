package api

import (
	"testing"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/utils"
	"github.com/golang/mock/gomock"
)

// Server serves HTTP requests for our banking service.

func TestGetAccount(t *testing.T) {
	account := generateRandomAccount()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := M

	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

}

func generateRandomAccount() db.Accounts {
	return db.Accounts{
		ID:       int64(utils.RandomInt(1, 1000)),
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
}
