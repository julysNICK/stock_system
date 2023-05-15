package db

import (
	"context"
	"testing"
	"time"

	"github.com/julysNICK/stock_system/utils"
	"github.com/stretchr/testify/require"
)


func createRandomStore(t *testing.T) Store {
	hashedPassword := utils.RandomPassword()
	require.NotEmpty(t, hashedPassword)

	arg := CreateStoreParams{
		Name: 				 utils.RandomName(),
		Address: 			 utils.RandomAddress(),
		ContactEmail: 		 utils.RandomEmail(),
		ContactPhone: 		 utils.RandomPhone(),
		HashedPassword: 	 hashedPassword,
	}

	store, err := testQueries.CreateStore(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, store)

	require.Equal(t, arg.Name, store.Name)
	require.Equal(t, arg.Address, store.Address)
	require.Equal(t, arg.ContactEmail, store.ContactEmail)
	require.Equal(t, arg.ContactPhone, store.ContactPhone)
	require.Equal(t, arg.HashedPassword, store.HashedPassword)

	require.NotZero(t, store.ID)
	require.NotZero(t, store.CreatedAt)

	return store
}


func TestCreateStore(t *testing.T) {
	createRandomStore(t)
}

func TestGetStore(t *testing.T){
	store1 := createRandomStore(t)

	store2, err := testQueries.GetStore(context.Background(), store1.ID)

	require.NoError(t, err)

	require.NotEmpty(t, store2)

	require.Equal(t, store1.ID, store2.ID)
	require.Equal(t, store1.Name, store2.Name)
	require.Equal(t, store1.Address, store2.Address)
	require.Equal(t, store1.ContactEmail, store2.ContactEmail)
	require.Equal(t, store1.ContactPhone, store2.ContactPhone)
	require.Equal(t, store1.HashedPassword, store2.HashedPassword)
	require.WithinDuration(t, store1.CreatedAt, store2.CreatedAt, time.Second)

}

func TestGetStoreForUpdate(t *testing.T){

	store1 := createRandomStore(t)

	store2, err := testQueries.GetStoreForUpdate(context.Background(), store1.ID)

	require.NoError(t, err)

	require.NotEmpty(t, store2)

  require.Equal(t, store1.ID, store2.ID)
	require.Equal(t, store1.Name, store2.Name)
	require.Equal(t, store1.Address, store2.Address)
	require.Equal(t, store1.ContactEmail, store2.ContactEmail)
	require.Equal(t, store1.ContactPhone, store2.ContactPhone)
	require.Equal(t, store1.HashedPassword, store2.HashedPassword)
	require.WithinDuration(t, store1.CreatedAt, store2.CreatedAt, time.Second)

}

func TestListStore(t *testing.T){
	


	arg := ListStoresParams{
		Limit:  10,
		Offset: 0,
	}
	stores, err := testQueries.ListStores(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, stores)

	require.Len(t, stores, 10)
	
	require.Equal(t, int32(1), stores[0].ID)

}