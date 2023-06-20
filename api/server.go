package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_system/db/sqlc"
	"github.com/julysNICK/stock_system/token"
	"github.com/julysNICK/stock_system/utils"
)

type Server struct {
	config utils.Config
	store  db.StoreDB
	router *gin.Engine
	token  token.Maker
}

func NewServer(config utils.Config, store db.StoreDB) (*Server, error) {
	tokenMaker, err := token.NewJwtMaker(config.TOKEN_SYMMETRIC_KEY)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config: config,
		token:  tokenMaker,
		store:  store}

	server.setupRouter()

	return server, nil

}

func (server *Server) setupRouter() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	router.POST("/stores/login", server.LoginStore)
	router.GET("/stores", server.ListStores)
	router.POST("/stores", server.CreateStore)
	router.GET("/products", server.ListProducts)
	router.GET("/chat/:room", server.HandlerMessage)
	
	router.GET("/products/:product_id", server.GetProduct)
	router.GET("/products/supplier/:supplier_id", server.GetProductsBySupplierId)
	router.POST("/products", server.CreateProduct)
	router.PATCH("/products/:product_id", server.UpdateProduct)
	router.GET("/products/category", server.GetProductsByCategory)
	router.GET("/products/search", server.GetProductsBySearch)


	router.PATCH("/stores/:store_id", server.UpdateStore)
	router.GET("/stores/:store_id", server.GetStore)

	router.POST("/suppliers", server.CreateSupplier)
	router.GET("/suppliers/:supplier_id", server.GetSupplier)
	router.GET("/suppliers", server.ListSuppliers)
	router.PATCH("/suppliers/:supplier_id", server.UpdateSupplier)

	router.POST("/sales", server.CreateSale)
	router.GET("/sales/:sale_id", server.GetSale)
	router.GET("/sales", server.ListSales)
	router.DELETE("/sales/:sale_id", server.DeleteSale)

	router.POST("/stock_alerts", server.CreateStockAlert)
	router.GET("/stock_alerts/:stock_alert_id", server.GetStockAlert)
	router.PATCH("/stock_alerts/:stock_alert_id", server.UpdateStockAlert)
	router.DELETE("/stock_alerts/:stock_alert_id", server.DeleteStockAlert)
	// authRoutes := router.Group("/").Use(authMiddleware(server.token))

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
