package connection

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("dfcloud_connection", func(r *config.Resource) {
		r.ShortGroup = "connection"
		r.References["network_id"] = config.Reference{
			Type: "github.com/Bluesboy/provider-dfcloud/apis/cluster/network/v1alpha1.Network",
		}
	})
}
