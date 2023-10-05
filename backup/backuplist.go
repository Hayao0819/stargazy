package backup

import (
	"path"
	"time"

	"os"

	//"github.com/Hayao0819/stargazy/kernel"
	//errutils "github.com/Hayao0819/stargazy/utils/error"
	fileutils "github.com/Hayao0819/stargazy/utils/file"
	jsonutils "github.com/Hayao0819/stargazy/utils/json"
	pathutils "github.com/Hayao0819/stargazy/utils/path"
	strutils "github.com/Hayao0819/stargazy/utils/str"
	"github.com/thanhpk/randstr"
)

type List struct {
	List    []Bak
	Id      string
	BaseDir string
}

// ファイルをバックアップします
func (l *List) Add(files ...string) error {
	bak := Bak{
		Files: []Files{},
		Date:  time.Now(),
	}
	now := bak.Date.Format("20060102")
	dir := pathutils.FromSlash(l.BaseDir, now)

	// make dir
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	// jsonを作成
	for _, f := range files {
		_, noext, ext := pathutils.Split(f)
		id := randstr.String(5)

		bak.Files = append(bak.Files, Files{
			Original: f,
			Backup:   strutils.Join(noext, "-", id, ext),
			Id:       id,
		})
	}

	// copy files
	for _, f := range bak.Files {
		if err := fileutils.Copy(f.Original, path.Join(dir, f.Backup)); err != nil {
			return err
		}
	}

	// リストアもとのJSONを作成
	if err := jsonutils.Create(bak, path.Join(dir, "meta.json")); err != nil {
		return err
	}

	// List
	l.List = append(l.List, bak)
	return nil
}

func (l *List) Remove(name string) error {
	return nil
}

func (l *List) GetList() ([]Bak, error) {
	return nil, nil
}

func (l *List) Find(name string) (*Bak, error) {
	return nil, nil
}

func (l *List) Initilize(name string) error {
	// Make BaseDir
	l.BaseDir = pathutils.FromSlash(GetBackupDir(), name, l.Id)
	if err := os.MkdirAll(l.BaseDir, 0750); err != nil {
		return err
	}

	// Read list
	return nil
}
