package service

import (
	"fmt"
	"line-proj/instance"
	"line-proj/line_api"
	"line-proj/request"
)

func (sv *service) WebHookActionTypeFollow(event request.Event) error {
	sv.ctx.Logger().Debugf("userID %s followed!", event.Source.UserID)

	instance.FollowerUserID = append(instance.FollowerUserID, event.Source.UserID)

	userProfile, err := line_api.GetUserProfileByUserID(event.Source.UserID)
	if err != nil {
		return err
	}

	body := line_api.PushMessageRequest{
		To: event.Source.UserID,
		Messages: []interface{}{
			// สูงสุด 5 message
			line_api.TextMessage{
				Type: MessageTypeText,
				Text: fmt.Sprintf("Welcome %s to LINE OA (follow)", userProfile.DisplayName),
			},
		},
	}

	if err := line_api.PushMessage(body); err != nil {
		return err
	}

	return nil
}
