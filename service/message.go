package service

import (
	"line-proj/line_api"
	"line-proj/request"
)

func (sv *service) ExampleReplyMessage(event request.Event) error {
	body := line_api.ReplyMessageRequest{
		ReplyToken: event.ReplyToken,
		Messages: []interface{}{
			// สูงสุด 5 message
			line_api.TextMessage{
				Type: MessageTypeText,
				Text: "reply message",
			},
		},
	}

	if err := line_api.ReplyMessage(body); err != nil {
		return err
	}

	return nil
}

func (sv *service) ExamplePushMessage(event request.Event) error {
	body := line_api.PushMessageRequest{
		To: event.Source.UserID,
		Messages: []interface{}{
			// สูงสุด 5 message
			// ตัวอย่างการส่ง message ธรรมดา
			line_api.TextMessage{
				Type: MessageTypeText,
				Text: "push text message",
			},
			// ตัวอย่างการใส่ emoji ใน text ปกติ
			line_api.TextMessage{
				Type: MessageTypeText,
				Text: "push text message with emoji",
				Emojis: []line_api.Emoji{
					{
						Index:     0,
						ProductID: "5ac21e6c040ab15980c9b444",
						EmojiID:   "001",
					},
					{
						Index:     1,
						ProductID: "5ac1bfd5040ab15980c9b435",
						EmojiID:   "002",
					},
				},
			},
			// ตัวอย่างการส่ง sticker
			line_api.StickerMessage{
				Type:      MessageTypeSticker,
				PackageID: "11539",
				StickerID: "52114110",
			},
			// ตัวอย่างการส่งภาพ
			line_api.ImageMessage{
				Type:               MessageTypeImage,
				OriginalContentURL: "https://picsum.photos/1000",
				PreviewImageURL:    "https://picsum.photos/300",
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}
