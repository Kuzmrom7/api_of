package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/store"
	"database/sql"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/log"
	"context"
	"errors"
	"encoding/json"
)

func (s *MenuStorage) CreateMenu(ctx context.Context, menu *types.Menu) error {

	log.Debug("Storage: Menu: Insert: insert menu: %#v", menu)

	if menu == nil {
		err := errors.New("place can not be nil")
		log.Errorf("Storage: Menu: Insert: insert menu err: %s", err)
		return err
	}

	var (
		err error
		id  store.NullString
	)

	const sqlCreateMenu = `
		INSERT INTO menu (name_menu, id_place, url)
		VALUES ($1, $2, $3)
		RETURNING id_menu;
	`

	err = s.client.QueryRow(sqlCreateMenu, menu.Meta.Name, menu.Meta.PlaceID, menu.Meta.Url).Scan(&id)
	if err != nil {
		log.Errorf("Storage: Menu: Insert: insert menu query err: %s", err)
		return err
	}

	menu.Meta.ID = id.String

	return err
}

func (s *MenuStorage) List(ctx context.Context, placeid string) (map[string]*types.Menu, error) {

	menus := make(map[string]*types.Menu)

	log.Debug("Storage: Menu: List: get list menu")

	const sqlstrListMenu = `
					SELECT menu.id_menu, menu.name_menu, menu.url, menu.created, menu.updated
					FROM menu
					WHERE menu.id_place = $1;`

	rows, err := s.client.Query(sqlstrListMenu, placeid)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Menu: List: get list menu query err: %s", err)
		return nil, err
	}

	for rows.Next() {

		di := new(menuModel)

		if err := rows.Scan(&di.id, &di.name, &di.url, &di.created, &di.updated); err != nil {
			log.Errorf("Storage: Menu: List: get list menu scan rows err: %s", err)
			return nil, err
		}

		c := di.convert()
		menus[c.Meta.ID] = c
	}

	return menus, nil
}

func (s *MenuStorage) InsertDishInMenu(ctx context.Context, menuid, dishid string) error {

	log.Debugf("Storage: Menu: Dish: Insert: insert dish by id %s in menu by id %s", dishid, menuid)

	if menuid == "" {
		err := errors.New("menuid can not be nil")
		log.Errorf("Storage: Menu: Dish: Insert: insert dish in menu err: %s", err)
		return err
	}
	if dishid == "" {
		err := errors.New("dishid can not be nil")
		log.Errorf("Storage: Menu: Dish: Insert: insert dish in menu err: %s", err)
		return err
	}

	var (
		err error
		id  store.NullString
	)

	const sqlCreateMenuDish = `
		INSERT INTO menudish (id_menu, id_dish)
		VALUES ($1, $2)
		RETURNING id_menu;
	`

	err = s.client.QueryRow(sqlCreateMenuDish, menuid, dishid).Scan(&id)
	if err != nil {
		log.Errorf("Storage: Menu: Dish: Insert: insert dish in menu query err: %s", err)
		return err
	}

	return err
}

func (s *MenuStorage) DeleteDishInMenu(ctx context.Context, menuid, dishid string) error {

	log.Debugf("Storage: Menu: Dish: Delete: delete dish by id %s from menu by id %s", dishid, menuid)

	if menuid == "" {
		err := errors.New("menuid can not be nil")
		log.Errorf("Storage: Menu: Dish: Delete: delete dish from menu err: %s", err)
		return err
	}
	if dishid == "" {
		err := errors.New("dishid can not be nil")
		log.Errorf("Storage: Menu: Dish: Delete: delete dish from menu err: %s", err)
		return err
	}

	_, err := s.client.Exec(sqlMenuDishRemove, menuid, dishid)
	if err != nil {
		log.Errorf("Storage: Menu: Dish: Delete: delete dish from menu exec err: %s", err)
		return err
	}

	return nil
}

func (s *MenuStorage) Fetch(ctx context.Context, id string) (*types.Menu, error) {

	var (
		err error
		mn  = new(menuModel)
	)

	log.Debugf("Storage: Menu: Info: get menu by id: %s ", id)

	const sqlFetchMenu = `
		SELECT menu.id_menu, menu.name_menu ,menu.url, menu.created, menu.updated
		FROM menu
		WHERE menu.id_menu = $1;`

	err = s.client.QueryRow(sqlFetchMenu, id).Scan(&mn.id, &mn.name, &mn.url, &mn.created, &mn.updated)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Menu: Info: get menu by id query err: %s", err)
		return nil, err
	}

	men := mn.convert()

	return men, nil

}

func (s *MenuStorage) ListDishesInMenu(ctx context.Context, menuid, usrid string) ([]*types.Dish, error) {

	var dishes []*types.Dish

	log.Debug("Storage: Menu: Dish: List: get list dishes in menu")

	const sqlstrListMenuDishes = `
					SELECT to_json(
				json_build_object(
					'meta', json_build_object(
					'id', id_dish,
					'name', name_dish,
					'description', description,
					'timemin', time_min,
					'updated', updated
					'created', created
				),
				'urls', url
				)
			)
					FROM dish
							INNER JOIN menudish on menudish.id_dish = dish.id_dish
							INNER JOIN menu on menu.id_menu = menudish.id_menu
					WHERE menu.id_menu = $1 AND dish.user_id = $2;`

	rows, err := s.client.Query(sqlstrListMenuDishes, menuid, usrid)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Menu: Dish: List: get list dishes in menu query err: %s", err)
		return nil, err
	}

	for rows.Next() {

		var buf string

		if err := rows.Scan(&buf); err != nil {
			log.Errorf("Storage: Menu: Dish: List: get list dishes in menu scan rows err: %s", err)
			return nil, err
		}

		di := new(types.Dish)

		if err := json.Unmarshal([]byte(buf), &di); err != nil {
			return nil, err
		}

		dishes = append(dishes, di)
	}

	return dishes, nil
}

func (s *MenuStorage) ListDishesNotMenu(ctx context.Context, menuid, userid string) ([]*types.Dish, error) {

	var dishes []*types.Dish

	log.Debug("Storage: Menu: Dish: List: get list dishes not menu")

	const sqlstrListDishNotMenu = `
					SELECT to_json(
				json_build_object(
					'meta', json_build_object(
					'id', id_dish,
					'name', name_dish,
					'description', description,
					'timemin', time_min,
					'updated', updated
					'created', created
				),
				'urls', url
				)
			)
					FROM dish
					WHERE dish.user_id = $2 AND dish.id_dish NOT IN
								(
									SELECT dish.id_dish
									FROM dish
										INNER JOIN menudish on menudish.id_dish = dish.id_dish
										INNER JOIN menu on menu.id_menu = menudish.id_menu
									WHERE menu.id_menu = $1 AND dish.user_id = $2
								);`

	rows, err := s.client.Query(sqlstrListDishNotMenu, menuid, userid)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Menu: Dish: List: get list dishes not menu query err: %s", err)
		return nil, err
	}

	for rows.Next() {

		var buf string

		if err := rows.Scan(&buf); err != nil {
			log.Errorf("Storage: Menu: Dish: List: get list dishes not menu scan rows err: %s", err)
			return nil, err
		}

		di := new(types.Dish)

		if err := json.Unmarshal([]byte(buf), &di); err != nil {
			return nil, err
		}

		dishes = append(dishes, di)
	}

	return dishes, nil
}

func (nm *menuModel) convert() *types.Menu {
	c := new(types.Menu)

	c.Meta.ID = nm.id.String
	c.Meta.Name = nm.name.String
	c.Meta.Url = nm.url.String
	c.Meta.Created = nm.created
	c.Meta.Updated = nm.updated

	return c
}
