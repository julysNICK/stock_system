package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/julysNICK/stock_system/db/mock"
	db "github.com/julysNICK/stock_system/db/sqlc"
	"github.com/julysNICK/stock_system/utils"
	"github.com/stretchr/testify/require"
)

func randomSale(t *testing.T, product db.Product) (store db.Sale) {

	return db.Sale{
		ID:           int64(utils.RandomInt(1, 100)),
		ProductID:    product.ID,
		SaleDate:     time.Now(),
		QuantitySold: int32(utils.RandomInt(1, 100)),
	}

}

func randomProduct(t *testing.T) (store db.Product) {

	return db.Product{
		ID:          int64(utils.RandomInt(1, 100)),
		Name:        utils.RandomString(10),
		Quantity:    int32(utils.RandomInt(1, 100)),
		Description: utils.RandomString(10),
		Price:       "10",
		StoreID:     1,
	}

}

func requireBodyMatchSale(t *testing.T, body *bytes.Buffer, sale db.Sale) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotSale db.Sale
	err = json.Unmarshal(data, &gotSale)
	require.NoError(t, err)
	require.Equal(t, sale, gotSale)
}

func requireBodyMatchSaleTX(t *testing.T, body *bytes.Buffer, sale db.SaleTxResult) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotSale db.SaleTxResult
	err = json.Unmarshal(data, &gotSale)
	require.NoError(t, err)
	require.Equal(t, sale, gotSale)
}

func TestDeleteProduct(t *testing.T) {

	productRandom := randomProduct(t)
	salesRandom := randomSale(t, productRandom)

	testCase := []struct {
		name   string
		SaleId int64

		buildStubs    func(store *mockdb.MockStoreDB)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			SaleId: salesRandom.ID,

			buildStubs: func(store *mockdb.MockStoreDB) {

				store.EXPECT().DeleteSale(gomock.Any(), gomock.Eq(salesRandom.ID)).Times(1).
					Return(nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {

				require.Equal(t, http.StatusOK, recorder.Code)
				// requireBodyMatchSale(t, recorder.Body, updateStore)
			},
		},
		{
			name:   "NOT FOUND",
			SaleId: salesRandom.ID,

			buildStubs: func(store *mockdb.MockStoreDB) {

				store.EXPECT().DeleteSale(gomock.Any(), gomock.Eq(salesRandom.ID)).Times(1).
					Return(sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:   "INTERNAL ERROR",
			SaleId: salesRandom.ID,

			buildStubs: func(store *mockdb.MockStoreDB) {

				store.EXPECT().DeleteSale(gomock.Any(), gomock.Eq(salesRandom.ID)).Times(1).
					Return(sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
		{
			name:   "PARAMS ERROR URI",
			SaleId: 0,

			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().DeleteSale(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},

		{
			name:   "PARAMS ERROR Uri",
			SaleId: 0,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().DeleteSale(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(
			tc.name,
			func(t *testing.T) {
				ctrl := gomock.NewController(t)

				defer ctrl.Finish()

				store := mockdb.NewMockStoreDB(ctrl)
				tc.buildStubs(store)

				server := NewServer(store)

				recorder := httptest.NewRecorder()

				url := fmt.Sprintf("/sales/%d", tc.SaleId)

				request, err := http.NewRequest(http.MethodDelete, url, nil)
				require.NoError(t, err)

				server.router.ServeHTTP(recorder, request)
				tc.checkResponse(t, recorder)
			},
		)

	}

}

func TestGetSales(t *testing.T) {
	productRandom := randomProduct(t)
	saleRandom := randomSale(t, productRandom)

	testCase := []struct {
		name          string
		SalesID       int64
		buildStubs    func(store *mockdb.MockStoreDB)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			SalesID: saleRandom.ID,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().GetSale(gomock.Any(), gomock.Eq(saleRandom.ID)).Times(1).
					Return(saleRandom, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				// requireBodyMatchSale(t, recorder.Body, saleRandom)
			},
		},
		{
			name:    "NOT FOUND",
			SalesID: saleRandom.ID,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().GetSale(gomock.Any(), gomock.Eq(saleRandom.ID)).Times(1).
					Return(db.Sale{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:    "INTERNAL ERROR",
			SalesID: saleRandom.ID,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().GetSale(gomock.Any(), gomock.Eq(saleRandom.ID)).Times(1).
					Return(db.Sale{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
		{
			name:    "PARAMS ERROR",
			SalesID: 0,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().GetSale(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(
			tc.name,
			func(t *testing.T) {
				ctrl := gomock.NewController(t)

				defer ctrl.Finish()

				store := mockdb.NewMockStoreDB(ctrl)
				tc.buildStubs(store)

				server := NewServer(store)

				recorder := httptest.NewRecorder()

				url := fmt.Sprintf("/sales/%d", tc.SalesID)

				request, err := http.NewRequest(http.MethodGet, url, nil)
				require.NoError(t, err)

				server.router.ServeHTTP(recorder, request)
				tc.checkResponse(t, recorder)
			},
		)

	}

}

func TestCreateSale(t *testing.T) {
	productRandom := randomProduct(t)
	saleRandom := randomSale(t, productRandom)
	// saleTxResult := db.SaleTxResult{
	// 	Sale: saleRandom,
	// 	Product: db.Product{
	// 		ID:          saleRandom.ProductID,
	// 		Name:        productRandom.Name,
	// 		Description: productRandom.Description,
	// 		Price:       productRandom.Price,
	// 		Quantity:    productRandom.Quantity,
	// 		StoreID:     productRandom.StoreID,
	// 	},
	// }

	testCase := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStoreDB)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"product_id":    saleRandom.ProductID,
				"quantity":      saleRandom.QuantitySold,
				"quantity_sold": saleRandom.QuantitySold,
			},
			buildStubs: func(store *mockdb.MockStoreDB) {

				arg := db.SaleTxParams{
					ProductID:    saleRandom.ProductID,
					QuantitySold: saleRandom.QuantitySold,
				}

				log.Println(arg)

				store.EXPECT().SaleTx(gomock.Any(), gomock.Eq(arg)).Times(1)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {

				require.Equal(t, http.StatusOK, recorder.Code)
				// requireBodyMatchSaleTX(t, recorder.Body, saleTxResult)
			},
		},

		// {
		// 	name: "INTERNAL ERROR",
		// 	body: gin.H{
		// 		"product_id":    saleRandom.ProductID,
		// 		"quantity":      saleRandom.QuantitySold,
		// 		"quantity_sold": saleRandom.QuantitySold,
		// 	},
		// 	buildStubs: func(store *mockdb.MockStoreDB) {
		// 		arg := db.CreateSaleParams{
		// 			ProductID:    saleRandom.ProductID,
		// 			QuantitySold: saleRandom.QuantitySold,
		// 			SaleDate:     time.Now(),
		// 		}
		// 		store.EXPECT().CreateSale(gomock.Any(), gomock.Eq(arg)).Times(1).
		// 			Return(db.Store{}, sql.ErrConnDone)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusInternalServerError, recorder.Code)

		// 	},
		// },
		// {
		// 	name: "PARAMS ERROR",
		// 	body: gin.H{
		// 		"product_id":    saleRandom.ProductID,
		// 		"quantity":      saleRandom.QuantitySold,
		// 		"quantity_sold": saleRandom.QuantitySold,
		// 	},
		// 	buildStubs: func(store *mockdb.MockStoreDB) {
		// 		store.EXPECT().CreateSale(gomock.Any(), gomock.Any()).Times(0)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusBadRequest, recorder.Code)

		// 	},
		// },
	}

	for i := range testCase {
		tc := testCase[i]

		t.Run(
			tc.name,
			func(t *testing.T) {
				ctrl := gomock.NewController(t)

				defer ctrl.Finish()

				store := mockdb.NewMockStoreDB(ctrl)
				tc.buildStubs(store)

				server := NewServer(store)

				recorder := httptest.NewRecorder()

				data, err := json.Marshal(tc.body)
				require.NoError(t, err)

				url := "/sales"

				request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
				require.NoError(t, err)

				server.router.ServeHTTP(recorder, request)
				tc.checkResponse(t, recorder)
			},
		)

	}

}
