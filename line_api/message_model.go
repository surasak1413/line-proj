package line_api

/*
โครงสร้างของ object ของแต่ละ type จะมีความแตกต่างกัน
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#message-objects
*/

type ReplyMessageRequest struct {
	ReplyToken string        `json:"replyToken"`
	Messages   []interface{} `json:"messages"`
}

type PushMessageRequest struct {
	To       string        `json:"to"` // userId, groupId หรือ roomId
	Messages []interface{} `json:"messages"`
}

// สูงสุด 5
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

type StickerMessage struct {
	Type      string `json:"type"`
	PackageID string `json:"packageId"`
	StickerID string `json:"stickerId"`
}

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

type LocationMessage struct {
	Type      string  `json:"type"`
	Title     string  `json:"title"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// TODO template

// TODO flex
