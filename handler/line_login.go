package handler

import (
	"line-proj/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LineLoginGotoAuthPage(ctx echo.Context) error {
	sv, err := service.New(ctx)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	url, err := sv.LineLoginGetAuthPage()
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.Redirect(http.StatusSeeOther, *url)
}

func LineLoginCallback(ctx echo.Context) error {
	code := ctx.FormValue("code")
	state := ctx.FormValue("state")

	sv, err := service.New(ctx)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	err = sv.LineLoginCallback(code, state)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.String(200, "success")
}
