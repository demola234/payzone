package db

import (
	"context"
	"testing"
	"time"

	"github.com/demola234/payzone/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) Users {
	arg := CreateUserParams{
		Username: utils.RandomOwner(),
		HashedPassword: "secret",
		FullName: utils.RandomOwner(),
		Email: utils.RandomEmail(),
	}

	account, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.FullName, account.FullName)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.HashedPassword, account.HashedPassword)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateUser(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}
