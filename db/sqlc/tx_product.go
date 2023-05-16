package db

import (
	"context"
	"fmt"
)

type ProductTxParams struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Price int64 `json:"price"`
	StoreID int64 `json:"store_id"`
	Quantity int32 `json:"quantity"`
}

type ProductTxResult struct {
	Product Product `json:"product"`
	StoreID Store `json:"store_id"`
	
}

func (store *StoreDB) ProductTx(ctx context.Context, arg ProductTxParams) (ProductTxResult, error) {
	var result ProductTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Product, err = q.CreateProduct(ctx, CreateProductParams{
			Name: arg.Name,
			Description: arg.Description,
			Price: fmt.Sprintf("%d", arg.Price),
			StoreID: arg.StoreID,
			Quantity: arg.Quantity,
		})

		if err != nil {
			return err
		}

		result.StoreID, err = q.GetStore(ctx, arg.StoreID)

		if err != nil {
			return err
		}

		return nil

	})
	return result, err
} 