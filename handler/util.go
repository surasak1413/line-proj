package handler

import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v4"
)

func printRequestBody(ctx echo.Context) error {
	body, err := io.ReadAll(ctx.Request().Body)
	if err != nil {
		return err
	}

	var jsonData interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return err
	}

	prettyJSON, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {

		return err
	}

	ctx.Logger().Debug(string(prettyJSON))

	return nil
}
