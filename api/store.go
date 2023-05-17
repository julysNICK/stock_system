package api

import (
	"log"
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

type getSalesRequest struct {
	StoreID int64 `uri:"store_id" binding:"required,min=1"`
}

func (server *Server) GetStore(ctx *gin.Context) {
	var req getSalesRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		log.Println("error")
		validatorErrorParserInParams(ctx, err)
		return
	}

	sale, err := server.store.GetStore(ctx, req.StoreID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sale)

}
