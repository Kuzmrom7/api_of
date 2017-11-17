package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/store"
	"github.com/orderfood/api_of/pkg/storage/storage"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	*UserStorage
}

func (s *Storage) User() storage.User {
	if s == nil {
		return nil
	}
	return s.UserStorage
}

func New(c store.Config) (*Storage, error) {

	log.Println(c.Connection)

	client, err := sql.Open("postgres", c.Connection)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(); err != nil {
		return nil, err
	}

	s := new(Storage)
	s.UserStorage = newUserStorage(client)
	return s, nil
}
