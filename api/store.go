package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_system/db/sqlc"
)

type CreateStoreRequest struct {
	Name           string `json:"name" binding:"required, min=1"`
	Address        string `json:"address" binding:"required, min=8"`
	ContactEmail   string `json:"contact_email" binding:"required, email"`
	ContactPhone   string `json:"contact_phone" binding:"required "`
	HashedPassword string `json:"hashed_password" binding:"required, min=8"`
}

func (server *Server) CreateStore(ctx *gin.Context) {
	var req CreateStoreRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, store)
}
