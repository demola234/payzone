package db

import (
	"context"

	"testing"
	"time"

	"github.com/demola234/payzone/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Accounts {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
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

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountBalanceParams{
		ID:      account.ID,
		Balance: utils.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccountBalance(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}

func TestGetAccountForUpdate(t *testing.T) {
	account := createRandomAccount(t)

	account2, err := testQueries.GetAccountForUpdate(context.Background(), account.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, account.Owner, account2.Owner)
	require.Equal(t, account.Balance, account2.Balance)
	require.Equal(t, account.Currency, account2.Currency)
	require.WithinDuration(t, account.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, account2)
}


func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func createAListRandomAccounts(t *testing.T, n int) []Accounts {
	var accounts []Accounts
	for i := 0; i < n; i++ {
		accounts = append(accounts, createRandomAccount(t))
	}
	return accounts
}

// Create a user and generate a list of USD, EUR, NGN, GBP accounts for that user
func createRandomAccountsForUser(t *testing.T, user Users) []Accounts {
	var accounts []Accounts
	for i := 0; i < 4; i++ {
		arg := CreateAccountParams{
			Owner:    user.Username,
			Balance:  utils.RandomMoney(),
			Currency: utils.RandomCurrency(),
		}

		account, err := testQueries.CreateAccount(context.Background(), arg)
		if err != nil {
			t.Fatal(err)
		}

		require.NoError(t, err)
		require.NotEmpty(t, account)
		require.Equal(t, arg.Owner, account.Owner)
		require.Equal(t, arg.Balance, account.Balance)
		require.Equal(t, arg.Currency, account.Currency)
		require.NotZero(t, account.CreatedAt)

		accounts = append(accounts, account)
	}
	return accounts
}

func TestListAccountsForUser(t *testing.T) {
	
}