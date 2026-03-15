# provider-dfcloud

`provider-dfcloud` is a Crossplane provider for Dragonfly Cloud. It is built
with Upjet from the Terraform provider
[`funnelflux/dfcloud`](https://registry.terraform.io/providers/funnelflux/dfcloud/latest).

The provider currently exposes managed resources for:
- `Network`
- `Datastore`
- `Connection`

## Install

Install the provider package into Crossplane:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
name: provider-dfcloud
spec:
package: xpkg.upbound.io/deus/provider-dfcloud:v0.0.1
```

Apply it with:

```bash
kubectl apply -f provider.yaml
```

Replace `v0.0.1` with the package version you want to install.

## Credentials Secret

The provider reads credentials from a Secret referenced by `ProviderConfig` or
`ClusterProviderConfig`.

The secret must contain a key named `credentials`. The value of this key must be
JSON with:
- required field: `api_key`

Example Secret:

```yaml
apiVersion: v1
kind: Secret
metadata:
name: dfcloud-creds
namespace: crossplane-system
type: Opaque
stringData:
credentials: |
  {
    "api_key": "dfc_XXXXXXXXXXXXXXXX"
  }
```

If you use the default Dragonfly Cloud API endpoint, omit `api_host`.

## ProviderConfig

Cluster-scoped configuration:

```yaml
apiVersion: dfcloud.funnelflux.pro/v1beta1
kind: ProviderConfig
metadata:
name: default
spec:
credentials:
  source: Secret
  secretRef:
    name: dfcloud-creds
    namespace: crossplane-system
    key: credentials
```

Namespaced configuration:

```yaml
apiVersion: dfcloud.m.funnelflux.pro/v1beta1
kind: ProviderConfig
metadata:
name: default
namespace: crossplane-system
spec:
credentials:
  source: Secret
  secretRef:
    name: dfcloud-creds
    namespace: crossplane-system
    key: credentials
```

## Example Resource

Example `Network` using the cluster-scoped `ProviderConfig`:

```yaml
apiVersion: network.dfcloud.funnelflux.pro/v1alpha1
kind: Network
metadata:
name: example-network
spec:
forProvider:
  name: example-network
  cidrBlock: 10.0.0.0/16
  location:
    provider: gcp
    region: europe-west1
providerConfigRef:
  name: default
```

## Source

Crossplane provider source:
- https://github.com/funnelflux/provider-dfcloud

Terraform provider source:
- https://github.com/funnelflux/terraform-provider-dfcloud
