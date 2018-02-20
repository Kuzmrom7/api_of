package pgsql

import (
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/storage/store"
	"github.com/orderfood/api_of/pkg/log"
	"errors"
	"database/sql"
	"time"
)

func (s *PlaceStorage) GetPlaceByIDUser(ctx context.Context, id string) (*types.Place, error) {

	var (
		err error
		pl  = new(placeModel)
	)

	log.Debugf("Storage: Place: Get: get place by user id: %s ", id)

	err = s.client.QueryRow(sqlPlaceGetByIDUsr, id).Scan(&pl.name, &pl.phone, &pl.adress,
		&pl.city, &pl.url, &pl.id)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Place: Get: get place by user id query err: %s", err)
		return nil, err
	}

	plc := pl.convert()

	return plc, nil

}

func (s *PlaceStorage) CreatePlace(ctx context.Context, place *types.Place) error {

	log.Debug("Storage: Place: Insert: insert place: %#v", place)

	if place == nil {
		err := errors.New("place can not be nil")
		log.Errorf("Storage: Place: Insert: insert place err: %s", err)
		return err
	}

	var (
		err error
		id  store.NullString
	)

	err = s.client.QueryRow(sqlCreatePlace, place.Meta.Name, place.Meta.Phone, place.Meta.Url, place.Meta.City, place.Meta.Adress, place.Meta.UserID, place.Meta.TypePlaceID).Scan(&id)
	if err != nil {
		log.Errorf("Storage: Place: Insert: insert place query err: %s", err)
		return err
	}
	place.Meta.ID = id.String

	return err
}

func (s *PlaceStorage) GetTypePlaceByName(ctx context.Context, name string) (string, error) {
	var (
		err error
		pl  = new(typeModel)
	)

	log.Debugf("Storage: Place: GetType: get type place by name: %s ", name)

	err = s.client.QueryRow(sqlTypePlaceIDGetByName, name).Scan(&pl.id)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return "", nil
	default:
		log.Errorf("Storage: Place: GetType: get type place by name query err: %s", err)
		return "", err
	}

	typeplaceID := pl.id.String

	return typeplaceID, nil
}

func (s *PlaceStorage) ListType(ctx context.Context) (map[string]*types.TypePlaces, error) {

	places := make(map[string]*types.TypePlaces)

	log.Debug("Storage: Place: ListType: get type place list")

	rows, err := s.client.Query(sqlstrListTypePlace)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Place: ListType: get type place list query err: %s", err)
		return nil, err
	}

	for rows.Next() {
		tp := new(typeModel)

		if err := rows.Scan(&tp.id, &tp.name); err != nil {
			log.Errorf("Storage: Place: ListType: get type place list scan rows err: %s", err)
			return nil, err
		}

		c := tp.convert()
		places[c.ID] = c
	}

	return places, nil
}

func (s *PlaceStorage) Update(ctx context.Context, place *types.Place) error {

	log.Debugf("Storage: Place: Update: update place: %#v", place)

	if place == nil {
		err := errors.New("place can not be nil")
		log.Errorf("Storage: Place: Update: update place err: %s", err)
		return err
	}

	place.Meta.Updated = time.Now()

	err := s.client.QueryRow(sqlstrPlaceUpdate, place.Meta.Phone, place.Meta.Adress,
		place.Meta.City, place.Meta.Url, place.Meta.Name).Scan(&place.Meta.Updated)
	if err != nil {
		log.Errorf("Storage: Place: Update: update place query err: %s", err)
		return err
	}
	return nil
}

func (nm *typeModel) convert() *types.TypePlaces {
	c := new(types.TypePlaces)

	c.ID = nm.id.String
	c.NameType = nm.name.String
	return c
}

func (pl *placeModel) convert() *types.Place {
	c := new(types.Place)

	c.Meta.Name = pl.name.String
	c.Meta.Phone = pl.phone.String
	c.Meta.Adress = pl.adress.String
	c.Meta.City = pl.city.String
	c.Meta.Url = pl.url.String
	c.Meta.ID = pl.id.String

	return c
}
