-- name: CreateProduct :one
INSERT INTO products (
  name, 
  description, 
  price, 
  store_id,
  quantity
) values  (
  $1, $2, $3, $4, $5)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: GetProductForUpdate :one
SELECT * FROM products WHERE id = $1 LIMIT 1 FOR UPDATE FOR NO KEY UPDATE;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE products SET
name = COALESCE(sqlc.narg(name), name),
description = COALESCE(sqlc.narg(description), description),
price = COALESCE(sqlc.narg(price), price),
quantity = COALESCE(sqlc.narg(quantity), quantity)
WHERE id = $1
RETURNING *;

-- name: GetProductsWithJoinWithStore :many
  SELECT products.*, stores.* FROM products INNER JOIN stores ON products.store_id = stores.id 
  ORDER BY products.id
  LIMIT $1
  OFFSET $2;


