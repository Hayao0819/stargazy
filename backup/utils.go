package backup

import "github.com/Hayao0819/stargazy/conf"

func GetBackupDir() string {
	return conf.Config.BackUpDir
}
