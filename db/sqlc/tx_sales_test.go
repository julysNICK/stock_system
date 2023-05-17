package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTxSale(t *testing.T) {
	store := NewStoreDB(testDb)
	storeRandom := CreateRandomStore(t)
	productRandom := CreateRandomProduct(t, storeRandom)

	arg, err := store.SaleTx(context.Background(), SaleTxParams{
		ProductID:    productRandom.ID,
		SaleDate:     productRandom.CreatedAt,
		QuantitySold: 10,
	})

	require.NoError(t, err)

	require.NotEmpty(t, arg.Sale)
}
