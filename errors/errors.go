package errors

import (
	"encoding/json"
	"fmt"
)

type Code string

const (
	CodeValidation                 Code = "VALIDATION_ERROR"
	CodeUnsupportedContractVersion Code = "UNSUPPORTED_CONTRACT_VERSION"
	CodeAccessDenied               Code = "ACCESS_DENIED"
	CodeThrottled                  Code = "THROTTLED"
	CodeGeneric                    Code = "ERROR"
)

const (
	ErrorMsgMalformedInput     string = "Input is not a valid JSON"
	ErrorMsgMalformedOutputFmt string = "Failed to generate response. Error: %s"
)

// PluginError is used when the signature associated is no longer
// valid.
type PluginError struct {
	ErrCode Code   `json:"errorCode"`
	Msg     string `json:"errorMessage"`
}

func NewError(code Code, msg string) *PluginError {
	return &PluginError{
		ErrCode: code,
		Msg:     msg,
	}
}

func NewGenericError(msg string) *PluginError {
	return NewError(CodeGeneric, msg)
}

func NewGenericErrorf(format string, msg string) *PluginError {
	return NewError(CodeGeneric, fmt.Sprintf(format, msg))
}

func NewUnsupportedError(msg string) *PluginError {
	return NewError(CodeValidation, msg+"is not supported")
}

func NewValidationError(msg string) *PluginError {
	return NewError(CodeValidation, msg)
}

func NewValidationErrorf(format string, msg string) *PluginError {
	return NewError(CodeValidation, fmt.Sprintf(format, msg))
}

func NewUnsupportedContractVersionError(version string) *PluginError {
	return NewError(CodeUnsupportedContractVersion, fmt.Sprintf("%q is not a supported notary plugin contract version", version))
}

func NewJSONParsingError(msg string) *PluginError {
	return NewValidationError(msg)
}

// Error returns the formatted error message.
func (e *PluginError) Error() string {
	op, err := json.Marshal(e)
	if err != nil {
		return "something went wrong"
	}
	return string(op)
}
