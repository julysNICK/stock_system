package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_system/db/sqlc"
)

type CreateSaleRequest struct {
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int32 `json:"quantity" binding:"required"`
	// SaleDate  time.Time `json:"sale_date" binding:"required"`
	QuantitySold int32 `json:"quantity_sold" binding:"required"`
}

func (server *Server) CreateSale(ctx *gin.Context) {
	var req CreateSaleRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.SaleTxParams{
		ProductID:    req.ProductID,
		QuantitySold: req.QuantitySold,
		SaleDate:     time.Now(),
	}

	sale, err := server.store.SaleTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, sale)
}

type getSaleRequest struct {
	SaleID int64 `uri:"sale_id" binding:"required,min=1"`
}

func (server *Server) GetSale(ctx *gin.Context) {
	var req getSaleRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	sale, err := server.store.GetSale(ctx, req.SaleID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, sale)
}

type DeleteSaleRequest struct {
	SaleID int64 `uri:"sale_id" binding:"required,min=1"`
}

func (server *Server) DeleteSale(ctx *gin.Context) {
	var req DeleteSaleRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	err := server.store.DeleteSale(ctx, req.SaleID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Sale deleted successfully"})
}

type listSalesRequest struct {
	Limit  int32 `form:"limit" binding:"required,min=1"`
	Offset int32 `form:"offset" binding:"required,min=0"`
}

func (server *Server) ListSales(ctx *gin.Context) {
	var req listSalesRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.ListSalesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	sales, err := server.store.ListSales(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, sales)
}
