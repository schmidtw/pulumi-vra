---
title: VRA Installation & Configuration
meta_desc: Information on how to install the VRA provider.
layout: package
---

## Installation

The Pulumi VRA provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumiverse/vra`](https://www.npmjs.com/package/@pulumiverse/vra)
* Python: [`pulumiverse_vra`](https://pypi.org/project/pulumiverse-vra/)
* Go: [`github.com/pulumiverse/pulumi-vra/sdk/go/vra`](https://pkg.go.dev/github.com/pulumiverse/pulumi-vra/sdk)
* .NET: [`Pulumiverse.VRA`](https://www.nuget.org/packages/Pulumiverse.VRA)

### Provider Binary

The VRA provider binary is a third party binary. It can be installed using the pulumi plugin command.

```bash
pulumi plugin install resource vra <version> --server github://api.github.com/pulumiverse
```

Replace the version string with your desired version.

## Setup

To provision resources with the Pulumi VRA provider, you need to provide the `url` and
generally the `access_token`.

## Configuration Options

Use `pulumi config set vra:<option>`.

| Option                | Required/Optional | Description                                                                     |
|-----------------------|-------------------|---------------------------------------------------------------------------------|
| `access_token`        | Optional          | VRA access token for API operations (environment: `VRA_ACCESS_TOKEN`).   |
| `refresh_token`       | Optional          | VRA refresh token for API operations (environment: `VRA_REFRESH_TOKEN`). |
| `url`                 | Required          | VRA provider url for API operations (environment: `VRA_URL`).            |
| `insecure`            | Optional          | Specifies if TLS certificates are validated (environment: `VRA_INSECURE` / `VRA7_INSECURE`). |
| `reauthorize_timeout` | Optional          | Specifies the timeout for how often to reauthorize the access token. (environment: `VRA_REAUTHORIZE_TIMEOUT` / `VRA7_REAUTHORIZE_TIMEOUT`). |
| `api_timeout`         | Optional          | Specifies the timeout in seconds for API operations. (environment: `VRA_API_TIMEOUT`) |
