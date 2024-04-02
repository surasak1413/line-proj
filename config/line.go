package config

import "os"

type lineConfig struct {
	LineChannelID                   string `yaml:"line_channel_id"`
	LineChannelSecret               string `yaml:"line_channel_secret"`
	LineChannelMessagingAccessToken string `yaml:"line_channel_messaging_access_token"`
}

func LoadLineConfig() lineConfig {
	return lineConfig{
		LineChannelID:                   os.Getenv("LINE_CHANNEL_ID"),
		LineChannelSecret:               os.Getenv("LINE_CHANNEL_SECRET"),
		LineChannelMessagingAccessToken: os.Getenv("LINE_CHANNEL_MESSAGING_ACCESS_TOKEN"),
	}
}
