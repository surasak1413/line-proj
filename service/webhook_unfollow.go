package service

import (
	"encoding/json"
	"fmt"
	"line-proj/config"
	"line-proj/request"
	"os"
)

func (sv *service) WebHookActionTypeUnfollow(event request.Event) error {
	fmt.Printf("userID %s unfollowed!\n", event.Source.UserID)

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

	if userMap[config.Line.LineMessagingChannelID] == nil {
		userMap[config.Line.LineMessagingChannelID] = make(map[string]bool)
	}

	userMap[config.Line.LineMessagingChannelID][event.Source.UserID] = false

	updatedData, err := json.MarshalIndent(userMap, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile("users.json", updatedData, 0644); err != nil {
		return err
	}

	return nil
}
