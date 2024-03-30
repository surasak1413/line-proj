package service

import (
	"line-proj/request"
)

const (
	MessageTypeText    = "text"
	MessageTypeSticker = "sticker"
	MessageTypeImage   = "image"

	UserCommandHelp  = "help"
	UserCommandReply = "reply"
	UserCommandPush  = "push"
)

func (sv *service) WebHookActionTypeMessage(event request.Event) error {
	switch event.Message.Type {
	case MessageTypeText:
		if err := sv.InteractWithUserCommand(event); err != nil {
			return err
		}
	}

	return nil
}

func (sv *service) InteractWithUserCommand(event request.Event) error {
	switch event.Message.Text {
	case UserCommandHelp:
		// TODO
	case UserCommandReply:
		if err := sv.ExampleReplyMessage(event); err != nil {
			return err
		}
	case UserCommandPush:
		if err := sv.ExamplePushMessage(event); err != nil {
			return err
		}

	}

	return nil
}
