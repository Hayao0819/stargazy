package kernel

import "github.com/Hayao0819/stargazy/conf"

type Binary struct {
	Kernel    string
	Initramfs string
	Upstream  *conf.KernelUpstream
}

func NewBinary(kernel, initiramfs string) *Binary {
	return &Binary{
		Kernel:    kernel,
		Initramfs: initiramfs,
	}
}

type Config struct {
}
