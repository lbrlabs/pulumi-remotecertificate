package main

import (
	"fmt"
	remotecert "github.com/lbrlabs/pulumi-remotecertificate/sdk/go/remotecertificate"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		example, err := remotecert.NewCertThumbPrint(ctx, "example", &remotecert.CertThumbPrintArgs{
			Server: pulumi.String("api.pulumi.com"),
			Port:   pulumi.Int(443),
		})
		if err != nil {
			return fmt.Errorf("error creating public IP: %v", err)
		}

		ctx.Export("thumbprint", example.Hash)

		return nil

	})
}
