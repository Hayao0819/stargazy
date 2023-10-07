package kernel

import (
	"os/exec"

	"github.com/Hayao0819/stargazy/conf"
	fileutils "github.com/Hayao0819/stargazy/utils/file"
)

type Src struct {
	Path    string
	Version string
}

func (s *Src) RunWithStdIo(name string, args ...string) error {
	return s.Command(name, args...).Run()
}

func (s *Src) Command(name string, args ...string) *exec.Cmd {
	c := exec.Command(name, args...)
	c.Dir = s.Path
	return c
}

func (s *Src) GetVersion() error {
	cmd := s.Command("make", "kernelconfig")
	verbytes, err := cmd.Output()
	if err != nil {
		return err
	}
	s.Version = string(verbytes)
	return nil
}

func GetSrcList() (*[]Src, error) {
	basedir := conf.Config.KernelSourceDir
	dirs, _, err := fileutils.ReadDirFullPath(basedir)
	if err != nil {
		return nil, err
	}
	sl := []Src{}

	for _, d := range *dirs {
		s := Src{
			Path: d,
		}

		if err := s.GetVersion(); err != nil {
			return nil, err
		}

		sl = append(sl, s)
	}

	return &sl, nil
}
