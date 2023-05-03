package api

import "testing"

func TestAuthMiddleWare(t *testing.T) {
	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		setupAuth    func(t *testing.T) (string, error)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}
}
