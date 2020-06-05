package encoding

type encodingOptions struct {
	msgType      uint8
	reqType      uint8
	compressType uint8
}

type EncodingOption func(ops *encodingOptions)

func WithMsgType(msgType uint8) EncodingOption {
	return func(ops *encodingOptions) {
		ops.msgType = msgType
	}
}

func WithReqType(reqType uint8) EncodingOption {
	return func(ops *encodingOptions) {
		ops.reqType = reqType
	}
}

func WithCompressType(compressType uint8) EncodingOption {
	return func(ops *encodingOptions) {
		ops.compressType = compressType
	}
}
