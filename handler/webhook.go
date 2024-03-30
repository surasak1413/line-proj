package handler

import (
	"line-proj/request"
	"line-proj/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func MessagingAPIWebhook(ctx echo.Context) error {
	if err := printRequestBody(ctx); err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	req := new(request.MessagingAPIWebhookRequest)

	if err := ctx.Bind(req); err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	sv, err := service.New(ctx)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = sv.MessagingAPIWebhook(*req)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.String(200, "success")
}
