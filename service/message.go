package service

import (
	"fmt"
	"line-proj/instance"
	"line-proj/line_api"
	"line-proj/request"
	"strings"
)

func (sv *service) ExamplePushMessageAllUserCommand(event request.Event) error {
	body := line_api.PushMessageRequest{
		To: event.Source.UserID,
		Messages: []interface{}{
			// สูงสุด 5 message
			line_api.TextMessage{
				Type: MessageTypeText,
				Text: strings.Join(AllUserCommand, ", "),
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}

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

func (sv *service) ExamplePushMessageText(event request.Event) error {
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
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}

func (sv *service) ExamplePushMessageSticker(event request.Event) error {
	body := line_api.PushMessageRequest{
		To: event.Source.UserID,
		Messages: []interface{}{
			// สูงสุด 5 message
			// ตัวอย่างการส่ง sticker
			line_api.StickerMessage{
				Type:      MessageTypeSticker,
				PackageID: "11539",
				StickerID: "52114110",
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}

func (sv *service) ExamplePushMessageImage(event request.Event) error {
	body := line_api.PushMessageRequest{
		To: event.Source.UserID,
		Messages: []interface{}{
			// สูงสุด 5 message
			// ตัวอย่างการส่งภาพ
			line_api.ImageMessage{
				Type:               MessageTypeImage,
				OriginalContentURL: "https://superrich-exchange-dev-media.s3.ap-southeast-1.amazonaws.com/employee/profile/20231124/2Yc6V6TYa8sZpmorRxiLUPD1ynt.jpeg",
				PreviewImageURL:    "https://superrich-exchange-dev-media.s3.ap-southeast-1.amazonaws.com/employee/profile/20231124/2Yc6V6TYa8sZpmorRxiLUPD1ynt.jpeg",
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}

func (sv *service) ExamplePushMessageVideo(event request.Event) error {
	body := line_api.PushMessageRequest{
		To: event.Source.UserID,
		Messages: []interface{}{
			// สูงสุด 5 message
			// ตัวอย่างการส่งวิดีโอ
			line_api.VideoMessage{
				Type:               MessageTypeVideo,
				OriginalContentURL: "https://superrich-exchange-dev-media.s3.ap-southeast-1.amazonaws.com/company/profile/20240331/2eST7DDttdTODk3yNRXbDKjpslq.mp4",
				PreviewImageURL:    "https://superrich-exchange-dev-media.s3.ap-southeast-1.amazonaws.com/employee/profile/20231124/2Yc6V6TYa8sZpmorRxiLUPD1ynt.jpeg",
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}

func (sv *service) ExamplePushMessageLocation(event request.Event) error {
	body := line_api.PushMessageRequest{
		To: event.Source.UserID,
		Messages: []interface{}{
			// สูงสุด 5 message
			// ตัวอย่างการส่ง location
			line_api.LocationMessage{
				Type:      MessageTypeLocation,
				Title:     "my location",
				Address:   "1-3 Kioicho, Chiyoda-ku, Tokyo, 102-8282, Japan",
				Latitude:  35.67966,
				Longitude: 139.73669,
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}

func (sv *service) ExamplePushMessageTemplate(event request.Event) error {
	body := line_api.PushMessageRequest{
		To: event.Source.UserID,
		Messages: []interface{}{
			// สูงสุด 5 message
			// ตัวอย่างการส่ง template
			line_api.TemplateMessage{
				Type: MessageTypeTemplate,
				// TODO
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}

func (sv *service) ExampleMulticastMessage(event request.Event) error {
	userProfile, err := line_api.GetUserProfileByUserID(event.Source.UserID)
	if err != nil {
		return err
	}

	body := line_api.MulticastMessageRequest{
		To: instance.FollowerUserID,
		Messages: []interface{}{
			// สูงสุด 5 message
			line_api.TextMessage{
				Type: MessageTypeText,
				Text: fmt.Sprintf("Welcome %s to LINE OA (multicast)", userProfile.DisplayName),
			},
		},
	}

	if err := line_api.MulticastMessage(body); err != nil {
		return err
	}

	return nil
}

func (sv *service) ExampleBroadcastMessage(event request.Event) error {
	userProfile, err := line_api.GetUserProfileByUserID(event.Source.UserID)
	if err != nil {
		return err
	}

	body := line_api.BroadcastMessageRequest{
		Messages: []interface{}{
			// สูงสุด 5 message
			line_api.TextMessage{
				Type: MessageTypeText,
				Text: fmt.Sprintf("Welcome %s to LINE OA (broadcast)", userProfile.DisplayName),
			},
		},
	}

	if err := line_api.BroadcastMessage(body); err != nil {
		return err
	}

	return nil
}
