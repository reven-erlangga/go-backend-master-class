// Code testing transfer sql.
// versions:
//   sqlc v1.17.2
// source: transfer.sql

package db

import (
	"context"
	"reflect"
	"testing"

	"github.com/reven-erlangga/go-backend-master-class/utils"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams{
		FromAccountID: 1,
		ToAccountID: 1,
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

func TestQueries_CreateTransfer(t *testing.T) {
	type args struct {
		ctx context.Context
		arg CreateTransferParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    Transfer
		wantErr bool
	}{
		{
			name: "Create a new transfer",
			q: testQueries,
			args: args{
				ctx: context.Background(),
				arg: CreateTransferParams{
					FromAccountID: 12,
					ToAccountID: 12,
					Amount: 31,
				},
			},
			want: Transfer{
				FromAccountID: 12,
				ToAccountID: 12,
				Amount: 31,

			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.CreateTransfer(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.CreateTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.CreateTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_DeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

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
			name: "delete an transfer",
			q:    testQueries,
			args: args{
				ctx: context.Background(),
				id:  transfer.ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.q.DeleteTransfer(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Queries.DeleteTransfer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueries_GetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

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
				id:  transfer.ID,
			},
			want: Transfer{
				ID: transfer.ID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.GetTransfer(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.GetTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_ListTransfers(t *testing.T) {
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
					Limit:  5,
					Offset: 5,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.ListTransfers(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.ListTransfers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.ListTransfers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueries_UpdateTransfer(t *testing.T) {
	type args struct {
		ctx context.Context
		arg UpdateTransferParams
	}
	tests := []struct {
		name    string
		q       *Queries
		args    args
		want    Transfer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.UpdateTransfer(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.UpdateTransfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queries.UpdateTransfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
