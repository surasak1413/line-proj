package service

import (
	"fmt"
	"line-proj/config"
	"net/http"
	"net/url"
)

type lineLoginAction interface {
	LineLoginGetAuthPage() (*string, error)
	LineLoginCallback(code, state string) error
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
	queryParams.Set("redirect_uri", fmt.Sprintf("%s/v1/login", config.App.ServerURL))
	queryParams.Set("state", "")
	queryParams.Set("scope", "profile openid")
	queryParams.Set("nonce", "")

	req.URL.RawQuery = q.Encode()

	url := req.URL.String()

	return &url, nil
}

func (sv *service) LineLoginCallback(code, state string) error {
	// check state process
	// TODO implement

	// get access token
	// TODO

	return nil
}
