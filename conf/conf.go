package conf

import (
	"errors"

	"github.com/spf13/viper"
)

type AppConfig struct {
	BackUpDir string `mapstructure:"backup_dir"`
}

// Global config
var Config = AppConfig{}

func Initilize(configPath string) error {
	if configPath == "" {
		return nil
	}

	viper.SetConfigFile(configPath)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}

func Validate() error {
	if Config.BackUpDir == "" {
		return errors.New("Empty backup directory")
	}
	return nil
}

func ApplyDefault() {

}
