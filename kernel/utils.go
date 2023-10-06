package kernel

import "github.com/Hayao0819/stargazy/conf"

func GetUpstreamNameList() []string {
	r := []string{}
	for _, c := range conf.Config.KernelUpstream {
		r = append(r, c.Name)
	}
	return r
}
