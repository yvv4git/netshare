package server

import (
	"errors"
	"netshare/internal/config"
)

// Factory - factory for different types of servers
func Factory(config config.Config) (IServer, error) {
	if config.Type == "web" {
		return newWebServer(config.Host, config.Port, config.ShareDir), nil
	}

	return nil, errors.New("Wrong server type")
}
