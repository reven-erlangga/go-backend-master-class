// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"database/sql"
)

type Account struct {
	ID        int64          `json:"id"`
	Owner     sql.NullString `json:"owner"`
	Balance   sql.NullInt64  `json:"balance"`
	Currency  string         `json:"currency"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// cant be negative or positif
	Amount    int64        `json:"amount"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64        `json:"amount"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}