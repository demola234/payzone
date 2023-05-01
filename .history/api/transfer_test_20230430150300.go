package api

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateTransfer(t *testing.T) {
	// Create 2 new account
	account1 := generateRandomAccount()
	account2 := generateRandomAccount()

	testCases := []struct {
		name        string
		body        gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(recoder *httptest.ResponseRecorder)

	}{
		{
	}

}
