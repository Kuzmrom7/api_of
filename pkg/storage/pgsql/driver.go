package pgsql

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"

	"github.com/orderfood/api_of/pkg/log"
)

type Storage struct {
	*UserStorage
	*PlaceStorage
	*MenuStorage
	*DishStorage
	*PersonalStorage
	*AdressStorage
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

func (s *Storage) Dish() storage.Dish {
	if s == nil {
		return nil
	}
	return s.DishStorage
}

func (s *Storage) Personal() storage.Personal {
	if s == nil {
		return nil
	}
	return s.PersonalStorage
}

func (s *Storage) Adress() storage.Adress {
	if s == nil {
		return nil
	}
	return s.AdressStorage
}

func New(c store.Config) (*Storage, error) {

	/*log.Println(c.Connection)*/
	log.Debugf("Connection to DB: ", c.Connection)

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
	s.DishStorage = newDishStorage(client)
	s.PersonalStorage = newPersonalStorage(client)
	s.AdressStorage = newAdressStorage(client)
	return s, nil
}
