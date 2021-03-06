package store

import "database/sql"

type IDB interface {
	IClient
	Begin() (*sql.Tx, error)
}

type IClient interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Query(query string, args ...interface{}) (*sql.Rows, error)
}
