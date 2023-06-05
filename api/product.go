package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/julysNICK/stock_system/db/sqlc"
)

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	StoreID     int64  `json:"store_id" binding:"required"`
	Quantity    int32  `json:"quantity" binding:"required"`
	SupplierID  int64  `json:"supplier_id" binding:"required"`
	Category    string `json:"category" binding:"required"`
	ImageUrl    string `json:"image_url" binding:"required"`
}

type CreateProductResponse struct {
	Product db.Product `json:"product" binding:"required"`
	Store   db.Store   `json:"store" binding:"required"`
}

func (server *Server) CreateProduct(ctx *gin.Context) {
	var req CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("err", err)

		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.ProductTxParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		StoreID:     req.StoreID,
		Quantity:    req.Quantity,
		SupplierID:  req.SupplierID,
		Category:    req.Category,
		ImageUrl:    req.ImageUrl,
	}

	product, err := server.store.ProductTx(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, product)
}

type getProductRequest struct {
	ProductID int64 `uri:"product_id" binding:"required,min=1"`
}

func (server *Server) GetProduct(ctx *gin.Context) {
	var req getProductRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	product, err := server.store.GetProduct(ctx, req.ProductID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	getStore, err := server.store.GetStore(ctx, product.StoreID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, CreateProductResponse{
		Product: product,
		Store:   getStore,
	})
}

type listProductsRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	Limit  int32 `form:"limit" binding:"required,min=5,max=10"`
}

func (server *Server) ListProducts(ctx *gin.Context) {
	var req listProductsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.GetProductsWithJoinWithStoreParams{
		Limit:  req.Limit,
		Offset: (req.PageID - 1) * req.Limit,
	}

	products, err := server.store.GetProductsWithJoinWithStore(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, products)
}

type updateProductRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Price       int64  `json:"price,omitempty"`

	Quantity int32 `json:"quantity,omitempty"`
}

type updateProductResponseUri struct {
	ProductID int64 `uri:"product_id" binding:"required,min=1"`
}

type UpdateProductResponse struct {
	Product db.Product `json:"product" binding:"required"`
	Store   db.Store   `json:"store" binding:"required"`
}

func (server *Server) UpdateProduct(ctx *gin.Context) {
	var req updateProductRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	var reqUri updateProductResponseUri

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		validatorErrorParserInParams(ctx, err)
		return
	}

	arg := db.UpdateProductParams{
		ID: reqUri.ProductID,
		Name: sql.NullString{
			String: req.Name,
			Valid:  req.Name != "",
		},

		Description: sql.NullString{
			String: req.Description,
			Valid:  req.Description != "",
		},

		Price: sql.NullString{
			String: fmt.Sprintf("%d", req.Price),
			Valid:  req.Price != 0,
		},

		Quantity: sql.NullInt32{
			Int32: req.Quantity,
			Valid: req.Quantity != 0,
		},
	}

	product, err := server.store.UpdateProduct(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	getStore, err := server.store.GetStore(ctx, product.StoreID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, UpdateProductResponse{
		Product: product,
		Store:   getStore,
	})
}
