-- name: CreateProduct :one
INSERT INTO products (
  name, 
  description, 
  category,
  image_url,
  price, 
  store_id,
  supplier_id,
  quantity
) values  (
  $1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: GetProductForUpdate :one
SELECT * FROM products WHERE id = $1 LIMIT 1 FOR UPDATE FOR NO KEY UPDATE;


-- name: ListAllProducts :many
SELECT * FROM products ORDER BY id LIMIT $1 OFFSET $2;

-- name: ListProducts :many
SELECT * FROM products WHERE category = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: SearchProducts :many
SELECT * FROM products WHERE name ILIKE '%' || $1 || '%' ;

-- name: UpdateProduct :one
UPDATE products SET
name = COALESCE(sqlc.narg(name), name),
description = COALESCE(sqlc.narg(description), description),
category = COALESCE(sqlc.narg(category), category),
image_url = COALESCE(sqlc.narg(image_url), image_url),
price = COALESCE(sqlc.narg(price), price),
quantity = COALESCE(sqlc.narg(quantity), quantity)
WHERE id = $1
RETURNING *;

-- name: GetProductsWithJoinWithStore :many
  SELECT products.*, stores.* FROM products INNER JOIN stores ON products.store_id = stores.id WHERE products.category = $1
  ORDER BY products.id
  LIMIT $2
  OFFSET $3;


-- name: GetProductsWithJoinWithSupplierBySupplierId :many
  SELECT products.*, suppliers.* FROM products INNER JOIN suppliers ON products.supplier_id = suppliers.id WHERE products.supplier_id = $1
  ORDER BY products.id
  LIMIT $2
  OFFSET $3;


-- name: GetProductsByCategory :many
SELECT * FROM products WHERE category = $1
ORDER BY id
LIMIT $2
OFFSET $3;

