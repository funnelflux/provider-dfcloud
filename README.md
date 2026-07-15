# Provider DfCloud

`provider-dfcloud` is a [Crossplane](https://crossplane.io/) provider for
[Dragonfly Cloud](https://www.dragonflydb.io/cloud), generated with
[Upjet](https://github.com/crossplane/upjet) from the upstream Terraform
provider `dragonflydb/dfcloud`.

Upbound Marketplace:

- https://marketplace.upbound.io/providers/funnelflux/provider-dfcloud/v0.1.2
- Package: `xpkg.upbound.io/funnelflux/provider-dfcloud`

## Overview

- Upstream Terraform provider: `dragonflydb/dfcloud`
- Current upstream version: `0.0.25`
- Terraform version used for schema generation: `1.5.7`
- Cluster-scoped API group: `dfcloud.funnelflux.pro`
- Namespaced API group: `dfcloud.m.funnelflux.pro`

Managed resources currently exposed by this provider:

- `Network`
- `Datastore`
- `Connection`

The repository contains two generated API surfaces:

- cluster-scoped resources under `apis/cluster/...`
- namespaced resources under `apis/namespaced/...`

## Authentication

The upstream provider expects Dragonfly Cloud credentials via provider config.
The API key can be supplied with the `DFCLOUD_API_KEY` environment variable.

## Repository Layout

- `cmd/generator/main.go` - entrypoint for Upjet generation
- `config/` - provider configuration, embedded Terraform schema, provider metadata
- `config/cluster/...` - cluster-scoped resource configuration
- `config/namespaced/...` - namespaced resource configuration
- `apis/` - generated Crossplane API types
- `package/crds/` - generated CRDs
- `.work/terraform/` - temporary Terraform workspace for schema generation
- `.work/dragonflydb/dfcloud/` - sparse checkout of upstream provider docs

## Development

Install/update git submodules first:

```console
make submodules
```

Run full code generation:

```console
make generate
```

Run the generator directly:

```console
go run cmd/generator/main.go "$PWD"
```

Run locally against a Kubernetes cluster:

```console
make run
```

Build the provider binary:

```console
make build
```

Build, push, and install package artifacts:

```console
make all
```

Verify code after regeneration:

```console
go build ./...
go test ./...
```

## Resource Behavior Notes

- `dfcloud_datastore` and `dfcloud_network` use provider-assigned IDs as external names
- `dfcloud_connection` prefers `connection_id` as external name, with fallback to `id`
- `datastore.network_id` is configured as a Crossplane reference to `Network`

## Updating Upstream Provider

When Dragonfly Cloud releases a new Terraform provider version:

1. Update `TERRAFORM_PROVIDER_VERSION` in `Makefile`.
2. If `.work/dragonflydb/dfcloud/` already exists, switch it to the new tag.
3. Remove stale Terraform lock/cache from `.work/terraform/` if generation fails.
4. Run `make generate`.
5. Review changes in:
   - `config/schema.json`
   - `config/provider-metadata.yaml`
   - `apis/**/zz_*`
   - `package/crds/*.yaml`
6. Verify with `go build ./...` and `go test ./...`.

Useful cleanup when schema generation is pinned to an old upstream version:

```console
rm -rf .work/terraform/.terraform .work/terraform/.terraform.lock.hcl .work/terraform/terraform-logs.txt
```

## Report a Bug

For bugs, improvements, or feature requests, open an issue:

https://github.com/funnelflux/provider-dfcloud/issues
