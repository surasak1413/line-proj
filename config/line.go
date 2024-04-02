package config

import "os"

type lineConfig struct {
	LineMessageChannelID            string `yaml:"line_message_channel_id"`
	LineLoginChannelID              string `yaml:"line_login_channel_id"`
	LineLoginChannelSecret          string `yaml:"line_login_channel_secret"`
	LineLIFFID                      string `yaml:"line_liff_id"`
	LineMessageChannelSecret        string `yaml:"line_message_channel_secret"`
	LineChannelMessagingAccessToken string `yaml:"line_channel_messaging_access_token"`
}

func LoadLineConfig() lineConfig {
	return lineConfig{
		LineMessageChannelID:            os.Getenv("LINE_MESSAGE_CHANNEL_ID"),
		LineLoginChannelID:              os.Getenv("LINE_LOGIN_CHANNEL_ID"),
		LineLoginChannelSecret:          os.Getenv("LINE_LOGIN_CHANNEL_SECRET"),
		LineLIFFID:                      os.Getenv("LINE_LIFF_ID"),
		LineMessageChannelSecret:        os.Getenv("LINE_MESSAGE_CHANNEL_SECRET"),
		LineChannelMessagingAccessToken: os.Getenv("LINE_CHANNEL_MESSAGING_ACCESS_TOKEN"),
	}
}
