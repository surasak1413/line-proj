package line_api

import (
	"fmt"
	"line-proj/config"
	"net/http"

	"github.com/go-resty/resty/v2"
)

func ReplyMessage(body ReplyMessageRequest) error {
	resp, err := resty.New().R().
		SetAuthToken(config.Line.LineChannelAccessToken).
		SetBody(body).
		Post("https://api.line.me/v2/bot/message/reply")
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("%d : %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

func PushMessage(body PushMessageRequest) error {
	resp, err := resty.New().R().
		SetAuthToken(config.Line.LineChannelAccessToken).
		SetBody(body).
		Post("https://api.line.me/v2/bot/message/push")
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("%d : %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}
