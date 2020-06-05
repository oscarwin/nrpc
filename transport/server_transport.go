package transport

import (
	"fmt"
	"net"
	"sync"
)

type ServerTransport interface {
	Listen() error
	Accept() (Transport, error)
	Close()
}

type serverTransport struct {
	options  *ServerOptions
	listener net.Listener
	lo       sync.RWMutex
}

func NewServerTransport(ops ...ServerOption) *serverTransport {
	ret := &serverTransport{
		options: &ServerOptions{},
	}
	for _, option := range ops {
		option(ret.options)
	}
	return ret
}

func (t *serverTransport) Listen() error {
	t.lo.Lock()
	defer t.lo.Unlock()
	if t.IsListening() {
		return nil
	}
	l, err := net.Listen(t.options.Network, t.options.Address)
	if err != nil {
		return err
	}
	t.listener = l
	return nil
}

func (t *serverTransport) Accept() (Transport, error) {
	t.lo.RLock()
	listener := t.listener
	t.lo.RUnlock()

	if listener == nil {
		return nil, fmt.Errorf("listen is nil")
	}
	conn, err := listener.Accept()
	if err != nil {
		return nil, err
	}

	return NewNSocketWithConnTimeout(conn, t.options.Timeout), nil
}

func (t *serverTransport) IsListening() bool {
	return t.listener != nil
}
