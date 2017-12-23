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

func (s *MenuStorage) GetPlaceByName(ctx context.Context, name string) (string, error) {
	var (
		err error
		plc = new(idModel)
	)

	err = s.client.QueryRow(sqlPlaceIDGetByName, name).Scan(&plc.id)

	switch err {
	case nil:
	case sql.ErrNoRows:
		return "", nil
	default:
		return "", err
	}

	placeID := plc.id.String

	return placeID, nil
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

		if err := rows.Scan(&di.id, &di.name, &di.url); err != nil {

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

func (s *MenuStorage) CreateMenuDish(ctx context.Context, menuid, dishid string) error {

	log.Println("STORAGE--- CreateMenuDish()")

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

func (nm *menuModel) convert() *types.Menu {
	c := new(types.Menu)

	c.Meta.ID = nm.id.String
	c.Meta.Name = nm.name.String
	c.Meta.Url = nm.url.String

	return c
}