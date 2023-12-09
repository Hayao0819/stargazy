package kernel

import (
	"os"
	"strings"

	"github.com/Hayao0819/stargazy/conf"
	"github.com/go-git/go-git/v5"
)

func GetUpstreamNameList() []string {
	r := []string{}
	for _, c := range conf.Config.KernelUpstream {
		r = append(r, c.Name)
	}
	return r
}

type Upstream conf.KernelUpstream

func GetUpstreamsFromConfig() []*Upstream {
	list := []*Upstream{}
	for _, u := range conf.Config.KernelUpstream {
		casted := Upstream(u)
		list = append(list, &casted)
	}
	return list
}

func (u *Upstream) GetType() string {
	return strings.ToLower(u.Type)
}

func (u *Upstream) GitGet() error {
	sourceDir := conf.Config.KernelSourceDir
	_, err := git.PlainClone(sourceDir, false, &git.CloneOptions{
		URL:      u.Url,
		Progress: os.Stdout,
	})
	if err != nil {
		return err
	}

	return nil

}

func (u *Upstream) Get() error {
	switch u.GetType() {
	case "git":
		return u.GitGet()
	}
	return nil
}
