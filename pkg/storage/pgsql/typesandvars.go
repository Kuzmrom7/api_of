package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"
	"time"
)

//-----------------------STORAGEs------------------------//
type DishStorage struct {
	storage.Dish
	client store.IDB
}

type MenuStorage struct {
	storage.Menu
	client store.IDB
}

type AdressStorage struct {
	storage.Adress
	client store.IDB
}

type PersonalStorage struct {
	storage.Personal
	client store.IDB
}

type PlaceStorage struct {
	storage.Place
	client store.IDB
}

type UserStorage struct {
	storage.User
	client store.IDB
}

//-----------------------MODELs-------------------------//
type dichModel struct {
	id          store.NullString
	id_Type     store.NullString
	name        store.NullString
	description store.NullString
	url         store.NullString
	timemin     store.NullInt64
	created     time.Time
	updated     time.Time
}

//-----------------------MODELs-------------------------//
type menuModel struct {
	id      store.NullString
	name    store.NullString
	url     store.NullString
	created time.Time
	updated time.Time
}

type adressModel struct {
	id      store.NullString
	name    store.NullString
	idplace store.NullString
}

type personalModel struct {
	id      store.NullString
	fio     store.NullString
	phone   store.NullString
	created time.Time
	updated time.Time
}

type idModel struct {
	id store.NullString
}

type typeModel struct {
	id   store.NullString
	name store.NullString
}

type typeModelDishes struct {
	id   store.NullString
	name store.NullString
}

type typeModelPersonals struct {
	id   store.NullString
	name store.NullString
}

type Users struct {
	User_id  string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Created  string `json:"created"`
	Updated  string `json:"updated"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type userModel struct {
	id       store.NullString
	username store.NullString
	email    store.NullString
	gravatar store.NullString
	password store.NullString
	salt     store.NullString
}

type placeModel struct {
	id     store.NullString
	name   store.NullString
	phone  store.NullString
	url    store.NullString
	city   store.NullString
	adress store.NullString
}

//-------------------------------------------------------------------------------------------------------//
func newMenuStorage(client store.IDB) *MenuStorage {
	s := new(MenuStorage)
	s.client = client
	return s
}

func newDishStorage(client store.IDB) *DishStorage {
	s := new(DishStorage)
	s.client = client
	return s
}

func newPersonalStorage(client store.IDB) *PersonalStorage {
	s := new(PersonalStorage)
	s.client = client
	return s
}

func newPlaceStorage(client store.IDB) *PlaceStorage {
	s := new(PlaceStorage)
	s.client = client
	return s
}

func newUserStorage(client store.IDB) *UserStorage {
	s := new(UserStorage)
	s.client = client
	return s
}

func newAdressStorage(client store.IDB) *AdressStorage {
	s := new(AdressStorage)
	s.client = client
	return s
}
