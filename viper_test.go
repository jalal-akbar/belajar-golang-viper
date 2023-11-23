package main

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	assrt := assert.New(t)
	var config = viper.New()
	assrt.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	assert.Nil(t, err)
	assert.Equal(t, "belajar-golang-viper", config.GetString("app.name"))
	assert.Equal(t, true, config.GetBool("database.show_sql"))
	assert.Equal(t, "localhost", config.GetString("database.host"))
	assert.Equal(t, 3306, config.GetInt("database.port"))
}

func TestYAML(t *testing.T) {
	var asrt = assert.New(t)
	var config = viper.New()

	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	asrt.Nil(err)
	asrt.Equal("belajar-golang-viper", config.GetString("app.name"))
	asrt.Equal("1.0.0", config.GetString("app.version"))
	asrt.Equal("Jalaluddin Muh Akbar", config.GetString("app.author"))
	asrt.True(config.GetBool("database.show_sql"))
	asrt.Equal("localhost", config.GetString("database.host"))
	asrt.Equal(3306, config.GetInt("database.port"))
}

func TestENV(t *testing.T) {
	var assrt = assert.New(t)
	var config = viper.New()

	config.SetConfigFile("config.env")
	config.AddConfigPath(".")

	err := config.ReadInConfig()

	assrt.Nil(err)
	assrt.Equal("belajar-golang-viper", config.GetString("APP_NAME"))
}
