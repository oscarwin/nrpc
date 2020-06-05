package encoding

import (
	"bytes"
	"encoding/binary"
)

const (
	Magic        uint8 = 1
	Version      uint8 = 0
	FrameHeadLen       = 15
	CodecName          = "codec"
)

type Codec interface {
	Encode([]byte, ...EncodingOption) ([]byte, error)
	Decode([]byte) ([]byte, error)
}

type codec struct{}

type frameHeader struct {
	Magic        uint8  // Magic number
	Version      uint8  // Version
	MsgType      uint8  // message type, 0x0: request, 0x1: heartbeat
	ReqType      uint8  // request type, 0x0: call, 0x1: oneway
	CompressType uint8  // compression or not, 0x0: not compression, 0x1: compression
	StreamID     uint16 // stream ID
	Length       uint32 // size of package
	Reserved     uint32 // reserved bytes
}

var registedCodecs map[string]Codec
var defaultCodec = &codec{}

func init() {
	if registedCodecs == nil {
		registedCodecs = make(map[string]Codec)
	}
	registedCodecs[CodecName] = defaultCodec
}

func RegisterCodec(name string, c Codec) {
	if registedCodecs == nil {
		registedCodecs = make(map[string]Codec)
	}
	registedCodecs[name] = c
}

func GetCodec(name string) Codec {
	if v, ok := registedCodecs[name]; ok {
		return v
	}
	return defaultCodec
}

func (c *codec) Encode(b []byte, ops ...EncodingOption) ([]byte, error) {
	options := &encodingOptions{}
	for _, op := range ops {
		op(options)
	}
	frame := frameHeader{
		Magic:        Magic,
		Version:      Version,
		MsgType:      options.msgType,
		ReqType:      options.reqType,
		CompressType: options.compressType,
		StreamID:     0,
		Length:       uint32(len(b)),
	}

	totolLen := FrameHeadLen + len(b)
	buf := bytes.NewBuffer(make([]byte, 0, totolLen))
	if err := binary.Write(buf, binary.BigEndian, frame.Magic); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, frame.Version); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, frame.MsgType); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, frame.ReqType); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, frame.CompressType); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, frame.StreamID); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, frame.Length); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, frame.Reserved); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, b); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *codec) Decode(b []byte) ([]byte, error) {
	return b[FrameHeadLen:], nil
}
