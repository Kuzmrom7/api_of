package storage

import "github.com/orderfood/api_of/pkg/storage/storage"

type Storage interface {
	User() storage.User
}