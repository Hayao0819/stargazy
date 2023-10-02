package backup

import (
	"time"
	//"github.com/Hayao0819/stargazy/kernel"
)

type KernelBackup struct {
	Name string
	Path []string
	Date time.Time
}

func (k *KernelBackup) Remove() error {
	return nil
}

func (k *KernelBackup) Restore() error {
	return nil
}

func (K *KernelBackup) Files() []string {
	return nil
}
