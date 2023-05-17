package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_system/db/sqlc"
)

type Server struct {
	store  *db.StoreDB
	router *gin.Engine
}

func NewServer(store *db.StoreDB) *Server {
	server := &Server{store: store}

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.POST("/stores", server.CreateStore)
	router.GET("/stores/:store_id", server.GetStore)

	server.router = router

	return server

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
