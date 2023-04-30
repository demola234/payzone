package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// Server serves HTTP requests for our banking service.

func TestGetAccount(t *testing.T) {
	account := generateRandomAccount()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)

	server := NewServer(store)
	recorder := httptest.NewRecorder()

	// Create a new HTTP request that matches the route we want to test.
	url := fmt.Sprint("/accounts/%d", account.ID)
	request, err := http.NewRequest()
	require.NoError(t, err)
	server.router.ServeHTTP(recorder, request)

	


}

func generateRandomAccount() db.Accounts {
	return db.Accounts{
		ID:       int64(utils.RandomInt(1, 1000)),
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
}
