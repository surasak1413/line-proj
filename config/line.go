package config

import "os"

type lineConfig struct {
	LineMessagingChannelID     string `yaml:"line_messaging_channel_id"`
	LineMessagingChannelSecret string `yaml:"line_messaging_channel_secret"`
	LineMessagingAccessToken   string `yaml:"line_messaging_access_token"`

	LineLIFFID             string `yaml:"line_liff_id"`
	LineLoginChannelID     string `yaml:"line_login_channel_id"`
	LineLoginChannelSecret string `yaml:"line_login_channel_secret"`
}

func LoadLineConfig() lineConfig {
	return lineConfig{
		LineMessagingChannelID:     os.Getenv("LINE_MESSAGING_CHANNEL_ID"),
		LineMessagingChannelSecret: os.Getenv("LINE_MESSAGING_CHANNEL_SECRET"),
		LineMessagingAccessToken:   os.Getenv("LINE_MESSAGING_ACCESS_TOKEN"),

		LineLIFFID:             os.Getenv("LINE_LIFF_ID"),
		LineLoginChannelID:     os.Getenv("LINE_LOGIN_CHANNEL_ID"),
		LineLoginChannelSecret: os.Getenv("LINE_LOGIN_CHANNEL_SECRET"),
	}
}
