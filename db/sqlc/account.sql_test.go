// Code testing account sql.
// versions:
//   sqlc v1.17.2
// source: account.sql

package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueries_GetAccount(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.GetAccount(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			
			assert.Equal(t, tt.want, got)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.q.ListAccounts(tt.args.ctx, tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Queries.ListAccounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
