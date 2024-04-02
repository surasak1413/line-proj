package service

import (
	"encoding/json"
	"fmt"
	"line-proj/config"
	"line-proj/line_api"
	"line-proj/request"
	"os"
)

func (sv *service) WebHookActionTypeFollow(event request.Event) error {
	fmt.Printf("userID %s followed!\n", event.Source.UserID)

	file, err := os.ReadFile("users.json")
	if err != nil {
		file = []byte("{}")
	}

	if len(file) == 0 {
		file = []byte("{}")
	}

	userMap := make(map[string]map[string]bool)
	if err := json.Unmarshal(file, &userMap); err != nil {
		return err
	}

	if userMap[config.Line.LineChannelID] == nil {
		userMap[config.Line.LineChannelID] = make(map[string]bool)
	}

	userMap[config.Line.LineChannelID][event.Source.UserID] = true

	updatedData, err := json.MarshalIndent(userMap, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile("users.json", updatedData, 0644); err != nil {
		return err
	}

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
