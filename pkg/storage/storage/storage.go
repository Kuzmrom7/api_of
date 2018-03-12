package storage

import (
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
)

type User interface {
	GetByLogin(ctx context.Context, login string) (*types.User, error)
	GetById(ctx context.Context, id string) (*types.User, error)
	CreateUser(ctx context.Context, user *types.User) error
	CheckExistsByLogin(ctx context.Context, login string) (bool, error)
}

type Place interface {
	CreatePlace(ctx context.Context, place *types.Place) error
	ListType(ctx context.Context) (map[string]*types.TypePlaces, error)
	List(ctx context.Context) ([]*types.Place, error)
	GetPlaceByIDUser(ctx context.Context, id string) (*types.Place, error)
	GetPlaceByID(ctx context.Context, id string) (*types.Place, error)
	Update(ctx context.Context, place *types.Place) error
}

type Menu interface {
	CreateMenu(ctx context.Context, menu *types.Menu) error
	InsertDishInMenu(ctx context.Context, menuid, dishid string) error
	CheckUnique(ctx context.Context, menuid, dishid string) (bool, error)
	DeleteDishInMenu(ctx context.Context, menuid, dishid string) error
	Fetch(ctx context.Context, id string) (*types.Menu, error)
	List(ctx context.Context, placeid string) (map[string]*types.Menu, error)
	ListDishesInMenu(ctx context.Context, menuid, placeid string) ([]*types.Dish, error)
	ListDishesNotMenu(ctx context.Context, menuid, placeid string) ([]*types.Dish, error)
}

type Dish interface {
	CreateDish(ctx context.Context, dich *types.Dish) error
	RemoveDish(ctx context.Context, id string) error
	Fetch(ctx context.Context, id string) (*types.Dish, error)
	List(ctx context.Context, placeid string) ([]*types.Dish, error)
	TypeList(ctx context.Context) (map[string]*types.TypeDishes, error)
	Update(ctx context.Context, dish *types.Dish) error
}

type Personal interface {
	CreatePerson(ctx context.Context, personal *types.Personal) error
	GetTypePersonIDByName(ctx context.Context, name string) (string, error)
	ListType(ctx context.Context) (map[string]*types.TypePersonals, error)
	List(ctx context.Context, placeid string) (map[string]*types.Personal, error)
}

type Adress interface {
	CreateAdress(ctx context.Context, adress *types.Adress) error
	List(ctx context.Context, placeid string) (map[string]*types.Adress, error)
}
