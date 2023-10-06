package fileutils

import (
	"io"
	"os"
	"path"
)

func Copy(src string, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, srcFile); err != nil {
		return err
	}
	return nil
}

// ReadDirはディレクトリ内にあるファイル一覧を返します。
// 1つめのos.DirEntryにはディレクトリの一覧があり、2つめはファイルの一覧があります。
func ReadDir(name string) (*[]os.DirEntry, *[]os.DirEntry, error) {
	list, err := os.ReadDir(name)
	files, dirs := []os.DirEntry{}, []os.DirEntry{}
	if err != nil {
		return nil, nil, err
	}
	for _, l := range list {
		if l.IsDir() {
			dirs = append(dirs, l)
		} else {
			files = append(files, l)
		}
	}
	return &dirs, &files, nil
}

func DirEntryToFullPath(dir string, entry *[]os.DirEntry) *[]string {
	r := []string{}
	for _, e := range *entry {
		r = append(r, path.Join(dir, e.Name()))
	}
	return &r
}

// ReadDirはディレクトリ内にあるファイル一覧をフルパスで返します。
// 1つめにはディレクトリの一覧があり、2つめはファイルの一覧があります。
func ReadDirFullPath(name string) (*[]string, *[]string, error) {
	files, dirs, err := ReadDir(name)
	if err != nil {
		return nil, nil, err
	}

	files_str := DirEntryToFullPath(name, files)
	dirs_str := DirEntryToFullPath(name, dirs)
	return files_str, dirs_str, nil
}
