package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"
	"time"
)

const (
	//-----------------DISH-------------------//

	sqlstrListDish = `
					SELECT dish.id_dish, dish.name_dish, dish.description, dish.url, dish.updated, dish.created, dish.time_min
					FROM dish
					WHERE dish.user_id = $1;`

	sqlFetchDish = `
		SELECT dish.id_dish, dish.name_dish, dish.description, dish.url, dish.updated, dish.created, dish.time_min
		FROM dish
		WHERE dish.user_id = $1 AND dish.name_dish = $2;`

	sqlstrListDishNotMenu = `
					SELECT dish.id_dish, dish.name_dish, dish.description, dish.url, dish.updated, dish.created, dish.time_min
					FROM dish
					WHERE dish.user_id = $2 AND dish.id_dish NOT IN
								(
									SELECT dish.id_dish
									FROM dish
										INNER JOIN menudish on menudish.id_dish = dish.id_dish
										INNER JOIN menu on menu.id_menu = menudish.id_menu
									WHERE menu.id_menu = $1 AND dish.user_id = $2
								);`

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

	sqlstrDishUpdate = `
		UPDATE dish
		SET
			time_min = $1,
			description = $2,
			updated = now()
		WHERE name_dish = $3 AND user_id = $4
		RETURNING updated;`

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

	sqlMenuDishRemove = `DELETE FROM menudish WHERE id_menu = $1 AND id_dish = $2;`

	sqlFetchMenu = `
		SELECT menu.id_menu, menu.url, menu.created, menu.updated
		FROM menu
		WHERE menu.id_place = $1 AND menu.name_menu = $2;`

	sqlstrListMenu = `
					SELECT menu.id_menu, menu.name_menu, menu.url, menu.created, menu.updated
					FROM menu
					WHERE menu.id_place = $1;`

	sqlstrListMenuDishes = `
					SELECT dish.id_dish, dish.name_dish, dish.description, dish.url, dish.updated, dish.id_typeDish, dish.created, dish.time_min
					FROM dish
							INNER JOIN menudish on menudish.id_dish = dish.id_dish
							INNER JOIN menu on menu.id_menu = menudish.id_menu
					WHERE menu.id_menu = $1 AND dish.user_id = $2;`
	//-----------------PLACE-------------------//

	sqlstrListTypePlace = `
		SELECT type_place.id_typePlace, type_place.name_type
		FROM type_place;`

	sqlPlaceIDGetByUsr = `SELECT place.id_place
		FROM place
		WHERE place.user_id = $1;`

	sqlTypePlaceIDGetByName = `SELECT type_place.id_typePlace
		FROM type_place
		WHERE type_place.name_type = $1;`

	sqlPlaceGetByIDUsr = `
		SELECT place.name, place.phone_number, place.adress, place.city, place.url, place.id_place
		FROM place
		WHERE place.user_id = $1;`


	//-----------------USERS-------------------//

	sqlstrUserGetByLogin = `
		SELECT users.user_id, users.username, users.email, users.gravatar, users.password, users.salt
		FROM users
		WHERE users.username = $1 OR users.user_id = $1;`

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

	sqlstrListTypePersonal = `
		SELECT type_personal.id_typePersonal, type_personal.name_type
		FROM type_personal;`

	sqlstrListPersonal = `
					SELECT personal.id_personal, personal.fio, personal.phone, personal.updated, personal.created
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
