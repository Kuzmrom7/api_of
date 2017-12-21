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

	err = s.client.QueryRow(sqlCreateMenu, menu.Meta.Name, menu.Meta.PlaceID).Scan(&id)

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

