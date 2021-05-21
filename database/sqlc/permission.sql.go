// Code generated by sqlc. DO NOT EDIT.
// source: permission.sql

package database

import (
	"context"
)

const createPermission = `-- name: CreatePermission :one
INSERT INTO permissions (
    name
) VALUES (
    $1
)
RETURNING id, name, created_at, deleted_at, updated_at
`

func (q *Queries) CreatePermission(ctx context.Context, name string) (Permission, error) {
	row := q.db.QueryRowContext(ctx, createPermission, name)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePermission = `-- name: DeletePermission :exec
UPDATE permissions
SET deleted_at = now()
WHERE id = $1
`

func (q *Queries) DeletePermission(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePermission, id)
	return err
}

const getPermission = `-- name: GetPermission :one
SELECT id, name, created_at, deleted_at, updated_at FROM permissions WHERE id = 1 AND deleted_at is null
`

func (q *Queries) GetPermission(ctx context.Context) (Permission, error) {
	row := q.db.QueryRowContext(ctx, getPermission)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updatePermission = `-- name: UpdatePermission :one
UPDATE permissions
SET name = $2
WHERE id = $1
RETURNING id, name, created_at, deleted_at, updated_at
`

type UpdatePermissionParams struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdatePermission(ctx context.Context, arg UpdatePermissionParams) (Permission, error) {
	row := q.db.QueryRowContext(ctx, updatePermission, arg.ID, arg.Name)
	var i Permission
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.DeletedAt,
		&i.UpdatedAt,
	)
	return i, err
}