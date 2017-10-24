package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

)
const (
	DB_CONNECT_STRING = "host= 0.0.0.0 port=5432 user=orderfood  password=orderfood dbname=orderfood sslmode=disable"
)

var (
	DB *sql.DB
)

func DB_connect(){
	var err error
	DB, err = sql.Open("postgres", DB_CONNECT_STRING)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}

