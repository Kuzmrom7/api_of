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
	List(ctx context.Context) (map[string]*types.TypePlaces, error)
	GetPlaceByIDUser(ctx context.Context, id string) (*types.Place, error)
	Update(ctx context.Context, place *types.Place) error
}

type Menu interface {
	CreateMenu(ctx context.Context, menu *types.Menu) error
	CreateMenuDish(ctx context.Context, menuid, dishid string) error
	GetPlaceByName(ctx context.Context, name string) (string, error)
	GetIDmenuByName(ctx context.Context, name string) (string, error)
	Fetch(ctx context.Context, idplace, name string) (*types.Menu, error)
	List(ctx context.Context, placeid string) (map[string]*types.Menu, error)
	ListMenuDish(ctx context.Context, menuid, typedishid string) (map[string]*types.Dish, error)
}

type Dish interface {
	CreateDich(ctx context.Context, dich *types.Dish) error
	Remove(ctx context.Context, id string) error
	GetIDdichByName(ctx context.Context, name, usrid string) (string, error)
	List(ctx context.Context, userid string) (map[string]*types.Dish, error)
	TypeList(ctx context.Context) (map[string]*types.TypeDishes, error)
	GetTypeDishIDByName(ctx context.Context, name string) (string, error)
}

type Personal interface {
	CreatePerson(ctx context.Context, personal *types.Personal) error
	//Remove(ctx context.Context, id string) error
	GetTypePersonIDByName(ctx context.Context, name string) (string, error)
	GetPlaceIDByUsrid(ctx context.Context, id string) (string, error)
	ListType(ctx context.Context) (map[string]*types.TypePersonals, error)
	List(ctx context.Context, placeid string) (map[string]*types.Personal, error)
}
