package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/julysNICK/stock_system/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomSupplier(t *testing.T) Supplier {
	arg := CreateSupplierParams{
		Name:         utils.RandomName(),
		Email:        utils.RandomEmail(),
		ContactPhone: utils.RandomPhone(),
	}

	supplier, err := testQueries.CreateSupplier(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, supplier)

	require.Equal(t, arg.Name, supplier.Name)

	require.Equal(t, arg.Email, supplier.Email)

	require.Equal(t, arg.ContactPhone, supplier.ContactPhone)

	require.NotZero(t, supplier.ID)

	require.NotZero(t, supplier.CreatedAt)

	return supplier

}

func TestCreateSupplier(t *testing.T) {

	CreateRandomSupplier(t)

}

func TestGetSupplier(t *testing.T) {

	supplier1 := CreateRandomSupplier(t)

	supplier2, err := testQueries.GetSupplier(context.Background(), supplier1.ID)

	require.NoError(t, err)

	require.NotEmpty(t, supplier2)

	require.Equal(t, supplier1.ID, supplier2.ID)

	require.Equal(t, supplier1.Name, supplier2.Name)

	require.Equal(t, supplier1.Email, supplier2.Email)

	require.Equal(t, supplier1.ContactPhone, supplier2.ContactPhone)

	require.WithinDuration(t, supplier1.CreatedAt, supplier2.CreatedAt, time.Second)
}

func TestUpdateSupplier(t *testing.T) {
	supplier1 := CreateRandomSupplier(t)

	arg := UpdateSupplierParams{
		ID: supplier1.ID,
		Name: sql.NullString{
			String: utils.RandomName(),
			Valid:  true,
		},
		Email: sql.NullString{
			String: utils.RandomEmail(),
			Valid:  true,
		},
		ContactPhone: sql.NullString{
			String: utils.RandomPhone(),
			Valid:  true,
		},
	}

	supplier2, err := testQueries.UpdateSupplier(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, supplier2)
	require.Equal(t, arg.Name.String, supplier2.Name)

	require.Equal(t, arg.Email.String, supplier2.Email)

	require.Equal(t, arg.ContactPhone.String, supplier2.ContactPhone)

	require.Equal(t, supplier1.ID, supplier2.ID)

	require.WithinDuration(t, supplier1.CreatedAt, supplier2.CreatedAt, time.Second)

}
