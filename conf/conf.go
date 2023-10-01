package conf

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	BackUpDir string `mapstructure:"backup_dir"`
}

// Global config
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

func ApplyDefault(){

}
