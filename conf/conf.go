package conf

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	BackUpDir       string           `mapstructure:"backup_dir"`
	KernelSourceDir string           `mapstructure:"kernel_source_dir"`
	KernelUpstream  []KernelUpstream `mapstructure:"kernel_upstream"`
}

type KernelUpstream struct {
	Name string `mapstructure:"name"`
	Type string `mapstructure:"type"`
	Url  string `mapstructure:"url"`
}

// Global config
var Config = AppConfig{}

func Initilize(configPath string) error {
	vp := viper.New()

	/*
		if configPath == "" {
			return nil
		}
	*/

	// Set config file
	vp.SetConfigFile(configPath)

	// Set default
	vp.SetDefault("backup_dir", "/usr/share/stargazy/backup")
	vp.SetDefault("kernel_source_dir", "/usr/share/stargazy/kernel_source")

	// Read from env
	vp.AutomaticEnv()

	// Read from configFile
	if err := vp.ReadInConfig(); err != nil {
		return err
	}

	// Unmershal to global var
	if err := vp.Unmarshal(&Config); err != nil {
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
