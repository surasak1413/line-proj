package request

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#webhook-event-objects
*/
type MessagingAPIWebhookRequest struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#message-event
*/
type Event struct {
	Type            string          `json:"type"`
	Message         Message         `json:"message"`
	Timestamp       int64           `json:"timestamp"`
	Source          Source          `json:"source"`
	ReplyToken      string          `json:"replyToken"`
	Mode            string          `json:"mode"`
	WebhookEventID  string          `json:"webhookEventId"`
	DeliveryContext DeliveryContext `json:"deliveryContext"`
}

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#wh-text
*/
type Message struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Text string `json:"text"`
}

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#common-properties
*/
type Source struct {
	Type    string `json:"type"`
	GroupID string `json:"groupId"`
	RoomID  string `json:"roomId"`
	UserID  string `json:"userId"`
}

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/docs/messaging-api/receiving-messages/#redelivered-webhooks
*/
type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}
