package common

import "errors"

var (
	ErrInvalidAddress = errors.New("invalid address")
	ErrNotFound       = errors.New("not found")
	ErrRateLimit      = errors.New("rate limit exceeded")
)
