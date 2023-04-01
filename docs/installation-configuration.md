---
title: Remote Certificate Installation & Configuration
meta_desc: Information on how to install the Remote Certificate provider.
layout: installation
---

## Installation

The Pulumi Remote Certificate provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@lbrlabs/pulumi-remotecertificate`](https://www.npmjs.com/package/@lbrlabs/pulumi-remotecertificate)
* Python: [`lbrlabs_remotecertificate`](https://pypi.org/project/lbrlabs-pulumi-remotecertificate/)
* Go: [`github.com/lbrlabs/pulumi-remotecertificate/sdk/go/remotecertificate`](https://pkg.go.dev/github.com/lbrlabs/pulumi-remotecertificate/sdk)
* .NET: [`lbrlabs.PulumiPackage.RemoteCertificate`](https://www.nuget.org/packages/Lbrlabs.PulumiPackage.RemoteCertificate)

### Provider Binary

The Remote Certificate provider binary is a third party binary. It can be installed using the `pulumi plugin` command.

```bash
pulumi plugin install resource remotecertificate <version> --server github://api.github.com/lbrlabs
```

Replace the version string with your desired version.

## Configuration Options

The Remote Certificate provider requires no configuration options
