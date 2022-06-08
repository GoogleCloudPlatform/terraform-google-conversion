package conversion

import (
	"go.uber.org/zap"
)

// Single struct for options so that adding new options does not require
// updating function signatures all along the pipe
type CAIToHCLOptions struct {
	ErrorLogger *zap.Logger
}

// Returns HCL as a byte slice, using hclwrite.Tokens.Bytes() internally.
func CAIToHCL(assets []Asset, options *CAIToHCLOptions) ([]byte, error) {
	// For now, no assets are supported.
	var hclBytes []byte
	return hclBytes, nil
}
