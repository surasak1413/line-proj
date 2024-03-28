package handler

import (
	"line-proj/request"
	"line-proj/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MessagingAPIWebhook(ctx echo.Context) error {
	req := new(request.MessagingAPIWebhookRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	sv, err := service.New(ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = sv.MessagingAPIWebhook(*req)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.String(200, "success")
}
