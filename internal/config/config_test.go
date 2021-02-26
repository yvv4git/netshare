//+build integration

package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfigYml(t *testing.T) {
	t.Log("Testing yml config reading")

	file, err := os.Open("../../.netshare.yml")
	assert.Nil(t, err)

	viper.SetConfigType("yaml")
	err = viper.ReadConfig(file)
	assert.Nil(t, err)

	var config Config
	err = viper.Unmarshal(&config)
	assert.Nil(t, err)

	assert.Equal(t, "localhost", config.Host)
	assert.Equal(t, 8080, config.Port)
	assert.Equal(t, "data", config.ShareDir)
	assert.Equal(t, "web", config.Type)
}
