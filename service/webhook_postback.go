package service

import (
	"fmt"
	"line-proj/line_api"
	"line-proj/request"
)

func (sv *service) WebHookActionTypePostback(event request.Event) error {
	fmt.Printf("userID %s trigger postback!\n", event.Source.UserID)

	body := line_api.PushMessageRequest{
		To: event.Source.UserID,
		Messages: []interface{}{
			// สูงสุด 5 message
			line_api.TextMessage{
				Type: MessageTypeText,
				Text: fmt.Sprintf("data from postback '%s'", event.Postback.Data),
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}
