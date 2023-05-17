package db

import (
	"context"
	"time"
)

type SaleTxParams struct {
	ProductID int64 `json:"product_id"`
	SaleDate time.Time `json:"sale_date"`
	QuantitySold int32 `json:"quantity_sold"`
}

type SaleTxResult struct {
	Sale Sale `json:"sales"`
	Product Product `json:"product"`
}

func (store *StoreDB) SaleTx(ctx context.Context, arg SaleTxParams) (SaleTxResult, error){
	var result SaleTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Product, err = q.GetProduct(ctx, arg.ProductID)

		if err != nil {
			return err
		}

		result.Sale, err = q.CreateSale(ctx, CreateSaleParams{
			ProductID: arg.ProductID,
			SaleDate: arg.SaleDate,
			QuantitySold: arg.QuantitySold,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}