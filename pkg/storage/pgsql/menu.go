package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/store"
	"database/sql"
	"github.com/orderfood/api_of/pkg/common/types"
	"context"
	"log"
	"errors"
)

func (s *MenuStorage) CreateMenu(ctx context.Context, menu *types.Menu) error {

	log.Println("STORAGE--- CreateMenu()")

	if menu == nil {
		err := errors.New("menu can not be nil")
		return err
	}

	var (
		err error
		id  store.NullString
	)

	err = s.client.QueryRow(sqlCreateMenu, menu.Meta.Name, menu.Meta.PlaceID, menu.Meta.Url).Scan(&id)

	menu.Meta.ID = id.String

	return err
}

func (s *MenuStorage) List(ctx context.Context, placeid string) (map[string]*types.Menu, error) {

	menus := make(map[string]*types.Menu)

	rows, err := s.client.Query(sqlstrListMenu, placeid)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:

		return nil, err
	}

	for rows.Next() {

		di := new(menuModel)

		if err := rows.Scan(&di.id, &di.name, &di.url, &di.created, &di.updated); err != nil {

			return nil, err
		}

		c := di.convert()
		menus[c.Meta.ID] = c
	}

	return menus, nil
}

func (s *MenuStorage) GetIDmenuByName(ctx context.Context, name string) (string, error) {
	var (
		err error
		di  = new(dichModel)
	)

	err = s.client.QueryRow(sqlMenuIDGetByName, name).Scan(&di.id)

	switch err {
	case nil:
	case sql.ErrNoRows:
		return "", nil
	default:
		return "", err
	}

	menuID := di.id.String

	return menuID, nil
}

func (s *MenuStorage) InsertDishInMenu(ctx context.Context, menuid, dishid string) error {

	log.Println("STORAGE--- InsertDishInMenu()")

	if menuid == "" {
		err := errors.New("menuid can not be nil")
		return err
	}
	if dishid == "" {
		err := errors.New("dishid can not be nil")
		return err
	}

	var (
		err error
		id  store.NullString
	)

	err = s.client.QueryRow(sqlCreateMenuDish, menuid, dishid).Scan(&id)

	return err
}


func (s *MenuStorage) DeleteDishInMenu(ctx context.Context, menuid, dishid string) error {

	log.Println("STORAGE--- DeleteDishInMenu()")

	if menuid == "" {
		err := errors.New("menuid can not be nil")
		return err
	}
	if dishid == "" {
		err := errors.New("dishid can not be nil")
		return err
	}

	_, err := s.client.Exec(sqlMenuDishRemove, menuid, dishid)
	if err != nil {
		return err
	}
	return nil

}


func (s *MenuStorage) Fetch(ctx context.Context, idplace, name string) (*types.Menu, error) {

	var (
		err error
		mn  = new(menuModel)
	)

	err = s.client.QueryRow(sqlFetchMenu, idplace, name).Scan(&mn.id, &mn.url, &mn.created, &mn.updated)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}

	men := mn.convert()

	men.Meta.Name = name

	return men, nil

}

func (s *MenuStorage) ListDishesInMenu(ctx context.Context, menuid, typedishid string) (map[string]*types.Dish, error) {

	menudishes := make(map[string]*types.Dish)

	rows, err := s.client.Query(sqlstrListMenuDishes, menuid, typedishid)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:

		return nil, err
	}

	for rows.Next() {

		di := new(dichModel)

		if err := rows.Scan(&di.id, &di.name, &di.description, &di.url, &di.updated); err != nil {

			return nil, err
		}

		c := di.convert()
		menudishes[c.Meta.ID] = c
	}

	return menudishes, nil
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
