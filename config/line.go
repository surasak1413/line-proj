package config

import "os"

type lineConfig struct {
	LineChannelID          string
	LineChannelSecret      string
	LineChannelAccessToken string
}

func LoadLineConfig() lineConfig {
	return lineConfig{
		LineChannelID:          os.Getenv("LINE_CHANNEL_ID"),
		LineChannelSecret:      os.Getenv("LINE_CHANNEL_SECRET"),
		LineChannelAccessToken: os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	}
}
