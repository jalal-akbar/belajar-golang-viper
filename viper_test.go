package main

import (
	"os"
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

func TestAutomaticENV(t *testing.T) {
	var assrt = assert.New(t)
	var config = viper.New()
	config.SetEnvPrefix("from") // will be uppercased automatically
	config.BindEnv("env")       // will be uppercased automatically

	os.Setenv("FROM_ENV", "Hello")

	config.SetConfigFile("./config.env")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	err := config.ReadInConfig()

	assrt.Nil(err)
	assrt.Equal("belajar-golang-viper", config.GetString("APP_NAME"))
	assrt.Equal("1.0.0", config.GetString("APP_VERSION"))
	assrt.Equal("Jalaluddin Muh Akbar", config.GetString("APP_AUTHOR"))
	assrt.Equal(true, config.GetBool("DATABASE_SHOW_SQL"))
	assrt.Equal("localhost", config.GetString("DATABASE_HOST"))
	assrt.Equal(3306, config.GetInt("DATABASE_PORT"))

	assrt.Equal("Hello", config.GetString("env"))

}
