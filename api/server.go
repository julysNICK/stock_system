package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_system/db/sqlc"
	"github.com/julysNICK/stock_system/token"
	"github.com/julysNICK/stock_system/utils"
)

const (
	secretKey = "secrettestkey"
)

type Server struct {
	config utils.Config
	store  db.StoreDB
	router *gin.Engine
	token  token.Maker
}


func NewServer(config utils.Config, store db.StoreDB) (*Server, error) {
	tokenMaker, err := token.NewJwtMaker(secretKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config: config,
		token: tokenMaker,
		store: store}

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

	authRoutes := router.Group("/").Use(authMiddleware(server.token))

	authRoutes.PATCH("/stores/:store_id", server.UpdateStore)
	authRoutes.GET("/stores/:store_id", server.GetStore)

	authRoutes.POST("/suppliers", server.CreateSupplier)
	authRoutes.GET("/suppliers/:supplier_id", server.GetSupplier)
	authRoutes.PATCH("/suppliers/:supplier_id", server.UpdateSupplier)

	authRoutes.GET("/products/:product_id", server.GetProduct)
	authRoutes.GET("/products", server.ListProducts)
	authRoutes.POST("/products", server.CreateProduct)
	authRoutes.PATCH("/products/:product_id", server.UpdateProduct)

	authRoutes.POST("/sales", server.CreateSale)
	authRoutes.GET("/sales/:sale_id", server.GetSale)
	authRoutes.GET("/sales", server.ListSales)
	authRoutes.DELETE("/sales/:sale_id", server.DeleteSale)

	authRoutes.POST("/stock_alerts", server.CreateStockAlert)
	authRoutes.GET("/stock_alerts/:stock_alert_id", server.GetStockAlert)
	authRoutes.PATCH("/stock_alerts/:stock_alert_id", server.UpdateStockAlert)
	authRoutes.DELETE("/stock_alerts/:stock_alert_id", server.DeleteStockAlert)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
