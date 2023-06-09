package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/demola234/payzone/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) Users {
	hashPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.FullName, user2.FullName)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.WithinDuration(t, user.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, user2.CreatedAt, time.Second)
}

func TestChangePassword(t *testing.T) {
	user := createRandomUser(t)
	arg := ChangePasswordParams{
		Username:          user.Username,
		HashedPassword:    utils.RandomString(6),
		PasswordChangedAt: time.Now(),
	}

	user, err := testQueries.ChangePassword(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.NotZero(t, user.PasswordChangedAt)
	require.WithinDuration(t, user.PasswordChangedAt, time.Now(), time.Second)
}

func TestGetUserByEmail(t *testing.T) {
	user := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.Email, user2.Email)
	require.Equal(t, user.FullName, user2.FullName)
	require.Equal(t, user.HashedPassword, user2.HashedPassword)
	require.WithinDuration(t, user.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, user2.CreatedAt, time.Second)
}

func TestCheckUserExists(t *testing.T) {
	user := createRandomUser(t)
	exists, err := testQueries.CheckUsernameExists(context.Background(), user.Username)

	require.NoError(t, err)
	require.True(t, exists)
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)
	arg := UpdateUserParams{
		Username: user.Username,
		Email: sql.NullString{
			String: utils.RandomEmail(),
			Valid:  true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, updatedUser.Username)
	require.Equal(t,  user.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t,  user.FullName, updatedUser.FullName)
}
