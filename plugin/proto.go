// Package plugin defines the protocol layer for communication between notation
// and notation external plugin.
package plugin

// Capability is a feature available in the plugin contract.
type Capability string

const (
	// CapabilitySignatureGenerator is the name of the capability
	// for a plugin to support generating raw signatures.
	CapabilitySignatureGenerator Capability = "SIGNATURE_GENERATOR.RAW"

	// CapabilityEnvelopeGenerator is the name of the capability
	// for a plugin to support generating envelope signatures.
	CapabilityEnvelopeGenerator Capability = "SIGNATURE_GENERATOR.ENVELOPE"

	// CapabilityTrustedIdentityVerifier is the name of the
	// capability for a plugin to support verifying trusted identities.
	CapabilityTrustedIdentityVerifier Capability = "SIGNATURE_VERIFIER.TRUSTED_IDENTITY"

	// CapabilityRevocationCheckVerifier is the name of the
	// capability for a plugin to support verifying revocation checks.
	CapabilityRevocationCheckVerifier Capability = "SIGNATURE_VERIFIER.REVOCATION_CHECK"
)
