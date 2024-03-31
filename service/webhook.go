package service

import (
	"errors"
	"line-proj/request"
)

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/reference/messaging-api/#webhooks
*/

var (
	ErrUnsupportEventType  = errors.New("unsupport event type")
	ErrUnsupportSourceType = errors.New("unsupport source type")
)

const (
	SourceTypeUser = "user"

	EventTypeMessage  = "message"  // user ส่ง message
	EventTypeUnsend   = "unsend"   // user ยกเลิก message (เฉพาะ messageที่ส่งสำเร็จแล้ว)
	EventTypeFollow   = "follow"   // เพิ่ม, unblock line OA
	EventTypeUnfollow = "unfollow" // block line OA
	EventTypePostback = "postback"
)

type webhookAction interface {
	MessagingAPIWebhook(req request.MessagingAPIWebhookRequest) error
}

func (sv *service) MessagingAPIWebhook(req request.MessagingAPIWebhookRequest) error {
	for _, event := range req.Events {
		switch event.Source.Type {
		case SourceTypeUser:
			if err := sv.WebhookSourceTypeUser(event); err != nil {
				return err
			}
		default:
			return ErrUnsupportSourceType
		}

	}

	return nil
}

func (sv *service) WebhookSourceTypeUser(event request.Event) error {
	sv.ctx.Logger().Debugf("userID %s trigger event %s", event.Source.UserID, event.Type)

	switch event.Type {
	case EventTypeMessage:
		if err := sv.WebHookActionTypeMessage(event); err != nil {
			return err
		}
	case EventTypeUnsend:
		if err := sv.WebHookActionTypeUnsend(event); err != nil {
			return err
		}
	case EventTypeFollow:
		if err := sv.WebHookActionTypeFollow(event); err != nil {
			return err
		}
	case EventTypeUnfollow:
		if err := sv.WebHookActionTypeUnfollow(event); err != nil {
			return err
		}
	case EventTypePostback:
		if err := sv.WebHookActionTypePostback(event); err != nil {
			return err
		}
	default:
		return ErrUnsupportEventType
	}

	return nil
}
