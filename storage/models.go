// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package storage

import (
	"database/sql"
	"time"

	"github.com/sqlc-dev/pqtype"
)

type PasswordReset struct {
	ID        int32
	UserEmail string
	UserToken string
	CreatedAt time.Time
}

type User struct {
	ID           int32
	Username     sql.NullString
	Email        sql.NullString
	PasswordHash sql.NullString
	Profile      pqtype.NullRawMessage
}

type Workout struct {
	ID          int32
	UserID      int32
	Name        string
	Description sql.NullString
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
