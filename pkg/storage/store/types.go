package store

import "database/sql"

type NullString struct {
	sql.NullString
}
