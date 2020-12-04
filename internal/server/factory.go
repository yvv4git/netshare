package server

import "errors"

// Factory - factory for different types of servers
func Factory(serverType string, host string, port int, share string) (NetServerer, error) {
	if serverType == "web" {
		return newWebServer(host, port, share), nil
	}

	return nil, errors.New("Wrong server type")
}
