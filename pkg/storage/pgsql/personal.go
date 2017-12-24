package pgsql

import (
	"database/sql"
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
	"errors"
	"log"
	"github.com/orderfood/api_of/pkg/storage/store"
)


func (s *PersonalStorage) GetTypePersonIDByName(ctx context.Context, name string) (string, error) {
	var (
		err error
		pers = new(idModel)
	)

	err = s.client.QueryRow(ssqlTypePersonalIDGetByName, name).Scan(&pers.id)

	switch err {
	case nil:
	case sql.ErrNoRows:
		return "", nil
	default:
		return "", err
	}

	placeID := pers.id.String

	return placeID, nil
}

func (s *PersonalStorage) CreatePerson(ctx context.Context, personal *types.Personal) error {

	log.Println("STORAGE--- CreatePersonal()")

	if personal == nil {
		err := errors.New("personal can not be nil")
		return err
	}

	var (
		err error
		id  store.NullString
	)

	err = s.client.QueryRow(sqlCreatePerson, personal.Meta.Fio, personal.Meta.Phone, personal.Meta.PlaceID, personal.Meta.TypePersonalID).Scan(&id)

	personal.Meta.ID = id.String

	return err
}

func (s *PersonalStorage) ListType(ctx context.Context) (map[string]*types.TypePersonals, error) {

	personals := make(map[string]*types.TypePersonals)

	rows, err := s.client.Query(sqlstrListTypePersonal)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:

		return nil, err
	}

	for rows.Next() {
		tp := new(typeModelPersonals)

		if err := rows.Scan(&tp.id, &tp.name); err != nil {
			return nil, err
		}

		c := tp.convert()
		personals[c.ID] = c
	}

	return personals, nil
}

func (s *PersonalStorage) List(ctx context.Context, placeid string) (map[string]*types.Personal, error) {

	personals := make(map[string]*types.Personal)

	rows, err := s.client.Query(sqlstrListPersonal, placeid)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:

		return nil, err
	}


	for rows.Next() {

		di := new(personalModel)

		if err := rows.Scan(&di.id, &di.fio, &di.phone, &di.updated, &di.created); err != nil {

			return nil, err
		}

		c := di.convert()
		personals[c.Meta.ID] = c
	}

	return personals, nil
}


func (nm *typeModelPersonals) convert() *types.TypePersonals {
	c := new(types.TypePersonals)

	c.ID = nm.id.String
	c.NameType = nm.name.String
	return c
}

func (nm *personalModel) convert() *types.Personal {
	c := new(types.Personal)

	c.Meta.ID = nm.id.String
	c.Meta.Fio = nm.fio.String
	c.Meta.Phone = nm.phone.String
	c.Meta.Updated = nm.updated
	c.Meta.Created = nm.created

	return c
}