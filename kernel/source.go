package kernel

type LocalSource struct {
	Path    string `mapstructure:"path"`
	Version string `mapstructure:"version"`
}

type Upstream struct {
	Name string `mapstructure:"name"`
	Type string `mapstructure:"type"`
	Url  string `mapstructure:"url"`
}
