package db

import (
	"context"
	"database/sql"
)

type StockAlertTxParams struct {
	ProductID     int64 `json:"product_id"`
	SupplierID    int64 `json:"supplier_id"`
	AlertQuantity int32 `json:"alert_quantity"`
}

type StockAlertTxResult struct {
	StockAlert StockAlert `json:"stock_alert"`
	Product    Product    `json:"product"`
	Supplier   Supplier   `json:"supplier"`
}

func (store *StoreDB) StockAlertTx(ctx context.Context, arg StockAlertTxParams) (StockAlertTxResult, error) {
	var result StockAlertTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Product, err = q.GetProduct(ctx, arg.ProductID)
		if err != nil {

			return err
		}

		result.Supplier, err = q.GetSupplier(ctx, arg.SupplierID)
		if err != nil {

			return err
		}

		result.StockAlert, err = q.CreateStockAlert(ctx, CreateStockAlertParams{
			ProductID: sql.NullInt64{
				Int64: arg.ProductID,
				Valid: true,
			},
			SupplierID: sql.NullInt64{
				Int64: arg.SupplierID,
				Valid: true,
			},
			AlertQuantity: arg.AlertQuantity,
		})

		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
