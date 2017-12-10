package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"
	"database/sql"
	"github.com/orderfood/api_of/pkg/common/types"
	"context"
	"log"
	"errors"
)

const (

	sqlCreateMenu = `
		INSERT INTO place (name, phone_number, url, city, adress, user_id, id_typePlace)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_place;
	`

	sqlPlaceIDGetByName = `SELECT place.id_place
		FROM place
		WHERE place.name = $1;`
)

type MenuStorage struct {
	storage.Place
	client store.IDB
}

type placeModel struct {
	id 				store.NullString
}

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

	err = s.client.QueryRow(sqlCreateMenu, menu.Meta.Name, menu.Meta.PlaceID).Scan(&id)

	menu.Meta.ID = id.String

	return err
}

func (s *MenuStorage) GetPlaceByName (ctx context.Context, name string) (string, error) {
	var (
		err error
		plc = new(placeModel)
	)

	err = s.client.QueryRow(sqlPlaceIDGetByName, name).Scan(&plc.id)

	switch err{
	case nil:
	case sql.ErrNoRows:
		return "", nil
	default:
		return "", err
	}

	placeID := plc.id.String

	return placeID, nil
}

func newMenuStorage(client store.IDB) *MenuStorage {
	s := new(MenuStorage)
	s.client = client
	return s
}