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
	router.GET("/stores", server.ListStores)
	router.PATCH("/stores/:store_id", server.UpdateStore)

	router.POST("/suppliers", server.CreateSupplier)
	router.GET("/suppliers/:supplier_id", server.GetSupplier)
	router.PATCH("/suppliers/:supplier_id", server.UpdateSupplier)

	router.GET("/products/:product_id", server.GetProduct)
	router.GET("/products", server.ListProducts)
	router.POST("/products", server.CreateProduct)
	router.PATCH("/products/:product_id", server.UpdateProduct)

	router.POST("/sales", server.CreateSale)
	router.GET("/sales/:sale_id", server.GetSale)
	router.GET("/sales", server.ListSales)
	router.DELETE("/sales/:sale_id", server.DeleteSale)

	router.POST("/stock_alerts", server.CreateStockAlert)
	router.GET("/stock_alerts/:stock_alert_id", server.GetStockAlert)
	router.PATCH("/stock_alerts/:stock_alert_id", server.UpdateStockAlert)
	router.DELETE("/stock_alerts/:stock_alert_id", server.DeleteStockAlert)

	server.router = router

	return server

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
