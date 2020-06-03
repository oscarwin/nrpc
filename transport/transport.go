package transport

import "io"

type NTransport interface {
	io.ReadWriter
	Open() error
	IsOpen() bool
}
