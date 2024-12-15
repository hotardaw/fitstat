// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, password_hash, username)
VALUES ($1, $2, $3)
RETURNING user_id, email, password_hash, username, active, created_at, updated_at, last_login
`

type CreateUserParams struct {
	Email        string
	PasswordHash string
	Username     string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.PasswordHash, arg.Username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.PasswordHash,
		&i.Username,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1
RETURNING user_id, email, password_hash, username, active, created_at, updated_at, last_login
`

func (q *Queries) DeleteUser(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, userID)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT user_id, email, password_hash, username, active, created_at, updated_at, last_login
FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Email,
			&i.PasswordHash,
			&i.Username,
			&i.Active,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LastLogin,
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

const getUser = `-- name: GetUser :one

SELECT users.user_id, users.email, users.password_hash, users.username, users.active, users.created_at, users.updated_at, users.last_login
FROM users
WHERE users.user_id = $1
`

// The comments above each query are SQLc directives dictating the naming of the generated Go func & what type of query it is (:one, :many, or :exec).
//
//	:many returns a slice of records via QueryContext
//
// :one returns a single record via QueryRowContext
// :exec returns the error from ExecContext
// More: https://docs.sqlc.dev/en/latest/reference/query-annotations.html
func (q *Queries) GetUser(ctx context.Context, userID int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.PasswordHash,
		&i.Username,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET email = $2, password_hash = $3, username = $4
WHERE user_id = $1
RETURNING user_id, email, password_hash, username, active, created_at, updated_at, last_login
`

type UpdateUserParams struct {
	UserID       int32
	Email        string
	PasswordHash string
	Username     string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.UserID,
		arg.Email,
		arg.PasswordHash,
		arg.Username,
	)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.PasswordHash,
		&i.Username,
		&i.Active,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
	)
	return i, err
}
