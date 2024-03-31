package line_api

/*
โครงสร้างของ object ของแต่ละ type จะมีความแตกต่างกัน
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#message-objects
*/

// ตอบกลับข้อความ
type ReplyMessageRequest struct {
	ReplyToken string        `json:"replyToken"`
	Messages   []interface{} `json:"messages"` // สูงสุด 5 object, message_model.go
}

// ส่ง message แบบระบุ user, group หรือ room ที่ add LINE OA
type PushMessageRequest struct {
	To       string        `json:"to"`       // userId, groupId หรือ roomId
	Messages []interface{} `json:"messages"` // สูงสุด 5 object, message_model.go
}

// ส่ง message แบบระบุ user ที่ add LINE OA
type MulticastMessageRequest struct {
	To       []string      `json:"to"`       // userId สูงสุด 500
	Messages []interface{} `json:"messages"` // สูงสุด 5 object, message_model.go
}

// ส่ง message ไปยังทุกคนที่ add LINE OA
type BroadcastMessageRequest struct {
	Messages []interface{} `json:"messages"` // สูงสุด 5 object, message_model.go
}
