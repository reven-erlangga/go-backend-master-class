// Code testing account sql.
// versions:
//   sqlc v1.17.2
// source: account.sql

package db

import (
	"context"
	"testing"

	"github.com/reven-erlangga/go-backend-master-class/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	if err != nil {
		t.Errorf("Cannot create random account")
	}

	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestQueries_CreateAccount(t *testing.T) {
	type args struct {
		ctx context.Context
		arg CreateAccountParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    Account
		wantErr bool
	}{
		{
			name: "create a new account",
			q:    testQueries,
			args: args{
				ctx: context.Background(),
				arg: CreateAccountParams{
					Owner:    "Tom",
					Balance:  100,
					Currency: "USD",
				},
			},
			want: Account{
				Owner:    "Tom",
				Balance:  100,
				Currency: "USD",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := tt.q.CreateAccount(tt.args.ctx, tt.args.arg)

			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			actual := Account{
				Owner:    account.Owner,
				Balance:  account.Balance,
				Currency: account.Currency,
			}

			expected := Account{
				Owner:    tt.want.Owner,
				Balance:  tt.want.Balance,
				Currency: tt.want.Currency,
			}

			require.NotEmpty(t, account)
			require.EqualValues(t, actual, expected)
			require.NotZero(t, account.ID)
		})
	}
}

func TestQueries_GetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    Account
		wantErr bool
	}{
		{
			name: "get single account",
			q:    testQueries,
			args: args{
				ctx: context.Background(),
				id:  account1.ID,
			},
			want: Account{
				ID: account1.ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := tt.q.GetAccount(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.NotEmpty(t, account)
			require.Equal(t, account.ID, tt.want.ID)
		})
	}
}

func TestQueries_ListAccounts(t *testing.T) {
	type args struct {
		ctx context.Context
		arg ListAccountsParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    []Account
		wantErr bool
	}{
		{
			name: "list all account",
			q:    testQueries,
			args: args{
				ctx: context.Background(),
				arg: ListAccountsParams{
					Limit:  5,
					Offset: 5,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.ListAccounts(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.ListAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.NotEmpty(t, got)
		})
	}
}

func BenchmarkQueries_ListAccount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testQueries.ListAccounts(context.Background(), ListAccountsParams{
			Limit:  5,
			Offset: 5,
		})
	}
}

func TestQueries_DeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		wantErr bool
	}{
		{
			name: "delete an account",
			q:    testQueries,
			args: args{
				ctx: context.Background(),
				id:  account.ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.DeleteAccount(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Queries.DeleteAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueries_UpdateAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)

	type args struct {
		ctx context.Context
		arg UpdateAccountParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    Account
		wantErr bool
	}{
		{
			name: "update an account",
			q: testQueries,
			args: args{
				ctx: context.Background(),
				arg: UpdateAccountParams{
					ID: randomAccount.ID,
					Balance: 350,
				},
			},
			want: Account{
				ID: randomAccount.ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account, err := tt.q.UpdateAccount(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.UpdateAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			require.NotEmpty(t, account)
			require.Equal(t, account.ID, tt.want.ID)
		})
	}
}
