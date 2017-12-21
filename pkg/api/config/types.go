package config

import "github.com/orderfood/api_of/pkg/common/config"

type Config struct {
	APIServer config.APIServer
	Database struct {
		Connection *string
	}
}
