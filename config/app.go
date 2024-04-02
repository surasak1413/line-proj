package config

import "os"

type appConfig struct {
	Port      string `yaml:"port"`
	ServerURL string `yaml:"server_url"`
}

func LoadAppConfig() appConfig {
	return appConfig{
		Port:      os.Getenv("APP_PORT"),
		ServerURL: os.Getenv("SERVER_URL"),
	}
}
