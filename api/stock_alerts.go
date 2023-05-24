package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_system/db/sqlc"
)

type CreateStockAlertRequest struct {
	ProductID     int64 `json:"product_id" binding:"required,min=1"`
	SupplierID    int64 `json:"supplier_id" binding:"required,min=1"`
	AlertQuantity int32 `json:"alert_quantity" binding:"required,min=1"`
}

func (server *Server) CreateStockAlert(ctx *gin.Context) {
	var req CreateStockAlertRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.StockAlertTxParams{
		ProductID:     req.ProductID,
		SupplierID:    req.SupplierID,
		AlertQuantity: req.AlertQuantity,
	}

	stockAlert, err := server.store.StockAlertTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, stockAlert)
}

type getStockAlertRequest struct {
	StockAlertID int64 `uri:"stock_alert_id" binding:"required,min=1"`
}

func (server *Server) GetStockAlert(ctx *gin.Context) {
	var req getStockAlertRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	stockAlert, err := server.store.GetStockAlert(ctx, req.StockAlertID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, stockAlert)
}

type DeleteStockAlertRequest struct {
	StockAlertID int64 `uri:"stock_alert_id" binding:"required,min=1"`
}

func (server *Server) DeleteStockAlert(ctx *gin.Context) {
	var req DeleteStockAlertRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	err := server.store.DeleteStockAlert(ctx, req.StockAlertID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Stock alert with id %d deleted", req.StockAlertID))
}

type UpdateStockAlertRequest struct {
	AlertQuantity int64 `json:"alert_quantity,omitempty" binding:"min=1"`
}
type UpdateStockAlertRequestUri struct {
	StockAlertID int64 `uri:"stock_alert_id" binding:"required,min=1"`
}

func (server *Server) UpdateStockAlert(ctx *gin.Context) {
	var req UpdateStockAlertRequest
	var reqUri UpdateStockAlertRequestUri

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.UpdateStockAlertParams{
		ID: reqUri.StockAlertID,
		AlertQuantity: sql.NullInt32{
			Int32: int32(req.AlertQuantity),
			Valid: req.AlertQuantity != 0,
		},
	}

	stockAlert, err := server.store.UpdateStockAlert(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, stockAlert)
}
