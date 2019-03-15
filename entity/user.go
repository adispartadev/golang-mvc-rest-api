package entity

import (
	"database/sql"
)

type User struct {
	ID           int            `json:"id"`
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	RefreshToken sql.NullString `json:"-"`
}
