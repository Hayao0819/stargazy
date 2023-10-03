package backup

import (
	"os"
	"time"

	//"github.com/Hayao0819/stargazy/kernel"
	//errutils "github.com/Hayao0819/stargazy/utils/error"
	pathutils "github.com/Hayao0819/stargazy/utils/path"
)

type KernelBackupList struct {
	List    []KernelBackup
	BaseDir string
}

func (l *KernelBackupList) Add(name string, files ...string) error {
	now := time.Now().Format("20060102")
	if err := os.MkdirAll(pathutils.FromSlash(l.BaseDir, now), 0750); err != nil {
		return err
	}

	return nil
}

func (l *KernelBackupList) Remove(name string) error {
	return nil
}

func (l *KernelBackupList) GetList() ([]Backup, error) {
	return nil, nil
}

func (l *KernelBackupList) Find(name string) (*Backup, error) {
	return nil, nil
}

func (l *KernelBackupList) Initilize() error {
	l.BaseDir = pathutils.FromSlash(GetBackupDir() + "/kernel-binary/")
	if err := os.MkdirAll(l.BaseDir, 0750); err != nil {
		return nil
	}
	return nil
}
