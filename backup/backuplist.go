package backup

import (
	"encoding/json"
	"path"
	"time"

	"os"

	fileutils "github.com/Hayao0819/stargazy/utils/file"
	jsonutils "github.com/Hayao0819/stargazy/utils/json"
	pathutils "github.com/Hayao0819/stargazy/utils/path"
	strutils "github.com/Hayao0819/stargazy/utils/str"
	"github.com/thanhpk/randstr"
)

type List struct {
	List       []*Bak
	Id         string
	BaseDir    string
	Initilized bool
}

var timeLayout = "20060102150405"

// オリジナルのファイル
func NewBakFromOriginalFiles(date time.Time, files ...string) *Bak {
	b := Bak{
		Date:  date,
		Files: []Files{},
	}

	for _, f := range files {
		_, noext, ext := pathutils.Split(f)
		id := randstr.String(5)

		b.Files = append(b.Files, Files{
			Original: f,
			Backup:   strutils.Join(noext, "-", id, ext),
			Id:       id,
		})
	}

	return &b
}

// ファイルをバックアップします
func (l *List) Add(files ...string) error {
	now := time.Now()
	dir := pathutils.FromSlash(l.BaseDir, now.Format(timeLayout))

	// make dir
	if err := os.MkdirAll(dir, 0750); err != nil {
		return err
	}

	// Backupを作成
	bak := NewBakFromOriginalFiles(now, files...)

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

func (l *List) GetDateList() []*time.Time {
	list := []*time.Time{}
	for _, t := range l.List {
		list = append(list, &t.Date)
	}
	return list
}

// Read directory and get backup list
func (l *List) Initilize(name string) error {
	// Check
	if l.Initilized {
		return nil
	}

	// Make BaseDir
	l.BaseDir = pathutils.FromSlash(GetBackupDir(), name, l.Id)
	if err := os.MkdirAll(l.BaseDir, 0750); err != nil {
		return err
	}

	// Get directory list
	dirs, err := os.ReadDir(l.BaseDir)
	if err != nil {
		return err
	}
	for _, d := range dirs {
		jbytes, err := os.ReadFile(path.Join(l.BaseDir, d.Name(), "meta.json"))
		if err != nil {
			return err
		}
		bak := new(Bak)
		if err := json.Unmarshal(jbytes, bak); err != nil {
			return err
		}
		l.List = append(l.List, bak)
	}

	// Read list
	l.Initilized = true
	return nil
}
