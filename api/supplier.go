package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_system/db/sqlc"
)

type CreateSupplierRequest struct {
	Name         string `json:"name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	Email        string `json:"email" binding:"required"`
	ContactPhone string `json:"contact_phone" binding:"required"`
}

func (server *Server) CreateSupplier(ctx *gin.Context) {
	var req CreateSupplierRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.CreateSupplierParams{
		Name:         req.Name,
		Address:      req.Address,
		Email:        req.Email,
		ContactPhone: req.ContactPhone,
	}

	supplier, err := server.store.CreateSupplier(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, supplier)

}

type getSupplierRequest struct {
	SupplierId int64 `uri:"supplier_id" binding:"required,min=1"`
}

func (server *Server) GetSupplier(ctx *gin.Context) {
	var req getSupplierRequest

	if err := ctx.ShouldBindUri(&req); err != nil {

		validatorErrorParserInParams(ctx, err)
		return
	}

	supplier, err := server.store.GetSupplier(ctx, req.SupplierId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, supplier)
}

type UpdateSupplierRequest struct {
	Name         string `json:"name,omitempty"`
	Address      string `json:"address,omitempty"`
	Email        string `json:"email,omitempty"`
	ContactPhone string `json:"contact_phone,omitempty"`
}

type UpdateSupplierRequestUri struct {
	SupplierId int64 `uri:"supplier_id" binding:"required,min=1"`
}

func (server *Server) UpdateSupplier(ctx *gin.Context) {
	var req UpdateSupplierRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	var reqUri UpdateSupplierRequestUri

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.UpdateSupplierParams{
		ID: reqUri.SupplierId,
		Name: sql.NullString{
			String: req.Name,
			Valid:  req.Name != "",
		},

		Address: sql.NullString{
			String: req.Address,
			Valid:  req.Address != "",
		},

		Email: sql.NullString{
			String: req.Email,
			Valid:  req.Email != "",
		},

		ContactPhone: sql.NullString{
			String: req.ContactPhone,
			Valid:  req.ContactPhone != "",
		},
	}

	supplier, err := server.store.UpdateSupplier(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, supplier)
}
