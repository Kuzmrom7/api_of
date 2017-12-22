package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"
)

const (
	//-----------------DISH-------------------//

	sqlstrListDish = `
		SELECT dish.id_dish, dish.name_dish, dish.description
		FROM dish;`

	sqlCreateDich = `
		INSERT INTO dish (name_dish, description, time_min)
		VALUES ($1, $2, $3)
		RETURNING id_dish;
	`

	sqlDichIDGetByName = `SELECT dish.id_dish
		FROM dish
		WHERE dish.name_dish = $1;`

	sqlDichRemove = `DELETE FROM dish WHERE id_dish = $1;`

	//-----------------MENU-------------------//

	sqlCreateMenu = `
		INSERT INTO menu (name_menu, id_place)
		VALUES ($1, $2)
		RETURNING id_menu;
	`
	//for menu and personal
	sqlPlaceIDGetByName = `SELECT place.id_place
		FROM place
		WHERE place.name = $1;`

	//-----------------PLACE-------------------//

	sqlstrListType = `
		SELECT type_place.id_typePlace, type_place.name_type
		FROM type_place;`

	sqlCreatePlace = `
		INSERT INTO place (name, phone_number, url, city, adress, user_id, id_typePlace)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_place;
	`
	sqlTypePlaceIDGetByName = `SELECT type_place.id_typePlace
		FROM type_place
		WHERE type_place.name_type = $1;`

	sqlPlaceGetByIDUsr = `
		SELECT place.name, place.phone_number, place.adress, place.city, place.url
		FROM place
		WHERE place.user_id = $1;`

	sqlstrPlaceUpdate = `
		UPDATE place
		SET
			phone_number = $1,
			adress = $2,
			city = $3,
			url = $4,
			updated = now()
		WHERE name = $5
		RETURNING updated;`

	//-----------------USERS-------------------//

	sqlstrUserGetByLogin = `
		SELECT users.user_id, users.username, users.email, users.gravatar, users.password, users.salt
		FROM users
		WHERE users.username = $1;`

	sqlstrUserExistsByLogin = `
		SELECT TRUE
		FROM users
		WHERE users.username = $1 OR users.email = $1
	`

	sqlstrUserGetById = `
		SELECT users.user_id, users.username, users.email, users.gravatar, users.password, users.salt
		FROM users
		WHERE users.user_id = $1;
	`

	sqlCreateUser = `
		INSERT INTO users (username, email, gravatar, password, salt)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id;
	`
	//-----------------PERSONAL-------------------//
	ssqlTypePersonalIDGetByName = `SELECT type_personal.id_typePersonal
		FROM type_personal
		WHERE type_personal.name_type = $1;`

	sqlCreatePerson = `
		INSERT INTO personal (fio, phone, id_place, id_typePersonal)
		VALUES ($1, $2, $3, $4)
		RETURNING id_personal;
	`
	sqlPlaceIDGetByUsr = `SELECT place.id_place
		FROM place
		WHERE place.user_id = $1;`
)

//-----------------------STORAGEs------------------------//
type DichStorage struct {
	storage.Dich
	client store.IDB
}

type MenuStorage struct {
	storage.Menu
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
	name        store.NullString
	description store.NullString
}

type idModel struct {
	id store.NullString
}

type typeplaceModel struct {
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

func newDichStorage(client store.IDB) *DichStorage {
	s := new(DichStorage)
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
