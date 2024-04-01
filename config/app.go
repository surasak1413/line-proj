package config

import "os"

type appConfig struct {
	Port      string
	ServerURL string
}

func LoadAppConfig() appConfig {
	return appConfig{
		Port:      os.Getenv("APP_PORT"),
		ServerURL: os.Getenv("SERVER_URL"),
	}
}
