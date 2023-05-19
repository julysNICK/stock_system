package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_system/db/sqlc"
)

type CreateStoreRequest struct {
	Name           string `json:"name" binding:"required"`
	Address        string `json:"address" binding:"required"`
	ContactEmail   string `json:"contact_email" binding:"required"`
	ContactPhone   string `json:"contact_phone" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
}

func (server *Server) CreateStore(ctx *gin.Context) {
	var req CreateStoreRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.CreateStoreParams{
		Name:           req.Name,
		Address:        req.Address,
		ContactEmail:   req.ContactEmail,
		ContactPhone:   req.ContactPhone,
		HashedPassword: req.HashedPassword,
	}

	store, err := server.store.CreateStore(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, store)
}

type getStoreRequest struct {
	StoreID int64 `uri:"store_id" binding:"required,min=1"`
}

func (server *Server) GetStore(ctx *gin.Context) {
	var req getStoreRequest

	if err := ctx.ShouldBindUri(&req); err != nil {

		validatorErrorParserInParams(ctx, err)
		return
	}

	sale, err := server.store.GetStore(ctx, req.StoreID)

	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sale)

}

type ListStoresRequest struct {
	Limit  int32 `form:"limit" binding:"required,min=1,max=10"`
	Offset int32 `form:"offset" binding:"required,min=0"`
}

func (server *Server) ListStores(ctx *gin.Context) {
	var req ListStoresRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.ListStoresParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	stores, err := server.store.ListStores(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, stores)
}

type UpdateStoreRequest struct {
	Name           string `json:"name" binding:"required"`
	Address        string `json:"address" binding:"required"`
	ContactEmail   string `json:"contact_email" binding:"required"`
	ContactPhone   string `json:"contact_phone" binding:"required"`
	HashedPassword string `json:"hashed_password" binding:"required"`
}

type UpdateStoreRequestUri struct {
	StoreID int64 `uri:"store_id" binding:"required,min=1"`
}

func (server *Server) UpdateStore(ctx *gin.Context) {
	var req UpdateStoreRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	var reqUri UpdateStoreRequestUri

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.UpdateStoreParams{
		ID: reqUri.StoreID,
		Name: sql.NullString{
			String: req.Name,
			Valid:  true,
		},
		Address: sql.NullString{
			String: req.Address,
			Valid:  true,
		},

		ContactEmail: sql.NullString{
			String: req.ContactEmail,
			Valid:  true,
		},
		ContactPhone: sql.NullString{
			String: req.ContactPhone,
			Valid:  true,
		},
		HashedPassword: sql.NullString{
			String: req.HashedPassword,
			Valid:  true,
		},
	}

	store, err := server.store.UpdateStore(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, store)
}
