package tarutils

import (
	"archive/tar"
	"io"
	"os"
)

func AddFile(tw *tar.Writer, path ...string) error {
	for _, f := range path {
		file, err := os.Open(f)
		if err != nil {
			return err
		}
		defer file.Close()
		if stat, err := file.Stat(); err == nil {
			header := tar.Header{
				Name:    f,
				Size:    stat.Size(),
				Mode:    int64(stat.Mode()),
				ModTime: stat.ModTime(),
			}
			if err := tw.WriteHeader(&header); err != nil {
				return err
			}
			if _, err := io.Copy(tw, file); err != nil {
				return err
			}
		}
	}
	return nil
}

func Create(writer io.Writer, files ...string) (*tar.Writer, error) {
	tw := tar.NewWriter(writer)
	if err := AddFile(tw, files...); err != nil {
		return nil, err
	}
	return tw, nil
}
