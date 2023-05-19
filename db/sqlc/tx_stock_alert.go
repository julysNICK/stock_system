package db

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	ErrorProductNotFound  = "product not found"
	ErrorSupplierNotFound = "supplier not found"
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

func (store *SQLStore) StockAlertTx(ctx context.Context, arg StockAlertTxParams) (StockAlertTxResult, error) {
	var result StockAlertTxResult

	err := store.execTx(ctx, func(q *Queries) error {

		var err error

		if arg.ProductID <= 0 {
			return fmt.Errorf(ErrorProductNotFound)
		}

		if arg.SupplierID <= 0 {
			return fmt.Errorf(ErrorSupplierNotFound)
		}

		if arg.AlertQuantity <= 0 {
			return fmt.Errorf("alert quantity must be greater than 0")
		}

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

type GetStockAlertTxsParams struct {
	ProductID  int64 `json:"product_id"`
	SupplierID int64 `json:"supplier_id"`
}

type GetStockAlertTxsResult struct {
	StocksAlert []StockAlert `json:"stock_alert"`
	Product     Product      `json:"product"`
	Supplier    Supplier     `json:"supplier"`
}

func (store *SQLStore) GetStockAlertsByIdAndBySupplierTx(ctx context.Context, arg GetStockAlertTxsParams) (GetStockAlertTxsResult, error) {
	var result GetStockAlertTxsResult

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
		arg := GetStockAlertsByProductIdAndSupplierIdParams{
			ProductID: sql.NullInt64{
				Int64: arg.ProductID,
				Valid: true,
			},
			SupplierID: sql.NullInt64{
				Int64: arg.SupplierID,
				Valid: true,
			},
		}

		result.StocksAlert, err = q.GetStockAlertsByProductIdAndSupplierId(ctx, arg)

		if err != nil {
			return err
		}
		return nil
	})

	return result, err
}
