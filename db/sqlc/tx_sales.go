package db

import (
	"context"
	"fmt"
	"time"
)

type SaleTxParams struct {
	ProductID    int64     `json:"product_id"`
	SaleDate     time.Time `json:"sale_date"`
	QuantitySold int32     `json:"quantity_sold"`
}

type SaleTxResult struct {
	Sale    Sale    `json:"sales"`
	Product Product `json:"product"`
}

func (store *StoreDB) SaleTx(ctx context.Context, arg SaleTxParams) (SaleTxResult, error) {
	var result SaleTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		if arg.ProductID <= 0 {
			return fmt.Errorf(ErrorProductNotFound)
		}

		if arg.QuantitySold <= 0 {
			return fmt.Errorf("quantity sold must be greater than 0")
		}

		if arg.SaleDate.IsZero() {
			return fmt.Errorf("sale date can't be empty")
		}

		if arg.QuantitySold <= 0 {
			return fmt.Errorf("quantity sold must be greater than 0")
		}

		result.Product, err = q.GetProduct(ctx, arg.ProductID)

		if err != nil {
			return err
		}

		result.Sale, err = q.CreateSale(ctx, CreateSaleParams{
			ProductID:    arg.ProductID,
			SaleDate:     arg.SaleDate,
			QuantitySold: arg.QuantitySold,
		})

		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
