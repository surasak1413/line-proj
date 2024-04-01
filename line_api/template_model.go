package line_api

type ButtonsTemplate struct {
	Type                 string        `json:"type"`
	ThumbnailImageURL    string        `json:"thumbnailImageUrl"`
	ImageAspectRatio     string        `json:"imageAspectRatio"` // default rectangle
	ImageSize            string        `json:"imageSize"`        // default cover
	ImageBackgroundColor string        `json:"imageBackgroundColor"`
	Title                string        `json:"title"`
	Text                 string        `json:"text"`
	DefaultAction        interface{}   `json:"defaultAction"` // แบบเดียวกับ Action object จะทำงานเมื่อ image, title หรือ text ถูกกด
	Actions              []interface{} `json:"actions"`       // สูงสุด 4 object, action_model.go
}

type ConfirmTemplate struct {
	Type    string        `json:"type"`
	Text    string        `json:"text"`
	Actions []interface{} `json:"actions"` // 2 object เท่านั้น, action_model.go
}
