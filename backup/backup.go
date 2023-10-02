package backup

import "github.com/Hayao0819/stargazy/conf"

func RunAllBackup(l BackupList) error {
	list, err := (l).GetList()
	if err != nil {
		return err
	}
	for _, b := range list {
		if err := b.Restore(); err != nil {
			return err
		}
	}
	return nil
}

func RunAllKernelBackup()error{
	l := KernelBackupList{}

	return RunAllBackup(&l)
}

func GetBackupDir() string {
	return conf.Config.BackUpDir
}
