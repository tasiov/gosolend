package utils

import (
	"time"

	"github.com/tasiov/gosolend/internal/common"
)

// ValidateTimeout ensures timeout is within acceptable range
func ValidateTimeout(timeout time.Duration) time.Duration {
	if timeout <= 0 {
		return time.Duration(common.DefaultTimeout) * time.Second
	}
	return timeout
}
