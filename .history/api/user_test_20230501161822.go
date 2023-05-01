package api

import (
	"testing"

	db "github.com/demola234/payzone/db/sqlc"
	"github.com/demola234/payzone/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {

}

func generateRandomUser() db.Users {
	hashPassword, _ := utils.HashPassword(utils.RandomString(6))
	return db.Users{
		Username: utils.RandomOwner(),
		FullName: utils.RandomOwner(),
		Email: utils.RandomEmail(),
		HashedPassword: hashPassword,
	}
}
