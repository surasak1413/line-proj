package service

import (
	"fmt"
	"line-proj/request"
)

func (sv *service) WebHookActionTypeUnfollow(event request.Event) error {
	fmt.Printf("userID %s unfollowed!\n", event.Source.UserID)

	return nil
}
