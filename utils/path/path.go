package pathutils

import (
	"path/filepath"
	"strings"
)

func FromSlash(path ...string) string {

	return filepath.FromSlash(filepath.Clean(strings.Join(path, "/")))
}

// 拡張子を除いたファイル名を返します
func BaseWithoutExt(path string) string {
	return strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
}

// パス上で拡張子を除いたファイルの名前を変更します
// 実際に名前変更を実行するわけでは有りません
func RenamePathWithoutExt(path string, new string) string {
	ext := filepath.Base(path)
	dir := filepath.Dir(path)
	return filepath.Join(dir, new+ext)
}

// DirName, BaseName, Extに分割します
func Split(path string) (string, string, string) {
	return filepath.Dir(path), BaseWithoutExt(path), filepath.Ext(path)
}
