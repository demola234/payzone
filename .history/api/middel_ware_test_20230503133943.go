package api

import (
	"testing"

	"github.com/demola234/payzone/token"
)

func TestAuthMiddleWare(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		setupAuth    func(t *testing.T, request *http.Request, token token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	} {
		{
		},
	}

	for i := range testCases {
		tc := testCases[]
	}
}
