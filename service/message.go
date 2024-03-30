package service

import (
	"line-proj/line_api"
	"line-proj/request"
)

func (sv *service) ExampleReplyMessage(event request.Event) error {
	body := line_api.ReplyMessageRequest{
		ReplyToken: event.ReplyToken,
		Messages: []interface{}{
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
			line_api.TextMessage{
				Type: MessageTypeText,
				Text: "push message",
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}
