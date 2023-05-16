// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: suplier.sql

package db

import (
	"context"
	"database/sql"
)

const createSupplier = `-- name: CreateSupplier :one
INSERT INTO suppliers (
  name, 
  address, 
  email, 
  contact_phone
) values  (
  $1, $2, $3, $4)
RETURNING id, name, address, email, contact_phone, created_at
`

type CreateSupplierParams struct {
	Name         string `json:"name"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	ContactPhone string `json:"contactPhone"`
}

func (q *Queries) CreateSupplier(ctx context.Context, arg CreateSupplierParams) (Supplier, error) {
	row := q.db.QueryRowContext(ctx, createSupplier,
		arg.Name,
		arg.Address,
		arg.Email,
		arg.ContactPhone,
	)
	var i Supplier
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.Email,
		&i.ContactPhone,
		&i.CreatedAt,
	)
	return i, err
}

const getSupplier = `-- name: GetSupplier :one
  SELECT id, name, address, email, contact_phone, created_at FROM suppliers WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSupplier(ctx context.Context, id int64) (Supplier, error) {
	row := q.db.QueryRowContext(ctx, getSupplier, id)
	var i Supplier
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.Email,
		&i.ContactPhone,
		&i.CreatedAt,
	)
	return i, err
}

const updateSupplier = `-- name: UpdateSupplier :one
  UPDATE suppliers SET
  name = COALESCE($2, name),
  address = COALESCE($3, address),
  email = COALESCE($4, email),
  contact_phone = COALESCE($5, contact_phone)
  WHERE id = $1
  RETURNING id, name, address, email, contact_phone, created_at
`

type UpdateSupplierParams struct {
	ID           int64          `json:"id"`
	Name         sql.NullString `json:"name"`
	Address      sql.NullString `json:"address"`
	Email        sql.NullString `json:"email"`
	ContactPhone sql.NullString `json:"contactPhone"`
}

func (q *Queries) UpdateSupplier(ctx context.Context, arg UpdateSupplierParams) (Supplier, error) {
	row := q.db.QueryRowContext(ctx, updateSupplier,
		arg.ID,
		arg.Name,
		arg.Address,
		arg.Email,
		arg.ContactPhone,
	)
	var i Supplier
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Address,
		&i.Email,
		&i.ContactPhone,
		&i.CreatedAt,
	)
	return i, err
}
