package rpc

import "errors"

var (
	ErrNilRequest = errors.New("eboolkiq: request is nil")
	ErrNilQueue   = errors.New("eboolkiq: queue must not empty")
)
