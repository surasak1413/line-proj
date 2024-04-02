package service

import (
	"encoding/json"
	"fmt"
	"line-proj/request"
)

func (sv *service) WebHookActionTypeUnsend(event request.Event) error {
	prettyJSON, err := json.MarshalIndent(event, "", "    ")
	if err != nil {

		return err
	}

	fmt.Println(string(prettyJSON))

	return nil
}
