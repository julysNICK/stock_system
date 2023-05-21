package api

import (
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_system/db/sqlc"
	"github.com/julysNICK/stock_system/utils"
)

func NewTestServer(t *testing.T, store db.StoreDB) *Server {

	config := utils.Config{
		TOKEN_SYMMETRIC_KEY: utils.RandomString(36),
		ACCESS_TOKEN_DURATION: time.Minute,
	}

	server, err := NewServer(config,store)

	if err != nil {
		t.Fatal(err)
	}

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
