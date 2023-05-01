package api

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateTransfer(t *testing.T) {
	// Create 2 new account
	amount := int64(10)
	account1 := generateRandomAccount()
	account2 := generateRandomAccount()

	testCases := []struct {
		name        string
		body        gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)

	}{
		{
			name: "OK",
			body: gin.H{
				"from_account_id": account1.ID,
				"to_account_id":   account2.ID,
				"amount":       amount,
				"currency":     account1.Currency,
		},
		buildStubs: func(store *mockdb.MockStore) {
				
		},
	},

}
