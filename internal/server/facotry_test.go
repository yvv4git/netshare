package server

import (
	"fmt"
	"netshare/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactory(t *testing.T) {
	t.Log("Testing - factory")

	config := config.Config{
		NetServer: config.NetServer{
			Host:     "localhost",
			Port:     8183,
			ShareDir: "data",
			Type:     "web",
		},
	}

	webServer, err := Factory(config)

	assert.Nil(t, err)

	if _, ok := webServer.(*web); ok {
		fmt.Println("This is web server")
	}
}
