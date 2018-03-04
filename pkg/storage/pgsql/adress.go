package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/store"
	"database/sql"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/log"
	"context"
	"errors"
)

func (s *AdressStorage) CreateAdress(ctx context.Context, adress *types.Adress) error {

	log.Debug("Storage: Adress: Insert: insert adress: %#v", adress)

	if adress == nil {
		err := errors.New("place can not be nil")
		log.Errorf("Storage: Adress: Insert: insert adress err: %s", err)
		return err
	}

	var (
		err error
		id  store.NullString
	)

	const sqlCreateAdress = `
		INSERT INTO adressing (name_adress, id_place)
		VALUES ($1, $2)
		RETURNING id_adressing;
	`

	err = s.client.QueryRow(sqlCreateAdress, adress.Meta.Name, adress.Meta.PlaceID).Scan(&id)
	if err != nil {
		log.Errorf("Storage: Adress: Insert: insert adress query err: %s", err)
		return err
	}

	adress.Meta.ID = id.String

	return err
}

func (s *AdressStorage) List(ctx context.Context, placeid string) (map[string]*types.Adress, error) {

	adresses := make(map[string]*types.Adress)

	log.Debug("Storage: Adress: List: get list adress")

	const sqlstrListAdress = `
					SELECT adressing.id_adressing, adressing.name_adress, adressing.id_place
					FROM adressing
					WHERE adressing.id_place = $1;`

	rows, err := s.client.Query(sqlstrListAdress, placeid)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Adress: List: get list adress query err: %s", err)
		return nil, err
	}

	for rows.Next() {

		di := new(adressModel)

		if err := rows.Scan(&di.id, &di.name, &di.idplace); err != nil {
			log.Errorf("Storage: Adress: List: get list adress scan rows err: %s", err)
			return nil, err
		}

		c := di.convert()
		adresses[c.Meta.ID] = c
	}

	return adresses, nil
}

func (nm *adressModel) convert() *types.Adress {
	c := new(types.Adress)

	c.Meta.ID = nm.id.String
	c.Meta.Name = nm.name.String
	c.Meta.PlaceID = nm.idplace.String

	return c
}
