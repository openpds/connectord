package connector

import "errors"

var (
	ErrNotImplemented    = errors.New("not implemented")
	ErrConnectionRefused = errors.New("connection refused")
)
