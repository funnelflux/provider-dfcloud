package network

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("dfcloud_network", func(r *config.Resource) {
		r.ShortGroup = "network"
	})
}
