package line_api

type GetUserProfileByUserIDResponse struct {
	DisplayName   string `json:"displayName"`
	UserID        string `json:"userId"`
	Language      string `json:"language"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}
