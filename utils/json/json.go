package jsonutils

import (
	"encoding/json"
	"os"
)

func Create(v any, path string) error {
	j, err := json.Marshal(v)
	if err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if _, err := f.Write(j); err != nil {
		return err
	}
	return nil
}
