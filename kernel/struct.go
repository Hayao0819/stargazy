package kernel

type Binary struct {
	Kernel    string
	Initramfs string
}

func NewBinary(kernel, initiramfs string) *Binary {
	return &Binary{
		Kernel:    kernel,
		Initramfs: initiramfs,
	}
}

type Source struct {
	Path     string `mapstructure:"path"`
	Upstream string `mapstructure:"upstream"`
}

type Config struct {
}
