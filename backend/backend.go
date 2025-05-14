package backend

import "io"

type WriteSyncer interface {
	io.Writer
	// Sync flushes buffered logs
	Sync() error
}

type Backend interface {
	WriteSyncer
	io.Closer
	// AllowANSI determines if allow to show colorful log
	// If ANSI color is not supported, please return false to ensure performance is not affected
	AllowANSI() bool
}
