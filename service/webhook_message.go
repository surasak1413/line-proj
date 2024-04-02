package service

import (
	"fmt"
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
	MessageTypeFlex     = "flex"

	UserCommandHelp                     = "help"
	UserCommandPushReplyText            = "text"
	UserCommandPushReplySticker         = "sticker"
	UserCommandPushReplyImage           = "image"
	UserCommandPushReplyVideo           = "video"
	UserCommandPushReplyLocation        = "location"
	UserCommandPushReplyTemplateButtons = "buttons"
	UserCommandPushReplyTemplateConfirm = "confirm"
	UserCommandPushReplyFlex            = "flex"
	UserCommandPush                     = "push"
	UserCommandMulticast                = "multicast"
	UserCommandBroadcast                = "broadcast"

	TemplateTypeButtons = "buttons"
	TemplateTypeConfirm = "confirm"

	ImageSizeCover   = "cover"   // ภาพจะเติมให้เต็มกรอบ ส่วนที่เกินจะไม่แสดง
	ImageSizeContain = "contain" //  ภาพทั้งภาพจะอยู่ในกรอบ

	ImageAspectRatioRectangle = "rectangle" // 1.51 : 1
	ImageAspectRatioSquare    = "square"    // 1 : 1

	ActionTypePostback = "postback"
	ActionTypeMessage  = "message"
	ActionTypeURI      = "uri"
)

var AllUserCommand = []string{
	UserCommandHelp,
	UserCommandPushReplyText,
	UserCommandPushReplySticker,
	UserCommandPushReplyImage,
	UserCommandPushReplyVideo,
	UserCommandPushReplyLocation,
	UserCommandPushReplyTemplateButtons,
	UserCommandPushReplyTemplateConfirm,
	UserCommandPushReplyFlex,
	UserCommandPush,
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
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleReplyMessageAllUserCommand(event); err != nil {
			return err
		}
	case UserCommandPushReplyText:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleReplyMessageText(event); err != nil {
			return err
		}
	case UserCommandPushReplySticker:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleReplyMessageSticker(event); err != nil {
			return err
		}
	case UserCommandPushReplyImage:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleReplyMessageImage(event); err != nil {
			return err
		}
	case UserCommandPushReplyVideo:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleReplyMessageVideo(event); err != nil {
			return err
		}
	case UserCommandPushReplyLocation:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleReplyMessageLocation(event); err != nil {
			return err
		}
	case UserCommandPushReplyTemplateButtons:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleReplyMessageTemplateButtons(event); err != nil {
			return err
		}
	case UserCommandPushReplyTemplateConfirm:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleReplyMessageTemplateConfirm(event); err != nil {
			return err
		}
	case UserCommandPushReplyFlex:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleReplyMessageFlex(event); err != nil {
			return err
		}
	case UserCommandPush:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExamplePushMessage(event); err != nil {
			return err
		}
	case UserCommandMulticast:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleMulticastMessage(event); err != nil {
			return err
		}
	case UserCommandBroadcast:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleBroadcastMessage(event); err != nil {
			return err
		}
	}

	return nil
}
