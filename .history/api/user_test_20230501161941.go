package api

import (
	"testing"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/utils"
)

func TestCreateUser(t *testing.T) {
	user := generateRandomUser()

	testCases := []struct {
		name          string
		accountID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{


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
