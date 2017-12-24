package storage

import (
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
)

type User interface {
	GetByLogin(ctx context.Context, login string) (*types.User, error)
	GetUserByID(ctx context.Context, id string) (*types.User, error)
	CreateUser(ctx context.Context, user *types.User) error
	CheckExistsByLogin(ctx context.Context, login string) (bool, error)
}

type Place interface {
	CreatePlace(ctx context.Context, place *types.Place) error
	GetTypePlaceByName(ctx context.Context, name string) (string, error)
	ListType(ctx context.Context) (map[string]*types.TypePlaces, error)
	GetPlaceByIDUser(ctx context.Context, id string) (*types.Place, error)
	GetPlaceIDByUsrid(ctx context.Context, id string) (string, error)
	Update(ctx context.Context, place *types.Place) error
}

type Menu interface {
	CreateMenu(ctx context.Context, menu *types.Menu) error
	InsertDishInMenu(ctx context.Context, menuid, dishid string) error
	DeleteDishInMenu(ctx context.Context, menuid, dishid string) error
	GetIDmenuByName(ctx context.Context, name string) (string, error)
	Fetch(ctx context.Context, idplace, name string) (*types.Menu, error)
	List(ctx context.Context, placeid string) (map[string]*types.Menu, error)
	ListDishesInMenu(ctx context.Context, menuid, usrid string) (map[string]*types.Dish, error)
	ListDishesNotMenu(ctx context.Context, menuid, userid string) (map[string]*types.Dish, error)
}

type Dish interface {
	CreateDish(ctx context.Context, dich *types.Dish) error
	RemoveDish(ctx context.Context, id string) error
	Fetch(ctx context.Context, usrid, name string) (*types.Dish, error)
	GetIdDishByName(ctx context.Context, name, usrid string) (string, error)
	List(ctx context.Context, userid string) (map[string]*types.Dish, error)
	TypeList(ctx context.Context) (map[string]*types.TypeDishes, error)
	GetTypeDishIDByName(ctx context.Context, name string) (string, error)
	Update(ctx context.Context, usrid string ,dish *types.Dish) error
}

type Personal interface {
	CreatePerson(ctx context.Context, personal *types.Personal) error
	//Remove(ctx context.Context, id string) error
	GetTypePersonIDByName(ctx context.Context, name string) (string, error)
	ListType(ctx context.Context) (map[string]*types.TypePersonals, error)
	List(ctx context.Context, placeid string) (map[string]*types.Personal, error)
}
