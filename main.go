package main

import (
	"context"
	"fmt"
	"line-proj/config"
	"line-proj/registry"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	server := echo.New()
	server.Logger.SetLevel(log.DEBUG)

	// load config
	if err := config.NewConfig(".env"); err != nil {
		server.Logger.Fatalf("load config error : %s", err.Error())
	}

	// load config with json
	// if err := config.NewConfigWithYaml("config.yaml"); err != nil {
	// 	server.Logger.Fatalf("load config error : %s", err.Error())
	// }

	// register api route
	registry.APIRegister(server)

	// start server
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := server.Start(fmt.Sprintf(":%s", config.App.Port)); err != nil && err != http.ErrServerClosed {
			server.Logger.Fatalf("start server error : %s", err.Error())
		}
	}()

	// grateful shutdown
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err)
	}
}
