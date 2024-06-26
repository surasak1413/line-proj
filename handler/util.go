package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

func printRequestBody(ctx echo.Context) error {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	if err := ctx.Request().Body.Close(); err != nil {
		return err
	}

	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))

	var jsonData interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return err
	}

	prettyJSON, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {

		return err
	}

	fmt.Println(string(prettyJSON))

	return nil
}
