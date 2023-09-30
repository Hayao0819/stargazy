package backup

type Backup interface {
	List() []string
	Add(path string) error
	Remove(path string)error
}


