package conf

import (
	"errors"
	"reflect"

	"github.com/Hayao0819/stargazy/kernel"
	reflectutils "github.com/Hayao0819/stargazy/utils/reflect"
	"github.com/spf13/viper"
)

type AppConfig struct {
	BackUpDir       string          `mapstructure:"backup_dir"`
	KernelSourceDir string          `mapstructure:"kernel_source_dir"`
	KernelUpstream  kernel.Upstream `mapstructure:"kernel_upstream"`
}

// Global config
var Config = viper.New()

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
	rv := reflect.ValueOf(Config)
	for _, f := range reflectutils.GetFields(Config) {
		if f.Type == reflect.TypeOf("") {
			if rv.FieldByName(f.Name).String() == "" {
				return errors.New("Empty " + f.Name)
			}
		}
	}
	return nil
}

func ApplyDefault() {

}
