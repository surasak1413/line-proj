package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"line-proj/config"
	"line-proj/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	ErrSignatureHeaderMissing = "header x-line-signature is missing"
	ErrInvalidSignature       = "invalid signature"
)

func LineSignatureValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		defer ctx.Request().Body.Close()

		body, err := io.ReadAll(ctx.Request().Body)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		signature := ctx.Request().Header.Get("x-line-signature")
		if util.IsEmptyString(signature) {
			return ctx.String(http.StatusInternalServerError, ErrSignatureHeaderMissing)
		}

		if !ValidateSignature(config.Line.LineChannelSecret, signature, body) {
			return ctx.String(http.StatusInternalServerError, ErrInvalidSignature)
		}

		return nil
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
