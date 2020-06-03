package nrpc

import (
	"context"

	"github.com/oscarwin/nrpc/transport"
)

type Service interface {
	Register(string, Handler)
	Serve(*transport.ServerOptions)
	Close()
}

type service struct {
	ctx      context.Context
	handlers map[string]Handler
	opts     *transport.ServerOptions
}

type Handler func(context.Context, interface{}, func(interface{}) error) (interface{}, error)

func (s *service) Register(method string, handler Handler) {
	if s.handlers == nil {
		s.handlers = make(map[string]Handler)
	}
	s.handlers[method] = handler
}
