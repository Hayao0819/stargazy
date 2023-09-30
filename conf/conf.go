package conf

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	Name string `json:"name"`
}

var Config = AppConfig{}

func Initilize(configPath string) (*AppConfig, error) {
	if configPath == "" {
		return &Config, nil
	}

	viper.SetConfigFile(configPath)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return nil, err
	}

	return &Config, nil
}
