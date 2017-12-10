package pgsql

import (
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"
	"log"
	"errors"
)

const (
	sqlstrPlaceGetById = `SELECT place.name, place.phone_number, place.url, place.city, place.adress, place.user_id, place.id_typePlace
		FROM place
		WHERE users.user_id = $1;`

	sqlCreatePlace = `
		INSERT INTO users (name, phone_number, url, city, adress, user_id, id_typePlace)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id_place;
	`
)

type PlaceStorage struct {
	storage.Place
	client store.IDB
}

func (s *PlaceStorage) Create (ctx context.Context, place *types.Place) error {

	log.Println("STORAGE--- CreatePlace()")

	if place == nil {
		err := errors.New("place can not be nil")
		return  err
	}

	var (
		err error
		id  store.NullString
	)

	err = s.client.QueryRow( sqlCreatePlace, place.Meta.Name, place.Meta.Phone, place.Meta.Url, place.Meta.City, place.Meta.Adress, place.Meta.UserID, place.Meta.TypePlaceID).Scan(&id)

	place.Meta.ID = id.String

	return err
}

func newPlaceStorage(client store.IDB) *PlaceStorage {
	s := new(PlaceStorage)
	s.client = client
	return s
}
