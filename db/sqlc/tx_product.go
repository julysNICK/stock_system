package db

import (
	"context"
	"database/sql"
	"fmt"
)

type ProductTxParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
	ImageUrl    string `json:"image_url"`
	Price       int64  `json:"price"`
	StoreID     int64  `json:"store_id"`
	SupplierID  int64  `json:"supplier_id"`
	Quantity    int32  `json:"quantity"`
}

type ProductTxResult struct {
	Product    Product  `json:"product"`
	StoreID    Store    `json:"store_id"`
	SupplierId Supplier `json:"supplier_id"`
}

func (store *SQLStore) ProductTx(ctx context.Context, arg ProductTxParams) (ProductTxResult, error) {
	var result ProductTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		if arg.Name == "" {
			return fmt.Errorf("name can't be empty")
		}

		if arg.Description == "" {
			return fmt.Errorf("description can't be empty")
		}

		if arg.Price <= 0 {
			return fmt.Errorf("price must be greater than 0")
		}

		if arg.StoreID <= 0 {
			return fmt.Errorf("store_id must be greater than 0")
		}

		if arg.SupplierID <= 0 {
			return fmt.Errorf("supplier_id must be greater than 0")
		}

		if arg.Quantity <= 0 {
			return fmt.Errorf("quantity must be greater than 0")
		}

		result.Product, err = q.CreateProduct(ctx, CreateProductParams{
			Name:        arg.Name,
			Description: arg.Description,
			Price:       fmt.Sprintf("%d", arg.Price),
			StoreID:     arg.StoreID,
			Quantity:    arg.Quantity,
			ImageUrl:    arg.ImageUrl,
			Category:    arg.Category,
			SupplierID:  arg.SupplierID,
		})

		if err != nil {
			return err
		}

		result.StoreID, err = q.GetStore(ctx, arg.StoreID)
		if err != nil {
			return err
		}
		result.SupplierId, err = q.GetSupplier(ctx, arg.SupplierID)

		if err != nil {
			return err
		}

		return nil

	})
	return result, err
}

type ProductBuyTxParams struct {
	ProductID int32 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
	StoreID   int32 `json:"store_id"`
}

type ProductBuyTxResult struct {
	Product Product `json:"product"`
}

func (store *SQLStore) ProductBuyTx(ctx context.Context, argTX ProductBuyTxParams) (ProductBuyTxResult, error) {
	var result ProductBuyTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		if argTX.ProductID == 0 {
			return fmt.Errorf("product_id can't be empty")
		}

		if argTX.Quantity == 0 {
			return fmt.Errorf("quantity can't be empty")
		}

		productGet, err := q.GetProduct(ctx, int64(argTX.ProductID))

		if err != nil {
			if err == sql.ErrNoRows {
				return fmt.Errorf("product_id not found")
			}

			return err
		}

		if productGet.Quantity < argTX.Quantity {
			return fmt.Errorf("Don't have enough quantity")
		}

		quantityResult := int32(productGet.Quantity) - argTX.Quantity

		arg := UpdateProductParams{
			ID: int64(argTX.ProductID),
			Quantity: sql.NullInt32{
				Int32: quantityResult,
				Valid: true,
			},
		}

		UpdateProduct, err := q.UpdateProduct(ctx, arg)
		fmt.Println("UpdateProduct", UpdateProduct)

		if err != nil {
			fmt.Println("UpdateProduct", err)
			return err
		}

		argCreateProductBuy := CreateProductParams{
			Name:        productGet.Name,
			Description: productGet.Description,
			Price:       productGet.Price,
			StoreID:     int64(argTX.StoreID),
			Quantity:    argTX.Quantity,
			ImageUrl:    productGet.ImageUrl,
			Category:    productGet.Category,
			SupplierID:  productGet.SupplierID,
		}

		result.Product, err = q.CreateProduct(ctx, argCreateProductBuy)

		if err != nil {
			return err
		}

		return nil

	})
	return result, err
}
