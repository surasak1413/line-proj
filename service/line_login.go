package service

import (
	"fmt"
	"line-proj/config"
	"line-proj/line_api"
	"line-proj/response"
	"net/http"
	"net/url"
)

type lineLoginAction interface {
	LineLoginGetAuthPage() (*string, error)
	LineLoginAuth(code, state string) (*response.LineUserProfileWithTokenResponse, error)
	LineLoginRefreshToken(refreshToken string) (*response.LineUserProfileWithTokenResponse, error)
	LineLoginGetUserProfile(accessToken string) (*response.LineProfileResponse, error)
}

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/docs/line-login/integrate-line-login/#making-an-authorization-request
*/
func (sv *service) LineLoginGetAuthPage() (*string, error) {
	baseURL := "https://access.line.me/oauth2/v2.1/authorize"

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	queryParams := url.Values{}
	queryParams.Set("response_type", "code")
	queryParams.Set("client_id", config.Line.LineChannelID)
	queryParams.Set("redirect_uri", fmt.Sprintf("%s/v1/callback", config.App.ServerURL)) // TODO change to callback uri
	queryParams.Set("state", "")
	queryParams.Set("scope", "profile openid")
	queryParams.Set("nonce", "")

	req.URL.RawQuery = q.Encode()

	url := req.URL.String()

	return &url, nil
}

func (sv *service) LineLoginAuth(code, state string) (*response.LineUserProfileWithTokenResponse, error) {
	// check state process
	// TODO implement

	// get access token
	token, err := line_api.LineLoginIssueAccessToken(fmt.Sprintf("%s/v1/callback", config.App.ServerURL), code) // TODO change to callback uri
	if err != nil {
		return nil, err
	}

	// get profile
	userProfile, err := line_api.LineLoginGetUserProfile(token.AccessToken)
	if err != nil {
		return nil, err
	}

	resp := response.LineUserProfileWithTokenResponse{
		DisplayName:   userProfile.DisplayName,
		UserID:        userProfile.UserID,
		PictureURL:    userProfile.PictureURL,
		StatusMessage: userProfile.StatusMessage,
		AccessToken:   token.AccessToken,
		ExpiresIn:     token.ExpiresIn,
		RefreshToken:  token.RefreshToken,
	}

	return &resp, nil
}

func (sv *service) LineLoginRefreshToken(refreshToken string) (*response.LineUserProfileWithTokenResponse, error) {
	// check state process
	// TODO implement

	// get access token
	token, err := line_api.LineLoginRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	// get profile
	userProfile, err := line_api.LineLoginGetUserProfile(token.AccessToken)
	if err != nil {
		return nil, err
	}

	resp := response.LineUserProfileWithTokenResponse{
		DisplayName:   userProfile.DisplayName,
		UserID:        userProfile.UserID,
		PictureURL:    userProfile.PictureURL,
		StatusMessage: userProfile.StatusMessage,
		AccessToken:   token.AccessToken,
		ExpiresIn:     token.ExpiresIn,
		RefreshToken:  token.RefreshToken,
	}

	return &resp, nil
}

func (sv *service) LineLoginGetUserProfile(accessToken string) (*response.LineProfileResponse, error) {
	// check state process
	// TODO implement

	// get profile
	userProfile, err := line_api.LineLoginGetUserProfile(accessToken)
	if err != nil {
		return nil, err
	}

	resp := response.LineProfileResponse{
		DisplayName:   userProfile.DisplayName,
		UserID:        userProfile.UserID,
		PictureURL:    userProfile.PictureURL,
		StatusMessage: userProfile.StatusMessage,
	}

	return &resp, nil
}
