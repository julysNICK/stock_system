// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: store.sql

package db

import (
	"context"
	"database/sql"
)

const createStore = `-- name: CreateStore :one
INSERT into stores (
name, 
address, 
contact_email, 
contact_phone, 
hashed_password
) values  (
  $1, $2, $3, $4, $5) 
RETURNING id, name, address, contact_email, contact_phone, hashed_password, created_at
`

type CreateStoreParams struct {
	Name           string `json:"name"`
	Address        string `json:"address"`
	ContactEmail   string `json:"contactEmail"`
	ContactPhone   string `json:"contactPhone"`
	HashedPassword string `json:"hashedPassword"`
}

func (q *Queries) CreateStore(ctx context.Context, arg CreateStoreParams) (Store, error) {
	row := q.db.QueryRowContext(ctx, createStore,
		arg.Name,
		arg.Address,
		arg.ContactEmail,
		arg.ContactPhone,
		arg.HashedPassword,
	)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}

const getStore = `-- name: GetStore :one
SELECT id, name, address, contact_email, contact_phone, hashed_password, created_at FROM stores WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStore(ctx context.Context, id int64) (Store, error) {
	row := q.db.QueryRowContext(ctx, getStore, id)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}

const getStoreByEmail = `-- name: GetStoreByEmail :one
SELECT id, name, address, contact_email, contact_phone, hashed_password, created_at FROM stores WHERE contact_email = $1 LIMIT 1
`

func (q *Queries) GetStoreByEmail(ctx context.Context, contactEmail string) (Store, error) {
	row := q.db.QueryRowContext(ctx, getStoreByEmail, contactEmail)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}

const getStoreForUpdate = `-- name: GetStoreForUpdate :one
SELECT id, name, address, contact_email, contact_phone, hashed_password, created_at FROM stores WHERE id = $1 LIMIT 1 FOR UPDATE FOR NO KEY UPDATE
`

func (q *Queries) GetStoreForUpdate(ctx context.Context, id int64) (Store, error) {
	row := q.db.QueryRowContext(ctx, getStoreForUpdate, id)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}

const listStores = `-- name: ListStores :many
SELECT id, name, address, contact_email, contact_phone, hashed_password, created_at FROM stores
ORDER BY id
LIMIT $1 
OFFSET $2
`

type ListStoresParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListStores(ctx context.Context, arg ListStoresParams) ([]Store, error) {
	rows, err := q.db.QueryContext(ctx, listStores, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Store{}
	for rows.Next() {
		var i Store
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Address,
			&i.ContactEmail,
			&i.ContactPhone,
			&i.HashedPassword,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStore = `-- name: UpdateStore :one
UPDATE stores SET
name = COALESCE($2, name),
address = COALESCE($3, address),
contact_email = COALESCE($4, contact_email),
contact_phone = COALESCE($5, contact_phone),
hashed_password = COALESCE($6, hashed_password)
WHERE id = $1
RETURNING id, name, address, contact_email, contact_phone, hashed_password, created_at
`

type UpdateStoreParams struct {
	ID             int64          `json:"id"`
	Name           sql.NullString `json:"name"`
	Address        sql.NullString `json:"address"`
	ContactEmail   sql.NullString `json:"contactEmail"`
	ContactPhone   sql.NullString `json:"contactPhone"`
	HashedPassword sql.NullString `json:"hashedPassword"`
}

func (q *Queries) UpdateStore(ctx context.Context, arg UpdateStoreParams) (Store, error) {
	row := q.db.QueryRowContext(ctx, updateStore,
		arg.ID,
		arg.Name,
		arg.Address,
		arg.ContactEmail,
		arg.ContactPhone,
		arg.HashedPassword,
	)
	var i Store
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.ContactEmail,
		&i.ContactPhone,
		&i.HashedPassword,
		&i.CreatedAt,
	)
	return i, err
}
