package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

		var verr validator.ValidationErrors

		erroAs := errors.As(err, &verr)

		if erroAs {
			out := make([]ErrorMsg, len(verr))
			for i, fe := range verr {
				out[i] = ErrorMsg{
					Field:   fe.Field(),
					Message: getErrorMsg(fe),
				}

			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}

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
