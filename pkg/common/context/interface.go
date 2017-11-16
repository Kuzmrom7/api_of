package context

import (
	"context"
	"github.com/orderfood/api_of/pkg/storage"
)

type Context interface {
	context.Context

	GetStorage() storage.Storage
	Background() context.Context
}