package backup

import (
	"time"

	"github.com/Hayao0819/stargazy/kernel"
)

type KernelBinary struct {
	Name string
	Path *kernel.Binary
	Date time.Time
}



func (k *KernelBinary) Add(name string, path ...string) {
	
}
