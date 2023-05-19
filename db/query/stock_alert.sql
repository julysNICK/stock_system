
-- name: CreateStockAlert :one
INSERT INTO stock_alerts (
  product_id, 
  supplier_id,
  alert_quantity
) values  (
  $1, $2, $3)
  RETURNING *;

  -- name: GetStockAlert :one
  SELECT * FROM stock_alerts WHERE id = $1 LIMIT 1;


  -- name: GetStockAlertsByProductIdAndSupplierId :many
  SELECT * FROM stock_alerts WHERE product_id = $1 AND supplier_id = $2;


  -- name: DeleteStockAlert :exec
  DELETE FROM stock_alerts WHERE id = $1;

  -- name: UpdateStockAlert :one
  UPDATE stock_alerts SET
  alert_quantity = COALESCE(sqlc.narg(alert_quantity), alert_quantity)
  WHERE id = $1
  RETURNING *;
