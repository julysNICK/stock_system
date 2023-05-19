-- name: CreateStore :one
INSERT into stores (
name, 
address, 
contact_email, 
contact_phone, 
hashed_password
) values  (
  $1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetStore :one
SELECT * FROM stores WHERE id = $1 LIMIT 1;


-- name: GetStoreByEmail :one
SELECT * FROM stores WHERE contact_email = $1 LIMIT 1;

-- name: GetStoreForUpdate :one
SELECT * FROM stores WHERE id = $1 LIMIT 1 FOR UPDATE FOR NO KEY UPDATE;

-- name: ListStores :many
SELECT * FROM stores
ORDER BY id
LIMIT $1 
OFFSET $2;

-- name: UpdateStore :one
UPDATE stores SET
name = COALESCE(sqlc.narg(name), name),
address = COALESCE(sqlc.narg(address), address),
contact_email = COALESCE(sqlc.narg(contact_email), contact_email),
contact_phone = COALESCE(sqlc.narg(contact_phone), contact_phone),
hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password)
WHERE id = $1
RETURNING *;