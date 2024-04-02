package config

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

var Line lineConfig
var App appConfig

func NewConfig(fileName string) error {
	err := godotenv.Load(fileName)
	if err != nil {
		return err
	}

	Line = LoadLineConfig()
	App = LoadAppConfig()

	return nil
}

func NewConfigWithYaml(fileName string) error {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	// var appConf appConfig
	if err := yaml.Unmarshal(yamlFile, &App); err != nil {
		return err
	}

	// var lineConf lineConfig
	if err := yaml.Unmarshal(yamlFile, &Line); err != nil {
		return err
	}

	return nil
}
