package backup

import (
	"time"
	//"github.com/Hayao0819/stargazy/kernel"
	//errutils "github.com/Hayao0819/stargazy/utils/error"
)

type Files struct {
	Original string `json:"original"` // 元ファイルのフルパス
	Backup   string `json:"backup"`   // バックアップ先でのファイル名
	Id       string `json:"id"`       // バックアップファイルのID
}

type Bak struct {
	Files []Files
	Date  time.Time
}

func (b *Bak) Remove() error {
	return nil
}

func (b *Bak) Restore() error {
	return nil
}

func (b *Bak) List() []string {
	return nil
}
