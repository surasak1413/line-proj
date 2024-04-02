package line_api

import (
	"encoding/json"
	"fmt"
	"line-proj/config"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type IssueAccessTokenResponse struct {
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

func LineLoginIssueAccessToken(redirectURL, code string) (*IssueAccessTokenResponse, error) {
	queryParams := url.Values{}
	queryParams.Set("grant_type", "authorization_code")
	queryParams.Set("code", code)
	queryParams.Set("redirect_uri", redirectURL)
	queryParams.Set("client_id", config.Line.LineLoginChannelID)
	queryParams.Set("client_secret", config.Line.LineLoginChannelSecret)

	resp, err := resty.New().R().
		SetQueryParamsFromValues(queryParams).
		Post("https://api.line.me/oauth2/v2.1/token")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("%d : %s", resp.StatusCode(), string(resp.Body()))
	}

	var respData IssueAccessTokenResponse
	err = json.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return nil, err
	}

	return &respData, nil
}

type RefreshTokenResponse struct {
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func LineLoginRefreshToken(refreshToken string) (*RefreshTokenResponse, error) {
	queryParams := url.Values{}
	queryParams.Set("grant_type", "refresh_token")
	queryParams.Set("refresh_token", refreshToken)
	queryParams.Set("client_id", config.Line.LineLoginChannelID)
	queryParams.Set("client_secret", config.Line.LineLoginChannelSecret)

	resp, err := resty.New().R().
		SetQueryParamsFromValues(queryParams).
		Post("https://api.line.me/oauth2/v2.1/token")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("%d : %s", resp.StatusCode(), string(resp.Body()))
	}

	var respData RefreshTokenResponse
	err = json.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return nil, err
	}

	return &respData, nil
}

type GetUserProfileResponse struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}

func LineLoginGetUserProfile(accessToken string) (*GetUserProfileResponse, error) {
	resp, err := resty.New().R().
		SetAuthToken(accessToken).
		Get("https://api.line.me/v2/profile")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("%d : %s", resp.StatusCode(), string(resp.Body()))
	}

	var respData GetUserProfileResponse
	err = json.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return nil, err
	}

	return &respData, nil
}
