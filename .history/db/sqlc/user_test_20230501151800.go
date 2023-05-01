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

	users, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, users)
	require.Equal(t, arg.Username, users.Username)
	require.Equal(t, arg.FullName, users.FullName)
	require.Equal(t, arg.Email, users.Email)
	require.Equal(t, arg.HashedPassword, users.HashedPassword)
	require.NotZero(t, users.CreatedAt)

	return users
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	users := createRandomUser(t)
	users2, err := testQueries.Getusers(context.Background(), users.ID)

	require.NoError(t, err)
	require.NotEmpty(t, users2)
	require.Equal(t, users.ID, users2.ID)
	require.Equal(t, users.Owner, users2.Owner)
	require.Equal(t, users.Balance, users2.Balance)
	require.Equal(t, users.Currency, users2.Currency)
	require.WithinDuration(t, users.CreatedAt, users2.CreatedAt, time.Second)
}
