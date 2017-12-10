package pgsql

import (
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"
	"log"
	"errors"
	"database/sql"
)

const (

	sqlCreatePlace = `
		INSERT INTO place (name, phone_number, url, city, adress, user_id, id_typePlace)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_place;
	`
	sqlTypePlaceIDGetByName = `SELECT type_place.id_typePlace
		FROM type_place
		WHERE type_place.name_type = $1;`
)

type PlaceStorage struct {
	storage.Place
	client store.IDB
}

type typeplaceModel struct {
	id 				store.NullString
}

func (s *PlaceStorage) CreatePlace(ctx context.Context, place *types.Place) error {

	log.Println("STORAGE--- CreatePlace()")

	if place == nil {
		err := errors.New("place can not be nil")
		return err
	}

	var (
		err error
		id  store.NullString
	)

	err = s.client.QueryRow(sqlCreatePlace, place.Meta.Name, place.Meta.Phone, place.Meta.Url, place.Meta.City, place.Meta.Adress, place.Meta.UserID, place.Meta.TypePlaceID).Scan(&id)

	place.Meta.ID = id.String

	return err
}

func (s *PlaceStorage) GetTypePlaceByName (ctx context.Context, name string) (string, error) {
	var (
		err error
		pl = new(typeplaceModel)
	)

	err = s.client.QueryRow(sqlTypePlaceIDGetByName, name).Scan(&pl.id)

	switch err{
	case nil:
	case sql.ErrNoRows:
		return "", nil
	default:
		return "", err
	}

	typeplaceID := pl.id.String

	return typeplaceID, nil
}

func newPlaceStorage(client store.IDB) *PlaceStorage {
	s := new(PlaceStorage)
	s.client = client
	return s
}
