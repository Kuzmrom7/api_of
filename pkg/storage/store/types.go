package store

import "database/sql"

type NullString struct {
	sql.NullString
}

type NullInt64 struct {
	sql.NullInt64
}

