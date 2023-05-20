package cron

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/julysNICK/stock_system/api"
	db "github.com/julysNICK/stock_system/db/sqlc"
)

func newTestServer(t *testing.T, store db.StoreDB) *api.Server {

	server, err := api.NewServer(store)

	if err != nil {
		t.Fatal(err)
	}

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
