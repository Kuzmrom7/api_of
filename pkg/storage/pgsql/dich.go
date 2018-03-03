package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/store"
	"github.com/orderfood/api_of/pkg/log"
	"github.com/orderfood/api_of/pkg/common/types"
	"context"
	"errors"
	"database/sql"
	"time"
	"encoding/json"
)

func (nm *dichModel) convert() *types.Dish {
	c := new(types.Dish)

	c.Meta.ID = nm.id.String
	c.Meta.Name = nm.name.String
	c.Meta.Desc = nm.description.String
	c.Meta.Updated = nm.updated
	c.Meta.TypeDishID = nm.id_Type.String
	c.Meta.Created = nm.created
	c.Meta.Timemin = nm.timemin.Int64

	return c
}

func (s *DishStorage) CreateDish(ctx context.Context, dish *types.Dish) error {

	log.Debug("Storage: Dish: Insert: insert dish: %#v", dish)

	if dish == nil {
		err := errors.New("dish can not be nil")
		log.Errorf("Storage: Dish: Insert: insert dish err: %s", err)
		return err
	}

	var (
		err error
		id  store.NullString
	)

	urls, err := json.Marshal(dish.Urls)
	if err != nil {
		log.Errorf("Storage: Dish: prepare types struct to database write: %s", err)
		urls = []byte("{}")
	}

	specs, err := json.Marshal(dish.Specs)
	if err != nil {
		log.Errorf("Storage: Dish: prepare types struct to database write: %s", err)
		specs = []byte("{}")
	}

	const sqlCreateDich = `
		INSERT INTO dish (name_dish, description, time_min, id_typeDish, url, spec, user_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id_dish;
	`

	err = s.client.QueryRow(sqlCreateDich, dish.Meta.Name, dish.Meta.Desc, dish.Meta.Timemin, dish.Meta.TypeDishID, string(urls), string(specs), dish.Meta.UserID).Scan(&id)
	if err != nil {
		log.Errorf("Storage: Dish: Insert: insert dish query err: %s", err)
		return err
	}

	dish.Meta.ID = id.String

	return err
}

func (s *DishStorage) RemoveDish(ctx context.Context, id string) error {

	const sqlDichRemove = `DELETE FROM dish WHERE id_dish = $1;`

	log.Debugf("Storage: Dish: Delete: delete dish by id %s", id)

	if id == "" {
		err := errors.New("id can not be nil")
		log.Errorf("Storage: Dish: Delete: delete dish err: %s", err)
		return err
	}

	_, err := s.client.Exec(sqlDichRemove, id)
	if err != nil {
		log.Errorf("Storage: Dish: Delete: delete dish exec err: %s", err)
		return err
	}
	return nil
}

func (s *DishStorage) Update(ctx context.Context, dish *types.Dish) error {

	log.Debugf("Storage: Dish: Update: update dish: %#v", dish)

	if dish == nil {
		err := errors.New("dish can not be nil")
		log.Errorf("Storage: Dish: Update: update dish err: %s", err)
		return err
	}

	dish.Meta.Updated = time.Now()

	const sqlstrDishUpdate = `
		UPDATE dish
		SET
			time_min = $1,
			description = $2,
			updated = now()
		WHERE id_dish = $3
		RETURNING updated;`

	err := s.client.QueryRow(sqlstrDishUpdate, dish.Meta.Timemin, dish.Meta.Desc,
		dish.Meta.ID).Scan(&dish.Meta.Updated)
	if err != nil {
		log.Errorf("Storage: Dish: Update: update dish query err: %s", err)
		return err
	}
	return nil
}

func (s *DishStorage) Fetch(ctx context.Context, id string) (*types.Dish, error) {

	var (
		err error
	)

	log.Debugf("Storage: Dish: Get: get dish by id: %s ", id)

	const sqlFetchDish = `
			SELECT to_json(
				json_build_object(
					'meta', json_build_object(
					'id', id_dish,
					'name', name_dish,
					'description', description,
					'timemin', time_min
				),
				'urls', url,
				'specs', spec
				)
			)
			FROM dish
			WHERE dish.id_dish = $1;`

	var buf string

	err = s.client.QueryRow(sqlFetchDish, id).Scan(&buf)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Dish: Get: get dish by id query err: %s", err)
		return nil, err
	}

	di := new(types.Dish)

	if err := json.Unmarshal([]byte(buf), &di); err != nil {
		return nil, err
	}

	return di, nil

}

func (s *DishStorage) List(ctx context.Context, userid string) ([]*types.Dish, error) {

	var dishes []*types.Dish

	log.Debug("Storage: Dish: List: get list dishes by user id %s", userid)

	const sqlstrListDish = `
			SELECT to_json(
				json_build_object(
					'meta', json_build_object(
					'id', id_dish,
					'name', name_dish,
					'description', description,
					'timemin', time_min
				),
				'urls', url,
				'specs', spec
				)
			)
			FROM dish
			WHERE dish.user_id = $1;`

	//const sqlstrListDish = `
	//				SELECT dish.id_dish, dish.name_dish, dish.description, dish.url, dish.updated, dish.created, dish.time_min
	//				FROM dish
	//				WHERE dish.user_id = $1;`

	rows, err := s.client.Query(sqlstrListDish, userid)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Dish: List: get list dishes query err: %s", err)
		return nil, err
	}

	for rows.Next() {

		var buf string

		if err := rows.Scan(&buf); err != nil {
			log.Errorf("Storage: Dish: List: get list dishes scan rows err: %s", err)
			return nil, err
		}

		di := new(types.Dish)

		if err := json.Unmarshal([]byte(buf), &di); err != nil {
			return nil, err
		}

		dishes = append(dishes, di)
	}

	return dishes, nil
}

func (s *DishStorage) TypeList(ctx context.Context) (map[string]*types.TypeDishes, error) {

	tydishes := make(map[string]*types.TypeDishes)

	log.Debug("Storage: Dish: ListType: get type dish list")

	const sqlstrListTypeDish = `
		SELECT type_dish.id_typeDish, type_dish.name_typeDish
		FROM type_dish;`

	rows, err := s.client.Query(sqlstrListTypeDish)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Errorf("Storage: Dish: ListType: get type dish list query err: %s", err)
		return nil, err
	}
	for rows.Next() {
		tp := new(typeModelDishes)

		if err := rows.Scan(&tp.id, &tp.name); err != nil {
			log.Errorf("Storage: Dish: ListType: get type dish list scan rows err: %s", err)
			return nil, err
		}

		c := tp.convert()
		tydishes[c.ID] = c
	}

	return tydishes, nil
}

func (nm *typeModelDishes) convert() *types.TypeDishes {
	c := new(types.TypeDishes)

	c.ID = nm.id.String
	c.NameType = nm.name.String
	return c
}
