package main

import (
	"bytes"
	"crypto/sha1"
	"crypto/tls"
	"fmt"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"strings"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

func main() {
	p.RunProvider("remotecertificate", Version,
		// We tell the provider what resources it needs to support.
		// In this case, a single custom resource.
		infer.Provider(infer.Options{
			Resources: []infer.InferredResource{
				infer.Resource[CertThumbPrint, CertThumbPrintArgs, CertThumbPrintState](),
			},
			Metadata: schema.Metadata{
				PluginDownloadURL: "github://api.github.com/lbrlabs",
				Description:       "A Pulumi provider for retrieving the thumbprint of a remote certificate",
				Keywords:          []string{"pulumi", "provider", "remote", "certificate", "thumbprint"},
				License:           "Apache-2.0",
				Publisher:         "lbrlabs",
				Repository:        "https://github.com/lbrlabs/pulumi-remotecertificate/",
				Homepage:          "https://lbrlabs.com",
				LanguageMap: 		map[string]interface{}{
					"nodejs": map[string]interface{}{
						"packageName": "@lbrlabs/pulumi-remotecertificate",
						"dependencies": map[string]interface{}{
							"@pulumi/pulumi": "^3.0.0",
						},
					},
					"python": map[string]interface{}{
						"packageName": "lbrlabs_pulumiservice",
						"requires":map[string]interface{}{
							"pulumi": ">=3.0.0,<4.0.0",
						},
					},
					"csharp": map[string]interface{}{
						"rootNamespace": "Lbrlabs.PulumiPackage",
						"packageReferences": map[string]interface{}{
						  "Pulumi": "3.*",
						},
					},
					"go": map[string]interface{}{
						"generateResourceContainerTypes": true,
						"importBasePath": "github.com/lbrlabs/pulumi-remotecertificate/sdk/go/remotecertficate",
					},
				},
			},
		}))
}

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type CertThumbPrint struct{}

// Each resource has in input struct, defining what arguments it accepts.
type CertThumbPrintArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Server string `pulumi:"server"`

	Port int32 `pulumi:"port"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type CertThumbPrintState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	CertThumbPrintArgs
	// Here we define a required output called result.
	Hash string `pulumi:"hash"`
}

// All resources must implement Create at a minumum.
func (CertThumbPrint) Create(ctx p.Context, name string, input CertThumbPrintArgs, preview bool) (string, CertThumbPrintState, error) {
	state := CertThumbPrintState{CertThumbPrintArgs: input}
	if preview {
		return name, state, nil
	}

	value, err := retrieveCertThumbPrint(input.Server, input.Port)
	if err != nil {
		return "", CertThumbPrintState{}, err
	}
	state.Hash = value
	return name, state, nil
}

func retrieveCertThumbPrint(url string, port int32) (string, error) {

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", url, port), &tls.Config{})
	if err != nil {
		return "", fmt.Errorf("failed to dial %s:%d: %v", url, port, err)
	}

	certs := conn.ConnectionState().PeerCertificates
	thumbPrintCert := certs[len(certs)-1]

	fingerprint := sha1.Sum(thumbPrintCert.Raw)

	var buf bytes.Buffer
	for i, f := range fingerprint {
		if i > 0 {
			fmt.Fprintf(&buf, "")
		}
		fmt.Fprintf(&buf, "%02X", f)
	}
	return strings.ToLower(buf.String()), nil

}
