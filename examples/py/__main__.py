import lbrlabs_pulumi_remotecertificate as remotecert
import pulumi

example = remotecert.CertThumbPrint(
    "example",
    server="api.pulumi.com",
    port=443,
)

pulumi.export("thumbprint", example.hash)