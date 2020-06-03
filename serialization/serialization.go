package serialization

import (
	"github.com/golang/protobuf/proto"
)

const (
	Proto   = "proto"   // Protocol Buffer
	Msgpack = "msgpack" // MsgPack
	Json    = "json"    // JSON
)

var registeredSerialization = make(map[string]Serialization)

type Serialization interface {
	Marshal(interface{}) ([]byte, error)
	UnMarshal([]byte, interface{}) error
}

// default serialization type is protobuf
var defaultSerialization *protobufSerial = &protobufSerial{}

func init() {
	registeredSerialization[Proto] = defaultSerialization
}

func RegisterSerialization(name string, serialization Serialization) {
	if registeredSerialization == nil {
		registeredSerialization = make(map[string]Serialization)
	}
	registeredSerialization[name] = serialization
}

func GetSerialization(name string) Serialization {
	if v, ok := registeredSerialization[name]; ok {
		return v
	}
	return defaultSerialization
}

type protobufSerial struct{}

func (s *protobufSerial) Marshal(v interface{}) ([]byte, error) {
	msg := v.(proto.Message)
	return proto.Marshal(msg)
}

func (s *protobufSerial) UnMarshal(b []byte, v interface{}) error {
	msg := v.(proto.Message)
	return proto.Unmarshal(b, msg)
}
