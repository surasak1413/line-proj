package service

import (
	"encoding/json"
	"line-proj/request"
)

func (sv *service) WebHookActionTypeUnfollow(event request.Event) error {
	prettyJSON, err := json.MarshalIndent(event, "", "    ")
	if err != nil {

		return err
	}

	sv.ctx.Logger().Debug(string(prettyJSON))

	return nil
}
