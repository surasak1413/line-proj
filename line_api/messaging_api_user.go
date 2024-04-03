package line_api

import (
	"encoding/json"
	"fmt"
	"line-proj/config"
	"net/http"

	"github.com/go-resty/resty/v2"
)

// limit 2,000 requests per second
func GetUserProfileByUserID(userID string) (*GetUserProfileByUserIDResponse, error) {
	resp, err := resty.New().R().
		SetAuthToken(config.Line.LineMessagingAccessToken).
		Get(fmt.Sprintf("https://api.line.me/v2/bot/profile/%s", userID))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("%d : %s", resp.StatusCode(), string(resp.Body()))
	}

	var respData GetUserProfileByUserIDResponse
	err = json.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return nil, err
	}

	return &respData, nil
}
