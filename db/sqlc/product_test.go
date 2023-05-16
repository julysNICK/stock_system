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

func CreateRandomProduct(t *testing.T, store Store) Product {
	arg := CreateProductParams{
		StoreID:     store.ID,
		Name:        utils.RandomName(),
		Price:       "100",
		Description: utils.RandomString(10),
		Quantity:    10,
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, product)

	require.Equal(t, arg.StoreID, product.StoreID)
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Price, product.Price)
	require.Equal(t, arg.Description, product.Description)
	require.NotZero(t, product.ID)
	require.NotZero(t, product.CreatedAt)
	return product
}

func TestCreateProduct(t *testing.T) {
	store := CreateRandomStore(t)
	CreateRandomProduct(t, store)
}

func TestGetProducts(t *testing.T) {
	mockTimer := time.Now()

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "quantity", "store_id", "created_at"}).
		AddRow(1, "test", "test", "100", 10, 1, mockTimer)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, description, price, quantity, store_id, created_at FROM products WHERE id = $1 LIMIT 1`)).
		WithArgs(1).
		WillReturnRows(rows)
	mockDb := New(db)

	product, err := mockDb.GetProduct(context.Background(), 1)

	require.NoError(t, err)

	require.NotEmpty(t, product)

	require.Equal(t, int64(1), product.ID)

}

func TestGetProductForUpdate(t *testing.T) {
	mockTimer := time.Now()

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "quantity", "store_id", "created_at"}).
		AddRow(1, "test", "test", "100", 10, 1, mockTimer)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, description, price, quantity, store_id, created_at FROM products WHERE id = $1 LIMIT 1 FOR UPDATE`)).
		WithArgs(1).
		WillReturnRows(rows)
	mockDb := New(db)

	product, err := mockDb.GetProductForUpdate(context.Background(), 1)

	require.NoError(t, err)

	require.NotEmpty(t, product)

	require.Equal(t, int64(1), product.ID)
}

func TestUpdateProduct(t *testing.T) {

	mockTimer := time.Now()

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "quantity", "store_id", "created_at"}).
		AddRow(1, "test", "test", "100", 10, 1, mockTimer)

	mock.ExpectQuery(regexp.QuoteMeta(`UPDATE products SET name = COALESCE($2, name), description = COALESCE($3, description), price = COALESCE($4, price), quantity = COALESCE($5, quantity) WHERE id = $1 RETURNING id, name, description, price, quantity, store_id, created_at`)).
		WithArgs(1, "test", "test", "100", 10).
		WillReturnRows(rows)

	mockDb := New(db)

	arg := UpdateProductParams{
		ID: 1,
		Name: sql.NullString{
			String: "test",
			Valid:  true,
		},
		Description: sql.NullString{
			String: "test",
			Valid:  true,
		},
		Price: sql.NullString{
			String: "100",
			Valid:  true,
		},
		Quantity: sql.NullInt32{
			Int32: 10,
			Valid: true,
		},
	}

	product, err := mockDb.UpdateProduct(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, product)

}

func TestListProducts(t *testing.T) {

	mockTimer := time.Now()

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "quantity", "store_id", "created_at"}).
		AddRow(1, "test", "test", "100", 10, 1, mockTimer)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, description, price, quantity, store_id, created_at FROM products ORDER BY id LIMIT $1 OFFSET $2`)).
		WithArgs(1, 0).
		WillReturnRows(rows)

	mockDb := New(db)

	arg := ListProductsParams{
		Limit:  1,
		Offset: 0,
	}

	products, err := mockDb.ListProducts(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, products)

}

func TestListProduct_ErrorInQueryContext(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, description, price, quantity, store_id, created_at FROM products ORDER BY id LIMIT $1 OFFSET $2`)).
		WithArgs(1, 0).
		WillReturnError(fmt.Errorf("some error"))

	mockDb := New(db)

	arg := ListProductsParams{
		Limit:  1,
		Offset: 0,
	}

	products, err := mockDb.ListProducts(context.Background(), arg)

	require.Error(t, err)

	require.Empty(t, products)

}

func TestListProduct_ErrorInScan(t *testing.T) {

	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "quantity", "store_id", "created_at"}).
		AddRow(1, "test", "test", 100, 10, 1, "2020-01-01 00:00:00")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, description, price, quantity, store_id, created_at FROM products ORDER BY id LIMIT $1 OFFSET $2`)).
		WithArgs(1, 0).
		WillReturnRows(rows)

	mock.ExpectClose()

	mockDb := New(db)

	arg := ListProductsParams{
		Limit:  1,
		Offset: 0,
	}

	products, err := mockDb.ListProducts(context.Background(), arg)

	require.Error(t, err)

	require.Empty(t, products)
}

func TestListProduct_RowsErr(t *testing.T) {
	// mockTimer := time.Now()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "quantity", "store_id", "created_at"}).CloseError(fmt.Errorf("some error"))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, description, price, quantity, store_id, created_at FROM products ORDER BY id LIMIT $1 OFFSET $2`)).
		WithArgs(1, 0).
		WillReturnRows(rows)

	mockDb := New(db)

	arg := ListProductsParams{
		Limit:  1,
		Offset: 0,
	}

	products, err := mockDb.ListProducts(context.Background(), arg)

	require.Error(t, err)

	require.Empty(t, products)
}

func TestListProduct_RowsErrClose(t *testing.T) {
	// mockTimer := time.Now()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "quantity", "store_id", "created_at"}).CloseError(fmt.Errorf("some error close"))

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, description, price, quantity, store_id, created_at FROM products ORDER BY id LIMIT $1 OFFSET $2`)).
		WithArgs(1, 0).
		WillReturnRows(rows)

	mockDb := New(db)

	_, err = mockDb.ListProducts(context.Background(), ListProductsParams{
		Limit:  1,
		Offset: 0,
	})

	require.EqualError(t, err, "some error close")
}
