---
title: VMware VRA
meta_desc: Provides an overview of the VMware VRA Provider for Pulumi.
layout: package
---

The VMware VRA provider for Pulumi can be used to provision a VMware VRA based installation.

## Example

{{< chooser language "typescript,python,go,csharp,yaml" >}}
{{% choosable language typescript %}}

```typescript

// I'm not sure.

```

{{% /choosable %}}
{{% choosable language python %}}

```python

// I'm not sure.

```

{{% /choosable %}}
{{% choosable language go %}}

```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vra/sdk/go/vra/deployment"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vm, err := deployment.NewDeployment(ctx, "vra-vm", &deployment.DeploymentArgs{
			ProjectId: pulumi.ID("id"),
			Name:      pulumi.String("name"),
			CatalogItemId: pulumi.String("catalog-item-string"),
		})
		if err != nil {
			return err
		}
		ctx.Export("blueprint", vm.BlueprintId)
		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language yaml %}}

```yaml
name: vra
runtime: yaml
description: A minimal VRA program in Pulumi YAML
resources:
  newDeployment:
    type: vra:deployment:NewDeployment
    properties:
      projectId: id
      name: name
      catalogItemId: catalog-item-string
outputs:
```

{{% /choosable %}}

{{< /chooser >}}
