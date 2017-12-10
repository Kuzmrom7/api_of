package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/storage"
	"github.com/orderfood/api_of/pkg/storage/store"
	"github.com/orderfood/api_of/pkg/common/types"
	"context"
	"log"
	"errors"
	"database/sql"
)

const (
	sqlCreateDich = `
		INSERT INTO dish (name_dish, description, time_min)
		VALUES ($1, $2, $3)
		RETURNING id_dish;
	`

	sqlDichIDGetByName = `SELECT dish.id_dish
		FROM dish
		WHERE dish.name_dish = $1;`

	sqlDichRemove = `DELETE FROM dish WHERE id_dish = $1;`
)

type DichStorage struct {
	storage.Dich
	client store.IDB
}

type dichModel struct {
	id store.NullString
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

func newDichStorage(client store.IDB) *DichStorage {
	s := new(DichStorage)
	s.client = client
	return s
}
