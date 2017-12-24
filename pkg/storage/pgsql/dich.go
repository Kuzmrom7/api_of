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
	"time"
)

func (nm *dichModel) convert() *types.Dish {
	c := new(types.Dish)

	c.Meta.ID = nm.id.String
	c.Meta.Name = nm.name.String
	c.Meta.Desc = nm.description.String
	c.Meta.Url = nm.url.String
	c.Meta.Updated = nm.updated
	c.Meta.TypeDishID = nm.id_Type.String
	c.Meta.Created = nm.created
	c.Meta.Timemin = nm.timemin.Int64

	return c
}

func (s *DishStorage) CreateDish(ctx context.Context, dich *types.Dish) error {

	log.Println("STORAGE--- CreateDish()")

	if dich == nil {
		err := errors.New("dich can not be nil")
		return err
	}

	var (
		err error
		id  store.NullString
	)

	err = s.client.QueryRow(sqlCreateDich, dich.Meta.Name, dich.Meta.Desc, dich.Meta.Timemin, dich.Meta.TypeDishID, dich.Meta.Url, dich.Meta.UserID).Scan(&id)

	dich.Meta.ID = id.String

	return err
}

func (s *DishStorage) GetIdDishByName(ctx context.Context, name, usrid string) (string, error) {
	var (
		err error
		di  = new(dichModel)
	)

	err = s.client.QueryRow(sqlDichIDGetByName, name, usrid).Scan(&di.id)

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

func (s *DishStorage) RemoveDish(ctx context.Context, id string) error {

	_, err := s.client.Exec(sqlDichRemove, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *DishStorage) Update(ctx context.Context, usrid string, dish *types.Dish) error {

	if dish == nil {
		err := errors.New("dish can not be nil")
		return err
	}
	if usrid == "" {
		err := errors.New("usrid can not be nil")
		return err
	}

	dish.Meta.Updated = time.Now()

	err := s.client.QueryRow(sqlstrDishUpdate, dish.Meta.Timemin, dish.Meta.Desc,
		dish.Meta.Name, usrid).Scan(&dish.Meta.Updated)
	if err != nil {
		return err
	}
	return nil
}

func (s *DishStorage) Fetch(ctx context.Context, usrid, name string) (*types.Dish, error) {

	var (
		err error
		di  = new(dichModel)
	)

	err = s.client.QueryRow(sqlFetchDish, usrid, name).Scan(&di.id, &di.name, &di.description, &di.url, &di.updated, &di.created, &di.timemin)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}

	dish := di.convert()

	return dish, nil

}

func (s *DishStorage) List(ctx context.Context, userid string) (map[string]*types.Dish, error) {

	dishes := make(map[string]*types.Dish)

	rows, err := s.client.Query(sqlstrListDish, userid)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:

		return nil, err
	}
	log.Println(userid)

	for rows.Next() {

		di := new(dichModel)

		if err := rows.Scan(&di.id, &di.name, &di.description, &di.url, &di.updated, &di.created, &di.timemin); err != nil {

			return nil, err
		}

		c := di.convert()
		dishes[c.Meta.ID] = c
	}

	return dishes, nil
}

func (s *DishStorage) TypeList(ctx context.Context) (map[string]*types.TypeDishes, error) {

	tydishes := make(map[string]*types.TypeDishes)

	rows, err := s.client.Query(sqlstrListTypeDish)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:

		return nil, err
	}
	for rows.Next() {
		tp := new(typeModelDishes)

		if err := rows.Scan(&tp.id, &tp.name); err != nil {
			return nil, err
		}

		c := tp.convert()
		tydishes[c.ID] = c
	}

	return tydishes, nil
}

func (s *DishStorage) GetTypeDishIDByName(ctx context.Context, name string) (string, error) {
	var (
		err  error
		dish = new(idModel)
	)

	err = s.client.QueryRow(ssqlTypeDishlIDGetByName, name).Scan(&dish.id)

	switch err {
	case nil:
	case sql.ErrNoRows:
		return "", nil
	default:
		return "", err
	}

	dishID := dish.id.String

	return dishID, nil
}

func (nm *typeModelDishes) convert() *types.TypeDishes {
	c := new(types.TypeDishes)

	c.ID = nm.id.String
	c.NameType = nm.name.String
	return c
}
