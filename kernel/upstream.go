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

func GetUpstreamsFromConfig()([]*Upstream){
	list := []*Upstream{}
	for _,u := range conf.Config.KernelUpstream{
		list=append(list, (*Upstream)(&u))
	}
	return list
}

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
