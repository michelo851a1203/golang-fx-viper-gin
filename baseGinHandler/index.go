package baseGinHandler

import (
	"fxDemoProject/baseConfig"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type BaseGinHandlerModel struct {
	Config *baseConfig.BaseConfigModel
}

var Module = fx.Options(fx.Provide(NewBaseGinHandlerModel))

func NewBaseGinHandlerModel(
	config *baseConfig.BaseConfigModel,
) *BaseGinHandlerModel {
	return &BaseGinHandlerModel{
		Config: config,
	}
}

type BaseGinHandlerInterface interface {
	InitializeGinHandler() *gin.Engine
	HelloWorldRoute() gin.HandlerFunc
}

func (baseGinHandlerModel *BaseGinHandlerModel) InitializeGinHandler() *gin.Engine {
	router := gin.Default()
	router.GET("/", baseGinHandlerModel.HelloWorldRoute())
	return router
}

func (baseGinHandlerModel *BaseGinHandlerModel) HelloWorldRoute() gin.HandlerFunc {
	config := baseGinHandlerModel.Config
	cacheConfig := config.GetCurrentConfig()

	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, cacheConfig.Name)
	}
}
