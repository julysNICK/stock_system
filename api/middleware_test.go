package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/julysNICK/stock_system/token"
	"github.com/stretchr/testify/require"
)

func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	idStore int64,
	duration time.Duration,
) {
	token, payload, err := tokenMaker.CreateToken(idStore, username, duration)

	require.NoError(t, err)

	require.NotEmpty(t, payload)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)

	request.Header.Set("authorization", authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name      string
		setupAuth func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		checkAuth func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", 1, time.Minute)
			},
			checkAuth: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recoder.Code)

			},
		},

		{
			name: "NO AUTHORIZATION HEADER",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {

			},
			checkAuth: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recoder.Code)

			},
		},
		{
			name: "NO UNSUPPORTED AUTHORIZATION HEADER",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, "unsuported", "username", 1, time.Minute)
			},
			checkAuth: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recoder.Code)

			},
		},
		{
			name: "EXPIRED TOKEN",
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeToken, "username", 1, -time.Minute)
			},
			checkAuth: func(t *testing.T, recoder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recoder.Code)

			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(
			tc.name,
			func(t *testing.T) {
				server := NewTestServer(t, nil)

				authPath := "/auth"

				server.router.GET(
					authPath,
					authMiddleware(server.token),
					func(ctx *gin.Context) {
						ctx.JSON(http.StatusOK, gin.H{})
					},
				)

				recorder := httptest.NewRecorder()

				request, err := http.NewRequest(http.MethodGet, authPath, nil)

				require.NoError(t, err)

				tc.setupAuth(t, request, server.token)

				server.router.ServeHTTP(recorder, request)

				tc.checkAuth(t, recorder)
			},
		)
	}
}
