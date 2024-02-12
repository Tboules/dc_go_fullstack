// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createNewUser = `-- name: CreateNewUser :execresult
INSERT INTO ` + "`" + `user` + "`" + ` (
  name, email, image, provider_id
) VALUES (
  ?, ?, ?, ?
)
`

type CreateNewUserParams struct {
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Image      sql.NullString `json:"image"`
	ProviderID string         `json:"provider_id"`
}

func (q *Queries) CreateNewUser(ctx context.Context, arg CreateNewUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createNewUser,
		arg.Name,
		arg.Email,
		arg.Image,
		arg.ProviderID,
	)
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM ` + "`" + `user` + "`" + `
WHERE ` + "`" + `id` + "`" + ` = ?
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, provider_id, name, email, email_verified, image FROM ` + "`" + `user` + "`" + `
WHERE id = ? LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.Name,
		&i.Email,
		&i.EmailVerified,
		&i.Image,
	)
	return i, err
}

const getUserByProviderId = `-- name: GetUserByProviderId :one
SELECT id, provider_id, name, email, email_verified, image FROM ` + "`" + `user` + "`" + `
WHERE ` + "`" + `provider_id` + "`" + ` = ? LIMIT 1
`

func (q *Queries) GetUserByProviderId(ctx context.Context, providerID string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByProviderId, providerID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.Name,
		&i.Email,
		&i.EmailVerified,
		&i.Image,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, provider_id, name, email, email_verified, image FROM ` + "`" + `user` + "`" + `
ORDER BY ` + "`" + `name` + "`" + `
`

func (q *Queries) GetUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.ProviderID,
			&i.Name,
			&i.Email,
			&i.EmailVerified,
			&i.Image,
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

const updateUser = `-- name: UpdateUser :execresult
UPDATE ` + "`" + `user` + "`" + `
  SET name = ?,
  email = ?,
  email_verified = ?,
  provider_id = ?,
  image = ?
WHERE id = ?
`

type UpdateUserParams struct {
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	EmailVerified sql.NullBool   `json:"email_verified"`
	ProviderID    string         `json:"provider_id"`
	Image         sql.NullString `json:"image"`
	ID            int64          `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.EmailVerified,
		arg.ProviderID,
		arg.Image,
		arg.ID,
	)
}
