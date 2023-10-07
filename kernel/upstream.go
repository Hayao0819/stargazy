package kernel

import (
	"strings"

	"github.com/Hayao0819/stargazy/conf"
)

func GetUpstreamNameList() []string {
	r := []string{}
	for _, c := range conf.Config.KernelUpstream {
		r = append(r, c.Name)
	}
	return r
}

type Upstream conf.KernelUpstream

func (u *Upstream) GetType() string {
	return strings.ToLower(u.Type)
}

func (u *Upstream) GitGet() error {
	return nil
}

func (u *Upstream) Get() error {
	switch u.GetType() {
	case "git":
		return u.GitGet()
	}
	return nil
}
