package db

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/julysNICK/stock_system/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomSale(t *testing.T) Sale {
	storeRandom := CreateRandomStore(t)
	productRandom := CreateRandomProduct(t, storeRandom)

	arg := CreateSaleParams{

		SaleDate:     utils.RandomDate(),
		ProductID:    productRandom.ID,
		QuantitySold: int32(utils.RandomInt(1, 100)),
	}

	sale, err := testQueries.CreateSale(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, sale)
	return sale
}

func TestCreateSale(t *testing.T) {
	CreateRandomSale(t)
}

func TestGetSale(t *testing.T) {
	sale1 := CreateRandomSale(t)

	sale2, err := testQueries.GetSale(context.Background(), sale1.ID)

	require.NoError(t, err)

	require.NotEmpty(t, sale2)

	require.Equal(t, sale1.ID, sale2.ID)

	require.Equal(t, sale1.SaleDate, sale2.SaleDate)

	require.Equal(t, sale1.ProductID, sale2.ProductID)

	require.Equal(t, sale1.QuantitySold, sale2.QuantitySold)
}

func TestDeleteSale(t *testing.T) {
	sale1 := CreateRandomSale(t)

	err := testQueries.DeleteSale(context.Background(), sale1.ID)

	require.NoError(t, err)

	sale2, err := testQueries.GetSale(context.Background(), sale1.ID)

	require.Error(t, err)

	require.EqualError(t, err, sql.ErrNoRows.Error())

	require.Empty(t, sale2)
}

// func TestListSale(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		CreateRandomSale(t)
// 	}

// 	arg := ListSalesParams{
// 		Limit:  5,
// 		Offset: 5,
// 	}

// 	sales, err := testQueries.ListSales(context.Background(), arg)

// 	require.NoError(t, err)

// 	require.Len(t, sales, 5)

// 	for _, sale := range sales {
// 		require.NotEmpty(t, sale)
// 	}
// }

func TestListSaleWithMock(t *testing.T) {
	mockTimer := time.Now()
	mockDate := utils.RandomDate()

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	arg := ListSalesParams{
		Limit:  5,
		Offset: 5,
	}
	rows := sqlmock.NewRows([]string{"id", "product_id", "sale_date", "quantity_sold", "created_at"}).
		AddRow(1, 1, mockDate, 10, mockTimer)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, product_id, sale_date, quantity_sold, created_at FROM sales ORDER BY id LIMIT $1 OFFSET $2`)).
		WithArgs(arg.Limit, arg.Offset).
		WillReturnRows(rows)

	mockDb := New(db)

	sales, err := mockDb.ListSales(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, sales)

}

func TestListSale_ErrorInQueryContext(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, product_id, sale_date, quantity_sold, created_at FROM sales ORDER BY id LIMIT $1 OFFSET $2`)).
		WithArgs(1, 0).
		WillReturnError(fmt.Errorf("some error"))

	mockDb := New(db)

	arg := ListSalesParams{
		Limit:  1,
		Offset: 0,
	}

	products, err := mockDb.ListSales(context.Background(), arg)

	require.Error(t, err)

	require.Empty(t, products)

}

func TestListSale_ErrorInScan(t *testing.T) {

	mockDate := utils.RandomDate()
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "product_id", "sale_date", "quantity_sold", "created_at"}).
		AddRow(1, 1, mockDate, 10, "2020-01-01 00:00:00")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, product_id, sale_date, quantity_sold, created_at FROM sales ORDER BY id LIMIT $1 OFFSET $2`)).
		WithArgs(1, 0).
		WillReturnRows(rows)

	mock.ExpectClose()

	mockDb := New(db)

	arg := ListSalesParams{
		Limit:  1,
		Offset: 0,
	}

	products, err := mockDb.ListSales(context.Background(), arg)

	require.Error(t, err)

	require.Empty(t, products)
}

func TestListSale_RowsErr(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "product_id", "sale_date", "quantity_sold", "created_at"}).CloseError(fmt.Errorf("some error"))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, product_id, sale_date, quantity_sold, created_at FROM sales ORDER BY id LIMIT $1 OFFSET $2`)).
		WithArgs(1, 0).
		WillReturnRows(rows)

	mockDb := New(db)

	arg := ListSalesParams{
		Limit:  1,
		Offset: 0,
	}

	products, err := mockDb.ListSales(context.Background(), arg)

	require.Error(t, err)

	require.Empty(t, products)
}
