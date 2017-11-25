package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"
	"github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/common/types"
	"database/sql"
	"github.com/orderfood/api_of/pkg/common/errors"
)

const (
	sqlstrGetMenuById = `
		SELECT m.id_menu, m.name_menu, m.created, m.updated FROM menu as m  WHERE m.user_id = $1;
		`
	sqlstrCreateMenu = `

		INSERT INTO menu (name_menu, id_place, user_id)
    VALUES ($1, $2, $3);
		`
)

type MenuStorage struct {
	storage.Menu
	client store.IDB
}

type MenuModel struct {
	Id_menu 	store.NullString
	Name_menu 	store.NullString
	Created		store.NullString
	Updated 	store.NullString
}

func (menu *MenuModel) convert() *types.Menu{
	var m = new(types.Menu)

	m.Meta.ID_menu = menu.Id_menu.String
	m.Meta.Name_menu = menu.Name_menu.String
	m.Meta.Created = menu.Created.String
	m.Meta.Updated = menu.Updated.String

}


func (s *MenuStorage) GetMenuByUserId (ctx context.Context, user_id string)(*types.Menu, error) {

	var (
		err error
		menu = new(MenuModel)
	)
	err = s.client.QueryRow(sqlstrGetMenuById, user_id).Scan(&menu.Id_menu, &menu.Name_menu, &menu.Created, &menu.Updated)
	switch err{
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}

	mnu := menu.convert()

	return mnu, nil
	}

func (s *MenuStorage) CreateMenu (ctx context.Context, menu *types.MenuAdd)(error){

	if menu == nil {
		err := errors.New("Menu not be nil")
		return err
		}

		var (
			err error
			id 	store.NullString

		)

	err = s.client.QueryRow(sqlstrCreateMenu, menu.Name_menu, menu.User_id, menu.Id_Place).Scan(&id)

	menu.ID_menu = id.String

	return err
}

/*
func (s *MenuModel) UpdateMenu (ctx context.Context, menu types.Menu) error {

}*/
