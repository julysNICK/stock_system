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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	mockdb "github.com/julysNICK/stock_system/db/mock"
	db "github.com/julysNICK/stock_system/db/sqlc"
	"github.com/julysNICK/stock_system/token"
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

func requireBodyMatchStoreList(t *testing.T, body *bytes.Buffer, stores []db.Store) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotStores []db.Store
	err = json.Unmarshal(data, &gotStores)
	require.NoError(t, err)
	require.Equal(t, stores, gotStores)
}

func TestGetStore(t *testing.T) {
	storeRandom := randomStore(t)

	testCase := []struct {
		name          string
		AccountId     int64
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStoreDB)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			AccountId: storeRandom.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", time.Minute)
			},

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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", time.Minute)
			},

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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", time.Minute)
			},

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
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", time.Minute)
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

			server := NewTestServer(t,store)
				recorder := httptest.NewRecorder()

				url := fmt.Sprintf("/stores/%d", tc.AccountId)

				request, err := http.NewRequest(http.MethodGet, url, nil)
				require.NoError(t, err)
				tc.setupAuth(t, request, server.token)
				server.router.ServeHTTP(recorder, request)
				tc.checkResponse(t, recorder)
			},
		)

	}

}

func TestUpdateStore(t *testing.T) {
	storeRandom := randomStore(t)
	updateStore := db.Store{
		ID:             storeRandom.ID,
		Name:           "test",
		Address:        "postpost",
		ContactEmail:   "test@test.com",
		ContactPhone:   "99 991309493",
		HashedPassword: "2222222222222",
	}
	testCase := []struct {
		name          string
		AccountId     int64
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStoreDB)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			AccountId: storeRandom.ID,
			body: gin.H{
				"name":            updateStore.Name,
				"address":         updateStore.Address,
				"contact_email":   updateStore.ContactEmail,
				"contact_phone":   updateStore.ContactPhone,
				"hashed_password": updateStore.HashedPassword,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStoreDB) {

				arg := db.UpdateStoreParams{
					ID: storeRandom.ID,
					Name: sql.NullString{
						String: updateStore.Name,
						Valid:  true,
					},

					Address: sql.NullString{
						String: updateStore.Address,
						Valid:  true,
					},
					ContactEmail: sql.NullString{
						String: updateStore.ContactEmail,
						Valid:  true,
					},

					ContactPhone: sql.NullString{
						String: updateStore.ContactPhone,
						Valid:  true,
					},

					HashedPassword: sql.NullString{
						String: updateStore.HashedPassword,
						Valid:  true,
					},
				}

				store.EXPECT().UpdateStore(gomock.Any(), gomock.Eq(arg)).Times(1).
					Return(updateStore, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {

				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchStore(t, recorder.Body, updateStore)
			},
		},
		{
			name:      "NOT FOUND",
			AccountId: storeRandom.ID,
			body: gin.H{
				"name":            updateStore.Name,
				"address":         updateStore.Address,
				"contact_email":   updateStore.ContactEmail,
				"contact_phone":   updateStore.ContactPhone,
				"hashed_password": updateStore.HashedPassword,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStoreDB) {
				arg := db.UpdateStoreParams{
					ID: storeRandom.ID,
					Name: sql.NullString{
						String: updateStore.Name,
						Valid:  true,
					},

					Address: sql.NullString{
						String: updateStore.Address,
						Valid:  true,
					},
					ContactEmail: sql.NullString{
						String: updateStore.ContactEmail,
						Valid:  true,
					},

					ContactPhone: sql.NullString{
						String: updateStore.ContactPhone,
						Valid:  true,
					},

					HashedPassword: sql.NullString{
						String: updateStore.HashedPassword,
						Valid:  true,
					},
				}

				store.EXPECT().UpdateStore(gomock.Any(), gomock.Eq(arg)).Times(1).
					Return(db.Store{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

			},
		},
		{
			name:      "INTERNAL ERROR",
			AccountId: storeRandom.ID,
			body: gin.H{
				"name":            updateStore.Name,
				"address":         updateStore.Address,
				"contact_email":   updateStore.ContactEmail,
				"contact_phone":   updateStore.ContactPhone,
				"hashed_password": updateStore.HashedPassword,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStoreDB) {
				arg := db.UpdateStoreParams{
					ID: storeRandom.ID,
					Name: sql.NullString{
						String: updateStore.Name,
						Valid:  true,
					},

					Address: sql.NullString{
						String: updateStore.Address,
						Valid:  true,
					},
					ContactEmail: sql.NullString{
						String: updateStore.ContactEmail,
						Valid:  true,
					},

					ContactPhone: sql.NullString{
						String: updateStore.ContactPhone,
						Valid:  true,
					},

					HashedPassword: sql.NullString{
						String: updateStore.HashedPassword,
						Valid:  true,
					},
				}

				store.EXPECT().UpdateStore(gomock.Any(), gomock.Eq(arg)).Times(1).
					Return(db.Store{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)

			},
		},
		{
			name:      "PARAMS ERROR URI",
			AccountId: 0,
			body: gin.H{
				"name":            updateStore.Name,
				"address":         updateStore.Address,
				"contact_email":   updateStore.ContactEmail,
				"contact_phone":   updateStore.ContactPhone,
				"hashed_password": updateStore.HashedPassword,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)

			},
		},

		{
			name:      "PARAMS ERROR BODY",
			AccountId: storeRandom.ID,
			body: gin.H{

				"address":         updateStore.Address,
				"contact_email":   updateStore.ContactEmail,
				"contact_phone":   updateStore.ContactPhone,
				"hashed_password": updateStore.HashedPassword,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Times(0)
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

				server := NewTestServer(t,store)

				recorder := httptest.NewRecorder()
				data, err := json.Marshal(tc.body)
				require.NoError(t, err)
				url := fmt.Sprintf("/stores/%d", tc.AccountId)

				request, err := http.NewRequest(http.MethodPatch, url, bytes.NewReader(data))
				require.NoError(t, err)
				tc.setupAuth(t, request, server.token)
				server.router.ServeHTTP(recorder, request)
				tc.checkResponse(t, recorder)
			},
		)

	}

}

func TestCreateStore(t *testing.T) {
	storeRandom := randomStore(t)

	testCase := []struct {
		name          string
		body          gin.H
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStoreDB)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":            storeRandom.Name,
				"address":         storeRandom.Address,
				"contact_email":   storeRandom.ContactEmail,
				"contact_phone":   storeRandom.ContactPhone,
				"hashed_password": storeRandom.HashedPassword,
			},
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", time.Minute)
			},
			buildStubs: func(store *mockdb.MockStoreDB) {

				// arg := db.CreateStoreParams{
				// 	Name:           storeRandom.Name,
				// 	Address:        storeRandom.Address,
				// 	ContactEmail:   storeRandom.ContactEmail,
				// 	ContactPhone:   storeRandom.ContactPhone,
				// 	HashedPassword: storeRandom.HashedPassword,
				// }

				store.EXPECT().CreateStore(gomock.Any(), gomock.Any()).Times(1).
					Return(storeRandom, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchStore(t, recorder.Body, storeRandom)
			},
		},

		// {
		// 	name: "INTERNAL ERROR",
		// 	body: gin.H{
		// 		"name":            storeRandom.Name,
		// 		"address":         storeRandom.Address,
		// 		"contact_email":   storeRandom.ContactEmail,
		// 		"contact_phone":   storeRandom.ContactPhone,
		// 		"hashed_password": storeRandom.HashedPassword,
		// 	},
		// 	buildStubs: func(store *mockdb.MockStoreDB) {
		// 		arg := db.CreateStoreParams{
		// 			Name:           storeRandom.Name,
		// 			Address:        storeRandom.Address,
		// 			ContactEmail:   storeRandom.ContactEmail,
		// 			ContactPhone:   storeRandom.ContactPhone,
		// 			HashedPassword: storeRandom.HashedPassword,
		// 		}
		// 		store.EXPECT().CreateStore(gomock.Any(), gomock.Eq(arg)).Times(1).
		// 			Return(db.Store{}, sql.ErrConnDone)
		// 	},
		// 	checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
		// 		require.Equal(t, http.StatusInternalServerError, recorder.Code)

		// 	},
		// },
		// {
		// 	name: "PARAMS ERROR",
		// 	body: gin.H{

		// 		"address":         storeRandom.Address,
		// 		"contact_email":   storeRandom.ContactEmail,
		// 		"contact_phone":   storeRandom.ContactPhone,
		// 		"hashed_password": storeRandom.HashedPassword,
		// 	},
		// 	buildStubs: func(store *mockdb.MockStoreDB) {
		// 		store.EXPECT().GetStore(gomock.Any(), gomock.Any()).Times(0)
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

		server := NewTestServer(t,store)

				recorder := httptest.NewRecorder()

				data, err := json.Marshal(tc.body)
				require.NoError(t, err)

				url := "/stores"

				request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
				require.NoError(t, err)
				tc.setupAuth(t, request, server.token)
				server.router.ServeHTTP(recorder, request)
				tc.checkResponse(t, recorder)
			},
		)

	}

}

func TestListStore(t *testing.T) {
	storeRandom := randomStore(t)

	listsStores := []db.Store{storeRandom, storeRandom, storeRandom}

	testCase := []struct {
		name          string
		limit         int32
		offset        int64
		buildStubs    func(store *mockdb.MockStoreDB)
		checkResponse func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			limit:  2,
			offset: 1,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().ListStores(gomock.Any(), gomock.Any()).Times(1).
					Return(listsStores, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchStoreList(t, recorder.Body, listsStores)
			},
		},

		{
			name:   "INTERNAL ERROR",
			limit:  2,
			offset: 1,
			buildStubs: func(store *mockdb.MockStoreDB) {

				store.EXPECT().ListStores(gomock.Any(), gomock.Any()).Times(1).
					Return(nil, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "PARAMS ERROR LIMIT",
			limit:  -1,
			offset: 0,
			buildStubs: func(store *mockdb.MockStoreDB) {
				store.EXPECT().ListStores(gomock.Any(), gomock.Any()).Times(0)
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

				server := NewTestServer(t,store)

				recorder := httptest.NewRecorder()

				url := fmt.Sprintf("/stores?limit=%d&offset=%d", tc.limit, tc.offset)

				request, err := http.NewRequest(http.MethodGet, url, nil)
				require.NoError(t, err)

				server.router.ServeHTTP(recorder, request)
				tc.checkResponse(t, recorder)
			},
		)

	}

}
