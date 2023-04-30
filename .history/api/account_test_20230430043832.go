package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	mockdb "github.com/demola234/payzone/db/mock"
	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// Server serves HTTP requests for our banking service.

func TestGetAccount(t *testing.T) {
	testCases := []struct {
		name         string
		accountID    int64
		buildStubs   func(store *mockdb.MockStore)
		expectStatus int
	}{
		{
			name:      "OK",
			accountID: 1,
			buildStubs: func(store *mockdb.MockStore) {
				const id int64 = 1
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(id)).
					Return(generateRandomAccount(id), nil).
					Times(1)
			},
			expectStatus: http.StatusOK,
		},
		{
			name:      "NotFound",
			accountID: 2,
			buildStubs: func(store *mockdb.MockStore) {
				const id int64 = 2
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(id)).
					Return(db.Accounts{}, sql.ErrNoRows).
					Times(1)
			},
			expectStatus: http.StatusNotFound,
		},
		{
			name:      "InternalError",
			accountID: 3,
			buildStubs: func(store *mockdb.MockStore) {
				const id int64 = 3
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(id)).
					Return(db.Accounts{}, sql.ErrConnDone).
					Times(1)
			},
			expectStatus: http.StatusInternalServerError,
		},
		{
			name:      "BadRequest",
			accountID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			expectStatus: http.StatusBadRequest,
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

			url := fmt.Sprintf("/accounts/%d", tc.accountID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(t, tc.expectStatus, recorder.Code)
			// requiredBodyMatchAccount(t, recorder.Body, generateRandomAccount(tc.accountID))
		})
	}
}

func generateRandomAccount(id int64) db.Accounts {
	return db.Accounts{
		ID:        id,
		Owner:     utils.RandomOwner(),
		Balance:   utils.RandomMoney(),
		Currency:  utils.RandomCurrency(),
		CreatedAt: time.Now(),
	}
}

// func requiredBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Accounts) {
// 	data, err := io.ReadAll(body)
// 	require.NoError(t, err)

// 	var gotAccount db.Accounts
// 	err = json.Unmarshal(data, &gotAccount)
// 	require.NoError(t, err)
// 	require.Equal(t, account, gotAccount)
// }

func TestCreateAccount(t *testing.T) {
  accounts :=	generateRandomAccount(1)
	testCases := []struct {
		name         string
		body         string
		buildStubs   func(store *mockdb.MockStore)
		expectStatus int
	}{
		{

			name: "OK",
			body: `{"}`,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Any()).
					Times(1).
					Return(generateRandomAccount(1), nil)
			},
			expectStatus: http.StatusOK,
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

			url := fmt.Sprintf("/accounts")
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			require.Equal(t, tc.expectStatus, recorder.Code)
			// requiredBodyMatchAccount(t, recorder.Body, generateRandomAccount(tc.accountID))
		})
	}
}