package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/demola234/payzone/db/mock"
	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	// Create 2 new account
	amount := int64(10)
	account1 := generateRandomAcct(1)
	account2 := generateRandomAcct(3)

	account1.Currency = utils.USD
	account2.Currency = utils.USD

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"from_account_id": account1.ID,
				"to_account_id":   account2.ID,
				"amount":          amount,
				"currency":        utils.USD,
			},
			buildStubs: func(store *mockdb.MockStore) {

				arg := db.TransferTxParams{
					FromAccountID: account1.ID,
					ToAccountID:   account2.ID,
					Amount:        amount,
				}
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(account1.ID)).
					Return(account1, nil).
					Times(1)
				store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(1).Return(account2, nil)

				store.EXPECT().
					TransferTx(gomock.Any(), gomock.Eq(arg)).
					Times(1)
			},
			checkResponse: func(recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
			},
		},
	
	{
		name: "ToAccountNotFound",
		body: gin.H{
			"from_account_id": account1.ID,
			"to_account_id":   account2.ID,
			"amount":          amount,
			"currency":        util.USD,
		},
		setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
			addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user1.Username, time.Minute)
		},
		buildStubs: func(store *mockdb.MockStore) {
			store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account1.ID)).Times(1).Return(account1, nil)
			store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account2.ID)).Times(1).Return(db.Account{}, sql.ErrNoRows)
			store.EXPECT().TransferTx(gomock.Any(), gomock.Any()).Times(0)
		},
		checkResponse: func(recorder *httptest.ResponseRecorder) {
			require.Equal(t, http.StatusNotFound, recorder.Code)
		},
	}
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/transfers"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(recorder)
		})
	}
}

func generateRandomAcct(id int64) db.Accounts {
	return db.Accounts{
		ID:       id,
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
}
