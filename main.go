package main

import (
	"context"
	"fxDemoProject/baseConfig"
	"fxDemoProject/baseGinHandler"
	"fxDemoProject/webServer"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
)

func RegisterWebServer(
	lifeCycle fx.Lifecycle,
	webServerModel *webServer.WebServerModel,
) {

	lifeCycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			return webServerModel.StartServer()
		},
		OnStop: func(ctx context.Context) error {
			return webServerModel.StopServer(ctx)
		},
	})
}

func main() {
	app := fx.New(
		baseConfig.Module,
		baseGinHandler.Module,
		webServer.Module,
		fx.Invoke(RegisterWebServer),
	)

	go func() {
		err := app.Start(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit

	err := app.Stop(context.Background())
	if err != nil {
		panic(err)
	}
}
