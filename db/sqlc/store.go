package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier

	TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := store.WithTx(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("trx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTrxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTrxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTrx performs a money transfer from one account to the other.
// It creates a transfer record, add account entries, and update accounts' balance within a single transaction
func (store *SQLStore) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: args.FromAccountID,
			ToAccountID:   args.ToAccountID,
			Amount:        args.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.FromAccountID,
			Amount:    -args.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: args.ToAccountID,
			Amount:    args.Amount,
		})
		if err != nil {
			return err
		}

		// get account -> update its balance
		if args.FromAccountID < args.ToAccountID {
			result.FromAccount, result.ToAccount, err = addMoney(
				ctx,
				q,
				args.FromAccountID,
				-args.Amount,
				args.ToAccountID,
				args.Amount,
			)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(
				ctx,
				q,
				args.ToAccountID,
				args.Amount,
				args.FromAccountID,
				-args.Amount,
			)
		}

		return err
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	return
}
