package utils

import (
	"github.com/tasiov/gosolend/internal/common"
)

// ValidateAddress ensures a Solana address is valid
func ValidateAddress(address string) error {
	if len(address) != 44 {
		return common.ErrInvalidAddress
	}
	// Add more validation as needed
	return nil
}
