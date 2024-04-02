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

	UserCommandHelp                = "help"
	UserCommandReply               = "reply"
	UserCommandPushText            = "push_text"
	UserCommandPushSticker         = "push_sticker"
	UserCommandPushImage           = "push_image"
	UserCommandPushVideo           = "push_video"
	UserCommandPushLocation        = "push_location"
	UserCommandPushTemplateButtons = "push_template_buttons"
	UserCommandPushTemplateConfirm = "push_template_confirm"
	UserCommandPushFlex            = "push_flex"
	UserCommandMulticast           = "multicast"
	UserCommandBroadcast           = "broadcast"

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
	UserCommandReply,
	UserCommandPushText,
	UserCommandPushSticker,
	UserCommandPushImage,
	UserCommandPushVideo,
	UserCommandPushLocation,
	UserCommandPushTemplateButtons,
	UserCommandPushTemplateConfirm,
	UserCommandPushFlex,
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
		if err := sv.ExamplePushMessageAllUserCommand(event); err != nil {
			return err
		}
	case UserCommandReply:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExampleReplyMessage(event); err != nil {
			return err
		}
	case UserCommandPushText:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExamplePushMessageText(event); err != nil {
			return err
		}
	case UserCommandPushSticker:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExamplePushMessageSticker(event); err != nil {
			return err
		}
	case UserCommandPushImage:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExamplePushMessageImage(event); err != nil {
			return err
		}
	case UserCommandPushVideo:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExamplePushMessageVideo(event); err != nil {
			return err
		}
	case UserCommandPushLocation:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExamplePushMessageLocation(event); err != nil {
			return err
		}
	case UserCommandPushTemplateButtons:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExamplePushMessageTemplateButtons(event); err != nil {
			return err
		}
	case UserCommandPushTemplateConfirm:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExamplePushMessageTemplateConfirm(event); err != nil {
			return err
		}
	case UserCommandPushFlex:
		fmt.Printf("userID %s trigger command '%s'\n", event.Source.UserID, event.Message.Text)
		if err := sv.ExamplePushMessageFlex(event); err != nil {
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
