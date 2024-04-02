package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"
	"line-proj/config"
	"line-proj/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrSignatureHeaderMissing = errors.New("header x-line-signature is missing")
	ErrInvalidSignature       = errors.New("invalid signature")
)

/*
อ่านเพิ่มเติมได้ที่ https://developers.line.biz/en/docs/messaging-api/receiving-messages/#verify-signature
*/
func LineSignatureValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// read request body
		body, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			ctx.Logger().Error(err)
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		if err := ctx.Request().Body.Close(); err != nil {
			ctx.Logger().Error(err)
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		// make request body can read again
		ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))

		// get signataure from header
		signature := ctx.Request().Header.Get("x-line-signature")
		if util.IsEmptyString(signature) {
			ctx.Logger().Error(ErrSignatureHeaderMissing)
			return ctx.String(http.StatusInternalServerError, ErrSignatureHeaderMissing.Error())
		}

		// validate signature
		if !ValidateSignature(config.Line.LineMessageChannelSecret, signature, body) {
			ctx.Logger().Error(ErrInvalidSignature)
			return ctx.String(http.StatusInternalServerError, ErrInvalidSignature.Error())
		}

		return next(ctx)
	}
}

func ValidateSignature(channelSecret string, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false
	}
	hash := hmac.New(sha256.New, []byte(channelSecret))

	_, err = hash.Write(body)
	if err != nil {
		return false
	}

	return hmac.Equal(decoded, hash.Sum(nil))
}
