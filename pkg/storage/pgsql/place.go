package pgsql

import (
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/storage/store"
	"github.com/orderfood/api_of/pkg/log"
	"errors"
	"database/sql"
	"time"
	"encoding/json"
)

func (s *PlaceStorage) GetPlaceByIDUser(ctx context.Context, id string) (*types.Place, error) {

	var (
		err error
	)

	log.Debugf("Storage: Place: Get: get place by user id: %s ", id)

	const sqlPlaceGetByUserID = `
			SELECT to_json(
				json_build_object(
					'meta', json_build_object(
					'id', id_place,
					'name', name,
					'phone', phone_number,
					'url', url,
					'city', city
				),
				'typesplace', type,
				'adresses', adress
				)
			)
			FROM place
			WHERE place.user_id = $1;`

	var buf string

	err = s.client.QueryRow(sqlPlaceGetByUserID, id).Scan(&buf)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Place: Get: get place by id query err: %s", err)
		return nil, err
	}

	plc := new(types.Place)

	if err := json.Unmarshal([]byte(buf), &plc); err != nil {
		return nil, err
	}

	return plc, nil
}

func (s *PlaceStorage) GetPlaceByID(ctx context.Context, id string) (*types.Place, error) {

	var (
		err error
	)

	log.Debugf("Storage: Place: Get: get place by id: %s ", id)

	const sqlPlaceGetByID = `
			SELECT to_json(
				json_build_object(
					'meta', json_build_object(
					'id', id_place,
					'name', name,
					'phone', phone_number,
					'url', url,
					'city', city
				),
				'typesplace', type,
				'adresses', adress
				)
			)
			FROM place
			WHERE place.id_place = $1;`

	var buf string

	err = s.client.QueryRow(sqlPlaceGetByID, id).Scan(&buf)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Place: Get: get place by id query err: %s", err)
		return nil, err
	}

	plc := new(types.Place)

	if err := json.Unmarshal([]byte(buf), &plc); err != nil {
		return nil, err
	}

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

	typesplace, err := json.Marshal(place.TypesPlace)
	if err != nil {
		log.Errorf("Storage: Place: prepare types struct to database write: %s", err)
		typesplace = []byte("{}")
	}

	const sqlCreatePlace = `
		INSERT INTO place (name, user_id, type)
		VALUES ($1, $2, $3)
		RETURNING id_place;`

	err = s.client.QueryRow(sqlCreatePlace, place.Meta.Name, place.Meta.UserID, string(typesplace)).Scan(&id)
	if err != nil {
		log.Errorf("Storage: Place: Insert: insert place query err: %s", err)
		return err
	}
	place.Meta.ID = id.String

	return err
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

	adress, err := json.Marshal(place.Adresses)
	if err != nil {
		log.Errorf("Storage: Place: prepare types struct to database write: %s", err)
		adress = []byte("{}")
	}

	place.Meta.Updated = time.Now()

	err = s.client.QueryRow(sqlstrPlaceUpdate, place.Meta.Phone, string(adress),
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
	c.Meta.City = pl.city.String
	c.Meta.Url = pl.url.String
	c.Meta.ID = pl.id.String

	return c
}
