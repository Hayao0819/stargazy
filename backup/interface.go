package backup

type Backup interface {
	Add() error
	Remove() error
	Restore() error
	Files() []string
}

type BackupList interface {
	Add(name string, path ...string) error
	Remove(name string) error
	List() []Backup
	Find(name string) (*Backup, error)
}

func GetFromLocal(*BackupList) error {
	return nil
}
