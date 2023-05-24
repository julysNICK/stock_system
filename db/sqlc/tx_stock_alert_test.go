package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStockAlert(t *testing.T) {

	store := NewStoreDB(testDb)
	storeRandom := CreateRandomStore(t)
	randomProduct := CreateRandomProduct(t, storeRandom)

	randomSupplier := CreateRandomSupplier(t)
	log.Println(randomSupplier)
	log.Println(randomProduct)
	arg, err := store.StockAlertTx(context.Background(), StockAlertTxParams{
		ProductID:     randomProduct.ID,
		SupplierID:    randomSupplier.ID,
		AlertQuantity: 10,
	})

	require.NoError(t, err)

	require.NotEmpty(t, arg.StockAlert)
}

func TestGetStocksAlert(t *testing.T) {

	store := NewStoreDB(testDb)
	storeRandom := CreateRandomStockAlert(t)

	arg, err := store.GetStockAlertsByIdAndBySupplierTx(context.Background(), GetStockAlertTxsParams{
		ProductID:  storeRandom.ProductID.Int64,
		SupplierID: storeRandom.SupplierID.Int64,
	})

	require.NoError(t, err)

	require.NotEmpty(t, arg.StocksAlert)
}
func TestGetStocksAlertErrorNotFoundProduct(t *testing.T) {

	store := NewStoreDB(testDb)
	storeRandom := CreateRandomStockAlert(t)

	arg, err := store.GetStockAlertsByIdAndBySupplierTx(context.Background(), GetStockAlertTxsParams{
		ProductID:  0,
		SupplierID: storeRandom.SupplierID.Int64,
	})

	require.Error(t, err)

	require.Empty(t, arg.StocksAlert)
}

func TestGetStocksAlertErrorNotFoundSupplier(t *testing.T) {

	store := NewStoreDB(testDb)
	storeRandom := CreateRandomStockAlert(t)

	arg, err := store.GetStockAlertsByIdAndBySupplierTx(context.Background(), GetStockAlertTxsParams{
		ProductID:  storeRandom.ProductID.Int64,
		SupplierID: 0,
	})

	require.Error(t, err)

	require.Empty(t, arg.StocksAlert)
}
func TestGetStocksAlertErrorNotFoundSupplierAll(t *testing.T) {

	store := NewStoreDB(testDb)
	storeRandom := CreateRandomStockAlert(t)

	arg, err := store.GetStockAlertsByIdAndBySupplierTx(context.Background(), GetStockAlertTxsParams{
		ProductID:  storeRandom.ProductID.Int64,
		SupplierID: storeRandom.SupplierID.Int64,
	})

	require.NoError(t, err)

	require.NotEmpty(t, arg.StocksAlert)
}
