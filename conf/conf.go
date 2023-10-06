package conf

import (
	"github.com/Hayao0819/stargazy/kernel"
	"github.com/spf13/viper"
)

type AppConfig struct {
	BackUpDir       string          `mapstructure:"backup_dir"`
	KernelSourceDir string          `mapstructure:"kernel_source_dir"`
	KernelUpstream  kernel.Upstream `mapstructure:"kernel_upstream"`
}

// Global config
var Config = AppConfig{}
var Viper = viper.New()

func Initilize(configPath string) error {

	// Set config file
	Viper.SetConfigFile(configPath)

	// Set default
	Viper.SetDefault("backup_dir", "/usr/share/stargazy/backup")
	Viper.SetDefault("kernel_source_dir", "/usr/share/stargazy/kernel_source")

	// Read from env
	Viper.AutomaticEnv()

	// Read from configFile
	if configPath != "" {
		if err := Viper.ReadInConfig(); err != nil {
			return err
		}
	}

	// Unmershal to global var
	if err := Viper.Unmarshal(&Config); err != nil {
		return err
	}

	return nil
}

func Validate() error {
	/*
		rv := reflect.ValueOf(Config)
		for _, f := range reflectutils.GetFields(Config) {
			if f.Type == reflect.TypeOf("") {
				if rv.FieldByName(f.Name).String() == "" {
					return errors.New("Empty " + f.Name)
				}
			}
		}
	*/
	return nil
}

func ApplyDefault() {

}
