package api

import (
	"testing"

	db "github.com/demola234/payzone/db/sqlc"
)

// Server serves HTTP requests for our banking service.

func TestGetAccount(t *testing.T) {

}

func generateRandomAccount() Account {
	return db.Accounts
}