// Code testing transfer sql.
// versions:
//   sqlc v1.17.2
// source: transfer.sql

package db

import (
	"context"
	"testing"

	"github.com/reven-erlangga/go-backend-master-class/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount: utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	if err != nil {
		t.Errorf("Cannot create random transfer")
	}

	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestQueries_GetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transferData := createRandomTransfer(t, account1, account2)

	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    Transfer
		wantErr bool
	}{
		{
			name: "get single transfer",
			q:    testQueries,
			args: args{
				ctx: context.Background(),
				id:  transferData.ID,
			},
			want: Transfer{
				ID: transferData.ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transfer, err := tt.q.GetTransfer(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			require.NotEmpty(t, transfer)
		})
	}
}

func TestQueries_ListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	type args struct {
		ctx context.Context
		arg ListTransfersParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    []Transfer
		wantErr bool
	}{
		{
			name: "list all transfer",
			q:    testQueries,
			args: args{
				ctx: context.Background(),
				arg: ListTransfersParams{
					
		FromAccountID: account1.ID,
		ToAccountID:   account1.ID,
					Limit:         5,
					Offset:        5,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transfer, err := tt.q.ListTransfers(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.ListTransfers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			require.NotEmpty(t, transfer)
		})
	}
}