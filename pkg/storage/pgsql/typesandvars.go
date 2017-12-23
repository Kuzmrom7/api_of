package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"
	"time"
)

const (
	//-----------------DISH-------------------//

	sqlstrListDish = `
					SELECT dish.id_dish, dish.name_dish, dish.description, dish.url, dish.updated
					FROM dish
					WHERE dish.user_id = $1;`

	sqlCreateDich = `
		INSERT INTO dish (name_dish, description, time_min, id_typeDish, url, user_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_dish;
	`

	sqlDichIDGetByName = `SELECT dish.id_dish
		FROM dish
		WHERE dish.name_dish = $1 and dish.user_id = $2;`

	sqlDichRemove = `DELETE FROM dish WHERE id_dish = $1;`

	sqlstrListTypeDish = `
		SELECT type_dish.id_typeDish, type_dish.name_typeDish
		FROM type_dish;`

	ssqlTypeDishlIDGetByName = `SELECT type_dish.id_typeDish
		FROM type_dish
		WHERE type_dish.name_typeDish = $1;`

	//-----------------MENU-------------------//

	sqlMenuIDGetByName = `SELECT menu.id_menu
		FROM menu
		WHERE menu.name_menu = $1;`

	sqlCreateMenu = `
		INSERT INTO menu (name_menu, id_place, url)
		VALUES ($1, $2, $3)
		RETURNING id_menu;
	`

	sqlCreateMenuDish = `
		INSERT INTO menudish (id_menu, id_dish)
		VALUES ($1, $2)
		RETURNING id_menu;
	`
	sqlFetchMenu = `
		SELECT menu.id_menu, menu.name_menu, menu.url, menu.created, menu.updated
		FROM menu
		WHERE menu.id_place = $1 AND menu.name_menu = $2;`

	//for menu and personal
	sqlPlaceIDGetByName = `SELECT place.id_place
		FROM place
		WHERE place.name = $1;`

	sqlstrListMenu = `
					SELECT menu.id_menu, menu.name_menu, menu.url, menu.created, menu.updated
					FROM menu
					WHERE menu.id_place = $1;`
	//-----------------PLACE-------------------//

	sqlstrListTypePlace = `
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
		SELECT place.name, place.phone_number, place.adress, place.city, place.url, place.id_place
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

	sqlstrListTypePersonal = `
		SELECT type_personal.id_typePersonal, type_personal.name_type
		FROM type_personal;`

	sqlstrListPersonal = `
					SELECT personal.id_personal, personal.fio, personal.phone, personal.updated
					FROM personal
					WHERE personal.id_place = $1;`
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
	url         store.NullString
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

type personalModel struct {
	id      store.NullString
	fio     store.NullString
	phone   store.NullString
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
