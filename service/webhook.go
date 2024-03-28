package service

import (
	"encoding/json"
	"fmt"
	"line-proj/request"
)

type webhookAction interface {
	MessagingAPIWebhook(req request.MessagingAPIWebhookRequest) error
}

func (sv *service) MessagingAPIWebhook(req request.MessagingAPIWebhookRequest) error {
	jsonData, err := json.MarshalIndent(req, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(jsonData))

	return nil
}
