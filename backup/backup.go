package backup

import (
	"time"
	//"github.com/Hayao0819/stargazy/kernel"
	//errutils "github.com/Hayao0819/stargazy/utils/error"
)

type Bak struct {
	Files []Files   `json:"files"`
	Date  time.Time `json:"date"`
}
type Files struct {
	Original string `json:"original"` // 元ファイルのフルパス
	Backup   string `json:"backup"`   // バックアップ先でのファイル名
	Id       string `json:"id"`       // バックアップファイルのID
}

func (b *Bak) Remove() error {
	return nil
}

func (b *Bak) Restore() error {
	return nil
}
