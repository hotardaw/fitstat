// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"database/sql"
)

type User struct {
	UserID       int32
	Email        string
	PasswordHash string
	Username     string
	Active       sql.NullBool
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
	LastLogin    sql.NullTime
}

type UserProfile struct {
	ProfileID         int32
	UserID            sql.NullInt32
	FirstName         sql.NullString
	LastName          sql.NullString
	DateOfBirth       sql.NullTime
	Gender            sql.NullString
	HeightInches      sql.NullInt32
	ProfilePictureUrl sql.NullString
	CreatedAt         sql.NullTime
	UpdatedAt         sql.NullTime
}
