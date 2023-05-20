package db

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func CreateRandomStockAlert(t *testing.T) StockAlert {

	storeRandom := CreateRandomStore(t)
	productRandom := CreateRandomProduct(t, storeRandom)
	supplierRandom := CreateRandomSupplier(t)

	arg := CreateStockAlertParams{
		ProductID: sql.NullInt64{
			Int64: productRandom.ID,
			Valid: true,
		},
		SupplierID: sql.NullInt64{
			Int64: supplierRandom.ID,
			Valid: true,
		},
		AlertQuantity: 10,
	}

	stockAlert, err := testQueries.CreateStockAlert(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, stockAlert)

	require.Equal(t, arg.ProductID, stockAlert.ProductID)

	require.Equal(t, arg.SupplierID, stockAlert.SupplierID)

	require.Equal(t, arg.AlertQuantity, stockAlert.AlertQuantity)

	require.NotZero(t, stockAlert.ID)

	require.NotZero(t, stockAlert.CreatedAt)

	return stockAlert
}

func TestCreateStockAlert(t *testing.T) {
	CreateRandomStockAlert(t)
}

func TestGetStockAlert(t *testing.T) {
	mockTimer := time.Now()

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "product_id", "supplier_id", "alert_quantity", "created_at"}).
		AddRow(1, 1, 1, 10, mockTimer)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, product_id, supplier_id, alert_quantity, created_at FROM stock_alerts WHERE id = $1 LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(rows)

	testQueries := New(db)

	stockAlert, err := testQueries.GetStockAlert(context.Background(), 1)

	require.NoError(t, err)

	require.NotEmpty(t, stockAlert)
}

func TestDeleteStockAlert(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM stock_alerts WHERE id = $1`)).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	testQueries := New(db)

	err = testQueries.DeleteStockAlert(context.Background(), 1)

	require.NoError(t, err)

}

func TestUpdateStockAlert(t *testing.T) {
	mockTimer := time.Now()

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	arg := UpdateStockAlertParams{
		ID: int64(1),
		AlertQuantity: sql.NullInt32{
			Int32: int32(10),
			Valid: true,
		},
	}

	rows := sqlmock.NewRows([]string{"id", "product_id", "supplier_id", "alert_quantity", "created_at"}).
		AddRow(1, 1, 1, 10, mockTimer)

	mock.ExpectQuery(regexp.QuoteMeta(` UPDATE stock_alerts SET alert_quantity = COALESCE($2, alert_quantity) WHERE id = $1 RETURNING id, product_id, supplier_id, alert_quantity, created_at`)).
		WithArgs(arg.ID, arg.AlertQuantity).
		WillReturnRows(rows)

	testQueries := New(db)

	stockAlert, err := testQueries.UpdateStockAlert(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, stockAlert)

}

func TestListStockAlerts(t *testing.T) {
	StockAlertRandom := CreateRandomStockAlert(t)
	arg := GetStockAlertsByProductIdAndSupplierIdParams{
		ProductID: sql.NullInt64{
			Int64: StockAlertRandom.ProductID.Int64,
			Valid: true,
		},
		SupplierID: sql.NullInt64{
			Int64: StockAlertRandom.SupplierID.Int64,
			Valid: true,
		},
	}

	fmt.Println(arg.ProductID, arg.SupplierID)
	stores, err := testQueries.GetStockAlertsByProductIdAndSupplierId(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, stores)

}
func TestListStockAlerts_ErrorInContext(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, product_id, supplier_id, alert_quantity, created_at FROM stock_alerts WHERE product_id = $1 AND supplier_id = $2`)).
		WithArgs(1, 1).
		WillReturnError(fmt.Errorf("some error"))

	mockDb := New(db)

	arg := GetStockAlertsByProductIdAndSupplierIdParams{
		ProductID: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
		SupplierID: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
	}

	StocksAlerts, err := mockDb.GetStockAlertsByProductIdAndSupplierId(context.Background(), arg)

	require.Error(t, err)

	require.Empty(t, StocksAlerts)

}

func TestListStockAlerts_ErrorInScan(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "product_id", "supplier_id", "alert_quantity", "created_at"}).
		AddRow(1, 1, 1, 10, "2020-01-01 00:00:00")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, product_id, supplier_id, alert_quantity, created_at FROM stock_alerts WHERE product_id = $1 AND supplier_id = $2`)).
		WithArgs(1, 1).
		WillReturnRows(rows)

	mock.ExpectClose()

	mockDb := New(db)

	arg := GetStockAlertsByProductIdAndSupplierIdParams{
		ProductID: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
		SupplierID: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
	}

	stocksAlerts, err := mockDb.GetStockAlertsByProductIdAndSupplierId(context.Background(), arg)

	require.Error(t, err)

	require.Empty(t, stocksAlerts)

}

func TestListStockAlerts_RowsErr(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "product_id", "supplier_id", "alert_quantity", "created_at"}).
		CloseError(fmt.Errorf("bad rows"))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, product_id, supplier_id, alert_quantity, created_at FROM stock_alerts WHERE product_id = $1 AND supplier_id = $2`)).
		WithArgs(1, 1).
		WillReturnRows(rows)

	mockDb := New(db)

	arg := GetStockAlertsByProductIdAndSupplierIdParams{
		ProductID: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
		SupplierID: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
	}

	stocksAlerts, err := mockDb.GetStockAlertsByProductIdAndSupplierId(context.Background(), arg)

	require.Error(t, err)

	require.Empty(t, stocksAlerts)
}
