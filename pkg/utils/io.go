package utils

import (
	"io"
)

// A DynamicWriter is just the Writer and ByteWriter interfaces in one.
type DynamicWriter interface {
	io.Writer
	io.ByteWriter
}

// A DynamicReader is just the Reader and ByteReader interfaces in one.
type DynamicReader interface {
	io.Reader
	io.ByteReader
}
