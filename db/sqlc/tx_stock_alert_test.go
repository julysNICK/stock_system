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
