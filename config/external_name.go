package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// dfcloud_datastore is identified by the provider-assigned id field.
	"dfcloud_datastore": config.IdentifierFromProvider,
	// dfcloud_network is identified by the provider-assigned id field.
	"dfcloud_network": config.IdentifierFromProvider,
	// dfcloud_connection is identified by the provider-assigned connection_id field.
	"dfcloud_connection": connectionExternalName(),
}

func connectionExternalName() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		if v, ok := tfstate["connection_id"].(string); ok && v != "" {
			return v, nil
		}
		return config.IDAsExternalName(tfstate)
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
