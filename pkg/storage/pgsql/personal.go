package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/storage/store"
	"log"
)

func (s *PersonalStorage) GetTypePersonIDByName(ctx context.Context, name string) (string, error) {
	var (
		err  error
		pers = new(idModel)
	)

	const ssqlTypePersonalIDGetByName = `SELECT type_personal.id_typePersonal
		FROM type_personal
		WHERE type_personal.name_type = $1;`

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

	const sqlCreatePerson = `
		INSERT INTO personal (fio, phone, id_place, id_typePersonal)
		VALUES ($1, $2, $3, $4)
		RETURNING id_personal;
	`

	err = s.client.QueryRow(sqlCreatePerson, personal.Meta.Fio, personal.Meta.Phone, personal.Meta.PlaceID, personal.Meta.TypePersonalID).Scan(&id)

	personal.Meta.ID = id.String

	return err
}

func (s *PersonalStorage) ListType(ctx context.Context) (map[string]*types.TypePersonals, error) {

	personals := make(map[string]*types.TypePersonals)

	const sqlstrListTypePersonal = `
		SELECT type_personal.id_typePersonal, type_personal.name_type
		FROM type_personal;`

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

	const sqlstrListPersonal = `
		SELECT personal.id_personal, personal.fio, personal.phone, personal.updated, personal.created
		FROM personal
		WHERE personal.id_place = $1;`

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
