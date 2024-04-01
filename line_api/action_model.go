package line_api

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#action-objects

label max character limit อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#action-object-label-spec
*/
type PostbackAction struct {
	Type        string `json:"type"`
	Label       string `json:"label"` // max 20
	Data        string `json:"data"`  // data สำหรับ webhook postback event (max 300)
	DisplayText string `json:"displayText"`
	InputOption string `json:"inputOption"`
	FillInText  string `json:"fillInText"`
}

type MessageAction struct {
	Type  string `json:"type"`
	Label string `json:"label"` // max 20
	Text  string `json:"text"`  // max 300
}

type URIAction struct {
	Type   string `json:"type"`
	Label  string `json:"label"`  // max 20
	URI    string `json:"uri"`    // max 1000
	AltURI AltURI `json:"altUri"` // uri สำหรับการเปิด LINE ผ่านช่องทางอื่นนอกจาก mobile app ถ้าไม่เซ็ตจะใช้อันเดียวกับ URI
}

type AltURI struct {
	Desktop string `json:"desktop"` // uri สำหรับ LINE for macOS and Windows
}
