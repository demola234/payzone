package api

import (
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

func TestCreateUser(t *testing.T) {
	user := generateRandomUser()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				full_name:
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateAccountParams{
					Owner:    account.Owner,
					Currency: account.Currency,
					Balance:  0,
				}

				store.EXPECT().
					CreateAccount(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)
				requireBodyMatchAccount(t, recoder.Body, account)
			},
		},
	}
}

func generateRandomUser() db.Users {
	hashPassword, _ := utils.HashPassword(utils.RandomString(6))
	return db.Users{
		Username:       utils.RandomOwner(),
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
		HashedPassword: hashPassword,
	}
}
