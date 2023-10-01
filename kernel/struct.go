package kernel

type Binary struct {
	Kernel string
	Initramfs string
}

type Source struct {
	Path     string `mapstructure:"path"`
	Upstream string `mapstructure:"upstream"`
}

type Config struct {
}
