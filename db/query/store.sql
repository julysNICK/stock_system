-- name: CreateStore :one
INSERT into stores (name, 
address, 
contact_email, 
contact_phone, 
hashed_password) values 
($1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetStore :one
SELECT * FROM stores WHERE id = $1 LIMIT 1;

-- name: GetStoreForUpdate :one
SELECT * FROM stores WHERE id = $1 LIMIT 1 FOR UPDATE FOR NO KEY UPDATE;

-- name: ListStores :many
SELECT * FROM stores
ORDER BY id
LIMIT $1 
OFFSET $2;