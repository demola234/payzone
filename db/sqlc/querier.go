// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Accounts, error)
	ChangePassword(ctx context.Context, arg ChangePasswordParams) (Users, error)
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entries, error)
	CreateSessions(ctx context.Context, arg CreateSessionsParams) (Sessions, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfers, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (Users, error)
	DeleteAccount(ctx context.Context, arg DeleteAccountParams) error
	GetAccount(ctx context.Context, id int64) (Accounts, error)
	GetAccountForUpdate(ctx context.Context, id int64) (Accounts, error)
	GetEntry(ctx context.Context, id int64) (Entries, error)
	GetSessions(ctx context.Context, id uuid.UUID) (Sessions, error)
	GetTransfer(ctx context.Context, id int64) (Transfers, error)
	GetUser(ctx context.Context, username string) (Users, error)
	ListAccountsByOwner(ctx context.Context, arg ListAccountsByOwnerParams) ([]Accounts, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entries, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfers, error)
	UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) (Accounts, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (Users, error)
}

var _ Querier = (*Queries)(nil)
