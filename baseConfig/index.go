package baseConfig

import (
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type BaseConfigModel struct {
}

var Module = fx.Options(fx.Provide(NewBaseConfigModel))

func NewBaseConfigModel() *BaseConfigModel {
	return &BaseConfigModel{}
}

type NewBaseConfigInterface interface {
	InitializeConfig(currentPath string) (*BaseConfig, error)
	SetCurrentConfig(config *BaseConfig)
	GetCurrentConfig() *BaseConfig
}

type BaseConfig struct {
	Port string `mapStructure:"port"`
	Name string `mapStructure:"name"`
}

var cacheConfig *BaseConfig

func (baseConfigModel *BaseConfigModel) InitializeConfig(currentPath string) (*BaseConfig, error) {
	viper.AddConfigPath(currentPath)
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var baseConfig = &BaseConfig{}
	err = viper.Unmarshal(baseConfig)
	if err != nil {
		return nil, err
	}

	return baseConfig, err
}

func (BaseConfigModel *BaseConfigModel) SetCurrentConfig(config *BaseConfig) {
	cacheConfig = config
}

func (BaseConfigModel *BaseConfigModel) GetCurrentConfig() *BaseConfig {
	return cacheConfig
}
