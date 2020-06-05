package transport

import (
	"encoding/binary"
	"io"

	"github.com/oscarwin/nrpc/codes"
	"github.com/oscarwin/nrpc/encoding"
)

const MaxPayloadSize = 10 * 1024 * 1024

type Transport interface {
	io.ReadWriter
	Open() error
	IsOpen() bool
	Close() error
}

type Framer interface {
	WriteFrame()
	ReadFrame()
}

type framer struct {
}

func (f *framer) ReadFrame(t Transport) ([]byte, error) {
	frameHeader := make([]byte, encoding.FrameHeadLen)
	if num, err := io.ReadFull(t, frameHeader); num != encoding.FrameHeadLen || err != nil {
		return nil, err
	}

	// check Magic
	if magic := uint8(frameHeader[0]); magic != encoding.Magic {
		return nil, codes.NewFrameWorkError(codes.ClientMsgError, "invalid magic")
	}
	// check length
	length := binary.BigEndian.Uint32(frameHeader[7:11])
	if length > MaxPayloadSize {
		return nil, codes.NewFrameWorkError(codes.ClientMsgError, "payload too large")
	}
	// read payload
	buffer := make([]byte, length)
	if num, err := io.ReadFull(t, buffer); uint32(num) != length || err != nil {
		return nil, err
	}

	return append(frameHeader, buffer...), nil
}
