package pgsql

import (
	//"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"
	"github.com/orderfood/api_of/pkg/common/types"
	"context"
	"log"
	"errors"
	"database/sql"
	//"encoding/json"
)

func (nm *dichModel) convert() *types.Dich {
	c := new(types.Dich)

	c.Meta.ID = nm.id.String
	c.Meta.Name = nm.name.String
	c.Meta.Desc = nm.description.String

	return c
}

func (s *DichStorage) CreateDich(ctx context.Context, dich *types.Dich) error {

	log.Println("STORAGE--- CreateDich()")

	if dich == nil {
		err := errors.New("dich can not be nil")
		return err
	}

	var (
		err error
		id  store.NullString
	)

	err = s.client.QueryRow(sqlCreateDich, dich.Meta.Name, dich.Meta.Desc, dich.Meta.Timemin).Scan(&id)

	dich.Meta.ID = id.String

	return err
}

func (s *DichStorage) GetIDdichByName(ctx context.Context, name string) (string, error) {
	var (
		err error
		di  = new(dichModel)
	)

	err = s.client.QueryRow(sqlDichIDGetByName, name).Scan(&di.id)

	switch err {
	case nil:
	case sql.ErrNoRows:
		return "", nil
	default:
		return "", err
	}

	dishID := di.id.String

	return dishID, nil
}

func (s *DichStorage) Remove(ctx context.Context, id string) error {

	_, err := s.client.Exec(sqlDichRemove, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *DichStorage) List(ctx context.Context) (map[string]*types.Dich, error) {

	dishes := make(map[string]*types.Dich)

	rows, err := s.client.Query(sqlstrListDish)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:

		return nil, err
	}

	for rows.Next() {
		di := new(dichModel)

		if err := rows.Scan(&di.id, di.name, di.description); err != nil {
			return nil, err
		}

		c := di.convert()
		dishes[c.Meta.ID] = c
	}

	return dishes, nil
}


