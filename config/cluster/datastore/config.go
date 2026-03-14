package datastore

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("dfcloud_datastore", func(r *config.Resource) {
		r.ShortGroup = "datastore"
		r.References["network_id"] = config.Reference{
			Type: "github.com/funnelflux/provider-dfcloud/apis/cluster/network/v1alpha1.Network",
		}
	})
}
