package response

type LineUserProfileWithTokenResponse struct {
	DisplayName   string `json:"display_name"`
	UserID        string `json:"user_id"`
	PictureURL    string `json:"picture_url"`
	StatusMessage string `json:"status_message"`
	AccessToken   string `json:"access_token"`
	ExpiresIn     int    `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
}

type LineProfileResponse struct {
	DisplayName   string `json:"display_name"`
	UserID        string `json:"user_id"`
	PictureURL    string `json:"picture_url"`
	StatusMessage string `json:"status_message"`
}
