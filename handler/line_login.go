package handler

import (
	"errors"
	"line-proj/service"
	"net/http"
	"strings"

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

func LineLoginAuth(ctx echo.Context) error {
	code := ctx.QueryParam("code")
	state := ctx.QueryParam("state")

	sv, err := service.New(ctx)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	resp, err := sv.LineLoginAuth(code, state)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(200, resp)
}

func LineLoginRefreshToken(ctx echo.Context) error {
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		err := errors.New("Unauthorized")
		ctx.Logger().Error(err)
		return ctx.String(http.StatusUnauthorized, err.Error())
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		err := errors.New("Unauthorized")
		ctx.Logger().Error(err)
		return ctx.String(http.StatusUnauthorized, err.Error())
	}

	token := tokenParts[1]

	sv, err := service.New(ctx)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	resp, err := sv.LineLoginRefreshToken(token)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(200, resp)
}

func LineLoginGetUserProfile(ctx echo.Context) error {
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		err := errors.New("Unauthorized")
		ctx.Logger().Error(err)
		return ctx.String(http.StatusUnauthorized, err.Error())
	}

	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		err := errors.New("Unauthorized")
		ctx.Logger().Error(err)
		return ctx.String(http.StatusUnauthorized, err.Error())
	}

	token := tokenParts[1]

	sv, err := service.New(ctx)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	resp, err := sv.LineLoginGetUserProfile(token)
	if err != nil {
		ctx.Logger().Error(err)
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(200, resp)
}
