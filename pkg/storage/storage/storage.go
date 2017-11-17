package storage

import (
			"context"
			"github.com/orderfood/api_of/pkg/common/types"
)

type User interface {
	GetUserByID(ctx context.Context, id string) (*types.User, error)
	CreateUser(ctx context.Context, user *types.User) error
	CheckExistsByLogin(ctx context.Context, login string) (bool, error)
}
