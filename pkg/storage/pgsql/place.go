package pgsql

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/log"
	"github.com/orderfood/api_of/pkg/storage/store"
	"time"
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
				'typesplace', type
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
				'typesplace', type
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

	const sqlstrListTypePlace = `
		SELECT type_place.id_typePlace, type_place.name_type
		FROM type_place;`

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

func (s *PlaceStorage) List(ctx context.Context) ([]*types.Place, error) {

	var places []*types.Place

	log.Debug("Storage: Place: List: get list places")

	const sqlstrListPlace = `
			SELECT to_json(
				json_build_object(
					'meta', json_build_object(
					'id', id_place,
					'name', name,
					'phone', phone_number,
					'url', url,
					'city', city
				),
				'typesplace', type
				)
			)
			FROM place;`

	rows, err := s.client.Query(sqlstrListPlace)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Place: List: get list places query err: %s", err)
		return nil, err
	}

	for rows.Next() {

		var buf string

		if err := rows.Scan(&buf); err != nil {
			log.Errorf("Storage: Place: List: get list places scan rows err: %s", err)
			return nil, err
		}

		pl := new(types.Place)

		if err := json.Unmarshal([]byte(buf), &pl); err != nil {
			return nil, err
		}

		places = append(places, pl)
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

	const sqlstrPlaceUpdate = `
		UPDATE place
		SET
			phone_number = $1,
			city = $2,
			url = $3,
			updated = now()
		WHERE id_place = $4
		RETURNING updated;`

	err := s.client.QueryRow(sqlstrPlaceUpdate, place.Meta.Phone,
		place.Meta.City, place.Meta.Url, place.Meta.ID).Scan(&place.Meta.Updated)
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
