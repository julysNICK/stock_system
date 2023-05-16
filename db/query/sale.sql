-- name: CreateSale :one
INSERT INTO  sales (
  product_id,
  sale_date,
  quantity_sold

) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetSale :one
SELECT * FROM sales WHERE id = $1 LIMIT 1;

-- name: DeleteSale :exec
DELETE FROM sales WHERE id = $1;

-- name: ListSales :many
SELECT * FROM sales ORDER BY id LIMIT $1 OFFSET $2;