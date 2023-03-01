package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStore_TransferTx(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	type args struct {
		ctx context.Context
		arg TransferTxParams
	}
	tests := []struct {
		name    string
		store   *Store
		args    args
		want    TransferTxResult
		wantErr bool
	}{
		{
			name: "transfer",
			store: NewStore(testDB),
			args: args{
				ctx: context.Background(),
				arg: TransferTxParams{
					FromAccountID: account1.ID,
					ToAccountId: account2.ID,
					Amount: int64(10),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.store.TransferTx(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store.TransferTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			require.NotEmpty(t, got)

			transfer := got.Transfer
			require.NotEmpty(t, transfer)
			require.Equal(t, account1.ID, transfer.FromAccountID)
			require.Equal(t, account2.ID, transfer.ToAccountID)
			require.NotZero(t, transfer.ID)
			require.NotZero(t, transfer.CreatedAt)
		})
	}
}
