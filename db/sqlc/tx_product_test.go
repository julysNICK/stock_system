package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTxProduct(t *testing.T){
		store := NewStoreDB(testDb)

		storeRandom := CreateRandomStore(t)

	 arg, err := store.ProductTx(context.Background(), ProductTxParams{
		Name: "test product",
		Description: "test description",
		Price: 100,
		StoreID: storeRandom.ID,
		Quantity: 10,
	 })

	 require.NoError(t, err)

	 require.NotEmpty(t, arg.Product)
}

func TestTxProductErrorCreate(t *testing.T){
		store := NewStoreDB(testDb)

		storeRandom := CreateRandomStore(t)

	 arg, err := store.ProductTx(context.Background(), ProductTxParams{
		Name: "test product",
		Description: "test description",
		Price: 100,
		StoreID: storeRandom.ID,
		Quantity: 10,
	 })

	 require.NoError(t, err)

	 require.NotEmpty(t, arg.Product)
}