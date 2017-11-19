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

type TypePlace interface {

}

type Place interface {

}

type Menu interface {

}

type Sections interface {

}

type TypesDishes interface {

}

type Dish interface {


}

type Personal interface {


}

type TypePersonal interface {

}