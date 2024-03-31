package line_api

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#text-message
*/
type TextMessage struct {
	Type   string  `json:"type"`
	Text   string  `json:"text"`
	Emojis []Emoji `json:"emojis"`
}

/*
Emoji อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/docs/messaging-api/emoji-list/
*/
type Emoji struct {
	Index     int    `json:"index"`
	ProductID string `json:"productId"`
	EmojiID   string `json:"emojiId"`
}

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#sticker-message
*/
type StickerMessage struct {
	Type      string `json:"type"`
	PackageID string `json:"packageId"`
	StickerID string `json:"stickerId"`
}

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#image-message
*/
type ImageMessage struct {
	Type string `json:"type"`

	/*
		Image file URL (Max character limit: 2000) The URL should be percent-encoded using UTF-8,
		Protocol: HTTPS (TLS 1.2 or later),
		Image format: JPEG or PNG,
		Max file size: 10 MB
	*/
	OriginalContentURL string `json:"originalContentUrl"`

	/*
		Image file URL (Max character limit: 2000) The URL should be percent-encoded using UTF-8,
		Protocol: HTTPS (TLS 1.2 or later),
		Image format: JPEG or PNG,
		Max file size: 1 MB
	*/
	PreviewImageURL string `json:"previewImageUrl"`
}

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#video-message
*/
type VideoMessage struct {
	Type string `json:"type"`

	/*
		Video file URL (Max character limit: 2000)
		Protocol: HTTPS (TLS 1.2 or later)
		Video format: mp4
		Max file size: 200 MB
	*/
	OriginalContentURL string `json:"originalContentUrl"`

	/*
		Image file URL (Max character limit: 2000) The URL should be percent-encoded using UTF-8,
		Protocol: HTTPS (TLS 1.2 or later),
		Image format: JPEG or PNG,
		Max file size: 1 MB
	*/
	PreviewImageURL string `json:"previewImageUrl"`

	TrackingID string `json:"trackingId"`
}

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#location-message
*/
type LocationMessage struct {
	Type      string  `json:"type"`
	Title     string  `json:"title"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#template-messages
*/
type TemplateMessage struct {
	Type     string      `json:"type"`
	AltText  string      `json:"altText"`
	Template interface{} `json:"template"` // template_model.go
}

type FlexMessage struct {
	Type     string       `json:"type"` // Type of the message, e.g., "flex"
	AltText  string       `json:"altText"`
	Contents FlexContents `json:"contents"`
}
