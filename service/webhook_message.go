package service

import (
	"line-proj/request"
	"strings"
)

const (
	MessageTypeText     = "text"
	MessageTypeSticker  = "sticker"
	MessageTypeImage    = "image"
	MessageTypeVideo    = "video"
	MessageTypeLocation = "location"
	MessageTypeTemplate = "template"

	UserCommandHelp         = "help"
	UserCommandReply        = "reply"
	UserCommandPushText     = "push_text"
	UserCommandPushSticker  = "push_sticker"
	UserCommandPushImage    = "push_image"
	UserCommandPushVideo    = "push_video"
	UserCommandPushLocation = "push_location"
	UserCommandMulticast    = "multicast"
	UserCommandBroadcast    = "broadcast"
)

var AllUserCommand = []string{
	UserCommandHelp,
	UserCommandReply,
	UserCommandPushText,
	UserCommandPushSticker,
	UserCommandPushImage,
	UserCommandPushVideo,
	UserCommandPushLocation,
	UserCommandMulticast,
	UserCommandBroadcast,
}

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
	switch strings.TrimSpace(strings.ToLower(event.Message.Text)) {
	case UserCommandHelp:
		if err := sv.ExamplePushMessageAllUserCommand(event); err != nil {
			return err
		}
	case UserCommandReply:
		if err := sv.ExampleReplyMessage(event); err != nil {
			return err
		}
	case UserCommandPushText:
		if err := sv.ExamplePushMessageText(event); err != nil {
			return err
		}
	case UserCommandPushSticker:
		if err := sv.ExamplePushMessageSticker(event); err != nil {
			return err
		}
	case UserCommandPushImage:
		if err := sv.ExamplePushMessageImage(event); err != nil {
			return err
		}
	case UserCommandPushVideo:
		if err := sv.ExamplePushMessageVideo(event); err != nil {
			return err
		}
	case UserCommandPushLocation:
		if err := sv.ExamplePushMessageLocation(event); err != nil {
			return err
		}
	case UserCommandMulticast:
		if err := sv.ExampleMulticastMessage(event); err != nil {
			return err
		}
	case UserCommandBroadcast:
		if err := sv.ExampleBroadcastMessage(event); err != nil {
			return err
		}
	case "flex":
		if err := sv.ExamplePushFlexBoxUserCommand(event); err != nil {
			return err
		}
	}

	return nil
}
