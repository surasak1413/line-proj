package line_api

import (
	"encoding/json"
	"fmt"
	"line-proj/config"
	"net/http"

	"github.com/go-resty/resty/v2"
)

/*
API มีจำกัด	request ต่อช่วงเวลา อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#rate-limits
*/

// limit 2,000 requests per second
func ReplyMessage(body ReplyMessageRequest) error {
	prettyJSON, err := json.MarshalIndent(body, "", "    ")
	if err != nil {

		return err
	}

	fmt.Println(string(prettyJSON))

	resp, err := resty.New().R().
		SetAuthToken(config.Line.LineMessagingAccessToken).
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

// limit 2,000 requests per second
func PushMessage(body PushMessageRequest) error {
	prettyJSON, err := json.MarshalIndent(body, "", "    ")
	if err != nil {

		return err
	}

	fmt.Println(string(prettyJSON))

	resp, err := resty.New().R().
		SetAuthToken(config.Line.LineMessagingAccessToken).
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

// limit 2,000 requests per second
func MulticastMessage(body MulticastMessageRequest) error {
	resp, err := resty.New().R().
		SetAuthToken(config.Line.LineMessagingAccessToken).
		SetBody(body).
		Post("https://api.line.me/v2/bot/message/multicast")
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("%d : %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}

// limit 60 requests per hour
func BroadcastMessage(body BroadcastMessageRequest) error {
	resp, err := resty.New().R().
		SetAuthToken(config.Line.LineMessagingAccessToken).
		SetBody(body).
		Post("https://api.line.me/v2/bot/message/broadcast")
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("%d : %s", resp.StatusCode(), string(resp.Body()))
	}

	return nil
}
