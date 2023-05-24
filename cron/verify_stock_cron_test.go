package cron

import (
	"database/sql"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/julysNICK/stock_system/db/mock"
	db "github.com/julysNICK/stock_system/db/sqlc"
	"github.com/julysNICK/stock_system/utils"
)

func TestVerifyStockCron(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	store := mockdb.NewMockStoreDB(ctrl)
	stocks := []db.StockAlert{
		{
			ID: 1,
			ProductID: sql.NullInt64{
				Int64: 7,
				Valid: true,
			},
			SupplierID: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
			AlertQuantity: 10,
			CreatedAt:     utils.RandomDate(),
		},
	}

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

	store.EXPECT().GetStockAlertsByProductIdAndSupplierId(gomock.Any(), args).Times(1).Return(stocks, nil)

}
