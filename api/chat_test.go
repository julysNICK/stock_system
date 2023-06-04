package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/julysNICK/stock_system/db/mock"
	"github.com/stretchr/testify/require"
)

func TestHandlerMessage(t *testing.T){

		ctrl := gomock.NewController(t)

		defer ctrl.Finish()

		store := mockdb.NewMockStoreDB(ctrl)

		server := NewTestServer(t, store)


		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/chat/%s", "1")

		request, err := http.NewRequest("GET", url, nil)

		require.NoError(t, err)

		server.router.ServeHTTP(recorder, request)

		require.Equal(t, http.StatusOK, recorder.Code)
	
}