---
title: Remote Certificate
meta_desc: Retrieves information about a remote certificate on a website
layout: overview
---

The Remote Certificate provider for Pulumi can be used to create a resources that contains the details of a remote certificate from any website.

This is especially useful when working with OIDC providers, which require a certificate thumbprint which can often require multiple manual steps to calculate.

## Example

{{< chooser language "typescript,python,go,csharp" >}}
{{% choosable language typescript %}}

```typescript
import * as remotecert from "@lbrlabs/pulumi-remotecertificate";

const example = new remotecert.CertThumbPrint("example", {
    server: "api.pulumi.com",
    port: 443,
})

export const thumbprint = example.hash;
```

{{% /choosable %}}
{{% choosable language python %}}

```python
import lbrlabs_pulumi_remotecertificate as remotecert
import pulumi

example = remotecert.CertThumbPrint(
    "example",
    server="api.pulumi.com",
    port=443,
)

pulumi.export("thumbprint", example.hash)
```

{{% /choosable %}}
{{% choosable language go %}}

```go
import (
	"fmt"
	remotecert "github.com/lbrlabs/pulumi-remotecertificate/sdk/remotecertificate"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		example, err := remotecert.NewCertThumbPrint(ctx, "example", &remotecert.CertThumbPrintArgs{
			Server: pulumi.String("api.pulumi.com"),
			Port: pulumi.Int(443)
		})
		if err != nil {
			return fmt.Errorf("error creating public IP: %v", err)
		}


		ctx.Export("thumbprint", example.Hash)

		return nil
	})
}
```

{{% /choosable %}}
{{% choosable language csharp %}}

```csharp
using System.Collections.Generic;
using Pulumi;
using Lbrlabs.PulumiPackage.Remotecertificate;

return await Deployment.RunAsync(() =>

{
    var example = new CertThumbPrint("example", new CertThumbPrintArgs
    {
        Server = "api.pulumi.com",
        Port = 443,
    });

    return new Dictionary<string, object?>
      {
         { "thumbprint", example.Hash },
      };
});
```

{{% /choosable %}}

{{< /chooser >}}
