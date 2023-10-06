package backup

import "github.com/Hayao0819/stargazy/conf"

func GetBackupDir() string {
	return conf.Config.BackUpDir
}

func GetKernelBackupList() (*List, error) {
	list := new(List)
	err := list.Initilize("kernel-backup")
	return list, err
}
