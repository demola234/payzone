package api

import "testing"

func TestAuthMiddleWare(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		
		body          gin.H
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}
}
