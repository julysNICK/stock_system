package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/julysNICK/stock_system/db/mock"
	db "github.com/julysNICK/stock_system/db/sqlc"
	"github.com/julysNICK/stock_system/utils"
	"github.com/stretchr/testify/require"
)

func randomStore(t *testing.T) (store db.Store) {

	return db.Store{
		ID:             int64(utils.RandomInt(1, 100)),
		Name:           utils.RandomName(),
		Address:        utils.RandomAddress(),
		ContactEmail:   utils.RandomEmail(),
		ContactPhone:   utils.RandomPhone(),
		HashedPassword: utils.RandomPassword(),
	}

}

// func TestCreateStoreAPI(t *testing.T) {
// 	store := randomStore(t)

// 	testCase := []struct {
// 		name         string
// 		body         io.Reader
// 		buildStubs   func(store *mockdb.MockStoreDB)
// 		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: mustNewJSONBody(t, CreateStoreRequest{
// 				Name:           store.Name,
// 				Address:        store.Address,
// 				ContactEmail:   store.ContactEmail,
// 				ContactPhone:   store.ContactPhone,
// 				HashedPassword: store.HashedPassword,
// 			}),
// 			buildStubs: func(store *mockdb.MockStoreDB) {
// 				store.EXPECT().CreateStore(gomock.Any(), gomock.Any()).Times(1).Return(store, nil)
// 			},
// 			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchAccount(t, recorder.Body, store)
// 			},
// 		},
// 	}

// 	for _, tc := range testCase {
// 		t.Run(
// 			(tc.name), func(t *testing.T) {
// 				store := mockdb.NewMockStoreDB(gomock.NewController(t))

// 				server := NewServer(store)

// 				tc.buildStubs(store)

// 				recorder := httptest.NewRecorder()

// 				url := "/stores"

// 				request, err := http.NewRequest(http.MethodPost, url, tc.body)

// 				require.NoError(t, err)
// 				server.router.ServeHTTP(recorder, request)

// 				tc.checkResponse(t, recorder)

// 			},
// 		)
// 	}
// }

func mustNewJSONBody(t *testing.T, data interface{}) *bytes.Buffer {
	body := &bytes.Buffer{}

	err := json.NewEncoder(body).Encode(data)
	require.NoError(t, err)

	return body

}

func requireBodyMatchStore(t *testing.T, body *bytes.Buffer, store db.Store) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotStore db.Store
	err = json.Unmarshal(data, &gotStore)
	require.NoError(t, err)
	require.Equal(t, store, gotStore)
}

func TestGetStore(t *testing.T) {
	storeRandom := randomStore(t)

	testCase := []struct {
		name          string
		AccountId     int64
		buildStubs    func(store *mockdb.MockStoreDB)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			AccountId: storeRandom.ID,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().GetStore(gomock.Any(), gomock.Eq(storeRandom.ID)).Times(1).
					Return(storeRandom, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchStore(t, recorder.Body, storeRandom)
			},
		},
		{
			name:      "NOT FOUND",
			AccountId: storeRandom.ID,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().GetStore(gomock.Any(), gomock.Eq(storeRandom.ID)).Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:      "INTERNAL ERROR",
			AccountId: storeRandom.ID,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().GetStore(gomock.Any(), gomock.Eq(storeRandom.ID)).Times(1).
					Return(db.Store{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
		{
			name:      "PARAMS ERROR",
			AccountId: 0,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().GetStore(gomock.Any(), gomock.Any()).Times(0)
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

				url := fmt.Sprintf("/stores/%d", tc.AccountId)

				request, err := http.NewRequest(http.MethodGet, url, nil)
				require.NoError(t, err)

				server.router.ServeHTTP(recorder, request)
				tc.checkResponse(t, recorder)
			},
		)

	}

}
