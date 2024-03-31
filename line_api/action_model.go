package line_api

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#action-objects
*/
type PostbackAction struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Data        string `json:"data"` // data สำหรับ webhook postback event
	DisplayText string `json:"displayText"`
	InputOption string `json:"inputOption"`
	FillInText  string `json:"fillInText"`
}

type MessageAction struct {
	Type  string `json:"type"`
	Label string `json:"label"`
	Text  string `json:"text"`
}

type URIAction struct {
	Type   string `json:"type"`
	Label  string `json:"label"`
	URI    string `json:"uri"`
	AltURI AltURI `json:"altUri"` // uri สำหรับการเปิด LINE ผ่านช่องทางอื่นนอกจาก mobile app ถ้าไม่เซ็ตจะใช้อันเดียวกับ URI
}

type AltURI struct {
	Desktop string `json:"desktop"` // uri สำหรับ LINE for macOS and Windows
}
