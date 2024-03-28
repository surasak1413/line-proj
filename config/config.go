package config

import (
	"github.com/joho/godotenv"
)

var Line lineConfig
var App appConfig

func NewConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	Line = LoadLineConfig()
	App = LoadAppConfig()

	return nil
}
