package storage

import "github.com/orderfood/api_of/pkg/storage/storage"

type Storage interface {
	User() storage.User
	Place() storage.Place
	Menu() storage.Menu
	Dish() storage.Dish
	Personal() storage.Personal
	Adress() storage.Adress
}
