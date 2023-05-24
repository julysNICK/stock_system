package cron

import (
	"context"
	"database/sql"
	"fmt"

	db "github.com/julysNICK/stock_system/db/sqlc"
)

func CheckStockAlerts(store db.StoreDB) {

	threshHold := 11
	args := db.GetStockAlertsByProductIdAndSupplierIdParams{
		ProductID: sql.NullInt64{
			Int64: 7,
			Valid: true,
		},
		SupplierID: sql.NullInt64{
			Int64: 2,
			Valid: true,
		},
	}
	fmt.Println("here 24")
	StocksAlerts, err := store.GetStockAlertsByProductIdAndSupplierId(context.Background(), args)
	fmt.Println("here 26")
	if err != nil {
		fmt.Printf("Error getting stock alerts: %s", err.Error())
		return
	}
	fmt.Println("here 31")
	for _, product := range StocksAlerts {
		if int(product.AlertQuantity) <= threshHold {
			// send email
			fmt.Printf("Product %d is running low on stock. Current stock is %d sending email for supplier", product.ProductID.Int64, product.AlertQuantity)
		}
	}
}
