package db

import (
	"context"
	"database/sql"
	"fmt"
)


type StoreDB struct {
	db *sql.DB
	*Queries
}

func NewStoreDB(db *sql.DB) *StoreDB {
	return &StoreDB{
		db:      db,
		Queries: New(db), // this New() is for the queries we defined in sqlc
	}
}	

func (store *StoreDB) execTx(ctx context.Context, fn func(*Queries) error) error {
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







