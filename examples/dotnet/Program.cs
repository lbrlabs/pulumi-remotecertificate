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
