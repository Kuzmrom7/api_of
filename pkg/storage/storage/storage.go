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
	Update(ctx context.Context, place *types.Place, name string)  error
}

type Menu interface {
	CreateMenu(ctx context.Context, menu *types.Menu) error
	GetPlaceByName(ctx context.Context, name string) (string, error)
}

type Dich interface {
	CreateDich(ctx context.Context, dich *types.Dich) error
	Remove(ctx context.Context, id string) error
	GetIDdichByName(ctx context.Context, name string) (string, error)
	List(ctx context.Context) (map[string]*types.Dich, error)
}

type Personal interface {
	CreatePerson(ctx context.Context, personal *types.Personal) error
	//Remove(ctx context.Context, id string) error
	GetTypePersonIDByName(ctx context.Context, name string) (string, error)
	GetPlaceIDByUsrid(ctx context.Context, id string) (string, error)
	//List(ctx context.Context) (map[string]*types.Dich, error)

}
