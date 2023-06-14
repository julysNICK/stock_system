// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type Querier interface {
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateSale(ctx context.Context, arg CreateSaleParams) (Sale, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateStockAlert(ctx context.Context, arg CreateStockAlertParams) (StockAlert, error)
	CreateStore(ctx context.Context, arg CreateStoreParams) (Store, error)
	CreateSupplier(ctx context.Context, arg CreateSupplierParams) (Supplier, error)
	DeleteSale(ctx context.Context, id int64) error
	DeleteStockAlert(ctx context.Context, id int64) error
	GetAllSuppliers(ctx context.Context, arg GetAllSuppliersParams) ([]Supplier, error)
	GetProduct(ctx context.Context, id int64) (Product, error)
	GetProductForUpdate(ctx context.Context, id int64) (Product, error)
	GetProductsByCategory(ctx context.Context, arg GetProductsByCategoryParams) ([]Product, error)
	GetProductsWithJoinWithStore(ctx context.Context, arg GetProductsWithJoinWithStoreParams) ([]GetProductsWithJoinWithStoreRow, error)
	GetProductsWithJoinWithSupplierBySupplierId(ctx context.Context, arg GetProductsWithJoinWithSupplierBySupplierIdParams) ([]GetProductsWithJoinWithSupplierBySupplierIdRow, error)
	GetSale(ctx context.Context, id int64) (Sale, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetStockAlert(ctx context.Context, id int64) (StockAlert, error)
	GetStockAlertsByProductIdAndSupplierId(ctx context.Context, arg GetStockAlertsByProductIdAndSupplierIdParams) ([]StockAlert, error)
	GetStore(ctx context.Context, id int64) (Store, error)
	GetStoreByEmail(ctx context.Context, contactEmail string) (Store, error)
	GetStoreForUpdate(ctx context.Context, id int64) (Store, error)
	GetSupplier(ctx context.Context, id int64) (Supplier, error)
	ListAllProducts(ctx context.Context, arg ListAllProductsParams) ([]Product, error)
	ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error)
	ListSales(ctx context.Context, arg ListSalesParams) ([]Sale, error)
	ListStores(ctx context.Context, arg ListStoresParams) ([]Store, error)
	SearchProducts(ctx context.Context, dollar_1 sql.NullString) ([]Product, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error)
	UpdateStockAlert(ctx context.Context, arg UpdateStockAlertParams) (StockAlert, error)
	UpdateStore(ctx context.Context, arg UpdateStoreParams) (Store, error)
	UpdateSupplier(ctx context.Context, arg UpdateSupplierParams) (Supplier, error)
}

var _ Querier = (*Queries)(nil)
