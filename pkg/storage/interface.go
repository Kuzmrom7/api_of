package storage

import "github.com/orderfood/api_of/pkg/storage/storage"

type Storage interface {
	User() storage.User
	TypePlace() storage.TypePlace
	Place() storage.Place
	Menu()  storage.Menu
	Sections()  storage.Sections
	TypesDish() storage.TypesDishes
	Dish()   storage.Dish
	Personal() storage.Personal
	TypePersonal() storage.TypePersonal
}