// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	ID       int32
	Username pgtype.Text
	Email    pgtype.Text
	Password pgtype.Text
}
