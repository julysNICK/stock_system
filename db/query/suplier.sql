-- name: CreateSupplier :one
INSERT INTO suppliers (
  name, 
  address, 
  email, 
  contact_phone
) values  (
  $1, $2, $3, $4)
RETURNING *;

  -- name: GetSupplier :one
  SELECT * FROM suppliers WHERE id = $1 LIMIT 1;

  -- name: UpdateSupplier :one
  UPDATE suppliers SET
  name = COALESCE(sqlc.narg(name), name),
  address = COALESCE(sqlc.narg(address), address),
  email = COALESCE(sqlc.narg(email), email),
  contact_phone = COALESCE(sqlc.narg(contact_phone), contact_phone)
  WHERE id = $1
  RETURNING *;