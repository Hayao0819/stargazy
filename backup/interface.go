package backup

type Backup interface {
	Remove() error
	Restore() error
	Files() []string
}

type BackupList interface {
	Add(name string, files ...string) error
	Remove(name string) error
	GetList() ([]Backup, error)
	Find(name string) (*Backup, error)
	Initilize() error
}

