package registry

import (
	"line-proj/handler"
	"line-proj/middleware"

	"github.com/labstack/echo/v4"
)

func APIRegister(server *echo.Echo) {
	server.POST("/v1/webhook", handler.MessagingAPIWebhook, middleware.LineSignatureValidation)
	server.POST("/v1/login", handler.LineLoginGotoAuthPage)
	server.POST("/v1/login/auth", handler.LineLoginAuth)
	server.POST("/v1/login/refresh", handler.LineLoginRefreshToken)
	server.POST("/v1/login/profile", handler.LineLoginGetUserProfile)

	server.POST("/v1/products", handler.GetProductsApi)
	server.POST("/v1/buy-product", handler.BuyProductsAPI)
	server.POST("/v1/orders", handler.GetOrdersApi)
	server.Static("/webpages", "webpages")

}
