package transport

import "time"

type ServerOptions struct {
	Address   string        // listening address
	Network   string        // network type: tcp, udp
	Protocol  string        // protocol type: proto, json
	Transport string        // transport type: binary
	Timeout   time.Duration // timeout
}

type ServerOption func(*ServerOptions)

func WithAddress(address string) ServerOption {
	return func(ops *ServerOptions) {
		ops.Address = address
	}
}

func WithNetwork(network string) ServerOption {
	return func(ops *ServerOptions) {
		ops.Network = network
	}
}

func WithProtocol(protocol string) ServerOption {
	return func(ops *ServerOptions) {
		ops.Protocol = protocol
	}
}

func WithTransport(transport string) ServerOption {
	return func(ops *ServerOptions) {
		ops.Transport = transport
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(ops *ServerOptions) {
		ops.Timeout = timeout
	}
}
