package plugin

// KeySpec is type of the signing algorithm, including algorithm and size.
type KeySpec string

// one of the following supported key spec names.
//
// https://github.com/notaryproject/notaryproject/blob/main/specs/signature-specification.md#algorithm-selection
const (
	KeySpecRSA2048 KeySpec = "RSA-2048"
	KeySpecRSA3072 KeySpec = "RSA-3072"
	KeySpecRSA4096 KeySpec = "RSA-4096"
	KeySpecEC256   KeySpec = "EC-256"
	KeySpecEC384   KeySpec = "EC-384"
	KeySpecEC521   KeySpec = "EC-521"
)

// HashAlgorithm is the type of a hash algorithm.
type HashAlgorithm string

// one of the following supported hash algorithm names.
//
// https://github.com/notaryproject/notaryproject/blob/main/specs/signature-specification.md#algorithm-selection
const (
	HashAlgorithmSHA256 HashAlgorithm = "SHA-256"
	HashAlgorithmSHA384 HashAlgorithm = "SHA-384"
	HashAlgorithmSHA512 HashAlgorithm = "SHA-512"
)

// SignatureAlgorithm is the type of signature algorithm
type SignatureAlgorithm string

// one of the following supported signing algorithm names.
//
// https://github.com/notaryproject/notaryproject/blob/main/specs/signature-specification.md#algorithm-selection
const (
	SignatureAlgorithmECDSA_SHA256      SignatureAlgorithm = "ECDSA-SHA-256"
	SignatureAlgorithmECDSA_SHA384      SignatureAlgorithm = "ECDSA-SHA-384"
	SignatureAlgorithmECDSA_SHA512      SignatureAlgorithm = "ECDSA-SHA-512"
	SignatureAlgorithmRSASSA_PSS_SHA256 SignatureAlgorithm = "RSASSA-PSS-SHA-256"
	SignatureAlgorithmRSASSA_PSS_SHA384 SignatureAlgorithm = "RSASSA-PSS-SHA-384"
	SignatureAlgorithmRSASSA_PSS_SHA512 SignatureAlgorithm = "RSASSA-PSS-SHA-512"
)
