package db

import (
	"context"
	"database/sql"
	"fmt"
)

type StoreDB interface {
	Querier
	SaleTx(ctx context.Context, arg SaleTxParams) (SaleTxResult, error)
	StockAlertTx(ctx context.Context, arg StockAlertTxParams) (StockAlertTxResult, error)
	ProductTx(ctx context.Context, arg ProductTxParams) (ProductTxResult, error)
	ProductBuyTx(ctx context.Context, argTX ProductBuyTxParams) (ProductBuyTxResult, error)
	GetStockAlertsByIdAndBySupplierTx(ctx context.Context, arg GetStockAlertTxsParams) (GetStockAlertTxsResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStoreDB(db *sql.DB) StoreDB {
	return &SQLStore{
		db:      db,
		Queries: New(db), // this New() is for the queries we defined in sqlc
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q) // execute the function passed in
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rbErr) // if rollback fails, return both errors
		}
		return err
	}
	return tx.Commit()
}
