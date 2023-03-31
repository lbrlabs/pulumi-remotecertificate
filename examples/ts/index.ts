import * as pulumi from "@pulumi/pulumi";
import * as remotecert from "@lbrlabs/pulumi-remotecertificate";

const example = new remotecert.CertThumbPrint("example", {
    server: "api.pulumi.com",
    port: 443,
})

export const thumbprint = example.hash;
