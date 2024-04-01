package config

import "os"

type lineConfig struct {
	LineChannelID                   string
	LineChannelSecret               string
	LineChannelMessagingAccessToken string
}

func LoadLineConfig() lineConfig {
	return lineConfig{
		LineChannelID:                   os.Getenv("LINE_CHANNEL_ID"),
		LineChannelSecret:               os.Getenv("LINE_CHANNEL_SECRET"),
		LineChannelMessagingAccessToken: os.Getenv("LINE_CHANNEL_MESSAGING_ACCESS_TOKEN"),
	}
}
