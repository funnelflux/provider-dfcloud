package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	connectionCluster "github.com/Bluesboy/provider-dfcloud/config/cluster/connection"
	datastoreCluster "github.com/Bluesboy/provider-dfcloud/config/cluster/datastore"
	networkCluster "github.com/Bluesboy/provider-dfcloud/config/cluster/network"
	connectionNamespaced "github.com/Bluesboy/provider-dfcloud/config/namespaced/connection"
	datastoreNamespaced "github.com/Bluesboy/provider-dfcloud/config/namespaced/datastore"
	networkNamespaced "github.com/Bluesboy/provider-dfcloud/config/namespaced/network"
)

const (
	resourcePrefix = "dfcloud"
	modulePath     = "github.com/Bluesboy/provider-dfcloud"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("dfcloud.funnelflux.pro"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		datastoreCluster.Configure,
		networkCluster.Configure,
		connectionCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("dfcloud.m.funnelflux.pro"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		datastoreNamespaced.Configure,
		networkNamespaced.Configure,
		connectionNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
