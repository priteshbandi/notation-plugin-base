package main

import (
	"context"

	"github.com/priteshbandi/notation-plugin-base/plugin"
)

type DummyPlugin struct {
}

// NewExamplePlugin validate the metadata of the plugin and return a *CLIPlugin.
func NewExamplePlugin() (*DummyPlugin, error) {
	return &DummyPlugin{}, nil
}

// DescribeKey returns the KeySpec of a key.
func (p *DummyPlugin) DescribeKey(ctx context.Context, req *plugin.DescribeKeyRequest) (*plugin.DescribeKeyResponse, error) {
	return &plugin.DescribeKeyResponse{}, nil
}

// GenerateSignature generates the raw signature based on the request.
func (p *DummyPlugin) GenerateSignature(ctx context.Context, req *plugin.GenerateSignatureRequest) (*plugin.GenerateSignatureResponse, error) {
	return &plugin.GenerateSignatureResponse{}, nil
}

// GenerateEnvelope generates the Envelope with signature based on the
// request.
func (p *DummyPlugin) GenerateEnvelope(ctx context.Context, req *plugin.GenerateEnvelopeRequest) (*plugin.GenerateEnvelopeResponse, error) {
	return &plugin.GenerateEnvelopeResponse{}, nil
}

func (p *DummyPlugin) VerifySignature(ctx context.Context, req *plugin.VerifySignatureRequest) (*plugin.VerifySignatureResponse, error) {
	return &plugin.VerifySignatureResponse{
		ProcessedAttributes: []interface{}{},
		VerificationResults: map[plugin.Capability]*plugin.VerificationResult{},
	}, nil
}

func (p *DummyPlugin) GetMetadata(ctx context.Context, req *plugin.GetMetadataRequest) (*plugin.GetMetadataResponse, error) {
	return &plugin.GetMetadataResponse{
		Name:         "vola! you have new plugin",
		Description:  "🍺",
		URL:          "Reach out to me via url",
		Version:      "1.0.0",
		Capabilities: []plugin.Capability{plugin.CapabilitySignatureGenerator, plugin.CapabilityTrustedIdentityVerifier, plugin.CapabilityRevocationCheckVerifier},
	}, nil
}
