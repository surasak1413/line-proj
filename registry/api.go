package registry

import (
	"line-proj/handler"
	"line-proj/middleware"

	"github.com/labstack/echo/v4"
)

func APIRegister(server *echo.Echo) {
	server.POST("/v1/webhook", handler.MessagingAPIWebhook, middleware.LineSignatureValidation)

	server.Static("/webpages", "webpages")

}
