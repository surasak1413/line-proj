package service

import "line-proj/request"

func (sv *service) WebHookActionTypePostback(event request.Event) error {
	sv.ctx.Logger().Debugf("userID %s unfollowed!", event.Source.UserID)

	return nil
}
