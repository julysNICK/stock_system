package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/julysNICK/stock_system/db/sqlc"
	"github.com/julysNICK/stock_system/token"
	"github.com/julysNICK/stock_system/utils"
)

const (
	AuthKey = "auth"
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

	hash, err := utils.HashedPassword(req.HashedPassword)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateStoreParams{
		Name:           req.Name,
		Address:        req.Address,
		ContactEmail:   req.ContactEmail,
		ContactPhone:   req.ContactPhone,
		HashedPassword: hash,
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

type ProfileStoreResponse struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	ContactEmail string `json:"contactEmail"`
	ContactPhone string `json:"contactPhone"`

	CreatedAt time.Time `json:"createdAt"`
}

func (server *Server) ProfileStore(ctx *gin.Context) {
	authPayload := ctx.MustGet(AuthorizationPayloadKeyToken).(*token.Payload)

	store, err := server.store.GetStore(ctx, authPayload.IdStore)

	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, ProfileStoreResponse{
		ID:           store.ID,
		Name:         store.Name,
		Address:      store.Address,
		ContactEmail: store.ContactEmail,
		ContactPhone: store.ContactPhone,
		CreatedAt:    store.CreatedAt,
	})

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
	Name           string `json:"name,omitempty"`
	Address        string `json:"address,omitempty"`
	ContactEmail   string `json:"contact_email,omitempty"`
	ContactPhone   string `json:"contact_phone,omitempty"`
	HashedPassword string `json:"hashed_password,omitempty"`
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

	fmt.Println(reqUri.StoreID)

	arg := db.UpdateStoreParams{
		ID: reqUri.StoreID,
		Name: sql.NullString{
			String: req.Name,
			Valid:  req.Name != "",
		},
		Address: sql.NullString{
			String: req.Address,
			Valid:  req.Address != "",
		},

		ContactEmail: sql.NullString{
			String: req.ContactEmail,
			Valid:  req.ContactEmail != "",
		},
		ContactPhone: sql.NullString{
			String: req.ContactPhone,
			Valid:  req.ContactPhone != "",
		},
		HashedPassword: sql.NullString{
			String: req.HashedPassword,
			Valid:  req.HashedPassword != "",
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

type loginStoreRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginStoreResponse struct {
	SessionId             uuid.UUID `json:"session_id"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	Store                 db.Store  `json:"store"`
}

func (server *Server) LoginStore(ctx *gin.Context) {
	var loginReq loginStoreRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	store, err := server.store.GetStoreByEmail(ctx, loginReq.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(loginReq.Password, store.HashedPassword)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	//create variable duration with value 24 hours
	duration := 24 * time.Hour

	accessToken, accessPayload, err := server.token.CreateToken(store.ID, store.Name, time.Duration(duration))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.token.CreateToken(store.ID, store.Name, time.Duration(duration))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		IDStore:      store.ID,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginStoreResponse{
		SessionId:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		Store:                 store,
	}

	ctx.JSON(http.StatusOK, rsp)

}
