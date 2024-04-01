package line_api

import (
	"encoding/json"
	"fmt"
	"line-proj/config"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`

	/*
		อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/docs/line-login/integrate-line-login/#scopes
	*/
	Scope string `json:"scope"`

	TokenType string `json:"token_type"`
}

func LineLoginGetAccessToken(redirectURL, code string) (*TokenResponse, error) {
	queryParams := url.Values{}
	queryParams.Set("grant_type", "authorization_code")
	queryParams.Set("code", code)
	queryParams.Set("redirect_uri", redirectURL)
	queryParams.Set("client_id", config.Line.LineChannelID)
	queryParams.Set("client_secret", config.Line.LineChannelSecret)

	resp, err := resty.New().R().
		SetQueryParamsFromValues(queryParams).
		Post("https://api.line.me/oauth2/v2.1/token")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("%d : %s", resp.StatusCode(), string(resp.Body()))
	}

	var respData TokenResponse
	err = json.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return nil, err
	}

	return &respData, nil
}
