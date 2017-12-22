package pgsql

import (
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/storage/store"
	"log"
	"errors"
	"database/sql"
	"time"
)

func (s *PlaceStorage) GetPlaceByIDUser(ctx context.Context, id string) (*types.Place, error) {

	var (
		err error
		pl  = new(placeModel)
	)

	err = s.client.QueryRow(sqlPlaceGetByIDUsr, id).Scan(&pl.name, &pl.phone, &pl.adress,
		&pl.city, &pl.url)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}

	plc := pl.convert()

	return plc, nil

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

func (s *PlaceStorage) GetTypePlaceByName(ctx context.Context, name string) (string, error) {
	var (
		err error
		pl  = new(typeplaceModel)
	)

	err = s.client.QueryRow(sqlTypePlaceIDGetByName, name).Scan(&pl.id)

	switch err {
	case nil:
	case sql.ErrNoRows:
		return "", nil
	default:
		return "", err
	}

	typeplaceID := pl.id.String

	return typeplaceID, nil
}

func (s *PlaceStorage) List(ctx context.Context) (map[string]*types.TypePlaces, error) {

	places := make(map[string]*types.TypePlaces)

	rows, err := s.client.Query(sqlstrListType)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:

		return nil, err
	}

	for rows.Next() {
		tp := new(typeplaceModel)

		if err := rows.Scan(&tp.id, &tp.name); err != nil {
			return nil, err
		}

		c := tp.convert()
		places[c.ID] = c
	}

	return places, nil
}

func (s *PlaceStorage) Update(ctx context.Context, place *types.Place) error{

	if place == nil {
		err := errors.New("place can not be nil")
		return err
	}

	place.Meta.Updated = time.Now()

	if _, err := s.client.Exec(sqlstrPlaceUpdate, place.Meta.Phone, place.Meta.Adress,
		place.Meta.City, place.Meta.Url, place.Meta.Name); err != nil {
		return err
	}
	return nil
}


func (nm *typeplaceModel) convert() *types.TypePlaces {
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

	return c
}