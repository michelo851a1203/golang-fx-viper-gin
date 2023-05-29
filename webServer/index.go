package webServer

import (
	"context"
	"errors"
	"fmt"
	"fxDemoProject/baseConfig"
	"fxDemoProject/baseGinHandler"
	"net/http"
	"time"

	"go.uber.org/fx"
)

type WebServerModel struct {
	Server  *http.Server
	Config  *baseConfig.BaseConfigModel
	Handler *baseGinHandler.BaseGinHandlerModel
}

var Module = fx.Options(fx.Provide(NewWebServerModel))

func NewWebServerModel(
	config *baseConfig.BaseConfigModel,
	handler *baseGinHandler.BaseGinHandlerModel,
) *WebServerModel {
	baseConfig, err := config.InitializeConfig(".")
	if err != nil {
		panic(err)
	}
	config.SetCurrentConfig(baseConfig)

	return &WebServerModel{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%s", baseConfig.Port),
			Handler: handler.InitializeGinHandler(),
		},
		Config:  config,
		Handler: handler,
	}
}

type WebServerInterface interface {
	StartServer() error
	StopServer(ctx context.Context) error
}

func (webServerModel *WebServerModel) StartServer() error {
	server := webServerModel.Server
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("webServerModel >> start server error: %v", err)
	}
	return nil
}

func (WebServerModel *WebServerModel) StopServer(ctx context.Context) error {
	server := WebServerModel.Server
	shutDownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := server.Shutdown(shutDownCtx); err != nil {
		return fmt.Errorf("webServerModel >> stop server error : %v", err)
	}
	return nil
}
