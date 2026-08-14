package mctypes

import (
	"io"
)

// A dynamicWriter is just the Writer and ByteWriter interfaces in one.
type dynamicWriter interface {
	io.Writer
	io.ByteWriter
}

// A dynamicReader is just the Reader and ByteReader interfaces in one.
type dynamicReader interface {
	io.Reader
	io.ByteReader
}
