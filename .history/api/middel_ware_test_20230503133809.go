package api

import "testing"

func TestAuthMiddleWare(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		set
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}
}
