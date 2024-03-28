package config

import "os"

type appConfig struct {
	Port string
}

func LoadAppConfig() appConfig {
	return appConfig{
		Port: os.Getenv("APP_PORT"),
	}
}
