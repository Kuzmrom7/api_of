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
	*PlaceStorage
	*MenuStorage
	*DichStorage
	*PersonalStorage
}

func (s *Storage) User() storage.User {
	if s == nil {
		return nil
	}
	return s.UserStorage
}

func (s *Storage) Place() storage.Place {
	if s == nil {
		return nil
	}
	return s.PlaceStorage
}

func (s *Storage) Menu() storage.Menu {
	if s == nil {
		return nil
	}
	return s.MenuStorage
}

func (s *Storage) Dich() storage.Dich {
	if s == nil {
		return nil
	}
	return s.DichStorage
}

func (s *Storage) Personal() storage.Personal {
	if s == nil {
		return nil
	}
	return s.PersonalStorage
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
	s.PlaceStorage = newPlaceStorage(client)
	s.MenuStorage = newMenuStorage(client)
	s.DichStorage = newDichStorage(client)
	s.PersonalStorage = newPersonalStorage(client)
	return s, nil
}
